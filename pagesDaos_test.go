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
	a := new(abstractPageDao)
	a.data = readViaFc(path, file)
	a.ExtractFromJson()
	return a.Dto()
}

func TestReadFile(t *testing.T) {
	path := path.Join(currentDir(), "testResources", "posts")
	file := "version0.json"

	actual := readAndGetDto(path, file)
	p0 := NewPostDAO(readViaFc(path, file), path, "")
	p0.ExtractFromJson()
	expected := p0.Dto()

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected\n", actual, "to be\n", expected)
	}

	file = "version1.json"

	actual = readAndGetDto(path, file)
	expected = newVersion1Dto()

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected", actual, "to be", expected)
	}

	file = "version2.json"

	p1 := NewNarrativeDAO(readViaFc(path, file), path, "")
	p1.ExtractFromJson()
	expected = p1.Dto()
	actual = readAndGetDto(path, file)

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected", actual, "to be", expected)
	}

	file = "version3.json"

	p2 := NewMarginalDAO(readViaFc(path, file), path, "")
	p2.ExtractFromJson()
	expected = p2.Dto()
	actual = readAndGetDto(path, file)

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

/*
func TestFindVersion_returns_zero_if_no_version_defined(t *testing.T) {
	json := []byte(`{"noversiongiven":""}`)
	expected := 0

	actual := FindJsonVersion(json)

	if expected != actual {
		t.Error("Expected", expected, ", but got", actual)
	}
}

func TestFindVersion_returns_one_if_version_one_is_defined(t *testing.T) {
	json := []byte(`{"version":1}`)
	expected := 1

	actual := FindJsonVersion(json)

	if expected != actual {
		t.Error("Expected", expected, ", but got", actual)
	}
}

func TestFindVersion_returns_one_if_new_empty_dao_is_needed(t *testing.T) {
	expected := 1

	actual := FindJsonVersion(nil)

	if expected != actual {
		t.Error("Expected", expected, ", but got", actual)
	}
}

func TestNewPostDAO_without_json_data_returns_newest_version(t *testing.T) {
	expected := `{
	"version":1,
	"thumbImg":"",
	"postImg":"",
	"filename":"",
	"id":0,
	"date":"",
	"url":"",
	"title":"",
	"title_plain":"",
	"excerpt":"",
	"content":"",
	"dsq_thread_id":""
	"thumbBase64":""
}`
	dto := NewFilledDto(0,
		"", "", "", "", "",
		"", "", "", "", "",
		"", "", "", "")

	d := NewPostDAO(nil, "", "")
	d.Dto(dto)

	actual := string(d.FillJson())

	if actual != expected {
		t.Error("Expected", expected, ", but got", actual)
	}
}
*/
