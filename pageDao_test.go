package staticPersistence

import (
	"path"
	"reflect"
	"testing"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

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

func TestFillJson(t *testing.T) {
	dto := NewFilledDto(42,
		"titleValue",
		"titlePlainValue",
		"thumbUrlValue",
		"imageUrlValue",
		"descriptionValue",
		"disqusIdValue",
		"createDateValue",
		"contentValue",
		"urlValue",
		"domainValue",
		"pathValue",
		"fspathValue",
		"htmlfilenameValue",
		"thumbBase64Value",
		"categoryValue",
		"microThumbUrl")

	d := new(pageDaoReader)
	d.Dto(dto)

	actual := string(d.FillJson())

	expected := `{
	"version":2,
	"filename":"htmlfilenameValue",
	"path_from_doc_root":"pathValue",
	"category":"categoryValue",
	"tags":"",
	"create_date":"createDateValue",
	"title":"titleValue",
	"title_plain":"titlePlainValue",
	"excerpt":"descriptionValue",
	"content":"contentValue",
	"thumb_base64":"thumbBase64Value",
	"images_urls":[{"title":"titleValue","w_190":"microThumbUrl","w_390":"thumbUrlValue","w_800":"imageUrlValue","max_resolution":""}]
}`

	if actual != expected {
		t.Error("Expected", actual, "to be", expected)
	}
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
		"", "", "filenameValue", "", "", "")
}
