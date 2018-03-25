package staticPersistence

import (
	"fmt"

	"github.com/ingmardrewing/staticIntf"
)

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
	data []byte
	Json
	dto staticIntf.PageDto
}

func (p *pageDAOv1) Data(data []byte) {
	p.data = data
}

func (p *pageDAOv1) Dto(dto ...staticIntf.PageDto) staticIntf.PageDto {
	if len(dto) > 0 {
		p.dto = dto[0]
	}
	return p.dto
}

func (p *pageDAOv1) ExtractFromJson() {
	id := p.ReadInt(p.data, "id")
	title := p.ReadString(p.data, "title")
	titlePlain := p.ReadString(p.data, "title_plain")
	description := p.ReadString(p.data, "description")
	createDate := p.ReadString(p.data, "createDate")
	content := p.ReadString(p.data, "content")
	pathFromDocRoot := p.ReadString(p.data, "path")
	htmlFilename := p.ReadString(p.data, "filename")

	p.dto = NewFilledDto(
		id,
		title,
		titlePlain,
		"",
		"",
		description,
		"",
		createDate,
		content,
		"",
		"",
		pathFromDocRoot,
		p.dto.FsPath(),
		htmlFilename,
		"")
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
