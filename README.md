### Gotils

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)
[![Build Status](https://img.shields.io/travis/adelowo/gbowo/master.svg?style=flat-square)](https://travis-ci.org/adelowo/gotils.svg?branch=master)

Some useful set of packages (utilities) i think i'd be needing over and over while writing web services in go.

Current Packages include :

- `Registry`

Allows for service registration and deregistration with consul.

```go

	defaultCfg := consul.DefaultConfig()
	defaultCfg.Address = *discoveryAddr

	client, err := consul.NewClient(defaultCfg)
	if err != nil {
		fatalPrintln(fmt.Sprintf("could not connect to consul... %v", err))
	}

	registrar := registry.NewWithClient(client)
	if err != nil {
		fatalPrintln(fmt.Sprintf("could not build registrar %v\n", err))
	}

	ip, err := registrar.IP()
	if err != nil {
		fatalPrintln(fmt.Sprintf("Could not retrieve IP address %v\n", err))
	}

	svc := &consul.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    config.ServiceName,
		Port:    *httpPort,
		Address: ip.String(),
	}

	if err := registrar.RegisterService(svc); err != nil {
		fatalPrintln(fmt.Sprintf("Could not register service... %v", err))
	}
	
	<-shutdownChan
	registrar.DeRegister(svc)

```

- `Hasher`

Provides a simple wrapper for `crypto/bcrypt`.

```go

package main

import (
  "fmt"
  "github.com/adelowo/gotils/hasher"
)

func main() {
  h := hasher.NewBcryptHasher()

  hashedPassword, err := h.Hash(plainPassword)

  fmt.Println(hashedPassword, err)

  fmt.Println(h.Verify(hashed, plain))
}

```
