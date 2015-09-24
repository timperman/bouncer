package plugin

import (
	"encoding/json"
	"github.com/timperman/bouncer/driver"
	"log"
	"net/http"
)

type Handshake struct {
	Implements []string
}

func Start(addr string) {
	http.HandleFunc("/Plugin.Activate", activate)

	v := driver.New("/")

	m := map[string]map[string]func(http.ResponseWriter, *http.Request){
		"POST": {
			"/VolumeDriver.Create":  v.Create,
			"/VolumeDriver.Remove":  v.Remove,
			"/VolumeDriver.Mount":   v.Mount,
			"/VolumeDriver.Unmount": v.Unmount,
			"/VolumeDriver.Path":    v.Path,
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
	log.Println("Activate call")
	if b, err := json.Marshal(&Handshake{Implements: []string{"VolumeDriver"}}); err == nil {
		w.Write(b)
	}
}
