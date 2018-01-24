package staticPersistence

import "github.com/ingmardrewing/fs"

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
