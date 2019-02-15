package staticPersistence

import (
	"path"
	"runtime"
	"testing"
)

func currentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func TestCleanStringValue(t *testing.T) {
	dirtyString := `What
 ever you "say" or don't \"say\"`

	actual := cleanStringValue(dirtyString)
	expected := `What ever you \"say\" or don't \"say\"`

	if actual != expected {
		t.Error("Expected", expected, "dtos, but got", actual)
	}
}

/*
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

	WritePageDtoToJson(dto1, dir, "t2.json")
	WritePageDtoToJson(dto2, dir, "t2.json")

	actualSize := len(fs.ReadDirEntriesEndingWith(dir))
	expectedSize := 3

	if actual != expected {
		t.Error("Expected number of files was ", expectedSize, ", but found", actualSize)
	}

	os.RemoveAll(dir)
}
*/
