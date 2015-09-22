package plugin

import (
  "log"
  "encoding/json"
  "net/http"
  "github.com/timperman/bouncer/network"
  "github.com/timperman/bouncer/volume"
)

type Handshake struct {
  Implements []string
}

func Start(addr string) {
  http.HandleFunc("/Plugin.Activate", activate)

  v := volume.New()
  n := network.New()

  m := map[string]map[string]func(http.ResponseWriter, *http.Request){
    "POST": {
      "/NetworkDriver.CreateNetwork", n.CreateNetwork),
      "/VolumeDriver.Create", v.Create),
      "/VolumeDriver.Remove", v.Remove),
      "/VolumeDriver.Mount", v.Mount),
      "/VolumeDriver.Unmount", v.Unmount),
      "/VolumeDriver.Path", v.Path,
    },
  }

  for method, routes := range m {
    for route, f := range routes {
      http.HandleFunc(route, handleFuncByMethod(method, f))
    }
  }

  log.Fatal(http.ListenAndServe(addr, nil))
}

func activate(w http.ResponseWriter, r *http.Request) {
  if b, err := json.Marshal(&Handshake{ Implements: []string{ "VolumeDriver" }, }); err == nil {
    w.Write(b)
  }
}
