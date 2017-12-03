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

import s "strings"
import "strconv"

import (
	"fmt"
	"log"
	"net/http"
)

func printIP(w http.ResponseWriter, req *http.Request) {
	ip := req.Header.Get("X-Real-IP")
	if len(ip) == 0 {
		ip = req.Header.Get("X-Forwarded-For")
	}
	fmt.Fprintf(w, ip)

	length := req.URL.Query().Get("length")
	if len(length) > 0 {
	    i, err := strconv.Atoi(length)
        if err == nil {
            fmt.Fprintf(w, " ")
            if (i <= 10*1024*1024) {
	            fmt.Fprintf(w, s.Repeat("a", i))
	        }
	    }
	}
}

func main() {
	http.HandleFunc("/", printIP)
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
