package staticPersistence

import "github.com/ingmardrewing/staticIntf"

func NewFilledDto(
	title,
	description,
	content,
	category,
	createDate,
	pathFromDocRoot,
	filename string,
	tags []string,
	images []staticIntf.Image) staticIntf.PageDto {

	return &pageDto{
		title,
		description,
		content,
		category,
		createDate,
		pathFromDocRoot,
		filename,
		tags,
		images}
}

func NewPageDto(
	title,
	description,
	content,
	category,
	createDate,
	pathFromDocRoot,
	filename string,
	tags []string,
	images []staticIntf.Image) staticIntf.PageDto {

	return &pageDto{
		title,
		description,
		content,
		category,
		createDate,
		pathFromDocRoot,
		filename,
		tags,
		images}
}

// docDtO
type pageDto struct {
	title           string
	description     string
	content         string
	category        string
	createDate      string
	pathFromDocRoot string
	filename        string
	tags            []string
	images          []staticIntf.Image
}

func (p pageDto) PathFromDocRoot() string { return p.pathFromDocRoot }
func (p pageDto) Filename() string        { return p.filename }

func (p pageDto) Title() string       { return p.title }
func (p pageDto) Description() string { return p.description }
func (p pageDto) Content() string     { return p.content }
func (p pageDto) Category() string    { return p.category }
func (p pageDto) CreateDate() string  { return p.createDate }

func (p pageDto) Tags() []string {
	return p.tags
}
func (p pageDto) Images() []staticIntf.Image {
	return p.images
}
