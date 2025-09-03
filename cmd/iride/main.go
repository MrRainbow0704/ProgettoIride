// Entrypoint per l'applicazione Iride.exe
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	hc "github.com/MrRainbow0704/ProgettoIride/HarrierControl"
	"github.com/MrRainbow0704/ProgettoIride/internal/camera"
	"github.com/MrRainbow0704/ProgettoIride/internal/gui"
	"github.com/MrRainbow0704/ProgettoIride/internal/version"
	"github.com/MrRainbow0704/ProgettoIride/internal/video"
	"github.com/lxn/walk"
)

var v = flag.Bool("v", false, "query version")

func main() {
	flag.Parse()
	if *v {
		fmt.Println(version.Get())
		os.Exit(0)
	}

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	walk.Resources.SetRootDirPath(camera.TmpDir)
	if _, err := os.Stat(camera.Prog); errors.Is(err, os.ErrNotExist) {
		if err := os.CopyFS(camera.ProgDir, hc.Binaries); err != nil {
			return err
		}
	}

	cam, err := video.NewCamera()
	if err != nil {
		panic(err)
	}
	defer cam.Destroy()

	_, err = gui.Window(cam).Run()
	return err
}
