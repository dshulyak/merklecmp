package benchs

import (
	"flag"
	"path/filepath"
)

var Output = flag.String("output", "", "directory where test artifacts will be stored")

func path(sub string) string {
	return filepath.Join(*Output, sub)
}
