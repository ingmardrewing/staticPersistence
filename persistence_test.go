package staticPersistence

import (
	"path"
	"runtime"
)

func currentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
