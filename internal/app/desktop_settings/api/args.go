package api

import "github.com/maldan/go-restserver"

type ArgsEmpty struct {
}

type ArgsPath struct {
	Context *restserver.RestServerContext
	Path    string `json:"path"`
}

type ArgsBackupSettings struct {
	Path  string `json:"path"`
	Value string `json:"value"`
}
