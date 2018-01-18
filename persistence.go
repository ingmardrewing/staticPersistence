package staticPersistence

import (
	"fmt"
	"strings"

	"github.com/ingmardrewing/fs"
)

func ReadMarginals(marginalsDir string) []DTO {
	fileContainers := ReadJsonFilesFromDir(marginalsDir)
	dtos := []DTO{}
	for _, fc := range fileContainers {
		dao := NewMarginalDAO(fc.GetData(), fc.GetPath(), fc.GetFilename())
		dao.ExtractFromJson()
		dtos = append(dtos, dao.Dto())
	}
	return dtos
}

func ReadPosts(postsDir string) []DTO {
	fmt.Println("staticPersistence.ReadPosts")
	fileContainers := ReadJsonFilesFromDir(postsDir)
	dtos := []DTO{}
	for _, fc := range fileContainers {
		dao := NewPostDAO(fc.GetData(), fc.GetPath(), fc.GetFilename())
		dao.ExtractFromJson()
		dtos = append(dtos, dao.Dto())
	}
	return dtos
}

func ReadJsonFilesFromDir(path string) []fs.FileContainer {
	fmt.Println("staticPersistence.ReadJsonFilesFromDir")
	files := fs.ReadDirEntriesEndingWith(path, "json")
	fmt.Println("Number of json files read: ", len(files))
	fileContainers := []fs.FileContainer{}
	for _, filename := range files {
		fc := fs.NewFileContainer()
		fc.SetPath(path)
		fc.SetFilename(filename)
		fc.Read()
		fileContainers = append(fileContainers, fc)
	}
	return fileContainers
}

func WriteMarginalDtoToJson(dto DTO, path, filename string) {
	dao := NewMarginalDAO(nil, path, filename)
	dao.Dto(dto)
	writeJson(dao.FillJson(), path, filename)
}

func WritePostDtoToJson(dto DTO, path, filename string) {
	dao := NewPostDAO(nil, path, filename)
	dao.Dto(dto)
	writeJson(dao.FillJson(), path, filename)
}

func writeJson(json []byte, path, filename string) {
	fc := fs.NewFileContainer()
	fc.SetDataAsString(string(json))
	fc.SetPath(path)
	fc.SetFilename(filename)
	fc.Write()
}

type Persistence interface{}

var (
	persistence Persistence
)

func init() {}

type PostAdder interface {
	Read()
	GetDir() string
	GetMdInitContent() string
	GetJsonFileName() string
	GetJsonFilePath() string
	GetJsonFileContent() string
	GetMdContent() string
	GetMdFileName() string
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
	imgjsoncontent  string
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

func (p *postAdder) GetDir() string {
	return p.dirpath
}

func (p *postAdder) GetMdInitContent() string {
	return p.mdinitcontent
}

func (p *postAdder) GetMdContent() string {
	return p.mdcontent
}

func (p *postAdder) GetJsonFileContent() string {
	return p.imgjsoncontent
}

func (p *postAdder) GetMdFilePath() string {
	return p.dirpath + p.mdfilename
}

func (p *postAdder) GetMdFileName() string {
	return p.mdfilename
}

func (p *postAdder) GetJsonFileName() string {
	return p.imgjsonfilename
}

func (p *postAdder) GetJsonFilePath() string {
	return p.dirpath
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
	p.imgjsoncontent = p.readJsonContent()
}

func (p *postAdder) readJsonContent() string {
	if len(p.GetJsonFileName()) > 0 {
		return p.readFileContents(p.dirpath, p.imgjsonfilename)
	}
	return ""
}

func (p *postAdder) readMdContent() string {
	if len(p.GetMdFileName()) > 0 {
		return p.readFileContents(p.dirpath, p.mdfilename)
	}
	return ""
}

func (p *postAdder) readFileContents(path, file string) string {
	fc := fs.NewFileContainer()
	fc.SetPath(path)
	fc.SetFilename(file)
	fc.Read()

	content := fc.GetDataAsString()
	return strings.TrimSuffix(content, "\n")
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
