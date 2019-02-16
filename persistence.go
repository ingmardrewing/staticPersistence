package staticPersistence

import (
	"fmt"
	"strings"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/rx"
	"github.com/ingmardrewing/staticIntf"
	log "github.com/sirupsen/logrus"
)

// Also used by staticAdd / addJson.go
func JsonFileNameTemplate() string { return "doc%05d.json" }

// dao versions
type PageDao interface {
	Dto(...staticIntf.PageDto) staticIntf.PageDto
	ExtractFromJson()
	FillJson() []byte
	Data([]byte)
}

func ReadPagesFromDir(dir string, domain string) []staticIntf.PageDto {
	fileContainers := ReadJsonFilesFromDir(dir)
	dtos := []staticIntf.PageDto{}
	for _, fc := range fileContainers {
		dtos = append(dtos, getDto(fc, domain))
	}
	return dtos
}

func WritePagesToDir(dtos []staticIntf.PageDto, dirname string) {
	counter := 0
	for _, dto := range dtos {
		WritePageDtoToJson(fixcontent(dto),
			dirname,
			fmt.Sprintf(JsonFileNameTemplate(), counter))
		counter = counter + 1
	}
}

func fixcontent(dto staticIntf.PageDto) staticIntf.PageDto {
	cnew := cleanStringValue(dto.Content())
	regex, err := rx.NewRx("\n|\r|\n\r")
	if err != nil {
		log.Error(err)
	}
	cnew = regex.SubstituteAllOccurences(cnew, "")
	cdate := (strings.Split(dto.CreateDate(), " "))[0]
	dparts := strings.Split(cdate, "-")
	if len(dparts) > 2 {
		y := dparts[0]
		m := dparts[1]
		d := dparts[2]
		cdate = fmt.Sprintf("%04s-%02s-%02s", y, m, d)
	}

	return NewFilledDto(dto.Id(),
		dto.Title(), dto.TitlePlain(), dto.ThumbUrl(),
		dto.ImageUrl(), dto.Description(), dto.DisqusId(),
		cdate, cnew, dto.Url(), dto.Domain(),
		"/"+dto.PathFromDocRoot(), dto.FsPath(), dto.HtmlFilename(),
		dto.ThumbBase64(), dto.Category(), dto.MicroThumbUrl())
}

func getDto(fc fs.FileContainer, domain string) staticIntf.PageDto {
	dao := newPageDaoReader(fc.GetData(), fc.GetPath(), fc.GetFilename())
	log.Debug("reading: " + fc.GetPath() + "/" + fc.GetFilename())
	dao.ExtractFromJson(domain)
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

func WritePageDtoToJson(dto staticIntf.PageDto, path, filename string) {
	newDto := fixJsonValues(dto)
	fmt.Println(newDto.Content())

	dao := newPageDaoReader(nil, path, filename)
	dao.Dto(dto)

	fc := fs.NewFileContainer()
	fc.SetDataAsString(string(dao.FillJson()))
	fc.SetPath(path)
	fc.SetFilename(filename)
	fc.Write()
}

func fixJsonValues(dto staticIntf.PageDto) staticIntf.PageDto {
	return NewFilledDto(
		dto.Id(),
		cleanStringValue(dto.Title()),
		cleanStringValue(dto.TitlePlain()),
		cleanStringValue(dto.ThumbUrl()),
		cleanStringValue(dto.ImageUrl()),
		cleanStringValue(dto.Description()),
		cleanStringValue(dto.DisqusId()),
		cleanStringValue(dto.CreateDate()),
		cleanStringValue(dto.Content()),
		cleanStringValue(dto.Url()),
		cleanStringValue(dto.Domain()),
		cleanStringValue(dto.PathFromDocRoot()),
		cleanStringValue(dto.FsPath()),
		cleanStringValue(dto.HtmlFilename()),
		cleanStringValue(dto.ThumbBase64()),
		cleanStringValue(dto.Category()),
		cleanStringValue(dto.MicroThumbUrl()))
}

func cleanStringValue(dirty string) string {
	withoutLineBreaks := removeLineBreaks(dirty)
	return removeQuotes(withoutLineBreaks)
}

func removeLineBreaks(val string) string {
	lineBreakRx, err := rx.NewRx("\n|\r|\n\r")
	if err != nil {
		log.Error(err)
	}
	return lineBreakRx.SubstituteAllOccurences(val, "")
}

func removeQuotes(val string) string {
	unescapedDoubleQuoteRx, err := rx.NewRx(`""`)
	if err != nil {
		log.Error(err)
	}
	unescapedQuoteRx, err := rx.NewRx(`([^\\])"`)
	if err != nil {
		log.Error(err)
	}
	r1 := unescapedDoubleQuoteRx.SubstituteAllOccurences(val, `"\"`)
	return unescapedQuoteRx.SubstituteAllOccurences(r1, `${1}\"`)
}
