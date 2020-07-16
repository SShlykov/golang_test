package main

import "log"

func errHandler(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}