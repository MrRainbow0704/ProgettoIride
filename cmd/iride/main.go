// Entrypoint per l'applicazione
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/MrRainbow0704/ProgettoIride/internal/camera"
	"github.com/MrRainbow0704/ProgettoIride/internal/video"
	"github.com/MrRainbow0704/ProgettoIride/internal/config"
	"github.com/MrRainbow0704/ProgettoIride/internal/gui"
	"github.com/MrRainbow0704/ProgettoIride/internal/log"
	"github.com/MrRainbow0704/ProgettoIride/internal/version"
)

var conf *config.Config

func main() {
	versionFlag := flag.Bool("v", false, "query version")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("Iride Version: %s (Dev: %v)\n", version.Get(), version.IsDev())
		return
	}

	// Crea file di log con timestamp
	logFile, err := os.Create(
		filepath.Join(config.LogDir, fmt.Sprintf("log_%d.log", time.Now().Unix())),
	)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}
	defer logFile.Close()

	log.LoadLogger(logFile, os.Stderr)

	if err := run(); err != nil {
		log.Fatalf("application error: %v", err)
	}
}

func run() error {
	// Carica la configurazione
	log.Info("loading configuration...")
	if err := config.LoadConfig(); err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}
	conf = config.Get()
	log.Infof("configuration loaded successfully. Configuration: %+v", *conf)

	// Inizializza la fotocamera
	if conf.DeviceID != "" {
		log.Infof("initializing camera with device id: %s", conf.DeviceID)
		cam := video.NewCamera(conf.DeviceID, camera.GetResolution())
		if cam == nil {
			return fmt.Errorf("failed to initialize camera with device id: %s", conf.DeviceID)
		}
		log.Info("camera initialized successfully")

		if err := camera.FetchCameraStatuses(); err != nil {
			return fmt.Errorf("failed to fetch camera statuses: %w", err)
		}

		// Avvia loop di cattura video in background
		log.Infof("starting background video loop with device id: %s", conf.DeviceID)
		if err := video.StartBackgroundVideoLoop(cam, gui.EmbeddedVideo); err != nil {
			return fmt.Errorf("failed to start background video loop: %w", err)
		}
		log.Info("background video loop started successfully")
	}

	// Avvia l'interfaccia grafica
	gui.GUI().ShowAndRun()

	return nil
}
