package staticPersistence

import (
	"encoding/json"

	"github.com/ingmardrewing/fs"
)

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

type JsonConfig struct {
	Domain     string `json:"domain"`
	AddPostDir string `json: "addPostDir"`
	AddPageDir string `json: "addPageDir"`
	Src        struct {
		PostsDir    string `json: "postsDir"`
		MarginalDir string `json: "marginalDir"`
		MainPages   string `json: "mainPages"`
		Narrative   string `json: "narrative"`
	} `json: "src"`
	DefaultMeta struct {
		BlogExcerpt     string `json: "blogExcerpt"`
		NaviPageTitle   string `json: "naviPageTitle"`
		NaviPageExcerpt string `json: "naviPageExcerpt"`
	} `json: "defaultMeta"`
	Context struct {
		TwitterHandle   string `json: "twitterHandle"`
		Topic           string `json: "topic"`
		Tags            string `json: "tags"`
		CardType        string `json: "cardType"`
		Section         string `json: "section"`
		FbPage          string `json: "fbPage"`
		TwitterPage     string `json: "twitterPage"`
		DisqusShortname string `json: "disqusShortname"`
		FacebookShare   string `json: "facebookShare"`
		TellAFriend     string `json: "tellAFriend"`
	} `json: "context"`
	Deploy struct {
		TargetDir   string `json: "targetDir"`
		CssFileName string `json: "cssFileName"`
		JsFileName  string `json: "jsFileName"`
		BlogDir     string `json: "jsFileName"`
		Rss         string `json: "rss"`
	} `json: "deploy"`
}
