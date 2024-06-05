package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc"
)

type config struct {
	username string
	oauth    string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.username, "username", "", "your twitch username")
	flag.StringVar(&cfg.oauth, "oauth", "", "your twitch oauth")

	flag.Parse()

	client := twitch.NewClient(cfg.username, cfg.oauth)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("===============================")
	fmt.Printf("----------LITERKOWICZ----------\n\n")

	fmt.Println("Name of the channel to connect to:")
	fmt.Printf("-> ")

	channel, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	channel = strings.TrimSpace(channel)
	client.Join(channel)
	go (func() {
		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
			}
		}()

		err := client.Connect()
		if err != nil {
			panic(err)
		}
	})()

	fmt.Printf("\n===============================\n")
	fmt.Println("Prefix emote (leave empty if none):")
	fmt.Printf("-> ")

	prefix, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	prefix = strings.TrimSpace(prefix)

	fmt.Printf("\n===============================\n")
	fmt.Println("Suffix emote (leave empty if none):")
	fmt.Printf("-> ")

	suffix, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	suffix = strings.TrimSpace(suffix)

	fmt.Printf("\n===============================\n")
	fmt.Println("What to say:")
	fmt.Printf("-> ")

	word, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	word = strings.TrimSpace(word)

	fmt.Printf("\n===============================\n")
	fmt.Println("Time between messages(ms) [default: 100]:")
	fmt.Printf("-> ")

	timeoutStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	timeoutStr = strings.TrimSpace(timeoutStr)
	if timeoutStr == "" {
		timeoutStr = "100"
	}

	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Fatal(err)
	}

	literkuj(word, channel, prefix, suffix, timeout, client)

	fmt.Printf("\n===============================\n")
	fmt.Println("JOB DONE :)")
	fmt.Printf("===============================\n")
}

func literkuj(word string, channel string, prefix string, suffix string, timeout int, client *twitch.Client) {
	for _, c := range word {
		var message string

		if prefix != "" && suffix != "" {
			message = fmt.Sprintf("%s %s %s", prefix, string(c), suffix)
		} else if prefix != "" {
			message = fmt.Sprintf("%s %s", prefix, string(c))
		} else if suffix != "" {
			message = fmt.Sprintf("%s %s", string(c), suffix)
		} else {
			message = string(c)
		}

		client.Say(channel, message)
		time.Sleep(time.Millisecond * time.Duration(timeout))
	}
}
