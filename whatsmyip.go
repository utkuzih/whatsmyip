/*
 * Sal's What's My IP Daemon
 * --------------------------
 * It listens for requests, then it grabs
 * the IP from X-Real-IP and spits it back.
 * --------------------------
 * Copyright (c) 2015, Salvatore LaMendola <salvatore@lamendola.me>
 * All rights reserved.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func spitIP(w http.ResponseWriter, req *http.Request) {
	ip := req.Header.Get("X-Real-IP")
	if len(ip) == 0 {
		ip = req.Header.Get("X-Forwarded-For")
	}
	fmt.Fprintf(w, ip)
}

func main() {
	http.HandleFunc("/", spitIP)
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
