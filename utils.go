package hispeed2

import (
	"regexp"
	"runtime"
	"time"
)

var funcNameRe = regexp.MustCompile(`^.*\.(.*)$`)

// Utility to check the time it takes for a function to execute...
func (h *HiSpeed2) LoadTime(start time.Time) {
	elapsed := time.Since(start)
	pc, _, _, _ := runtime.Caller(1) // pc = program caller :)
	funcObj := runtime.FuncForPC(pc)
	name := funcNameRe.ReplaceAllString(funcObj.Name(), "$1")

	h.InfoLog.Printf("Load Time: %s took %s", name, elapsed)
}
