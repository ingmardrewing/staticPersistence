package staticPersistence

import (
	"fmt"
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

/*
func TestNewPageDaoReader(t *testing.T) {
	p := path.Join(currentDir(), "testResources", "posts")
	f := "version0.json"
	r := newPageDaoReader(readViaFc(p, f), p, f)
	r.ExtractFromJson()
	dto := r.Dto()
	expected := "Food Market Analysis"
	if dto.TitlePlain() == expected {
		t.Error("Expected", expected, "but got", dto.Title())
	}
}
*/

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

	expected := fmt.Sprintf(`{
	"version":1,
	"thumbImg":"%s",
	"postImg":"%s",
	"filename":"%s",
	"id":%d,
	"date":"%s",
	"url":"%s",
	"title":"%s",
	"title_plain":"%s",
	"excerpt":"%s",
	"content":"%s",
	"dsq_thread_id":"%s",
	"thumbBase64":"%s",
	"category":"%s",
	"microThumbUrl":"%s"
}`,
		dto.ThumbUrl(),
		dto.ImageUrl(),
		dto.HtmlFilename(),
		dto.Id(),
		dto.CreateDate(),
		dto.Url(),
		dto.Title(),
		dto.TitlePlain(),
		dto.Description(),
		dto.Content(),
		dto.DisqusId(),
		dto.ThumbBase64(),
		dto.Category(),
		dto.MicroThumbUrl())

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
