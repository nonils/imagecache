package dto

type AuthRequest struct {
	ApiKey string `json:"apiKey"`
}

type AuthResponse struct {
	Auth  bool   `json:"auth"`
	Token string `json:"token"`
}

type Picture struct {
	Id             string `json:"id"`
	CroppedPicture string `json:"cropped_picture"`
}

type GetPicturesResponse struct {
	Pictures  []Picture `json:"pictures"`
	Page      int       `json:"page"`
	PageCount int       `json:"pageCount"`
	HasMore   bool      `json:"hasMore"`
}
