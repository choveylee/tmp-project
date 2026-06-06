package data

type BannerAdminData struct {
	BannerId string `json:"banner_id"`

	Title    string `json:"title"`
	Abstract string `json:"abstract"`

	ImageUrl string `json:"image_url"`
	LinkUrl  string `json:"link_url"`

	Weight int `json:"weight"`

	Status int `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListBannersAdminRespData struct {
	Banners []*BannerAdminData `json:"banners"`

	Total int64 `json:"total"`
}

type CreateBannerAdminRequest struct {
	Title    string `json:"title"`
	Abstract string `json:"abstract"`

	ImageUrl string `json:"image_url"`
	LinkUrl  string `json:"link_url"`

	Weight int `json:"weight"`

	Status int `json:"status"`
}

type CreateBannerAdminRespData struct {
	BannerId string `json:"banner_id"`
}

type GetBannerAdminRespData struct {
	BannerId string `json:"banner_id"`

	Title    string `json:"title"`
	Abstract string `json:"abstract"`

	ImageUrl string `json:"image_url"`
	LinkUrl  string `json:"link_url"`

	Weight int `json:"weight"`

	Status int `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateBannerAdminRequest struct {
	Title    string `json:"title"`
	Abstract string `json:"abstract"`

	ImageUrl string `json:"image_url"`
	LinkUrl  string `json:"link_url"`

	Weight int `json:"weight"`

	Status int `json:"status"`
}
