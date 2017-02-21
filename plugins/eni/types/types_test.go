package types

import (
	"testing"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/stretchr/testify/assert"
)

var validConfig = `{"eni":"eni1", "ipv4-address":"10.11.12.13"}`
var configNoENIID = `{"ipv4-address":"10.11.12.13"}`
var configNoIPV4Address = `{"eni":"eni1"}`

func TestNewConfWithValidConfig(t *testing.T) {
	args := &skel.CmdArgs{
		StdinData: []byte(validConfig),
	}
	_, err := NewConf(args)
	assert.NoError(t, err)
}

func TestNewConfWithEmptyConfig(t *testing.T) {
	args := &skel.CmdArgs{
		StdinData: []byte(""),
	}
	_, err := NewConf(args)
	assert.Error(t, err)
}

func TestNewConfWithInvalidConfig(t *testing.T) {
	args := &skel.CmdArgs{
		StdinData: []byte(`{"foo":"eni1"}`),
	}
	_, err := NewConf(args)
	assert.Error(t, err)
}

func TestNewConfWithMissingENIIDConfig(t *testing.T) {
	args := &skel.CmdArgs{
		StdinData: []byte(configNoENIID),
	}
	_, err := NewConf(args)
	assert.Error(t, err)
}

func TestNewConfWithMissingIPV4AddressConfig(t *testing.T) {
	args := &skel.CmdArgs{
		StdinData: []byte(configNoIPV4Address),
	}
	_, err := NewConf(args)
	assert.Error(t, err)
}
