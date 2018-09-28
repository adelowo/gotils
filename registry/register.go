package registry

import (
	"net"

	consul "github.com/hashicorp/consul/api"
)

type Client struct {
	inner *consul.Client
}

func NewWithConfig(conf *consul.Config) (*Client, error) {
	inner, err := consul.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &Client{
		inner: inner,
	}, nil
}

func NewWithClient(client *consul.Client) *Client {
	return &Client{inner: client}
}

func New(addr string) (*Client, error) {
	conf := consul.DefaultConfig()
	conf.Address = addr

	return NewWithConfig(conf)
}

// RegisterService takes a service definition and adds it to consul
// It disregards whatever ID was given and uses a uuid internal so as to prevent clashes
func (c *Client) RegisterService(svc *consul.AgentServiceRegistration) error {
	return c.inner.Agent().ServiceRegister(svc)
}

// Deregister removes a client from consul's service list
func (c *Client) DeRegister(svc *consul.AgentServiceRegistration) error {
	return c.inner.Agent().ServiceDeregister(svc.ID)
}

func ipAddr() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil || ipnet.IP.To16() != nil {
				return ipnet.IP, nil
			}
		}
	}

	return nil, nil
}
