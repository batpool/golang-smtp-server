A SMTP Server Package With TLS [![](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/batpool/golang-smtp-server)
=============================

Quick Start
===========
> `go get github.com/batpool/golang-smtp-server`

```go
package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"


	"github.com/batpool/golang-smtp-server"
)

func main() {
	handler := func(c *smtpsrv.Context) error {
		username, password, _ := c.User()
		if username != "username" || password != "satya" {
			return errors.New("invalid username or password")
		}
		return nil
	}

	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	cfg := smtpsrv.ServerConfig{
		BannerDomain:    "mail.batpool.com",
		ListenAddr:      ":1587",
		MaxMessageBytes: 5 * 1024,
		Handler:         handler,
		TLSConfig:       config,
	}

	fmt.Println(smtpsrv.ListenAndServeTLS(&cfg))
}

```
