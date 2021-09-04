GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o app .
cp app.svg /mnt/orangepi/root/.gam-app/maldan-gam-app-desktop_settings-v0.0.2
cp app /mnt/orangepi/root/.gam-app/maldan-gam-app-desktop_settings-v0.0.2