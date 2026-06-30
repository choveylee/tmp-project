package data

type CourseCategoryAdminData struct {
	CategoryId string `json:"category_id"`

	Name string `json:"name"`

	Weight int `json:"weight"`

	Status int `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListCourseCategoriesAdminRespData struct {
	Categories []*CourseCategoryAdminData `json:"categories"`

	Total int64 `json:"total"`
}

type CreateCourseCategoryAdminRequest struct {
	Name string `json:"name"`

	Weight int `json:"weight"`

	Status int `json:"status"`
}

type CreateCourseCategoryAdminRespData struct {
	CategoryId string `json:"category_id"`
}

type GetCourseCategoryAdminRespData struct {
	CategoryId string `json:"category_id"`

	Name string `json:"name"`

	Weight int `json:"weight"`

	Status int `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateCourseCategoryAdminRequest struct {
	Name string `json:"name"`

	Weight int `json:"weight"`

	Status int `json:"status"`
}

type CourseAdminData struct {
	CourseId string `json:"course_id"`

	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`

	CourseType int `json:"course_type"`

	Author string `json:"author"`
	Source string `json:"source"`

	Title string `json:"title"`

	Abstract string `json:"abstract"`

	CoverUrl string `json:"cover_url"`
	LinkUrl  string `json:"link_url"`

	PublishAt string `json:"publish_at"`

	Status int `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListCoursesAdminRespData struct {
	Courses []*CourseAdminData `json:"courses"`

	Total int64 `json:"total"`
}

type CreateCourseAdminRequest struct {
	CategoryId string `json:"category_id"`
	CourseType int    `json:"course_type"`

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

	PublishAt string `json:"publish_at"`

	Status int `json:"status"`
}

type CreateCourseAdminRespData struct {
	CourseId string `json:"course_id"`
}

type GetCourseAdminRespData struct {
	CourseId string `json:"course_id"`

	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`

	CourseType int `json:"course_type"`

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

	PublishAt string `json:"publish_at"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateCourseAdminRequest struct {
	CategoryId string `json:"category_id"`

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

	PublishAt string `json:"publish_at"`

	Status int `json:"status"`
}

type CourseVideoAdminData struct {
	VideoId string `json:"video_id"`

	VideoUrl string `json:"video_url"`

	Format   string `json:"format"`
	Language string `json:"language"`
	Size     string `json:"size"`
	Duration string `json:"duration"`

	UploadAt string `json:"upload_at"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListCourseVideosAdminRespData struct {
	Videos []*CourseVideoAdminData `json:"videos"`
}

type CreateCourseVideoAdminRequest struct {
	VideoUrl string `json:"video_url"`

	Format   string `json:"format"`
	Language string `json:"language"`
	Size     string `json:"size"`
	Duration string `json:"duration"`

	UploadAt string `json:"upload_at"`

	Weight int `json:"weight"`
	Status int `json:"status"`
}

type CreateCourseVideoAdminRespData struct {
	VideoId string `json:"video_id"`
}

type UpdateCourseVideoAdminRequest struct {
	VideoUrl string `json:"video_url"`

	Format   string `json:"format"`
	Language string `json:"language"`
	Size     string `json:"size"`
	Duration string `json:"duration"`

	UploadAt string `json:"upload_at"`

	Weight int `json:"weight"`
	Status int `json:"status"`
}
