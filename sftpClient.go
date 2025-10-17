package main

import (
	"fmt"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func sftpConn(sshClient *ssh.Client) (*sftp.Client, error) {
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка FTP подключения: %v\n", err)
		return nil, err
	}
	return sftpClient, nil
}
