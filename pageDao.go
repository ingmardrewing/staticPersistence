package staticPersistence

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ingmardrewing/staticIntf"
)

func newPageDaoReader(data []byte, path, filename string) *pageDaoReader {
	d := new(pageDaoReader)
	dto := NewFilledDto(0,
		"", "", "", "", "",
		"", "", "", "", "",
		path, filename, "", "", "")

	d.Dto(dto)
	d.Data(data)
	return d
}

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

func newKeyCollection() *keyCollection {
	kc := new(keyCollection)
	kc.pathMap = make(map[string][]*keyPath)

	kc.addKeyPath("url", &keyPath{[]string{"post", "url"}})
	kc.addKeyPath("url", &keyPath{[]string{"url"}})

	kc.addKeyPath("domain", &keyPath{[]string{"domain"}})

	kc.addKeyPath("id", &keyPath{[]string{"post", "post_id"}})
	kc.addKeyPath("id", &keyPath{[]string{"page", "post_id"}})
	kc.addKeyPath("id", &keyPath{[]string{"id"}})

	kc.addKeyPath("title", &keyPath{[]string{"title"}})
	kc.addKeyPath("title", &keyPath{[]string{"post", "title"}})

	kc.addKeyPath("titlePlain", &keyPath{[]string{"title_plain"}})
	kc.addKeyPath("titlePlain", &keyPath{[]string{"title"}})

	kc.addKeyPath("thumbUrl", &keyPath{[]string{"thumbUrl"}})
	kc.addKeyPath("thumbUrl", &keyPath{[]string{"thumbImg"}})
	kc.addKeyPath("thumbUrl", &keyPath{[]string{"imgUrl"}})

	kc.addKeyPath("imageUrl", &keyPath{[]string{"imageUrl"}})
	kc.addKeyPath("imageUrl", &keyPath{[]string{"postImg"}})
	kc.addKeyPath("imageUrl", &keyPath{[]string{"imgUrl"}})

	kc.addKeyPath("description", &keyPath{[]string{"page", "excerpt"}})
	kc.addKeyPath("description", &keyPath{[]string{"description"}})
	kc.addKeyPath("description", &keyPath{[]string{"excerpt"}})

	kc.addKeyPath("disqusId", &keyPath{[]string{"page", "custom_fields", "dsq_thread_id", "[0]"}})
	kc.addKeyPath("disqusId", &keyPath{[]string{"dsq_thread_id"}})
	kc.addKeyPath("disqusId", &keyPath{[]string{"disqusId"}})

	kc.addKeyPath("createDate", &keyPath{[]string{"post", "date"}})
	kc.addKeyPath("createDate", &keyPath{[]string{"page", "date"}})
	kc.addKeyPath("createDate", &keyPath{[]string{"createDate"}})
	kc.addKeyPath("createDate", &keyPath{[]string{"date"}})

	kc.addKeyPath("category", &keyPath{[]string{"category"}})

	kc.addKeyPath("content", &keyPath{[]string{"content"}})
	kc.addKeyPath("content", &keyPath{[]string{"post", "content"}})
	kc.addKeyPath("content", &keyPath{[]string{"act"}})

	kc.addKeyPath("pathFromDocRoot", &keyPath{[]string{"path"}})

	kc.addKeyPath("fsPath", &keyPath{[]string{"fsPath"}})

	kc.addKeyPath("htmlFilename", &keyPath{[]string{"filename"}})

	kc.addKeyPath("thumbBase64", &keyPath{[]string{"thumbBase64"}})

	kc.addKeyPath("version", &keyPath{[]string{"version"}})

	return kc
}

type pageDaoReader struct {
	data []byte
	Json
	dto staticIntf.PageDto
}

func (a *pageDaoReader) ReadFirstString(key string) string {
	kc := newKeyCollection()
	keys := kc.getKeyCollection(key)
	for _, k := range keys {
		txt := a.ReadString(a.data, k.nodes...)
		if len(txt) > 0 {
			return txt
		}
	}
	return ""
}

func (a *pageDaoReader) ReadFirstInt(key string) int {
	kc := newKeyCollection()
	keys := kc.getKeyCollection(key)
	for _, k := range keys {
		number := a.ReadInt(a.data, k.nodes...)
		if number > 0 {
			return number
		}
	}
	return 0
}

func (a *pageDaoReader) checkHtmlFilename(htmlFilename string) string {
	if len(htmlFilename) == 0 {
		return "index.html"
	}
	return htmlFilename
}

func (a *pageDaoReader) generateDomainAndPathFromDocRoot(pathFromDocRoot, domain, url string) (string, string) {
	if len(pathFromDocRoot) == 0 && len(domain) == 0 && len(url) > 0 {
		parts := strings.Split(url, "/")
		if len(parts) > 3 {
			return strings.Join(parts[4:], "/"), parts[2]
		}
	}
	return pathFromDocRoot, domain
}

func (a *pageDaoReader) generateCreateDateFromPathFromDocRoot(createDate, pathFromDocRoot string) string {
	if len(createDate) == 0 && len(pathFromDocRoot) > 0 {
		parts := strings.Split(pathFromDocRoot, "/")
		if len(parts) > 3 {
			loc, _ := time.LoadLocation("Europe/Berlin")
			y, _ := strconv.Atoi(parts[1])
			m, _ := strconv.Atoi(parts[2])
			d, _ := strconv.Atoi(parts[3])
			dt := time.Date(y, time.Month(m), d, 20, 0, 0, 0, loc)
			return dt.Format(time.RFC1123Z)
		}
	}
	return createDate
}

func (a *pageDaoReader) extractFsPathDomainAndPathFromDocRootFromUrl(version int, fsPath, domain, pathFromDocRoot, url string) (string, string, string) {
	if version == 1 && len(fsPath) == 0 && len(url) > 0 {
		parts := strings.Split(url, "/")
		if len(parts) > 3 {
			fsPath = strings.Join(parts[4:], "/")
			domain = parts[2]
			return fsPath, domain, fsPath
		}
	}
	return fsPath, domain, pathFromDocRoot
}

func (a *pageDaoReader) ExtractFromJson() {
	id := a.ReadFirstInt("id")
	version := a.ReadFirstInt("version")
	title := a.ReadFirstString("title")
	titlePlain := a.ReadFirstString("titlePlain")
	thumbUrl := a.ReadFirstString("thumbUrl")
	imageUrl := a.ReadFirstString("imageUrl")
	description := a.ReadFirstString("description")
	disqusId := a.ReadFirstString("disqusId")
	createDate := a.ReadFirstString("createDate")
	content := a.ReadFirstString("content")
	pathFromDocRoot := a.ReadFirstString("pathFromDocRoot")
	fsPath := a.ReadFirstString("fsPath")
	htmlFilename := a.checkHtmlFilename(a.ReadFirstString("htmlFilename"))
	thumbBase64 := a.ReadFirstString("thumbBase64")
	url := a.ReadFirstString("url")
	domain := a.ReadFirstString("domain")
	category := a.ReadFirstString("category")

	pathFromDocRoot, domain = a.generateDomainAndPathFromDocRoot(pathFromDocRoot, domain, url)
	createDate = a.generateCreateDateFromPathFromDocRoot(createDate, pathFromDocRoot)
	fsPath, domain, pathFromDocRoot = a.extractFsPathDomainAndPathFromDocRootFromUrl(version, fsPath, domain, pathFromDocRoot, url)

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
		fsPath,
		htmlFilename,
		thumbBase64,
		category)
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
		a.dto.Category())
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
	"dsq_thread_id":"%s"
	"thumbBase64":"%s"
	"category":"%s"
}`
}
