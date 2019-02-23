package staticPersistence

import (
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"
)

func newPageDaoReader(data []byte, path, filename string) *pageDaoReader {
	d := new(pageDaoReader)
	dto := NewFilledDto(
		"",
		"",
		"",
		"",
		"",
		path,
		filename,
		[]string{},
		[]staticIntf.Image{})
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

	images := []staticIntf.Image{}
	if len(doc.ImagesUrls) > 0 {
		images = append(images,
			NewImageDto(
				doc.ImagesUrls[0].Title,
				doc.ImagesUrls[0].W85Square,
				doc.ImagesUrls[0].W190Square,
				doc.ImagesUrls[0].W390Square,
				doc.ImagesUrls[0].W800Square,
				doc.ImagesUrls[0].W800,
				doc.ImagesUrls[0].W1600,
				doc.ImagesUrls[0].MaxResolution))
	}
	log.Debug(doc.ImagesUrls)

	a.dto = NewFilledDto(
		doc.Title,
		doc.Description,
		doc.Content,
		doc.Category,
		doc.CreateDate,
		doc.PathFromDocRoot,
		doc.Filename,
		doc.Tags,
		images)
}

// Setter, accepts a splice of bytes
func (a *pageDaoReader) Data(data []byte) {
	a.data = data
}

// Getter / setter, optionally accepting
// a staticIntf.PageDto
// and returning dto (nil, if nothing was stored)
func (a *pageDaoReader) Dto(dto ...staticIntf.PageDto) staticIntf.PageDto {
	if len(dto) > 0 {
		a.dto = dto[0]
	}
	return a.dto
}

// Fills a string json template with the values
// of the page dto via fmt.Sprintf
// and returns a splice of byte
func (a *pageDaoReader) FillJson() []byte {
	json := fmt.Sprintf(a.Template(),
		a.dto.Filename(),
		a.dto.PathFromDocRoot(),
		a.dto.Category(),
		a.getTagsAsString(),
		a.dto.CreateDate(),
		a.dto.Title(),
		a.dto.Description(),
		a.dto.Content(),
		a.getImagesAsString())
	return []byte(json)
}

func (a *pageDaoReader) getTagsAsString() string {
	return staticUtil.JoinStrings(a.dto.Tags(), `,`, `"`)
}

func (a *pageDaoReader) getImagesAsString() string {
	imgStrings := []string{}
	for _, img := range a.dto.Images() {
		imgString := fmt.Sprintf(a.ImageTemplate(),
			img.Title(),
			img.W85Square(),
			img.W190Square(),
			img.W390Square(),
			img.W800Square(),
			img.W800(),
			img.W1600(),
			img.MaxResolution())
		imgStrings = append(imgStrings, imgString)
	}
	return strings.Join(imgStrings, ",\n")
}

func (a *pageDaoReader) ImageTemplate() string {
	return `{
		"title":"%s",
		"w_85":"%s",
		"w_190":"%s",
		"w_390":"%s",
		"w_800":"%s",
		"w_800_portrait":"%s",
		"w_1600_portrait":"%s",
		"max_resolution":"%s"
	}`
}

func (a *pageDaoReader) Template() string {
	return `{
	"version":2,
	"filename":"%s",
	"path_from_doc_root":"%s",
	"category":"%s",
	"tags":[%s],
	"create_date":"%s",
	"title":"%s",
	"excerpt":"%s",
	"content":"%s",
	"images_urls":[%s]
}`
}

type imageUrls struct {
	Title         string `json:"title"`
	W85Square     string `json:"w_85"`
	W190Square    string `json:"w_190"`
	W390Square    string `json:"w_390"`
	W800Square    string `json:"w_800"`
	W800          string `json:"w_800_portrait"`
	W1600         string `json:"w_800_portrait"`
	MaxResolution string `json:"max_resolution"`
}

type docJson struct {
	Version         int         `json:"version"`
	Filename        string      `json:"filename"`
	PathFromDocRoot string      `json:"path_from_doc_root"`
	Category        string      `json:"category"`
	CreateDate      string      `json:"create_date"`
	Title           string      `json:"title"`
	Description     string      `json:"excerpt"`
	Content         string      `json:"content"`
	Tags            []string    `json:"tags"`
	ImagesUrls      []imageUrls `json:"images_urls"`
}
