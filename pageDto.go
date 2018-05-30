package staticPersistence

import "github.com/ingmardrewing/staticIntf"

func NewFilledDto(id int,
	title, titlePlain, thumbUrl,
	imageUrl, description, disqusId,
	createDate, content, url, domain,
	path, fspath, htmlfilename,
	thumbBase64, category, microThumbUrl string) staticIntf.PageDto {

	return &pageDTO{id, title, titlePlain, thumbUrl,
		imageUrl, description, disqusId,
		createDate, content, url, domain,
		path, fspath, htmlfilename, thumbBase64,
		category, microThumbUrl}
}

// docDtO
type pageDTO struct {
	id int
	title, titlePlain, thumbUrl,
	imageUrl, description, disqusId,
	createDate, content, url, domain,
	path, fspath, htmlfilename,
	thumbBase64, category, microThumbUrl string
}

func (p pageDTO) FsPath() string { return p.fspath }

func (p pageDTO) HtmlFilename() string { return p.htmlfilename }

func (p pageDTO) Id() int { return p.id }

func (p pageDTO) Title() string { return p.title }

func (p pageDTO) Domain() string { return p.domain }

func (p pageDTO) TitlePlain() string { return p.titlePlain }

func (p pageDTO) ThumbUrl() string { return p.thumbUrl }

func (p pageDTO) MicroThumbUrl() string { return p.microThumbUrl }

func (p pageDTO) ImageUrl() string { return p.imageUrl }

func (p pageDTO) Description() string { return p.description }

func (p pageDTO) DisqusId() string { return p.disqusId }

func (p pageDTO) CreateDate() string { return p.createDate }

func (p pageDTO) Content() string { return p.content }

func (p pageDTO) Category() string { return p.category }

func (p pageDTO) ThumbBase64() string { return p.thumbBase64 }

func (p pageDTO) Url() string { return p.url }

func (p pageDTO) PathFromDocRoot() string { return p.path }
