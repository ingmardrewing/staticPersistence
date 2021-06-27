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
	img := NewImageDto(
		"titleValue",
		"blogThumbUrl",
		"blogThumbUrl",
		"microThumbUrl",
		"microThumbUrl",
		"thumbUrlValue",
		"thumbUrlValue",
		"imageUrlValue",
		"imageUrlValue",
		"largeImageUrlValue",
		"")
	dto := NewPageDto(
		"titleValue",
		"descriptionValue",
		"contentValue",
		"categoryValue",
		"createDateValue",
		"pathValue",
		"htmlfilenameValue",
		[]string{"tag1", "tag2"},
		[]staticIntf.Image{img})

	d := new(pageDaoReader)
	d.Dto(dto)

	actual := string(d.FillJson())

	expected := `{
	"version":2,
	"filename":"htmlfilenameValue",
	"path_from_doc_root":"pathValue",
	"category":"categoryValue",
	"tags":["tag1","tag2"],
	"create_date":"createDateValue",
	"title":"titleValue",
	"excerpt":"descriptionValue",
	"content":"contentValue",
	"images_urls":[{
		"title":"titleValue",
		"w_85":"blogThumbUrl",
		"w_100":"blogThumbUrl",
		"w_190":"microThumbUrl",
		"w_200":"microThumbUrl",
		"w_390":"thumbUrlValue",
		"w_400":"thumbUrlValue",
		"w_800":"imageUrlValue",
		"w_800_portrait":"imageUrlValue",
		"w_1600_portrait":"largeImageUrlValue",
		"max_resolution":""
	}]
}`

	if actual != expected {
		t.Error("Expected", actual, "to be", expected)
	}
}

func TestReadFile(t *testing.T) {
	path := path.Join(currentDir(), "testResources", "posts")
	file := "version2.json"

	actual := readAndGetDto(path, file)
	expected := NewPageDto(
		"titleValue",
		"excerptValue",
		"contentValue",
		"",
		"dateValue",
		"",
		"filenameValue",
		[]string{},
		[]staticIntf.Image{})

	if expected.Title() != actual.Title() ||
		expected.Description() != actual.Description() ||
		expected.Content() != actual.Content() ||
		expected.CreateDate() != actual.CreateDate() ||
		expected.Filename() != actual.Filename() ||
		expected.PathFromDocRoot() != actual.PathFromDocRoot() {
		t.Error("Expected", actual, "to be", expected)
	}
}
