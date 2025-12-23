package envv

import "os"

func RunMain() {
	if err := Execute(); err != nil {
		os.Exit(1)
	}
}
