package data

type ArticleCategoryClientData struct {
	CategoryId string `json:"category_id"`

	Name string `json:"name"`
}

type ListArticleCategoriesClientRespData struct {
	Categories []*ArticleCategoryClientData `json:"categories"`
}

type ArticleClientData struct {
	ArticleId string `json:"article_id"`

	Title string `json:"title"`

	Abstract string `json:"abstract"`
	Content  string `json:"content"`

	CoverUrl string `json:"cover_url"`
	LinkUrl  string `json:"link_url"`

	PublishAt string `json:"publish_at"`
}

type ListArticlesClientRespData struct {
	Articles []*ArticleClientData `json:"articles"`

	Total int64 `json:"total"`
}
