package version

import (
	"fmt"
	"os"
	"runtime/debug"
)

// set using the '-X github.com/containers/libhvee/pkg/version.gitVersion' linker flag
var gitVersion string = ""

func ModuleVersion() string {
	return fmt.Sprintf("%s version %s", os.Args[0], moduleVersion())
}

func moduleVersion() string {
	switch {
	// This will be set when building from git using make
	case gitVersion != "":
		return gitVersion
	// moduleVersionFromBuildInfo() will be set when using `go install`
	default:
		return moduleVersionFromBuildInfo()
	}
}

func moduleVersionFromBuildInfo() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	if info.Main.Version == "(devel)" {
		return ""
	}
	return info.Main.Version
}
