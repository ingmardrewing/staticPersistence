package staticPersistence

import "github.com/buger/jsonparser"

const (
	v0 = iota
	v1 = iota
)

func FindJsonVersion(data []byte) int {
	if data == nil {
		return v1
	}
	j := new(Json)
	v := j.ReadInt(data, "version")
	return v
}

// Post Dao
func NewPostDAO(data []byte, path, filename string) DAO {
	var d DAO
	switch FindJsonVersion(data) {
	case v1:
		d = new(postDAOv1)
	default:
		d = new(postDAOv0)
	}
	d.Data(data)
	d.FsPath(path)
	d.FsFilename(filename)
	return d
}

// marginalDAO
func NewMarginalDAO(data []byte, path, filename string) DAO {
	var d DAO
	switch FindJsonVersion(data) {
	case v1:
		d = new(postDAOv1)
	default:
		d = new(postDAOv0)
	}
	d.Data(data)
	d.FsPath(path)
	d.FsFilename(filename)
	return d
}

// Json
type Json struct{}

func (j *Json) ReadString(value []byte, keys ...string) string {
	v, err := jsonparser.GetString(value, keys...)
	if err != nil {
		return ""
	}
	return v
}

func (j *Json) ReadInt(value []byte, keys ...string) int {
	v, err := jsonparser.GetInt(value, keys...)
	if err != nil {
		return 0
	}
	return int(v)
}

type DAO interface {
	ExtractFromJson()
	FillJson() []byte
	Data([]byte)
	FsFilename(string)
	FsPath(string)
	Id(...int) int
	Title(...string) string
	TitlePlain(...string) string
	ThumbUrl(...string) string
	ImageUrl(...string) string
	Description(...string) string
	DisqusId(...string) string
	CreateDate(...string) string
	Content(...string) string
	Url(...string) string
	PathFromDocRoot(...string) string
	HtmlFilename(...string) string
}

// docDAO
type docDAO struct {
	data []byte
	id   int
	title, titlePlain, thumbUrl,
	imageUrl, description, disqusId,
	createDate, content, url,
	path, filename,
	fspath, fsfilename string
}

func (p *docDAO) FsPath(fspath string) {
	p.fspath = fspath
}

func (p *docDAO) FsFilename(fsfilename string) {
	p.fsfilename = fsfilename
}

func (p *docDAO) Data(data []byte) {
	p.data = data
}

func (p *docDAO) Id(id ...int) int {
	if len(id) > 0 {
		p.id = id[0]
	}
	return p.id
}

func (p *docDAO) Title(title ...string) string {
	if len(title) > 0 {
		p.title = title[0]
	}
	return p.title
}

func (p *docDAO) TitlePlain(titlePlain ...string) string {
	if len(titlePlain) > 0 {
		p.titlePlain = titlePlain[0]
	}
	return p.titlePlain
}

func (p *docDAO) ThumbUrl(thumbUrl ...string) string {
	if len(thumbUrl) > 0 {
		p.thumbUrl = thumbUrl[0]
	}
	return p.thumbUrl
}

func (p *docDAO) ImageUrl(imageUrl ...string) string {
	if len(imageUrl) > 0 {
		p.imageUrl = imageUrl[0]
	}
	return p.imageUrl
}

func (p *docDAO) Description(desc ...string) string {
	if len(desc) > 0 {
		p.description = desc[0]
	}
	return p.description
}

func (p *docDAO) DisqusId(disqusId ...string) string {
	if len(disqusId) > 0 {
		p.disqusId = disqusId[0]
	}
	return p.disqusId
}

func (p *docDAO) CreateDate(date ...string) string {
	if len(date) > 0 {
		p.createDate = date[0]
	}
	return p.createDate
}

func (p *docDAO) Content(content ...string) string {
	if len(content) > 0 {
		p.content = content[0]
	}
	return p.content
}

func (p *docDAO) Url(url ...string) string {
	if len(url) > 0 {
		p.url = url[0]
	}
	return p.url
}

func (p *docDAO) PathFromDocRoot(path ...string) string {
	if len(path) > 0 {
		p.path = path[0]
	}
	return p.path
}

func (p *docDAO) HtmlFilename(htmlFilename ...string) string {
	if len(htmlFilename) > 0 {
		p.filename = htmlFilename[0]
	}
	return p.filename
}

func (p *docDAO) ExtractFromJson() {}
func (p *docDAO) FillJson() []byte {
	return []byte{}
}
