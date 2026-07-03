package harriercontrol

import "embed"


//go:embed *.dll *.flash HarrierControl.exe
var Binaries embed.FS
