package api

import (
	"os"

	"github.com/maldan/gam-app-desktop_settings/internal/app/desktop_settings/core"
	"github.com/maldan/go-cmhp/cmhp_file"
)

type MainApi struct {
}

func (r MainApi) GetBackground(args ArgsBackground) *os.File {
	f, _ := os.Open(core.DataDir + "/../maldan-gam-app-desktop/background/" + args.Id + ".jpeg")
	return f
}

func (r MainApi) PostBackground(args ArgsBackground) {
	os.MkdirAll(core.DataDir+"/../maldan-gam-app-desktop/background/", 0777)
	cmhp_file.WriteBin(core.DataDir+"/../maldan-gam-app-desktop/background/"+args.Id+".jpeg", args.Files[0])
}
