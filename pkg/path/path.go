package path

import (
	"os"
	"path/filepath"
)

// heartbeat
var (
	HbHome   = os.Getenv("HEARTBEAT_HOME")
	HbConfig = filepath.FromSlash(HbHome + "/heartbeat.yml")
)

// mate
var (
	HbMatePath    = filepath.FromSlash(HbHome + "/mate")
	HbScripts     = filepath.FromSlash(HbMatePath + "/scripts")
	HbStartScript = filepath.FromSlash(HbScripts + "/start-heartbeat.sh")
)
