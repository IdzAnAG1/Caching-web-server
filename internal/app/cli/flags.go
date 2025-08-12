package cli

import "flag"

var (
	CfgPath = flag.String("cfg-path", "", "Path to configuration")
	CfgTest = flag.String("test-config", "", "Test configuration")
)
