package service

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os/exec"
)

const (
	iptablesNat = `sudo iptables -t nat -%s POSTROUTING -o eth+ -s %s -d %s -j MASQUERADE -m comment --comment "client %s"`
)

func (s *Service) addNATRule(name, ip, destCidr string) error {
	command := fmt.Sprintf(iptablesNat, "A", ip, destCidr, name)

	log.Info().Msgf("Adding NAT rule command is %s ", command)

	return exec.Command("/bin/sh", "-c", command).Run()
}

func (s *Service) deleteNATRule(name, ip, destCidr string) error {
	command := fmt.Sprintf(iptablesNat, "D", ip, destCidr, name)

	log.Info().Msgf("Deleting NAT rule command is %s ", command)

	return exec.Command("/bin/sh", "-c", command).Run()
}
