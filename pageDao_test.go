package staticPersistence

import (
	"path"
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
		"createDateValue",
		"contentValue",
		"pathValue",
		"fspathValue",
		"htmlfilenameValue",
		"thumbBase64Value",
		"categoryValue",
		"microThumbUrl",
		[]string{},
		[]staticIntf.Image{})

	d := new(pageDaoReader)
	d.Dto(dto)

	actual := string(d.FillJson())

	expected := `{
	"version":2,
	"filename":"htmlfilenameValue",
	"path_from_doc_root":"pathValue",
	"category":"categoryValue",
	"tags":[],
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
	file := "version2.json"

	actual := readAndGetDto(path, file)
	expected := NewFilledDto(0,
		"titleValue", "title_plainValue", "thumbImageValue",
		"postImageValue", "excerptValue", "dateValue", "contentValue",

		"", "", "filenameValue", "", "", "", []string{}, []staticIntf.Image{})

	if expected.Title() != actual.Title() ||
		expected.Description() != actual.Description() ||
		expected.ThumbUrl() != actual.ThumbUrl() ||
		expected.ImageUrl() != actual.ImageUrl() ||
		expected.TitlePlain() != actual.TitlePlain() ||
		expected.Content() != actual.Content() ||
		expected.Content() != actual.Content() ||
		expected.HtmlFilename() != actual.HtmlFilename() ||
		expected.PathFromDocRoot() != actual.PathFromDocRoot() {
		t.Error("Expected", actual, "to be", expected)
	}
}
