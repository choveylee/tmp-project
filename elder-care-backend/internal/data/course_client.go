package data

type CourseCategoryClientData struct {
	CategoryId string `json:"category_id"`

	ModuleId string `json:"module_id"`

	Name string `json:"name"`
}

type ListCourseCategoriesClientRespData struct {
	Categories []*CourseCategoryClientData `json:"categories"`
}

type CourseClientData struct {
	CourseId string `json:"course_id"`

	Title string `json:"title"`

	Abstract string `json:"abstract"`

	CoverUrl string `json:"cover_url"`
	LinkUrl  string `json:"link_url"`

	PublishAt string `json:"publish_at"`
}

type ListCoursesClientRespData struct {
	Courses []*CourseClientData `json:"courses"`

	Total int64 `json:"total"`
}

type GetCourseClientRespData struct {
	CourseId string `json:"course_id"`

	Author string `json:"author"`
	Source string `json:"source"`

	Title string `json:"title"`

	Tags []string `json:"tags"`

	Abstract string `json:"abstract"`

	CoverUrl string `json:"cover_url"`
	LinkUrl  string `json:"link_url"`

	Detail string `json:"detail"`

	Summary    string `json:"summary"`
	Objective  string `json:"objective"`
	Outline    string `json:"outline"`
	References string `json:"references"`

	FavouriteCount int `json:"favourite_count"`
	ViewCount      int `json:"view_count"`

	PublishAt string `json:"publish_at"`
}

type CourseCatalogVideoClientData struct {
	VideoId string `json:"video_id"`

	VideoUrl string `json:"video_url"`

	Format   string `json:"format"`
	Language string `json:"language"`
	Size     string `json:"size"`
	Duration string `json:"duration"`

	UploadAt string `json:"upload_at"`
}

type CourseCatalogClientData struct {
	CatalogId string `json:"catalog_id"`

	ParentId string `json:"parent_id"`

	Name string `json:"name"`

	Video *CourseCatalogVideoClientData `json:"video"`
}

type ListCourseCatalogsClientRespData struct {
	Catalogs []*CourseCatalogClientData `json:"catalogs"`
}
