# Progetto Iride

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white)](https://go.dev/)
[![Windows](https://custom-icon-badges.demolab.com/badge/Windows-0078D6?logo=windows11&logoColor=white)](https://www.microsoft.com)

<div align=center><img src="internal/gui/assets/iride.ico"></div>

Questo progetto punta alla creazione di un macchinario video-ingranditore per assistere persone con disabilità visive che influiscono sulla visione da lontano.
Utilizza una telecamera Harrier-10x-AF della ActiveSilicon e fa uso della relativa SDK per USB.

## Requisiti

Per il corretto funzionamento del programma sono richiesti i seguenti elementi:

- Una telecamera [Harrier-10x-AF](https://www.activesilicon.com/products/harrier-10x-af-zoom-camera/) (o simili, nonostante potrebbero non essere supportate) e la relativa [scheda per interfacciarcisi](https://www.activesilicon.com/products/harrier-usb-hdmi-camera-interface-board/). (Acquisibili anche come [pezzo unico](https://www.activesilicon.com/products/harrier-10x-af-zoom-camera-with-usb-hdmi-output/))
- Un cavo USB-C 3.0 o superiore con la capacità di trasferire corrente e contemporaneamente trasportare dati al di sopra dei 10 Mbps.
- [FFmpeg](https://www.ffmpeg.org/) installato sul proprio sistema e all'interno della variabile d'ambiente `PATH`. Installabile con [WinGet](https://github.com/microsoft/winget-cli) su windows (`winget install Gyan.FFmpeg`), [Homebrew](https://brew.sh/) su MacOS (`brew install ffmpeg`) e su Linux con il vostro package manager di fiducia.

## Compilazione

Dopo aver installato tutto il necessario (vedi [qui](#requisiti-per-la-compilazione)) clona questa repo in una cartella del tuo computer:

```cmd
C:\Users\[USER]\Somewhere> git clone https://github.com/MrRainbow0704/ProgettoIride.git
```

semplicemente lanciare il comando `make` o `mingw32-make`:

```cmd
C:\Users\[USER]\Somewhere\ProgettoIride> make
```

e troverai l'eseguibile appena compilato all'interno della cartella `./bin/` della cartella.

Per compilare il programma in modalità di rilascio è sufficente passare a `make` la variabile `RELEASE` con un qualunque valore:

```cmd
C:\Users\[USER]\Somewhere\ProgettoIride> make RELEASE=1
```


### Requisiti per la Compilazione

Se si vuole compilare il programma a partire dal codice sorgente, oltre ai sopraindicati requisiti saranno necessari anche:

- Il linguaggio di programmazione [GoLang](https://go.dev/).
- La toolchain GCC, su windows installabile attraverso [MinGW (MSYS2)](https://www.msys2.org/).

Entrambi i comandi devono essere all'interno della variabile d'ambiente `PATH`, verificabile con il comando `which` (per sistemi UNIX) o `where.exe` (per sistemi windows)