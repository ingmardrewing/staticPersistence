package staticPersistence

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// narrativeDAO
func NewNarrativeDAO(data []byte, path, filename string) DAO {
	dto := NewDto()
	dto.FsPath(path)
	dto.HtmlFilename(filename)

	d := new(narrativeDAOv0)
	d.Dto(dto)
	d.Data(data)

	return d
}

// Original data structure from wordpress migration
// still having an unneccessary complex structure
// staying here for historical reasons
type narrativeDAOv0 struct {
	data []byte
	Json
	dto DTO
}

func (p *narrativeDAOv0) Data(data []byte) {
	p.data = data
}

func (p *narrativeDAOv0) Dto(dto ...DTO) DTO {
	if len(dto) > 0 {
		p.dto = dto[0]
	}
	return p.dto
}

func (p *narrativeDAOv0) ExtractFromJson() {
	p.dto.Id(p.ReadInt(p.data, "id"))
	p.dto.Title(p.ReadString(p.data, "title"))
	p.dto.ThumbUrl(p.ReadString(p.data, "imgUrl"))
	p.dto.ImageUrl(p.ReadString(p.data, "imgUrl"))
	p.dto.Description(p.ReadString(p.data, "description"))
	p.dto.DisqusId(p.ReadString(p.data, "disqusId"))
	p.dto.CreateDate(p.getDateFromFSPath())
	p.dto.Content(p.ReadString(p.data, "act"))
	p.dto.Category(p.ReadString(p.data, "act"))

	p.dto.PathFromDocRoot(p.ReadString(p.data, "path"))
	p.dto.ThumbBase64(p.ReadString(p.data, "thumbBase64"))
	p.dto.HtmlFilename("index.html")
}

func (p *narrativeDAOv0) getDateFromFSPath() string {
	fp := p.ReadString(p.data, "path")
	parts := strings.Split(fp, "/")
	loc, _ := time.LoadLocation("Europe/Berlin")
	y, _ := strconv.Atoi(parts[1])
	m, _ := strconv.Atoi(parts[2])
	d, _ := strconv.Atoi(parts[3])
	date := time.Date(y, time.Month(m), d, 20, 0, 0, 0, loc)
	return date.Format(time.RFC1123Z)
}

func (p *narrativeDAOv0) FillJson() []byte {
	json := fmt.Sprintf(p.Template(),
		p.dto.Id(),
		p.dto.Id(),
		p.dto.Title(),
		p.dto.Description(),
		p.dto.Url(),
		p.dto.ImageUrl(),
		p.dto.DisqusId(),
		p.dto.Id())
	return []byte(json)
}

func (p *narrativeDAOv0) Template() string {
	return `{
	 "id":%d,
	 "pageNumber":%d,
	 "title":"%s",
	 "description":"%s",
	 "path":"%s",
	 "imgUrl":"%s",
	 "disqusId":"%s",
	 "act":"%s"
 }`
}
