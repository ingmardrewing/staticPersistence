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
	dto := NewFilledDto(0, "", "", "",
		"", "", "", "", path, filename, "", "", "", "",
		[]string{}, []staticIntf.Image{})
	d.Dto(dto)
	d.Data(data)
	return d
}

type pageDaoReader struct {
	data []byte
	Json
	dto staticIntf.PageDto
}

func (a *pageDaoReader) ExtractFromJson() {
	var doc docJson
	json.Unmarshal(a.data, &doc)

	thumbUrl := ""
	imageUrl := ""
	microThumbUrl := ""
	images := []staticIntf.Image{}
	if len(doc.ImagesUrls) > 0 {
		microThumbUrl = doc.ImagesUrls[0].W190
		thumbUrl = doc.ImagesUrls[0].W390
		imageUrl = doc.ImagesUrls[0].W800

		images = append(images,
			NewImageDto(
				doc.ImagesUrls[0].Title,
				doc.ImagesUrls[0].W190,
				doc.ImagesUrls[0].W390,
				doc.ImagesUrls[0].W800,
				doc.ImagesUrls[0].MaxResolution))
	}
	log.Debug(doc.ImagesUrls)

	a.dto = NewFilledDto(
		0,
		doc.Title,
		doc.TitlePlain,
		thumbUrl,
		imageUrl,
		doc.Description,
		doc.CreateDate,
		doc.Content,
		doc.PathFromDocRoot,
		doc.PathFromDocRoot,
		doc.Filename,
		doc.ThumbBase64,
		doc.Category,
		microThumbUrl,
		doc.Tags,
		images)
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
	img := fmt.Sprintf(a.TemplateImageUrls(),
		a.dto.Title(),
		a.dto.MicroThumbUrl(),
		a.dto.ThumbUrl(),
		a.dto.ImageUrl(),
		"")

	tags := "[]"
	if len(a.dto.Tags()) > 0 {
		tags = `["` + strings.Join(a.dto.Tags(), `","`) + `"]`
	}
	json2 := fmt.Sprintf(a.Template(),
		a.dto.HtmlFilename(),
		a.dto.PathFromDocRoot(),
		a.dto.Category(),
		tags,
		a.dto.CreateDate(),
		a.dto.Title(),
		a.dto.TitlePlain(),
		a.dto.Description(),
		a.dto.Content(),
		a.dto.ThumbBase64(),
		img)

	return []byte(json2)
}

func (a *pageDaoReader) TemplateImageUrls() string {
	return `{"title":"%s","w_190":"%s","w_390":"%s","w_800":"%s","max_resolution":"%s"}`
}

func (a *pageDaoReader) Template() string {
	return `{
	"version":2,
	"filename":"%s",
	"path_from_doc_root":"%s",
	"category":"%s",
	"tags":%s,
	"create_date":"%s",
	"title":"%s",
	"title_plain":"%s",
	"excerpt":"%s",
	"content":"%s",
	"thumb_base64":"%s",
	"images_urls":[%s]
}`
}

type imageUrls struct {
	Title         string `json:"title"`
	W190          string `json:"w_190"`
	W390          string `json:"w_390"`
	W800          string `json:"w_800"`
	MaxResolution string `json:"max_resolution"`
}

type docJson struct {
	Version         int         `json:"version"`
	Filename        string      `json:"filename"`
	PathFromDocRoot string      `json:"path_from_doc_root"`
	Category        string      `json:"category"`
	Tags            []string    `json:"tags"`
	CreateDate      string      `json:"create_date"`
	Title           string      `json:"title"`
	TitlePlain      string      `json:"title_plain"`
	Description     string      `json:"excerpt"`
	Content         string      `json:"content"`
	ThumbBase64     string      `json:"thumb_base64"`
	ImagesUrls      []imageUrls `json:"images_urls"`
}
