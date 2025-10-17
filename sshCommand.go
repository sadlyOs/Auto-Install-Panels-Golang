package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/ssh"
)

func runCommand(client *ssh.Client, command string, stdout, stderr io.Writer) error {
	session, err := client.NewSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не удалось создать новую сессию", err)
	}
	defer session.Close()
	session.Stdout = stdout
	session.Stderr = stderr
	return session.Run(command)
}
