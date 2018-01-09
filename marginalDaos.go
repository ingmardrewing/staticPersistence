package staticPersistence

import "fmt"

// Marginal DAOs

// Original data structure from wordpress migration
// still having an unneccessary complex structure
// staying here for historical reasons
type marginalDAOv0 struct {
	data []byte
	Json
	dto DTO
}

func (p *marginalDAOv0) Data(data []byte) {
	p.data = data
}

func (p *marginalDAOv0) Dto(dto ...DTO) DTO {
	if len(dto) > 0 {
		p.dto = dto[0]
	}
	return p.dto
}

func (p *marginalDAOv0) ExtractFromJson() {
	p.dto.Id(p.ReadInt(p.data, "page", "post_id"))
	p.dto.Title(p.ReadString(p.data, "title"))
	p.dto.Filename(p.ReadString(p.data, "filename"))
	p.dto.ThumbUrl(p.ReadString(p.data, "thumbImg"))
	p.dto.ImageUrl(p.ReadString(p.data, "postImg"))
	p.dto.Description(p.ReadString(p.data, "page", "excerpt"))
	p.dto.DisqusId(p.ReadString(p.data, "page", "custom_fields", "dsq_thread_id", "[0]"))
	p.dto.CreateDate(p.ReadString(p.data, "page", "date"))
	p.dto.Content(p.ReadString(p.data, "content"))
	p.dto.PathFromDocRoot(p.ReadString(p.data, "path"))
}

func (p *marginalDAOv0) FillJson() []byte {
	json := fmt.Sprintf(p.Template(),
		p.dto.ThumbUrl(),
		p.dto.ImageUrl(),
		p.dto.Filename(),
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
	dto DTO
}

func (p *marginalDAOv1) Data(data []byte) {
	p.data = data
}

func (p *marginalDAOv1) Dto(dto ...DTO) DTO {
	if len(dto) > 0 {
		p.dto = dto[0]
	}
	return p.dto
}

func (p *marginalDAOv1) ExtractFromJson() {
	p.dto.Id(p.ReadInt(p.data, "id"))
	p.dto.Title(p.ReadString(p.data, "title"))
	p.dto.TitlePlain(p.ReadString(p.data, "title_plain"))
	p.dto.Filename(p.ReadString(p.data, "filename"))
	p.dto.Description(p.ReadString(p.data, "description"))
	p.dto.CreateDate(p.ReadString(p.data, "createDate"))
	p.dto.Content(p.ReadString(p.data, "content"))
	p.dto.PathFromDocRoot(p.ReadString(p.data, "path"))
}

func (p *marginalDAOv1) FillJson() []byte {
	json := fmt.Sprintf(p.Template(),
		p.dto.Id(),
		p.dto.Filename(),
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
	"version":1,
	"id":%d,
	"filename":"%s",
	"createDate":"%s",
	"url":"%s",
	"title":"%s",
	"title_plain":"%s",
	"descriptions":"%s",
	"content":"%s"
}`
}
