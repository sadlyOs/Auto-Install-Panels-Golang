package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	// Библиотека для SSH
	// Библиотека для SFTP
	// для загрузки файла .env
)

func sendRun(ip string, port string, login string, pass string) {
	var output bytes.Buffer
	logFile, err := os.Create("install.log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка создания лог файла: %v\n", err)
		return
	}
	defer logFile.Close()
	// строка успеха
	const success = "Now I will install the best control panel for you!"

	bashScript, errRead := readBash("testBash.bash")
	fmt.Println(bashScript)
	if errRead != nil {
		return
	}

	client, err := sshConn(ip, port, login, pass)
	defer client.Close()

	// Создание SFTP сервера

	sftpClient, err := sftpConn(client)
	defer sftpClient.Close()
	var remotePath string = "/tmp/install.bash"

	remoteFile, err := sftpClient.Create(remotePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка создания удаленного файла: %v\n", err)
		return
	}
	defer remoteFile.Close()

	_, err = io.Copy(remoteFile, strings.NewReader(bashScript))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to make script executable: %v\n", err)
		return
	}

	// Подготовка к захвату вывода

	fmt.Println("Скрипт успешно загружен!")
	command := fmt.Sprintf("chmod +x %s && bash +x %s", remotePath, remotePath)
	err = runCommand(client, command, io.MultiWriter(os.Stdout, &output, logFile), io.MultiWriter(os.Stderr, &output, logFile))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не удалось выполнить скрипт: %v\n", err)
		return
	}

	// Проверка ключевой строки в выводе
	logContent := output.String()
	if !strings.Contains(logContent, success) {
		fmt.Fprintf(os.Stderr, "Script execution failed: Key phrase '%s' not found in logs\n", success)
		return
	}
	fmt.Println("Установка прошла успешно")
}
