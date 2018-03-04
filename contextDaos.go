package staticPersistence

import "github.com/ingmardrewing/staticIntf"

type ContextDao interface {
	Dto() staticIntf.ContextDto
}

func NewContextDao(config JsonConfig) ContextDao {
	dao := new(contextDao)
	dao.config = config
	return dao
}

type contextDao struct {
	config JsonConfig
}

func (c *contextDao) Dto() staticIntf.ContextDto {
	return NewContextDto(
		c.config.Context.TwitterHandle,
		c.config.Context.Topic,
		c.config.Context.Tags,
		c.config.Domain,
		c.config.Context.CardType,
		c.config.Context.Section,
		c.config.Context.FbPage,
		c.config.Context.TwitterPage,
		c.config.Deploy.Rss,
		c.config.Deploy.CssFileName,
		c.config.Context.DisqusShortname,
		c.config.Deploy.TargetDir)
}
