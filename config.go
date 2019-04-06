package staticPersistence

import "encoding/json"

func NewConfigs(data []byte) []Config {
	siteConfigs := make([]Config, 0)
	json.Unmarshal(data, &siteConfigs)
	return siteConfigs
}

type Link struct {
	Label        string `json:"label"`
	ExternalLink string `json:"externalLink"`
	FileName     string `json:"filename"`
	Path         string `json:"path"`
}

type Src struct {
	Dir      string `json:"dir"`
	Type     string `json:"type"`
	Headline string `json:"headline"`
	SubDir   string `json:"subDir"`
}

type DefaultByTag struct {
	Tag     string `json:"tag"`
	Excerpt string `json:"excerpt"`
	Content string `json:"content"`
}

type Config struct {
	Domain       string `json:"domain"`
	SvgLogo      string `json:"svgLogo"`
	BasePath     string `json:"basePath"`
	HomeText     string `json:"homeText"`
	HomeHeadline string `json:"homeHeadline"`
	AddPostDir   string `json:"addPostDir"`
	WritePostDir string `json:"writePostDir"`
	AddPageDir   string `json:"addPageDir"`
	WritePageDir string `json:"writePageDir"`
	Src          []Src  `json:"src"`
	DefaultMeta  struct {
		BlogExcerpt     string         `json:"blogExcerpt"`
		KeyWords        string         `json:"key_words"`
		Subject         string         `json:"subject"`
		Author          string         `json:"author"`
		NaviPageTitle   string         `json:"naviPageTitle"`
		NaviPageExcerpt string         `json:"naviPageExcerpt"`
		DefaultByTags   []DefaultByTag `json:"defaultByTags"`
	} `json:"defaultMeta"`
	Context struct {
		TwitterHandle   string `json:"twitterHandle"`
		Topic           string `json:"topic"`
		Tags            string `json:"tags"`
		CardType        string `json:"cardType"`
		Section         string `json:"section"`
		FbPage          string `json:"fbPage"`
		TwitterPage     string `json:"twitterPage"`
		DisqusShortname string `json:"disqusShortname"`
		FacebookShare   string `json:"facebookShare"`
		TellAFriend     string `json:"tellAFriend"`
		MainLinks       []Link `json:"header"`
		MarginalLinks   []Link `json:"footer"`
	} `json:"context"`
	Deploy struct {
		TargetDir   string `json:"targetDir"`
		CssFileName string `json:"cssFileName"`
		JsFileName  string `json:"jsFileName"`
		BlogDir     string `json:"blog"`
		RssPath     string `json:"rssPath"`
		RssFilename string `json:"rssFilename"`
	} `json:"deploy"`
}
