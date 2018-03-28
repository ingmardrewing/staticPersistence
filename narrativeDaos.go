package staticPersistence

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// narrativeDAO
func NewNarrativeDAO(data []byte, path, filename string) PageDao {
	dto := NewFilledDto(0,
		"", "", "", "", "",
		"", "", "", "", "",
		path, filename, "", "")

	d := new(narrativeDAOv0)
	d.Dto(dto)
	d.Data(data)

	return d
}

// Original data structure from wordpress migration
// still having an unneccessary complex structure
// staying here for historical reasons
type narrativeDAOv0 struct {
	abstractPageDao
}

func (p *narrativeDAOv0) ExtractFromJson() {
	id := p.ReadInt(p.data, "id")
	title := p.ReadString(p.data, "title")
	thumbUrl := p.ReadString(p.data, "imgUrl")
	imageUrl := p.ReadString(p.data, "imgUrl")
	description := p.ReadString(p.data, "description")
	disqusId := p.ReadString(p.data, "disqusId")
	createDate := p.getDateFromFSPath()
	content := p.ReadString(p.data, "act")
	pathFromDocRoot := p.ReadString(p.data, "path")
	thumbBase64 := p.ReadString(p.data, "thumbBase64")
	htmlFilename := "index.html"

	p.dto = NewFilledDto(
		id,
		title,
		title,
		thumbUrl,
		imageUrl,
		description,
		disqusId,
		createDate,
		content,
		"",
		"",
		pathFromDocRoot,
		p.dto.FsPath(),
		htmlFilename,
		thumbBase64)

}

func (p *narrativeDAOv0) getDateFromFSPath() string {

	fp := p.ReadString(p.data, "path")
	parts := strings.Split(fp, "/")
	if len(parts) > 3 {
		loc, _ := time.LoadLocation("Europe/Berlin")
		y, _ := strconv.Atoi(parts[1])
		m, _ := strconv.Atoi(parts[2])
		d, _ := strconv.Atoi(parts[3])
		date := time.Date(y, time.Month(m), d, 20, 0, 0, 0, loc)
		return date.Format(time.RFC1123Z)
	}
	return ""
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
