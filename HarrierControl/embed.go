package harriercontrol

import "embed"

//go:embed *.dll HarrierControl.exe
var Binaries embed.FS
