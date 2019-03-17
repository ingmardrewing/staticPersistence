package staticPersistence

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func currentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func TestJsonFileNameTemplate(t *testing.T) {
	expected := "doc00234.json"
	actual := fmt.Sprintf(JsonFileNameTemplate(), 234)

	if actual != expected {
		t.Error("Expected", expected, "dtos, but got", actual)
	}
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

func TestReadPagesFromDir(t *testing.T) {
	p := path.Join(currentDir(), "testResources/v2docs/")
	dtos := ReadPagesFromDir(p)

	actual := len(dtos)
	expected := 5

	if actual != expected {
		t.Error("Expected", expected, "dtos, but got", actual)
	}
}

func TestWritePageDtoToJson(t *testing.T) {
	dto1 := getPageDto("1")
	dto2 := getPageDto("2")

	dir := path.Join(currentDir(), "testResources/v2docs_write")
	os.Mkdir(dir, 0755)

	WritePageDtoToJson(dto1, dir, "dto1.json")
	WritePageDtoToJson(dto2, dir, "dto2.json")

	actualSize := len(fs.ReadDirEntriesEndingWith(dir, "json"))
	expectedSize := 2

	if actualSize != expectedSize {
		t.Error("Expected number of files was ", expectedSize, ", but found", actualSize)
	}

	os.RemoveAll(dir)
}

func TestWritePagesToDir(t *testing.T) {
	dir := path.Join(currentDir(), "testResources/v2docs_write")
	os.Mkdir(dir, 0755)

	dtos := []staticIntf.PageDto{getPageDto("1"), getPageDto("2")}
	WritePagesToDir(dtos, dir)

	actualSize := len(fs.ReadDirEntriesEndingWith(dir, "json"))
	expectedSize := 2

	if actualSize != expectedSize {
		t.Error("Expected number of files was ", expectedSize, ", but found", actualSize)
	}

	os.RemoveAll(dir)
}

func getPageDto(name string) staticIntf.PageDto {
	img := NewImageDto("imageTitle"+name,
		"w80SquareUrl"+name,
		"w185SquareUrl"+name,
		"w390SquareUrl"+name,
		"w800SquareUrl"+name,
		"w800Url"+name,
		"w1600Url"+name,
		"maxResolutionUrl"+name)

	dto := NewFilledDto(
		"title"+name,
		"description"+name,
		"content"+name,
		"category"+name,
		"2019-03-01-17-30"+name,
		"path1"+name,
		"filename"+name,
		[]string{"tag1" + name, "tag2" + name},
		[]staticIntf.Image{img})
	return dto
}
