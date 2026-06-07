package data

type ArticleCategoryAdminData struct {
	CategoryId string `json:"category_id"`

	Name string `json:"name"`

	Weight int `json:"weight"`

	Status int `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListArticleCategoriesAdminRespData struct {
	Categories []*ArticleCategoryAdminData `json:"categories"`

	Total int64 `json:"total"`
}

type CreateArticleCategoryAdminRequest struct {
	Name string `json:"name"`

	Weight int `json:"weight"`

	Status int `json:"status"`
}

type CreateArticleCategoryAdminRespData struct {
	CategoryId string `json:"category_id"`
}

type GetArticleCategoryAdminRespData struct {
	CategoryId string `json:"category_id"`

	Name string `json:"name"`

	Weight int `json:"weight"`

	Status int `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateArticleCategoryAdminRequest struct {
	Name string `json:"name"`

	Weight int `json:"weight"`

	Status int `json:"status"`
}

type ArticleAdminData struct {
	ArticleId string `json:"article_id"`

	Title string `json:"title"`

	Abstract string `json:"abstract"`
	Content  string `json:"content"`

	CoverUrl string `json:"cover_url"`
	LinkUrl  string `json:"link_url"`

	PublishAt string `json:"publish_at"`

	Status int `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListArticlesAdminRespData struct {
	Articles []*ArticleAdminData `json:"articles"`

	Total int64 `json:"total"`
}

type CreateArticleAdminRequest struct {
	CategoryId string `json:"category_id"`

	Title string `json:"title"`

	Abstract string `json:"abstract"`
	Content  string `json:"content"`

	CoverUrl string `json:"cover_url"`
	LinkUrl  string `json:"link_url"`

	PublishAt string `json:"publish_at"`

	Status int `json:"status"`
}

type CreateArticleAdminRespData struct {
	ArticleId string `json:"article_id"`
}

type GetArticleAdminRespData struct {
	ArticleId string `json:"article_id"`

	Title string `json:"title"`

	Abstract string `json:"abstract"`
	Content  string `json:"content"`

	CoverUrl string `json:"cover_url"`
	LinkUrl  string `json:"link_url"`

	PublishAt string `json:"publish_at"`

	Status int `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateArticleAdminRequest struct {
	Title string `json:"title"`

	Abstract string `json:"abstract"`
	Content  string `json:"content"`

	CoverUrl string `json:"cover_url"`
	LinkUrl  string `json:"link_url"`

	PublishAt string `json:"publish_at"`

	Status int `json:"status"`
}
