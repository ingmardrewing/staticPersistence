package staticPersistence

import (
	"strconv"
	"strings"
	"time"

	"github.com/ingmardrewing/staticIntf"
)

type keyPath struct {
	nodes []string
}

type keyCollection struct {
	pathMap map[string][]*keyPath
}

func (k *keyCollection) addKeyPath(key string, path *keyPath) {
	if val, ok := k.pathMap[key]; ok {
		val = append(val, path)
		k.pathMap[key] = val
	} else {
		k.pathMap[key] = []*keyPath{path}
	}
}

func (k *keyCollection) getKeyCollection(key string) []*keyPath {
	return k.pathMap[key]
}

func NewKeyCollection() *keyCollection {
	kc := new(keyCollection)
	kc.pathMap = make(map[string][]*keyPath)

	kc.addKeyPath("id", &keyPath{[]string{"page", "post_id"}})
	kc.addKeyPath("id", &keyPath{[]string{"id"}})

	kc.addKeyPath("title", &keyPath{[]string{"title"}})
	kc.addKeyPath("titlePlain", &keyPath{[]string{"title_plain"}})

	kc.addKeyPath("thumbUrl", &keyPath{[]string{"thumbUrl"}})
	kc.addKeyPath("thumbUrl", &keyPath{[]string{"thumbImg"}})

	kc.addKeyPath("imageUrl", &keyPath{[]string{"imageUrl"}})
	kc.addKeyPath("imageUrl", &keyPath{[]string{"postImg"}})

	kc.addKeyPath("description", &keyPath{[]string{"page", "excerpt"}})
	kc.addKeyPath("disqusId", &keyPath{[]string{"page", "custom_fields", "dsq_thread_id", "[0]"}})

	kc.addKeyPath("createDate", &keyPath{[]string{"page", "date"}})
	kc.addKeyPath("createDate", &keyPath{[]string{"createDate"}})

	kc.addKeyPath("content", &keyPath{[]string{"content"}})
	kc.addKeyPath("pathFromDocRoot", &keyPath{[]string{"path"}})

	kc.addKeyPath("htmlFilename", &keyPath{[]string{"filename"}})

	kc.addKeyPath("ThumbBase64", &keyPath{[]string{"thumbBase64"}})
	return kc
}

type abstractPageDao struct {
	data []byte
	Json
	dto staticIntf.PageDto
}

func (a *abstractPageDao) ReadFirstString(key string) string {
	kc := NewKeyCollection()
	keys := kc.getKeyCollection(key)
	for _, k := range keys {
		txt := a.ReadString(a.data, k.nodes...)
		if len(txt) > 0 {
			return txt
		}
	}
	return ""
}

func (a *abstractPageDao) ReadFirstInt(key string) int {
	kc := NewKeyCollection()
	keys := kc.getKeyCollection(key)
	for _, k := range keys {
		number := a.ReadInt(a.data, k.nodes...)
		if number > 0 {
			return number
		}
	}
	return 0
}

func (a *abstractPageDao) ExtractFromJson() {
	id := a.ReadFirstInt("id")
	title := a.ReadFirstString("title")
	titlePlain := a.ReadFirstString("titlePlain")
	thumbUrl := a.ReadFirstString("thumbImg")
	imageUrl := a.ReadFirstString("postImg")
	description := a.ReadFirstString("description")
	disqusId := a.ReadFirstString("disqusId")
	createDate := a.ReadFirstString("createDate")
	content := a.ReadFirstString("content")
	pathFromDocRoot := a.ReadFirstString("pathFromDocRoot")
	htmlFilename := a.ReadFirstString("htmlFilename")
	thumbBase64 := a.ReadFirstString("thumbBase64")

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
		"",
		"",
		pathFromDocRoot,
		a.dto.FsPath(),
		htmlFilename,
		thumbBase64)
}

func (a *abstractPageDao) getDateFromPath(fp string) string {
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

func (a *abstractPageDao) Data(data []byte) {
	a.data = data
}

func (a *abstractPageDao) Dto(dto ...staticIntf.PageDto) staticIntf.PageDto {
	if len(dto) > 0 {
		a.dto = dto[0]
	}
	return a.dto
}
