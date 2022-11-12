package social_media_resource

type SocialMediaCreateResponse struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         int64  `json:"user_id"`
	CreatedAt      string `json:"created_at"`
}

type SocialMediaUserGetResponse struct {
	Id              int64  `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type SocialMediaGetResponse struct {
	Id             int64                       `json:"id"`
	Name           string                      `json:"name"`
	SocialMediaUrl string                      `json:"social_media_url"`
	UserId         int64                       `json:"user_id"`
	CreatedAt      string                      `json:"created_at"`
	UpdatedAt      string                      `json:"updated_at"`
	User           *SocialMediaUserGetResponse `json:"user"`
}

type SocialMediaUpdateResponse struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	SocialMediaId string `json:"social_media_url"`
	UserId        int64  `json:"user_id"`
	UpdatedAt     string `json:"updated_at"`
}
