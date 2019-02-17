package staticPersistence

import "github.com/ingmardrewing/staticIntf"

func NewFilledDto(id int,
	title, titlePlain, thumbUrl,
	imageUrl, description,
	createDate, content,
	path, fspath, htmlfilename,
	thumbBase64, category, microThumbUrl string,
	tags []string,
	images []staticIntf.Image) staticIntf.PageDto {

	return &pageDTO{id, title, titlePlain, thumbUrl,
		imageUrl, description,
		createDate, content,
		path, fspath, htmlfilename, thumbBase64,
		category, microThumbUrl, tags, images}
}

func NewImageDto(title, w190, w390, w800, maxResolution string) *imageDTO {
	return &imageDTO{title, w190, w390, w800, maxResolution}
}

// imageDTO
type imageDTO struct {
	title         string
	w190          string
	w390          string
	w800          string
	maxResolution string
}

func (i imageDTO) W190() string { return i.w190 }

func (i imageDTO) W390() string { return i.w390 }

func (i imageDTO) W800() string { return i.w800 }

func (i imageDTO) MaxResoultion() string { return i.maxResolution }

func (i imageDTO) Title() string { return i.title }

// docDtO
type pageDTO struct {
	id int
	title, titlePlain, thumbUrl,
	imageUrl, description,
	createDate, content,
	path, fspath, htmlfilename,
	thumbBase64, category, microThumbUrl string
	tags   []string
	images []staticIntf.Image
}

func (p pageDTO) FsPath() string { return p.fspath }

func (p pageDTO) HtmlFilename() string { return p.htmlfilename }

func (p pageDTO) Id() int { return p.id }

func (p pageDTO) Title() string { return p.title }

func (p pageDTO) TitlePlain() string { return p.titlePlain }

func (p pageDTO) ThumbUrl() string { return p.thumbUrl }

func (p pageDTO) MicroThumbUrl() string {
	return p.microThumbUrl
}

func (p pageDTO) Tags() []string {
	return p.tags
}

func (p pageDTO) Images() []staticIntf.Image {
	return p.images
}

func (p pageDTO) ImageUrl() string { return p.imageUrl }

func (p pageDTO) Description() string { return p.description }

func (p pageDTO) CreateDate() string { return p.createDate }

func (p pageDTO) Content() string { return p.content }

func (p pageDTO) Category() string { return p.category }

func (p pageDTO) ThumbBase64() string { return p.thumbBase64 }

func (p pageDTO) PathFromDocRoot() string { return p.path }
