package data

type ImgUploadRequest struct {
	ImageType string `json:"image_type" validate:"required,oneof=jpg png jpeg"`
}

type ImgUploadResponse struct {
	ImageId   string `json:"image_id"`
	UploadURL string `json:"upload_url"`
}
