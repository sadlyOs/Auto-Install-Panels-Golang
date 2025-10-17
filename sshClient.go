package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func sshConn(ip string, port string, login string, pass string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: login,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", ip+":"+port, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка подключения: %v\n", err)
		return nil, err
	}
	return client, nil
}
