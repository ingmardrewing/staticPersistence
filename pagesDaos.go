package staticPersistence

import "fmt"

// Page Dao
func NewPageDAO(data []byte, path, filename string) DAO {
	dto := NewDto()
	dto.FsPath(path)
	dto.HtmlFilename(filename)

	var d DAO
	switch FindJsonVersion(data) {
	case v1:
		d = new(pageDAOv1)
	default:
		d = new(pageDAOv1)
	}
	d.Dto(dto)
	d.Data(data)

	return d
}

// Pages DAOs
type pageDAOv1 struct {
	data []byte
	Json
	dto DTO
}

func (p *pageDAOv1) Data(data []byte) {
	p.data = data
}

func (p *pageDAOv1) Dto(dto ...DTO) DTO {
	if len(dto) > 0 {
		p.dto = dto[0]
	}
	return p.dto
}

func (p *pageDAOv1) ExtractFromJson() {
	p.dto.Id(p.ReadInt(p.data, "id"))
	p.dto.Title(p.ReadString(p.data, "title"))
	p.dto.TitlePlain(p.ReadString(p.data, "title_plain"))
	p.dto.Description(p.ReadString(p.data, "description"))
	p.dto.CreateDate(p.ReadString(p.data, "createDate"))
	p.dto.Content(p.ReadString(p.data, "content"))

	p.dto.PathFromDocRoot(p.ReadString(p.data, "path"))
	p.dto.HtmlFilename(p.ReadString(p.data, "filename"))
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
