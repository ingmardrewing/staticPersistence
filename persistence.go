package staticPersistence

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

// dao versions
type PageDao interface {
	Dto(...staticIntf.PageDto) staticIntf.PageDto
	ExtractFromJson()
	FillJson() []byte
	Data([]byte)
}

func ReadPagesFromDir(dir string) []staticIntf.PageDto {
	fileContainers := ReadJsonFilesFromDir(dir)
	dtos := []staticIntf.PageDto{}
	for _, fc := range fileContainers {
		dtos = append(dtos, getDto(fc))
	}
	return dtos
}

func getDto(fc fs.FileContainer) staticIntf.PageDto {
	dao := newPageDaoReader(fc.GetData(), fc.GetPath(), fc.GetFilename())
	dao.ExtractFromJson()
	return dao.Dto()
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
	writeDtoToJson(dto, path, filename)
}

func WritePostDtoToJson(dto staticIntf.PageDto, path, filename string) {
	writeDtoToJson(dto, path, filename)
}

func writeDtoToJson(dto staticIntf.PageDto, path, filename string) {
	dao := newPageDaoReader(nil, path, filename)
	dao.Dto(dto)

	fc := fs.NewFileContainer()
	fc.SetDataAsString(string(dao.FillJson()))
	fc.SetPath(path)
	fc.SetFilename(filename)
	fc.Write()
}
