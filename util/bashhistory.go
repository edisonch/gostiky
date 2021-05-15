package util

import "github.com/tredoe/osutil"

func ClearBashHistory() error {
	return osutil.Exec("history -c && history -w")
}
