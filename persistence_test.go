package staticPersistence

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/ingmardrewing/fs"
)

func currentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func TestReadPagesFromDir(t *testing.T) {
	p := path.Join(currentDir(), "testResources/posts/")
	dtos := ReadPagesFromDir(p)

	actual := len(dtos)
	expected := 5

	if actual != expected {
		t.Error("Expected", expected, "dtos, but got", actual)
	}

	dto1 := dtos[0]
	dto2 := dtos[1]
	dir := path.Join(currentDir(), "testResources/writeTest")
	os.Mkdir(dir, 0755)

	WriteMarginalDtoToJson(dto1, dir, "t1.json")

	WritePostDtoToJson(dto2, dir, "t2.json")

	actualSize := len(fs.ReadDirEntriesEndingWith(dir))
	expectedSize := 3

	if actual != expected {
		t.Error("Expected number of files was ", expectedSize, ", but found", actualSize)
	}

	os.RemoveAll(dir)
}
