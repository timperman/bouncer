package network

import (
  "net/http"

  "github.com/docker/libnetwork"
  "github.com/docker/libnetwork/config"
  "github.com/docker/libnetwork/options"
  "github.com/docker/libnetwork/netlabel"
)

type Driver libnetwork.NetworkController

func New() (*Driver, error) {
  opt := map[string]interface{}{
    netlabel.GenericData: options.Generic{},
  }
  if controller, err := libnetwork.New(config.OptionDriverConfig("host", opt)); err != nil {
    return nil, err
  }

  return controller, nil
}

func (d *Driver) CreateNetwork(w http.ResponseWriter, r *http.Request) {
  
}
