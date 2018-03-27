package staticPersistence

import "fmt"

// marginalDAO
func NewMarginalDAO(data []byte, path, filename string) PageDao {
	var d PageDao
	switch FindJsonVersion(data) {
	case v1:
		d = new(marginalDAOv1)
	default:
		d = new(marginalDAOv0)
	}

	dto := NewFilledDto(0,
		"", "", "", "", "",
		"", "", "", "", "",
		path, filename, "", "")

	d.Dto(dto)
	d.Data(data)

	return d
}

// Original data structure from wordpress migration
// still having an unneccessary complex structure
// staying here for historical reasons
type marginalDAOv0 struct {
	abstractPageDao
}

func (p *marginalDAOv0) FillJson() []byte {
	json := fmt.Sprintf(p.Template(),
		p.dto.ThumbUrl(),
		p.dto.ImageUrl(),
		p.dto.HtmlFilename(),
		p.dto.Id(),
		p.dto.CreateDate(),
		p.dto.Url(),
		p.dto.Title(),
		p.dto.TitlePlain(),
		p.dto.Description(),
		p.dto.Content(),
		p.dto.DisqusId())
	return []byte(json)
}

func (p *marginalDAOv0) Template() string {
	return `{
  "title":"%s",
  "title_plain":"%s",
  "filename":"%s",
  "content":"%s",
  "page":{
	  "id":%d,
	  "slug":"ingmars-booklist",
	  "url":"https:\/\/www.drewing.de\/blog\/",
	  "status":"publish",
	  "excerpt":"%s",
    "date":"%s",
  }
}`
}

// Marginal DAOs

// Original data structure from wordpress migration
// still having an unneccessary complex structure
// staying here for historical reasons
type marginalDAOv1 struct {
	abstractPageDao
}

func (p *marginalDAOv1) FillJson() []byte {
	json := fmt.Sprintf(p.Template(),
		p.dto.Id(),
		p.dto.PathFromDocRoot(),
		p.dto.HtmlFilename(),
		p.dto.CreateDate(),
		p.dto.Url(),
		p.dto.Title(),
		p.dto.TitlePlain(),
		p.dto.Description(),
		p.dto.Content())
	return []byte(json)
}

func (p *marginalDAOv1) Template() string {
	return `{
	"version": 1,
	"id": %d,
	"path": "%s",
	"filename": "%s",
	"createDate": "%s",
	"url": "%s",
	"title": "%s",
	"title_plain": "%s",
	"description": "%s",
	"content": "%s"
}`
}
