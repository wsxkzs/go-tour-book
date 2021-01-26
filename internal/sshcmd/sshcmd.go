package sshcmd

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

func SshExcuteCmd(host string, user string, password string, port int, command string) {
	sshHost := host
	sshUser := user
	sshPassword := password
	sshPort := port
	// 创建ssh登陆
	config := &ssh.ClientConfig{
		Timeout: time.Second,
		User:    sshUser,

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("创建ssh Client 失败,", err)
	}
	defer sshClient.Close()
	//创建ssh session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("创建ssh session失败，", err)
	}
	defer session.Close()

	//执行远程命令
	combo, err := session.CombinedOutput(command)
	if err != nil {
		log.Fatal("远程执行命令失败，", err)
	}
	log.Println("命令输出：", string(combo))

}
