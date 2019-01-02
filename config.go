package staticPersistence

import (
	"encoding/json"

	"github.com/ingmardrewing/fs"
)

// Reads the json config for the sites
// returning a JsonConfig
func ReadConfig(path, file string) []JsonConfig {
	fc := fs.NewFileContainer()
	fc.SetFilename(file)
	fc.SetPath(path)
	fc.Read()

	configBytes := fc.GetData()
	configStructs := make([]JsonConfig, 0)
	json.Unmarshal(configBytes, &configStructs)

	return configStructs
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

type JsonConfig struct {
	Domain       string `json:"domain"`
	HomeText     string `json:"homeText"`
	HomeHeadline string `json:"homeHeadline"`
	AddPostDir   string `json:"addPostDir"`
	WritePostDir string `json:"writePostDir"`
	AddPageDir   string `json:"addPageDir"`
	WritePageDir string `json:"writePageDir"`
	Src          []Src  `json:"src"`
	DefaultMeta  struct {
		BlogExcerpt     string `json:"blogExcerpt"`
		NaviPageTitle   string `json:"naviPageTitle"`
		NaviPageExcerpt string `json:"naviPageExcerpt"`
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
