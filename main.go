package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-desktop_settings/internal/app/desktop_settings"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
