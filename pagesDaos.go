package staticPersistence

import "fmt"

// Page Dao
func NewPageDAO(data []byte, path, filename string) PageDao {
	var d PageDao
	switch FindJsonVersion(data) {
	case v1:
		d = new(pageDAOv1)
	default:
		d = new(pageDAOv1)
	}

	dto := NewFilledDto(0,
		"", "", "", "", "",
		"", "", "", "", "",
		path, filename, "", "")

	d.Dto(dto)
	d.Data(data)

	return d
}

// Pages DAOs
type pageDAOv1 struct {
	abstractPageDao
}

func (p *pageDAOv1) FillJson() []byte {
	json := fmt.Sprintf(p.Template(),
		p.dto.Id(),
		p.dto.PathFromDocRoot(),
		p.dto.HtmlFilename(),
		p.dto.CreateDate(),
		p.dto.Url(),
		p.dto.Title(),
		p.dto.TitlePlain(),
		p.dto.Description(),
		p.dto.Content())
	return []byte(json)
}

func (p *pageDAOv1) Template() string {
	return `{
	"version":1,
	"id":%d,
	"path":"%s",
	"filename":"%s",
	"createDate":"%s",
	"url":"%s",
	"title":"%s",
	"title_plain":"%s",
	"description":"%s",
	"content":"%s"
}`
}
