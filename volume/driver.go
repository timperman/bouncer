package volume

import (
  "net/http"

  "github.com/docker/docker/volume/local"

  "github.com/timperman/bouncer/util"
)

type Request struct {
  Name string
}

type Driver struct {
  root *local.Root
  names []string
}

func (d *Driver) indexOf(n string) int {
  for i, v := range d.names {
    if v == n {
      return i
    }
  }
  return -1
}

func New() (*Driver) {
  return &Driver{}
}

func (d *Driver) Create(w http.ResponseWriter, r *http.Request) {
  req, err := util.JSONDecode(r)
  if err != nil {
    util.JSONResponse(w, map[string]interface{}{ "Err": err, })
  }

  d.root = local.New("/")
  if d.indexOf(req.Name) < 0 {
    d.names = append(d.names, req.Name)
  }
  util.JSONResponse(w, map[string]interface{}{ "Mountpoint": req.Name, "Err": nil, })
}

func (d *Driver) Remove(w http.ResponseWriter, r *http.Request) {
  req, err := util.JSONDecode(r)

  if i := d.indexOf(req.Name); i >= 0 {
    d.names = append(d.names[:i], d.names[i+1:]...)
    util.JSONResponse(w, map[string]interface{}{ "Err": err, })
  } else {
    util.JSONResponse(w, map[string]interface{}{ "Err": "Volume name not found", })
  }
}

func (d *Driver) Mount(w http.ResponseWriter, r *http.Request) {
  util.JSONResponse(w, map[string]interface{}{ "Mountpoint": util.JSONDecode(r).Name, "Err": nil, })
}

func (d *Driver) Unmount(w http.ResponseWriter, r *http.Request) {
  util.JSONResponse(w, map[string]interface{}{ "Err": nil, })
}

func (d *Driver) Path(w http.ResponseWriter, r *http.Request) {
  util.JSONResponse(w, map[string]interface{}{ "Mountpoint": util.JSONDecode(r).Name, "Err": nil, })
}
