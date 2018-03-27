package staticPersistence

import (
	"fmt"
	"strings"
)

// Post DAOs
func NewPostDAO(data []byte, path, filename string) PageDao {
	var d PageDao
	switch FindJsonVersion(data) {
	case v1:
		d = new(postDAOv1)
	default:
		d = new(postDAOv0)
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
type postDAOv0 struct {
	abstractPageDao
}

func (p *postDAOv0) ExtractFromJson() {
	id := p.ReadInt(p.data, "post", "post_id")
	title := p.ReadString(p.data, "post", "title")
	thumbUrl := p.ReadString(p.data, "thumbImg")
	imageUrl := p.ReadString(p.data, "postImg")
	description := p.ReadString(p.data, "post", "excerpt")
	disqusId := p.ReadString(p.data, "post", "custom_fields", "dsq_thread_id", "[0]")
	createDate := p.ReadString(p.data, "post", "date")
	content := p.ReadString(p.data, "post", "content")

	url := p.ReadString(p.data, "post", "url")
	parts := strings.Split(url, "/")
	path := strings.Join(parts[4:], "/")
	domain := parts[2]
	htmlFilename := "index.html"

	p.dto = NewFilledDto(
		id,
		title,
		"",
		thumbUrl,
		imageUrl,
		description,
		disqusId,
		createDate,
		content,
		url,
		domain,
		path,
		p.dto.FsPath(),
		htmlFilename,
		"")
}

func (p *postDAOv0) FillJson() []byte {
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
	abstractPageDao
}

func (p *postDAOv1) ExtractFromJson() {
	id := p.ReadInt(p.data, "id")
	title := p.ReadString(p.data, "title")
	thumbUrl := p.ReadString(p.data, "thumbImg")
	imageUrl := p.ReadString(p.data, "postImg")
	description := p.ReadString(p.data, "excerpt")
	disqusId := p.ReadString(p.data, "dsq_thread_id")
	createDate := p.ReadString(p.data, "date")
	content := p.ReadString(p.data, "content")

	url := p.ReadString(p.data, "url")
	parts := strings.Split(url, "/")
	path := strings.Join(parts[4:], "/")

	domain := parts[2]
	pathFromDocRoot := path
	htmlFilename := "index.html"

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
		url,
		domain,
		pathFromDocRoot,
		path,
		htmlFilename,
		"")
}

func (p *postDAOv1) FillJson() []byte {
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
