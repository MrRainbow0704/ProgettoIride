package version

import "strings"

var version = "v0.0.0"

func Get() string {
	return version
}

func IsDev() bool {
	v := strings.ToLower(version)
	return strings.Contains(v, "dev") || strings.Contains(v, "snapshot")
}
