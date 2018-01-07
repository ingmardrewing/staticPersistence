package staticPersistence

import (
	"strings"

	"github.com/ingmardrewing/fs"
)

type PostAdder interface {
	Read()
	GetMdInitContent() string
	GetMdContent() string
	GetImgFilePath() string
	GetImgFileName() string
	GetMdFilePath() string
}

type postAdder struct {
	dirpath         string
	imgfilename     string
	mdfilename      string
	mdinitcontent   string
	mdcontent       string
	imgjsonfilename string
}

func NewPostAdder(path string) PostAdder {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	p := new(postAdder)
	p.init(path)
	return p
}

func (p *postAdder) GetMdInitContent() string {
	return p.mdinitcontent
}

func (p *postAdder) GetMdContent() string {
	return p.mdcontent
}

func (p *postAdder) GetMdFilePath() string {
	return p.dirpath + p.mdinitcontent
}

func (p *postAdder) GetMdFileName() string {
	return p.mdinitcontent
}

func (p *postAdder) GetJsonFileName() string {
	return p.imgjsonfilename
}

func (p *postAdder) GetJsonFilePath() string {
	return p.dirpath + p.imgjsonfilename
}

func (p *postAdder) GetImgFileName() string {
	return p.imgfilename
}

func (p *postAdder) GetImgFilePath() string {
	return p.dirpath + p.imgfilename
}

func (p *postAdder) init(path string) {
	p.dirpath = path
	p.mdfilename = p.readMdFileNameFromFs()
	p.mdinitcontent = p.readMdContent()
}

func (p *postAdder) Read() {
	p.imgfilename = p.readImageFileNameFromFs()
	p.mdfilename = p.readMdFileNameFromFs()
	p.mdcontent = p.readMdContent()
	p.imgjsonfilename = p.readJsonFileNameFromFs()
}

func (p *postAdder) readMdContent() string {
	if len(p.mdfilename) > 0 {
		content := fs.ReadFileAsString(p.dirpath + p.mdfilename)
		return strings.TrimSuffix(content, "\n")
	}
	return ""
}

func (p *postAdder) readImageFileNameFromFs() string {
	imgs := fs.ReadDirEntriesEndingWith(p.dirpath, "png", "jpg")
	for _, i := range imgs {
		if !strings.Contains(i, "-w") {
			return i
		}
	}
	return ""
}

func (p *postAdder) readJsonFileNameFromFs() string {
	return p.getFirstFileEndingWith("json")
}

func (p *postAdder) readMdFileNameFromFs() string {
	return p.getFirstFileEndingWith("md")
}

func (p *postAdder) getFirstFileEndingWith(suffix string) string {
	fns := fs.ReadDirEntriesEndingWith(p.dirpath, suffix)
	if len(fns) > 0 {
		return fns[0]
	}
	return ""
}
