package onlinelabs

import (
	"fmt"
	"io/ioutil"
	"os"

	"code.google.com/p/go.crypto/ssh"
	"github.com/mitchellh/multistep"
)

func sshAddress(state multistep.StateBag) (string, error) {
	config := state.Get("config").(*config)
	ipAddress := state.Get("server_ip").(string)
	return fmt.Sprintf("%s:%d", ipAddress, config.SSHPort), nil
}

func sshConfig(state multistep.StateBag) (*ssh.ClientConfig, error) {
	config := state.Get("config").(*config)
	// privateKey := state.Get("privateKey").(string)

	// signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	pkFile, err := os.Open(config.SSHPrivateKeyFile)
	if err != nil {
		return nil, err
	}
	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(pkBytes)
	if err != nil {
		return nil, fmt.Errorf("Error setting up SSH config: %s", err)
	}

	return &ssh.ClientConfig{
		User: config.SSHUsername,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}, nil
}
