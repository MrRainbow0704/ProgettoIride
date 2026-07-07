// Il pacchetto config gestisce la configurazione dell'applicazione,
// inclusi i percorsi dei file e le impostazioni della telecamera.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	hc "github.com/MrRainbow0704/ProgettoIride/third_party/HarrierControl"
)

type Config struct {
	DeviceID   string `json:"device_id"`
	FFPath     string `json:"ffmpeg_path"`
}

var conf Config
var defaultConf = Config{
	DeviceID:   "",
	FFPath:     "ffmpeg",
}

var (
	RootDir string
	CnfPath string
	BinDir  string
	LogDir  string
	HCPath  string
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("failed to get user home directory: " + err.Error())
	}

	RootDir = filepath.Join(homeDir, ".iride")
	CnfPath = filepath.Join(RootDir, "config.json")
	BinDir = filepath.Join(RootDir, "bin")
	LogDir = filepath.Join(RootDir, "log")
	HCPath = filepath.Join(BinDir, "HarrierControl.exe")

	for _, dir := range []string{RootDir, BinDir, LogDir} {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				panic(fmt.Sprintf("failed to create directory: %v", err))
			}

			if dir == BinDir {
				if err := os.CopyFS(BinDir, hc.Binaries); err != nil {
					panic(fmt.Sprintf("failed to copy HarrierControl binary: %v", err))
				}
			}
		}
	}
}

func LoadConfig() error {
	cnfFile, err := os.Open(CnfPath)
	if err != nil {
		if os.IsNotExist(err) {
			if cnfFile, err = os.Create(CnfPath); err != nil {
				return fmt.Errorf("failed to create config file: %w", err)
			}
			defer cnfFile.Close()

			conf = defaultConf
			if err := json.NewEncoder(cnfFile).Encode(&conf); err != nil {
				return fmt.Errorf("failed to write default config: %w", err)
			}
			return nil
		} else {
			return fmt.Errorf("failed to open config file: %w", err)
		}
	}
	defer cnfFile.Close()

	if err := json.NewDecoder(cnfFile).Decode(&conf); err != nil {
		return fmt.Errorf("failed to decode config file: %w", err)
	}
	return nil
}

func Get() *Config {
	return &conf
}

func SaveConfig() error {
	cnfFile, err := os.Create(CnfPath)
	if err != nil {
		return fmt.Errorf("failed to open config file for writing: %w", err)
	}
	defer cnfFile.Close()

	if err := json.NewEncoder(cnfFile).Encode(&conf); err != nil {
		return fmt.Errorf("failed to write default config: %w", err)
	}
	return nil
}
