package version

import (
	"fmt"
	"os"
)

// set using the '-X github.com/containers/libhvee/pkg/version.gitVersion' linker flag
var gitVersion string = ""

func ModuleVersion() string {
	return fmt.Sprintf("%s version %s", os.Args[0], moduleVersion())
}

func moduleVersion() string {
	// This will be set when building from git using make
	return gitVersion
}
