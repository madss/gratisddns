package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func updateDns(username, password, domain, host string) {
	transport := &http.Transport{}
	client := &http.Client{Transport: transport}
	url := fmt.Sprintf(
		"https://ssl.gratisdns.dk/ddns.phtml?u=%s&p=%s&d=%s&h=%s",
		username,
		password,
		domain,
		host,
	)

	log.Printf("Requesting %s\n", url)

	res, err := client.Get(url)
	if err != nil {
		log.Printf("Failed to send request: %q\n", err)
		return
	}
	msg, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Printf("Failed to read response: %q\n", err)
	}
	log.Printf("Got response: %s\n", msg)
}

func main() {
	var username, password, domain, host string
	var schedule int

	flag.StringVar(&username, "u", "", "Your gratisdns `username`")
	flag.StringVar(&password, "p", "", "Your gratisdns ddns `password`")
	flag.StringVar(&domain, "d", "", "Your `domain` (e.g. example.com)")
	flag.StringVar(&host, "h", "", "A `host` from your A-records (e.g. www.example.com)")
	flag.IntVar(&schedule, "s", 0, "Schedule a dns update every `n` hours")

	flag.Parse()

	if username == "" || password == "" || domain == "" || host == "" {
		flag.PrintDefaults()
		return
	}

	updateDns(username, password, domain, host)
	if schedule > 0 {
		for _ = range time.Tick(time.Duration(schedule) * time.Hour) {
			updateDns(username, password, domain, host)
		}
	}
}
