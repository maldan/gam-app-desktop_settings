package core

import "encoding/json"

type Application struct {
	Name    string `json:"name"`
	Author  string `json:"author"`
	Version string `json:"version"`
	Path    string `json:"path"`
	Backup  string `json:"backup"`
}

type Config struct {
	Backup map[string]string `json:"backup"`
}

func To(m map[string]string, v interface{}) {
	out, _ := json.Marshal(m)
	json.Unmarshal(out, v)
}

var DataDir = ""
