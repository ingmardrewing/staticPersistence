package staticPersistence

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

// dao versions
const (
	v0 = iota
	v1 = iota
)

type PageDao interface {
	Dto(...staticIntf.PageDto) staticIntf.PageDto
	ExtractFromJson()
	FillJson() []byte
	Data([]byte)
}

func ReadNarrativePages(pagesDir string) []staticIntf.PageDto {
	fileContainers := ReadJsonFilesFromDir(pagesDir)
	dtos := []staticIntf.PageDto{}
	for _, fc := range fileContainers {
		dao := NewNarrativeDAO(fc.GetData(), fc.GetPath(), fc.GetFilename())
		dao.ExtractFromJson()
		dtos = append(dtos, dao.Dto())
	}
	return dtos
}

func ReadPages(pagesDir string) []staticIntf.PageDto {
	fileContainers := ReadJsonFilesFromDir(pagesDir)
	dtos := []staticIntf.PageDto{}
	for _, fc := range fileContainers {
		dao := NewPageDAO(fc.GetData(), fc.GetPath(), fc.GetFilename())
		dao.ExtractFromJson()
		dtos = append(dtos, dao.Dto())
	}
	return dtos
}

func ReadMarginals(marginalsDir string) []staticIntf.PageDto {
	fileContainers := ReadJsonFilesFromDir(marginalsDir)
	dtos := []staticIntf.PageDto{}
	for _, fc := range fileContainers {
		dao := NewMarginalDAO(fc.GetData(), fc.GetPath(), fc.GetFilename())
		dao.ExtractFromJson()
		dtos = append(dtos, dao.Dto())
	}
	return dtos
}

func ReadPosts(postsDir string) []staticIntf.PageDto {
	fileContainers := ReadJsonFilesFromDir(postsDir)
	dtos := []staticIntf.PageDto{}
	for _, fc := range fileContainers {
		dao := NewPostDAO(fc.GetData(), fc.GetPath(), fc.GetFilename())
		dao.ExtractFromJson()
		dtos = append(dtos, dao.Dto())
	}
	return dtos
}

func ReadJsonFilesFromDir(path string) []fs.FileContainer {
	files := fs.ReadDirEntriesEndingWith(path, "json")
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

func WriteMarginalDtoToJson(dto staticIntf.PageDto, path, filename string) {
	dao := NewMarginalDAO(nil, path, filename)
	dao.Dto(dto)
	writeJson(dao.FillJson(), path, filename)
}

func WritePostDtoToJson(dto staticIntf.PageDto, path, filename string) {
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
