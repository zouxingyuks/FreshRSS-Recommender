// Package version provides version info.
package version

import (
	"fmt"
)

const (
	// LOGO is system logo.
	LOGO = `
======================================================================================
  _____              _     _____   _____ _____ 
 |  ___| __ ___  ___| |__ |  __ \ / ____/ ____|
 | |_ | '__/ _ \/ __| '_ \| |__) | (___| (___  
 |  _|| | |  __/\__ \ | | |  _  / \___ \\___ \ 
 |_|  |_|  \___||___/_| |_|_| \_\ ____) |___) |
  _____                                      _           
 |  __ \                                    | |          
 | |__) |___  ___ ___  _ __ ___  _ __ ___  | | ___ _ __ 
 |  _  / _ \/ __/ _ \| '_ \ _ \| '_ \ _ \ | |/ _ \ '__|
 | | \ \  __/ (_| (_) | | | | | | | | | | || |  __/ |   
 |_|  \_\___|\___\___/|_| |_| |_|_| |_| |_||_|\___|_|   
======================================================================================`
)

var (
	// VERSION is version info.
	VERSION = "debug"

	// BUILDTIME  build time.
	BUILDTIME = "unknown"

	// GITHASH git hash for release.
	GITHASH = "unknown"
)

// ShowVersion shows the version info.
func ShowVersion() {
	fmt.Println(FormatVersion())
}

// FormatVersion returns service's version.
func FormatVersion() string {
	return fmt.Sprintf("\nVersion: %s\nBuildTime: %s\nGitHash: %s\n", VERSION, BUILDTIME, GITHASH)
}

// GetStartInfo returns start info that includes version and logo.
func GetStartInfo() string {
	startInfo := fmt.Sprintf("%s\n\n%s\n", LOGO, FormatVersion())
	return startInfo
}

// Version ...
func Version() *SysVersion {
	return &SysVersion{
		Version: VERSION,
		Hash:    GITHASH,
		Time:    BUILDTIME,
	}
}

// SysVersion describe a binary version
type SysVersion struct {
	Version string `json:"version"`
	Hash    string `json:"hash"`
	Time    string `json:"time"`
}
