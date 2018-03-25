package staticPersistence

import (
	"fmt"

	"github.com/ingmardrewing/staticIntf"
)

// marginalDAO
func NewMarginalDAO(data []byte, path, filename string) PageDao {
	var d PageDao
	switch FindJsonVersion(data) {
	case v1:
		d = new(marginalDAOv1)
	default:
		d = new(marginalDAOv0)
	}

	dto := NewFilledDto(0,
		"", "", "", "", "",
		"", "", "", "", "",
		path, filename, "", "")

	d.Dto(dto)
	d.Data(data)

	return d
}

// Original data structure from wordpress migration
// still having an unneccessary complex structure
// staying here for historical reasons
type marginalDAOv0 struct {
	data []byte
	Json
	dto staticIntf.PageDto
}

func (p *marginalDAOv0) Data(data []byte) {
	p.data = data
}

func (p *marginalDAOv0) Dto(dto ...staticIntf.PageDto) staticIntf.PageDto {
	if len(dto) > 0 {
		p.dto = dto[0]
	}
	return p.dto
}

func (p *marginalDAOv0) ExtractFromJson() {
	id := p.ReadInt(p.data, "page", "post_id")
	title := p.ReadString(p.data, "title")
	thumbUrl := p.ReadString(p.data, "thumbImg")
	imageUrl := p.ReadString(p.data, "postImg")
	description := p.ReadString(p.data, "page", "excerpt")
	disqusId := p.ReadString(p.data, "page", "custom_fields", "dsq_thread_id", "[0]")
	createDate := p.ReadString(p.data, "page", "date")
	content := p.ReadString(p.data, "content")
	pathFromDocRoot := p.ReadString(p.data, "path")
	htmlFilename := p.ReadString(p.data, "filename")

	p.dto = NewFilledDto(
		id,
		title,
		title,
		thumbUrl,
		imageUrl,
		description,
		disqusId,
		createDate,
		content,
		"",
		"",
		pathFromDocRoot,
		p.dto.FsPath(),
		htmlFilename,
		"")

}

func (p *marginalDAOv0) FillJson() []byte {
	json := fmt.Sprintf(p.Template(),
		p.dto.ThumbUrl(),
		p.dto.ImageUrl(),
		p.dto.HtmlFilename(),
		p.dto.Id(),
		p.dto.CreateDate(),
		p.dto.Url(),
		p.dto.Title(),
		p.dto.TitlePlain(),
		p.dto.Description(),
		p.dto.Content(),
		p.dto.DisqusId())
	return []byte(json)
}

func (p *marginalDAOv0) Template() string {
	return `{
  "title":"%s",
  "title_plain":"%s",
  "filename":"%s",
  "content":"%s",
  "page":{
	  "id":%d,
	  "slug":"ingmars-booklist",
	  "url":"https:\/\/www.drewing.de\/blog\/",
	  "status":"publish",
	  "excerpt":"%s",
    "date":"%s",
  }
}`
}

// Marginal DAOs

// Original data structure from wordpress migration
// still having an unneccessary complex structure
// staying here for historical reasons
type marginalDAOv1 struct {
	data []byte
	Json
	dto staticIntf.PageDto
}

func (p *marginalDAOv1) Data(data []byte) {
	p.data = data
}

func (p *marginalDAOv1) Dto(dto ...staticIntf.PageDto) staticIntf.PageDto {
	if len(dto) > 0 {
		p.dto = dto[0]
	}
	return p.dto
}

func (p *marginalDAOv1) ExtractFromJson() {
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

func (p *marginalDAOv1) FillJson() []byte {
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

func (p *marginalDAOv1) Template() string {
	return `{
	"version": 1,
	"id": %d,
	"path": "%s",
	"filename": "%s",
	"createDate": "%s",
	"url": "%s",
	"title": "%s",
	"title_plain": "%s",
	"description": "%s",
	"content": "%s"
}`
}
