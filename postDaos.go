package staticPersistence

import "fmt"

// Post DAOs

// Original data structure from wordpress migration
// still having an unneccessary complex structure
// staying here for historical reasons
type postDAOv0 struct {
	data []byte
	Json
	dto *docDTO
}

func (p *postDAOv0) ExtractFromJson() {
	p.dto.id = p.ReadInt(p.data, "post", "post_id")
	p.dto.title = p.ReadString(p.data, "post", "title")
	p.dto.thumbUrl = p.ReadString(p.data, "thumbImg")
	p.dto.imageUrl = p.ReadString(p.data, "postImg")
	p.dto.description = p.ReadString(p.data, "post", "excerpt")
	p.dto.disqusId = p.ReadString(p.data, "post", "custom_fields", "dsq_thread_id", "[0]")
	p.dto.createDate = p.ReadString(p.data, "post", "date")
	p.dto.content = p.ReadString(p.data, "post", "content")
	p.dto.url = p.ReadString(p.data, "post", "url")
	p.dto.path = p.ReadString(p.data, "path")
	p.dto.filename = p.ReadString(p.data, "filename")
}

func (p *postDAOv0) FillJson() []byte {
	json := fmt.Sprintf(p.Template(),
		p.dto.thumbUrl,
		p.dto.imageUrl,
		p.dto.filename,
		p.dto.id,
		p.dto.createDate,
		p.dto.url,
		p.dto.title,
		p.dto.titlePlain,
		p.dto.description,
		p.dto.content,
		p.dto.disqusId)
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
	dto *docDTO
}

func (p *postDAOv1) ExtractFromJson() {
	p.dto.id = p.ReadInt(p.data, "id")
	p.dto.title = p.ReadString(p.data, "title")
	p.dto.thumbUrl = p.ReadString(p.data, "thumbImg")
	p.dto.imageUrl = p.ReadString(p.data, "postImg")
	p.dto.description = p.ReadString(p.data, "excerpt")
	p.dto.disqusId = p.ReadString(p.data, "dsq_thread_id")
	p.dto.createDate = p.ReadString(p.data, "date")
	p.dto.content = p.ReadString(p.data, "content")
	p.dto.url = p.ReadString(p.data, "url")
	p.dto.path = p.ReadString(p.data, "path")
	p.dto.filename = p.ReadString(p.data, "filename")
}

func (p *postDAOv1) FillJson() []byte {
	json := fmt.Sprintf(p.Template(),
		p.dto.thumbUrl,
		p.dto.imageUrl,
		p.dto.filename,
		p.dto.id,
		p.dto.createDate,
		p.dto.url,
		p.dto.title,
		p.dto.titlePlain,
		p.dto.description,
		p.dto.content,
		p.dto.disqusId)
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
