package staticPersistence

import "strings"

func NewDto() *docDTO {
	return new(docDTO)
}

// docDtO
type docDTO struct {
	id int
	title, titlePlain, thumbUrl,
	imageUrl, description, disqusId,
	createDate, content, url,
	path, fspath, htmlfilename string
}

func (p *docDTO) FsPath(fspath string) {
	p.fspath = fspath
}

func (p *docDTO) HtmlFilename(htmlfilename ...string) string {
	if len(htmlfilename) > 0 {
		p.htmlfilename = htmlfilename[0]
	}
	return p.htmlfilename
}

func (p *docDTO) Id(id ...int) int {
	if len(id) > 0 {
		p.id = id[0]
	}
	return p.id
}

func (p *docDTO) Title(title ...string) string {
	if len(title) > 0 {
		p.title = title[0]
	}
	return p.title
}

func (p *docDTO) Domain() string {
	if len(p.url) > 0 {
		parts := strings.Split(p.url, "/")
		return strings.Join(parts[0:3], "/")
	}
	return ""
}

func (p *docDTO) TitlePlain(titlePlain ...string) string {
	if len(titlePlain) > 0 {
		p.titlePlain = titlePlain[0]
	}
	return p.titlePlain
}

func (p *docDTO) ThumbUrl(thumbUrl ...string) string {
	if len(thumbUrl) > 0 {
		p.thumbUrl = thumbUrl[0]
	}
	return p.thumbUrl
}

func (p *docDTO) ImageUrl(imageUrl ...string) string {
	if len(imageUrl) > 0 {
		p.imageUrl = imageUrl[0]
	}
	return p.imageUrl
}

func (p *docDTO) Description(desc ...string) string {
	if len(desc) > 0 {
		p.description = desc[0]
	}
	return p.description
}

func (p *docDTO) DisqusId(disqusId ...string) string {
	if len(disqusId) > 0 {
		p.disqusId = disqusId[0]
	}
	return p.disqusId
}

func (p *docDTO) CreateDate(date ...string) string {
	if len(date) > 0 {
		p.createDate = date[0]
	}
	return p.createDate
}

func (p *docDTO) Content(content ...string) string {
	if len(content) > 0 {
		p.content = content[0]
	}
	return p.content
}

func (p *docDTO) Url(url ...string) string {
	if len(url) > 0 {
		p.url = url[0]
	}
	return p.url
}

func (p *docDTO) PathFromDocRoot(path ...string) string {
	if len(path) > 0 {
		p.path = path[0]
	}
	return p.path
}
