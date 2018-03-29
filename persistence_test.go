package staticPersistence

import (
	"path"
	"reflect"
	"runtime"
	"testing"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func currentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func readViaFc(path, file string) []byte {
	fc := fs.NewFileContainer()
	fc.SetPath(path)
	fc.SetFilename(file)
	fc.Read()
	return fc.GetData()
}

func readAndGetDto(path, file string) staticIntf.PageDto {
	a := new(pageDaoReader)
	a.data = readViaFc(path, file)
	a.ExtractFromJson()
	return a.Dto()
}

func TestReadFile(t *testing.T) {
	path := path.Join(currentDir(), "testResources", "posts")
	file := "version1.json"

	actual := readAndGetDto(path, file)
	expected := newVersion1Dto()

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected", actual, "to be", expected)
	}
}

func newVersion1Dto() staticIntf.PageDto {
	return NewFilledDto(1,
		"titleValue", "title_plainValue", "thumbImageValue",
		"postImageValue", "excerptValue", "dsq_thread_idValue",
		"dateValue", "contentValue", "urlValue", "",
		"", "", "filenameValue", "")
}
