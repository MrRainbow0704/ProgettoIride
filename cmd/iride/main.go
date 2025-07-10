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
	if _, err := os.Stat(camera.Prog); errors.Is(err, os.ErrNotExist) {
		if err := os.CopyFS(camera.ProgDir, hc.Binaries); err != nil {
			return err
		}
	}
	_, err := gui.Window().Run()
	return err
}
