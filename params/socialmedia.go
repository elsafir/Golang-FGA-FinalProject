package params

type CreateSocialMedia struct {
	Name           string `json:"name" valid:"required~Field Name is required,maxstringlength(255)~Field Name maximum length is 255 characters,type(string)~Field Name must be string"`
	SocialMediaURL string `json:"social_media_url" valid:"required~Field Social Media URL is required,maxstringlength(255)~Field Social Media URL maximum length is 255 characters,type(string)~Field Social Media URL must be string"`
	UserID         uint   `json:"user_id" valid:"required~Field User ID is required"`
}

type UpdateSocialMedia struct {
	ID             uint   `json:"id" valid:"required~Field ID is required"`
	Name           string `json:"name" valid:"required~Field Name is required,maxstringlength(255)~Field Name maximum length is 255 characters,type(string)~Field Name must be string"`
	SocialMediaURL string `json:"social_media_url" valid:"required~Field Social Media URL is required,maxstringlength(255)~Field Social Media URL maximum length is 255 characters,type(string)~Field Social Media URL must be string"`
	UserID         uint   `json:"user_id" valid:"required~Field User ID is required"`
}
