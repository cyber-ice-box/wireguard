package service

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"os/exec"
	"text/template"
)

const (
	wgQuickBin     = "sudo wg-quick"
	wgManageBin    = "sudo wg"
	nic            = "wg0"
	configPath     = "/etc/wireguard"
	keepalive      = 25
	configTemplate = `[Interface]
Address = {{.Address}}
ListenPort = {{.Port}}
PrivateKey = {{.PrivateKey}}
SaveConfig = true

PostUp = iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT;
PostDown = iptables -D FORWARD -i %i -j ACCEPT; iptables -D FORWARD -o %i -j ACCEPT;`
)

func wireGuardCmd(cmd string) ([]byte, error) {
	log.Debug().Msgf("Executing command [ %s ]", cmd)
	c := exec.Command("/bin/sh", "-c", cmd)
	out, err := c.CombinedOutput()
	if err != nil {
		log.Error().Err(err).Str("Output:", string(out))
		return nil, err
	}
	return out, nil
}

func (s *Service) addPeer(ip, publicKey string) error {
	cmd := fmt.Sprintf("%s set %s peer %s persistent-keepalive %d allowed-ips %s", wgManageBin, nic, publicKey, keepalive, ip)
	log.Info().Msgf("Adding peer command is %s ", cmd)

	out, err := wireGuardCmd(cmd)
	if err != nil {
		log.Error().Msgf("Error on setting peer on interface %v", err)
		return err
	}
	log.Info().Msgf("Add peer output %s", string(out))
	return nil
}

func (s *Service) deletePeer(ip, publicKey string) error {
	log.Debug().Msgf("Peer with publickey [ %s ] is deleting from %s", publicKey, ip)
	cmd := wgManageBin + " rm " + publicKey + " allowed-ips " + ip

	if _, err := wireGuardCmd(cmd); err != nil {
		return err
	}

	return nil
}

func generateConfig(address, port, privateKey string) (string, error) {
	var tpl bytes.Buffer
	cfg := struct {
		Address    string
		PrivateKey string
		Port       string
	}{
		Address:    address,
		PrivateKey: privateKey,
		Port:       port,
	}
	t, err := template.New("config").Parse(configTemplate)
	if err != nil {
		panic(err)
	}
	err = t.Execute(&tpl, cfg)
	if err != nil {
		panic(err)
	}
	return tpl.String(), nil
}

func writeToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal().Err(err).Msg("failed to close file")
		}
	}()

	if _, err = io.WriteString(file, data); err != nil {
		return err
	}
	return file.Sync()
}

func upInterface() error {
	command := wgQuickBin + " up " + nic
	log.Info().Msgf("Interface %s is called to be up", nic)
	_, err := wireGuardCmd(command)

	return err
}
