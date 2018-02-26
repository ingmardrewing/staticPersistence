package staticPersistence

import "github.com/ingmardrewing/staticIntf"

type ContextDao interface {
	Dto() staticIntf.ContextDto
	NarrativeDto() staticIntf.ContextDto
}

func NewContextDao(config *Config) ContextDao {
	dao := new(contextDao)
	dao.config = config
	return dao
}

type contextDao struct {
	config *Config
}

func (c *contextDao) Dto() staticIntf.ContextDto {
	return NewContextDto(
		c.config.Read("context", "twitterHandle"),
		c.config.Read("context", "topic"),
		c.config.Read("context", "tags"),
		c.config.Read("context", "site"),
		c.config.Read("context", "cardType"),
		c.config.Read("context", "section"),
		c.config.Read("context", "fbPage"),
		c.config.Read("context", "twitterPage"),
		c.config.Read("context", "rss"),
		c.config.Read("deploy", "cssFileName"),
		c.config.Read("domain"),
		c.config.Read("context", "disqusShortname"))
}

func (c *contextDao) NarrativeDto() staticIntf.ContextDto {
	return NewContextDto(
		c.config.Read("context", "twitterHandle"),
		c.config.Read("context", "topic"),
		c.config.Read("context", "tags"),
		c.config.Read("context", "site"),
		c.config.Read("context", "cardType"),
		c.config.Read("context", "section"),
		c.config.Read("context", "fbPage"),
		c.config.Read("context", "twitterPage"),
		c.config.Read("context", "rss"),
		c.config.Read("deploy", "cssFileName"),
		c.config.Read("domain"),
		"devabode")
}
