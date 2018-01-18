package staticPersistence

import (
	"fmt"
	"strings"
)

// Post DAOs

// Original data structure from wordpress migration
// still having an unneccessary complex structure
// staying here for historical reasons
type postDAOv0 struct {
	data []byte
	Json
	dto DTO
}

func (p *postDAOv0) ExtractFromJson() {
	p.dto.Id(p.ReadInt(p.data, "post", "post_id"))
	p.dto.Title(p.ReadString(p.data, "post", "title"))
	p.dto.ThumbUrl(p.ReadString(p.data, "thumbImg"))
	p.dto.ImageUrl(p.ReadString(p.data, "postImg"))
	p.dto.Description(p.ReadString(p.data, "post", "excerpt"))
	p.dto.DisqusId(p.ReadString(p.data, "post", "custom_fields", "dsq_thread_id", "[0]"))
	p.dto.CreateDate(p.ReadString(p.data, "post", "date"))
	p.dto.Content(p.ReadString(p.data, "post", "content"))

	p.dto.Url(p.ReadString(p.data, "post", "url"))

	parts := strings.Split(p.dto.Url(), "/")
	path := strings.Join(parts[3:], "/")
	p.dto.PathFromDocRoot(path)

	p.dto.Filename("index.html")
}

func (p *postDAOv0) Data(data []byte) {
	p.data = data
}

func (p *postDAOv0) Dto(dto ...DTO) DTO {
	if len(dto) > 0 {
		p.dto = dto[0]
	}
	return p.dto
}

func (p *postDAOv0) FillJson() []byte {
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

func (p *postDAOv0) Template() string {
	return `{
	"thumbImg":"%s",
	"postImg":"%s",
	"filename":"%s",
	"post":{
		"post_id":"%s",
		"date":"%s",
		"url":"%s",
		"title":"%s",
		"title_plain":"%s",
		"excerpt":"%s",
		"content":"%s",
		"custom_fields":{
			"dsq_thread_id":["%s"]
		}
	}
}`
}

// New json format v1
type postDAOv1 struct {
	data []byte
	Json
	dto DTO
}

func (p *postDAOv1) ExtractFromJson() {
	p.dto.Id(p.ReadInt(p.data, "id"))
	p.dto.Title(p.ReadString(p.data, "title"))
	p.dto.ThumbUrl(p.ReadString(p.data, "thumbImg"))
	p.dto.ImageUrl(p.ReadString(p.data, "postImg"))
	p.dto.Description(p.ReadString(p.data, "excerpt"))
	p.dto.DisqusId(p.ReadString(p.data, "dsq_thread_id"))
	p.dto.CreateDate(p.ReadString(p.data, "date"))
	p.dto.Content(p.ReadString(p.data, "content"))
	p.dto.Url(p.ReadString(p.data, "url"))
	p.dto.PathFromDocRoot(p.ReadString(p.data, "path"))
	p.dto.Filename(p.ReadString(p.data, "filename"))
}

func (p *postDAOv1) Data(data []byte) {
	p.data = data
}

func (p *postDAOv1) Dto(dto ...DTO) DTO {
	if len(dto) > 0 {
		p.dto = dto[0]
	}
	return p.dto
}

func (p *postDAOv1) FillJson() []byte {
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

func (p *postDAOv1) Template() string {
	return `{
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
	"dsq_thread_id":"%s"
}`
}
