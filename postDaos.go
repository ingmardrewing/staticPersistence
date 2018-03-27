package staticPersistence

import "strings"

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
