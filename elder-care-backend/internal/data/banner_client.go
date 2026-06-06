package data

type BannerClientData struct {
	BannerId string `json:"banner_id"`

	Title    string `json:"title"`
	Abstract string `json:"abstract"`

	ImageUrl string `json:"image_url"`
	LinkUrl  string `json:"link_url"`
}

type ListBannersClientRespData struct {
	Banners []*BannerClientData `json:"banners"`
}
