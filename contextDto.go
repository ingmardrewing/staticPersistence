package staticPersistence

import "github.com/ingmardrewing/staticIntf"

func NewContextDto(twitterHandle, topic, tags, site, cardType, section, fbPage, twitterPage, rss, css, domain, disqus string) staticIntf.ContextDto {
	dto := new(contextDto)
	dto.TwitterHandle(twitterHandle)
	dto.Topic(topic)
	dto.Tags(tags)
	dto.Site(site)
	dto.CardType(cardType)
	dto.Section(section)
	dto.FBPage(fbPage)
	dto.TwitterPage(twitterPage)
	dto.Rss(rss)
	dto.Css(css)
	dto.Domain(domain)
	dto.DisqusId(disqus)
	return dto
}

type contextDto struct {
	twitterHandle string
	topic         string
	tags          string
	site          string
	cardType      string
	section       string
	fbPage        string
	twitterPage   string
	rss           string
	css           string
	domain        string
	disqusId      string
}

func (c *contextDto) TwitterHandle(twitterHandle ...string) string {
	if len(twitterHandle) > 0 {
		c.twitterHandle = twitterHandle[0]
	}
	return c.twitterHandle
}

func (c *contextDto) Topic(topic ...string) string {
	if len(topic) > 0 {
		c.topic = topic[0]
	}
	return c.topic
}

func (c *contextDto) Tags(tags ...string) string {
	if len(tags) > 0 {
		c.tags = tags[0]
	}
	return c.tags
}

func (c *contextDto) Site(site ...string) string {
	if len(site) > 0 {
		c.site = site[0]
	}
	return c.site
}

func (c *contextDto) CardType(cardType ...string) string {
	if len(cardType) > 0 {
		c.cardType = cardType[0]
	}
	return c.cardType
}

func (c *contextDto) Section(section ...string) string {
	if len(section) > 0 {
		c.section = section[0]
	}
	return c.section
}

func (c *contextDto) FBPage(fbPage ...string) string {
	if len(fbPage) > 0 {
		c.fbPage = fbPage[0]
	}
	return c.fbPage
}

func (c *contextDto) TwitterPage(twitterPage ...string) string {
	if len(twitterPage) > 0 {
		c.twitterPage = twitterPage[0]
	}
	return c.twitterPage
}

func (c *contextDto) Rss(rss ...string) string {
	if len(rss) > 0 {
		c.rss = rss[0]
	}
	return c.rss
}

func (c *contextDto) Css(css ...string) string {
	if len(css) > 0 {
		c.css = css[0]
	}
	return c.css
}

func (c *contextDto) Domain(domain ...string) string {
	if len(domain) > 0 {
		c.domain = domain[0]
	}
	return c.domain
}

func (c *contextDto) DisqusId(disqusId ...string) string {
	if len(disqusId) > 0 {
		c.disqusId = disqusId[0]
	}
	return c.disqusId
}
