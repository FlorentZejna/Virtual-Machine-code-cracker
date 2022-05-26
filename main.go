package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"os"

	"log"
)

func readFile(f string) (data []string, err error) {
	b, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, err
}

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

	passwords, err := readFile("/home/beriflapp/rootPass/words_alpha.txt")
	if err != nil {
		log.Println("Can't read file!")
		os.Exit(1)
	}

	for _, pass := range passwords {
		fmt.Println(pass)
	}

}
