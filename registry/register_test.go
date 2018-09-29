// +build integration

package registry

import (
	"os"
	"testing"

	consul "github.com/hashicorp/consul/api"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/require"
)

const (
	CONSUL_ADDRESS = "CONSUL_HTTP_ADDR"
)

func consulAddr() string {
	return os.Getenv(CONSUL_ADDRESS)
}

func TestClient_Register(t *testing.T) {

	client, err := New(consulAddr())
	require.NoError(t, err)

	svc := &consul.AgentServiceRegistration{
		ID:   uuid.New(),
		Port: 3000,
		Name: "oops",
	}

	err = client.RegisterService(svc)
	require.NoError(t, err)
}

func TestClient_DeRegister(t *testing.T) {

	client, err := New(consulAddr())
	require.NoError(t, err)

	svc := &consul.AgentServiceRegistration{
		ID:   uuid.New(),
		Port: 3000,
		Name: "oops",
	}

	err = client.RegisterService(svc)
	require.NoError(t, err)

	require.NoError(t, client.DeRegister(svc))
}
