package staticPersistence

import "github.com/ingmardrewing/staticIntf"

func NewFilledDto(
	title,
	description,
	content,
	category,
	createDate,
	path,
	filename string,
	tags []string,
	images []staticIntf.Image) staticIntf.PageDto {

	return &pageDTO{
		title,
		description,
		content,
		category,
		createDate,
		path,
		filename,
		tags,
		images}
}

// docDtO
type pageDTO struct {
	title       string
	description string
	content     string
	category    string
	createDate  string
	path        string
	filename    string
	tags        []string
	images      []staticIntf.Image
}

func (p pageDTO) PathFromDocRoot() string { return p.path }
func (p pageDTO) Filename() string        { return p.filename }

func (p pageDTO) Title() string       { return p.title }
func (p pageDTO) Description() string { return p.description }
func (p pageDTO) Content() string     { return p.content }
func (p pageDTO) Category() string    { return p.category }
func (p pageDTO) CreateDate() string  { return p.createDate }

func (p pageDTO) Tags() []string {
	return p.tags
}
func (p pageDTO) Images() []staticIntf.Image {
	return p.images
}
