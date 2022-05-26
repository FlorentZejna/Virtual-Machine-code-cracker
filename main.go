package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"

	"log"
)

var username string = "me"
var password string = "1234"
var host string = "192.168.100.145"

func main() {
	hostKeyCallback, err := knownhosts.New("/home/beriflapp/.ssh/known_hosts")
	if err != nil {
		log.Fatal(err)
	}
	
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: hostKeyCallback,
	}

	conn, errCon := ssh.Dial("tcp", host+":22", config)
	if errCon != nil {
		log.Fatal("unable to connect: ", errCon)
	}
	fmt.Println("access")
	defer conn.Close()

}
