package data

type CreateSessionAdminRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`

	CaptchaId   string `json:"captcha_id"`
	CaptchaCode string `json:"captcha_code"`
}

type CreateSessionAdminRespData struct {
	AccessToken string `json:"access_token"`

	ExpiresIn int `json:"expires_in"`

	// User *UserBriefAdminData `json:"user"`
}
