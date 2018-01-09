package staticPersistence

import "fmt"

// Post DAOs

// Original data structure from wordpress migration
// still having an unneccessary complex structure
// staying here for historical reasons
type postDAOv0 struct {
	Json
	docDAO
}

func (p *postDAOv0) ExtractFromJson() {
	p.id = p.ReadInt(p.data, "post", "post_id")
	p.title = p.ReadString(p.data, "post", "title")
	p.thumbUrl = p.ReadString(p.data, "thumbImg")
	p.imageUrl = p.ReadString(p.data, "postImg")
	p.description = p.ReadString(p.data, "post", "excerpt")
	p.disqusId = p.ReadString(p.data, "post", "custom_fields", "dsq_thread_id", "[0]")
	p.createDate = p.ReadString(p.data, "post", "date")
	p.content = p.ReadString(p.data, "post", "content")
	p.url = p.ReadString(p.data, "post", "url")
	p.path = p.ReadString(p.data, "path")
	p.filename = p.ReadString(p.data, "filename")
}

func (p *postDAOv0) FillJson() []byte {
	json := fmt.Sprintf(p.Template(),
		p.thumbUrl,
		p.imageUrl,
		p.filename,
		p.id,
		p.createDate,
		p.url,
		p.title,
		p.titlePlain,
		p.description,
		p.content,
		p.disqusId)
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
	Json
	docDAO
}

func (p *postDAOv1) ExtractFromJson() {
	p.id = p.ReadInt(p.data, "id")
	p.title = p.ReadString(p.data, "title")
	p.thumbUrl = p.ReadString(p.data, "thumbImg")
	p.imageUrl = p.ReadString(p.data, "postImg")
	p.description = p.ReadString(p.data, "excerpt")
	p.disqusId = p.ReadString(p.data, "dsq_thread_id")
	p.createDate = p.ReadString(p.data, "date")
	p.content = p.ReadString(p.data, "content")
	p.url = p.ReadString(p.data, "url")
	p.path = p.ReadString(p.data, "path")
	p.filename = p.ReadString(p.data, "filename")
}

func (p *postDAOv1) FillJson() []byte {
	json := fmt.Sprintf(p.Template(),
		p.thumbUrl,
		p.imageUrl,
		p.filename,
		p.id,
		p.createDate,
		p.url,
		p.title,
		p.titlePlain,
		p.description,
		p.content,
		p.disqusId)
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
