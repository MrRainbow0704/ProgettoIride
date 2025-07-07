package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	hc "github.com/MrRainbow0704/ProgettoIride/HarrierControl"
	"github.com/MrRainbow0704/ProgettoIride/internal/camera"
	"github.com/MrRainbow0704/ProgettoIride/internal/version"
)

func main() {
	v := flag.Bool("v", false, "query version")
	flag.Parse()
	if *v {
		fmt.Println(version.Get())
		return
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
	return nil
}
