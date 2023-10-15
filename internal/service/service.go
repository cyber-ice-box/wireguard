package service

import (
	"context"
	"fmt"
	"github.com/cyber-ice-box/wireguard/internal/delivery/repository/postgres"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog/log"
	"path"
	"sync"
)

type client struct {
	name            string
	ip              string
	publicKey       string
	allowedDestCIDR string
}

type Service struct {
	m                sync.RWMutex
	serverIp         string
	serverPrivateKey string
	clients          map[string]*client
}

func NewService() *Service {
	return &Service{clients: make(map[string]*client)}
}

func (s *Service) InitServer(address, port, privateKey string) error {

	config, err := generateConfig(address, port, privateKey)
	if err != nil {
		return fmt.Errorf("failed to create wireguard config %v", err)
	}

	if err = writeToFile(path.Join(configPath, nic+".conf"), config); err != nil {
		return fmt.Errorf("failed to write wireguard config to file %v", err)
	}

	if err = upInterface(); err != nil {
		return fmt.Errorf("failed to make wireguard interface up %v", err)
	}

	log.Debug().Str("Address: ", address).
		Str("ListenPort: ", port).
		Str("PrivateKey: ", privateKey).Msgf("Interface %s created and it is up", nic)

	return nil
}

func (s *Service) InitServerClients(repository *postgres.Queries) error {
	defer func() {
		if err := repository.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close pg connection")
		}
	}()

	participantsInRunningEvents, err := repository.GetParticipantsInRunningEvents(context.Background())
	if err != nil {
		return err
	}

	var errs error

	for _, p := range participantsInRunningEvents {
		if err = s.AddClient(p.ID.String(), p.VpnIpAddress.IPNet.String(), p.VpnPublicKey); err != nil {
			errs = multierror.Append(err)
			continue
		}
		if p.LabPermitted {
			if err = s.AllowClientAccess(p.ID.String(), p.LabCidr.IPNet.String()); err != nil {
				errs = multierror.Append(err)
			}
		}
	}

	if errs != nil {
		return errs
	}

	return nil
}

func (s *Service) AddClient(name, ip, publicKey string) error {
	s.m.Lock()

	if _, ex := s.clients[name]; ex {
		s.m.Unlock()
		return fmt.Errorf("client with name [ %s ] does exist", name)
	}
	s.clients[name] = &client{
		name:            name,
		ip:              ip,
		publicKey:       publicKey,
		allowedDestCIDR: "",
	}

	s.m.Unlock()

	if err := s.addPeer(ip, publicKey); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteClient(name string) error {
	s.m.Lock()
	defer s.m.Unlock()

	cl, ex := s.clients[name]
	if !ex {
		s.m.Unlock()
		return fmt.Errorf("cl with name [ %s ] does not exist", name)
	}

	if err := s.deletePeer(cl.ip, cl.publicKey); err != nil {
		return err
	}

	if err := s.deleteNATRule(cl.name, cl.ip, cl.allowedDestCIDR); err != nil {
		return err
	}

	delete(s.clients, name)

	return nil
}

func (s *Service) AllowClientAccess(name string, destCIDR string) error {
	s.m.Lock()
	defer s.m.Unlock()

	cl, ex := s.clients[name]
	if !ex {
		s.m.Unlock()
		return fmt.Errorf("cl with name [ %s ] does not exist", name)
	}

	if cl.allowedDestCIDR == destCIDR {
		return nil
	}

	if cl.allowedDestCIDR != "" {
		if err := s.deleteNATRule(cl.name, cl.ip, cl.allowedDestCIDR); err != nil {
			return err
		}
	}

	if err := s.addNATRule(cl.name, cl.ip, destCIDR); err != nil {
		return err
	}

	s.clients[name].allowedDestCIDR = destCIDR

	return nil
}

func (s *Service) DenyClientAccess(name string) error {
	s.m.Lock()
	defer s.m.Unlock()

	cl, ex := s.clients[name]
	if !ex {
		s.m.Unlock()
		return fmt.Errorf("cl with name [ %s ] does not exist", name)
	}

	if cl.allowedDestCIDR != "" {
		if err := s.deleteNATRule(cl.name, cl.ip, cl.allowedDestCIDR); err != nil {
			return err
		}
	}

	s.clients[name].allowedDestCIDR = ""

	return nil
}
