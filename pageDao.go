package staticPersistence

import (
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/ingmardrewing/staticIntf"
)

func newPageDaoReader(data []byte, path, filename string) *pageDaoReader {
	d := new(pageDaoReader)
	dto := NewFilledDto(0, "", "", "", "", "", "",
		"", "", "", "", path, filename, "", "", "", "")
	d.Dto(dto)
	d.Data(data)
	return d
}

type pageDaoReader struct {
	data []byte
	Json
	dto staticIntf.PageDto
}

func (a *pageDaoReader) extractUrlParts(url string) (string, string) {
	if len(url) > 0 {
		parts := strings.Split(url, "/")
		if len(parts) > 3 {
			return strings.Join(parts[3:], "/"), parts[2]
		} else {
			log.Debug(parts)
		}
	}
	log.Debug("URL missing while reading from json")
	return "", ""
}

func (a *pageDaoReader) ExtractFromJson() {
	var doc docJson
	json.Unmarshal(a.data, &doc)

	thumbUrl := doc.ThumbImg
	imageUrl := doc.PostImg
	htmlFilename := doc.Filename
	id := doc.Id
	createDate := doc.Date
	url := doc.Url
	title := doc.Title
	titlePlain := doc.Title_plain
	description := doc.Excerpt
	content := doc.Content
	disqusId := doc.Dsq_thread_id
	thumbBase64 := doc.ThumbBase64
	category := doc.Category
	microThumbUrl := doc.MicroThumbUrl

	log.Debug(url)
	pathFromDocRoot, domain := a.extractUrlParts(url)

	a.dto = NewFilledDto(
		id,
		title,
		titlePlain,
		thumbUrl,
		imageUrl,
		description,
		disqusId,
		createDate,
		content,
		url,
		domain,
		pathFromDocRoot,
		pathFromDocRoot,
		htmlFilename,
		thumbBase64,
		category,
		microThumbUrl)

	if len(pathFromDocRoot) == 0 && len(htmlFilename) == 0 {
		log.Fatalln(doc)
	}

}

func (a *pageDaoReader) Data(data []byte) {
	a.data = data
}

func (a *pageDaoReader) Dto(dto ...staticIntf.PageDto) staticIntf.PageDto {
	if len(dto) > 0 {
		a.dto = dto[0]
	}
	return a.dto
}

func (a *pageDaoReader) FillJson() []byte {
	json := fmt.Sprintf(a.Template(),
		a.dto.ThumbUrl(),
		a.dto.ImageUrl(),
		a.dto.HtmlFilename(),
		a.dto.Id(),
		a.dto.CreateDate(),
		a.dto.Url(),
		a.dto.Title(),
		a.dto.TitlePlain(),
		a.dto.Description(),
		a.dto.Content(),
		a.dto.DisqusId(),
		a.dto.ThumbBase64(),
		a.dto.Category(),
		a.dto.MicroThumbUrl())
	return []byte(json)
}

func (a *pageDaoReader) Template() string {
	return `{
	"version":1,
	"thumbImg":"%s",
	"postImg":"%s",
	"filename":"%s",
	"id":%d,
	"date":"%s",
	"url":"%s",
	"title":"%s",
	"title_plain":"%s",
	"excerpt":"%s",
	"content":"%s",
	"dsq_thread_id":"%s",
	"thumbBase64":"%s",
	"category":"%s",
	"microThumbUrl":"%s"
}`
}

type docJson struct {
	Version       int    `json:"version"`
	ThumbImg      string `json:"thumbImg"`
	PostImg       string `json:"postImg"`
	Filename      string `json:"filename"`
	Id            int    `json:"id"`
	Date          string `json:"date"`
	Url           string `json:"url"`
	Title         string `json:"title"`
	Title_plain   string `json:"title_plain"`
	Excerpt       string `json:"excerpt"`
	Content       string `json:"content"`
	Dsq_thread_id string `json:"dsq_thread_id"`
	ThumbBase64   string `json:"thumbBase64"`
	Category      string `json:"category"`
	MicroThumbUrl string `json:"microThumbUrl"`
}
