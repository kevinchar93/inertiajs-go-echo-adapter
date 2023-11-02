package page

type PageHeader struct {
	Title string `json:"title"`

	Meta struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	} `json:"meta"`
}

func NewPageHeader(title, metaName, metaContent string) (p *PageHeader) {
	p = new(PageHeader)
	p.Title = title
	p.Meta.Name = metaName
	p.Meta.Content = metaContent
	return p
}
