package params

type CreatePhoto struct {
	Title    string `json:"title" valid:"required~Field Title is required,maxstringlength(255)~Field Title maximum length is 255 characters"`
	Caption  string `json:"caption" valid:"required~Field Caption is required"`
	PhotoURL string `json:"photo_url" valid:"required~Field Photo URL is required,maxstringlength(255)~Field Photo URL maximum length is 255 characters"`
	UserID   uint   `json:"user_id" valid:"required~Field User ID is required"`
}

type UpdatePhoto struct {
	ID       uint   `json:"id" valid:"required~Field ID is required"`
	Title    string `json:"title" valid:"required~Field Title is required,maxstringlength(255)~Field Title maximum length is 255 characters"`
	Caption  string `json:"caption" valid:"required~Field Caption is required"`
	PhotoURL string `json:"photo_url" valid:"required~Field Photo URL is required,maxstringlength(255)~Field Photo URL maximum length is 255 characters"`
	UserID   uint   `json:"user_id" valid:"required~Field User ID is required"`
}
