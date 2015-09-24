package driver

import (
	"log"
	"net/http"

	"github.com/timperman/bouncer/util"
)

type Request struct {
	Name string
}

type VolumeDriver struct {
	root  string
	names []string
}

func (d *VolumeDriver) indexOf(n string) int {
	for i, v := range d.names {
		if v == n {
			return i
		}
	}
	return -1
}

func New(root string) *VolumeDriver {
	return &VolumeDriver{root: root}
}

func (d *VolumeDriver) Create(w http.ResponseWriter, r *http.Request) {
	req, err := util.JSONDecode(r)
	if err != nil {
		util.JSONResponse(w, map[string]interface{}{"Err": err})
		return
	}

	log.Printf("Create request: %v\n", req)

	name := req["Name"].(string)
	if d.indexOf(name) < 0 {
		d.names = append(d.names, name)
	}
	util.JSONResponse(w, map[string]interface{}{"Err": err})
}

func (d *VolumeDriver) Remove(w http.ResponseWriter, r *http.Request) {
	req, err := util.JSONDecode(r)
	if err != nil {
		util.JSONResponse(w, map[string]interface{}{"Err": err})
		return
	}

	log.Printf("Remove request: %v\n", req)

	if i := d.indexOf(req["Name"].(string)); i >= 0 {
		d.names = append(d.names[:i], d.names[i+1:]...)
		util.JSONResponse(w, map[string]interface{}{"Err": err})
	} else {
		util.JSONResponse(w, map[string]interface{}{"Err": "Volume name not found"})
	}
}

func (d *VolumeDriver) Mount(w http.ResponseWriter, r *http.Request) {
	req, err := util.JSONDecode(r)
	log.Printf("Mount request: %v\n", req)
	util.JSONResponse(w, map[string]interface{}{"Mountpoint": req["Name"], "Err": err})
}

func (d *VolumeDriver) Unmount(w http.ResponseWriter, r *http.Request) {
	req, err := util.JSONDecode(r)
	log.Printf("Unmount request: %v\n", req)
	util.JSONResponse(w, map[string]interface{}{"Err": err})
}

func (d *VolumeDriver) Path(w http.ResponseWriter, r *http.Request) {
	req, err := util.JSONDecode(r)
	log.Printf("Path request: %v\n", req)
	util.JSONResponse(w, map[string]interface{}{"Mountpoint": req["Name"], "Err": err})
}
