package model

type Image struct {
	Id             string `json:"id"`
	Author         string `json:"author"`
	Camera         string `json:"camera"`
	Tags           string `json:"tags"`
	CroppedPicture string `json:"cropped_picture"`
	FullPicture    string `json:"full_picture"`
}
