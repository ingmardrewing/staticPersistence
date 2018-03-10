package staticPersistence

import "github.com/ingmardrewing/fs"

// dao versions
const (
	v0 = iota
	v1 = iota
)

type DTO interface {
	Id(...int) int
	Title(...string) string
	TitlePlain(...string) string
	ThumbUrl(...string) string
	ImageUrl(...string) string
	Description(...string) string
	DisqusId(...string) string
	CreateDate(...string) string
	Content(...string) string
	Category(...string) string
	Url(...string) string
	Domain(...string) string
	PathFromDocRoot(...string) string
	HtmlFilename(...string) string
	ThumbBase64(...string) string
}

type DAO interface {
	Dto(...DTO) DTO
	ExtractFromJson()
	FillJson() []byte
	Data([]byte)
}

func ReadNarrativePages(pagesDir string) []DTO {
	fileContainers := ReadJsonFilesFromDir(pagesDir)
	dtos := []DTO{}
	for _, fc := range fileContainers {
		dao := NewNarrativeDAO(fc.GetData(), fc.GetPath(), fc.GetFilename())
		dao.ExtractFromJson()
		dtos = append(dtos, dao.Dto())
	}
	return dtos
}

func ReadPages(pagesDir string) []DTO {
	fileContainers := ReadJsonFilesFromDir(pagesDir)
	dtos := []DTO{}
	for _, fc := range fileContainers {
		dao := NewPageDAO(fc.GetData(), fc.GetPath(), fc.GetFilename())
		dao.ExtractFromJson()
		dtos = append(dtos, dao.Dto())
	}
	return dtos
}

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
