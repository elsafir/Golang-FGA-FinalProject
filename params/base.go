package params

type Response struct {
	Status              int         `json:"status"`
	Message             string      `json:"message,omitempty"`
	Error               string      `json:"error,omitempty"`
	AdditionalInfo      interface{} `json:"additional_info,omitempty"`
	Data                interface{} `json:"data,omitempty"`
}

type UserResponse struct {
	ID        int         `json:"id,omitempty"`
	Email     string      `json:"email,omitempty"`
	Username  string      `json:"username,omitempty"`
	Age       int         `json:"age,omitempty"`
	Token     string      `json:"token,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
	UpdatedAt interface{} `json:"updated_at,omitempty"`
}

type SocialMediaResponse struct {
	ID             int         `json:"id,omitempty"`
	Name           string      `json:"name,omitempty"`
	SocialMediaURL string      `json:"social_media_url,omitempty"`
	UserID         int         `json:"user_id,omitempty"`
	CreatedAt      interface{} `json:"created_at,omitempty"`
	UpdatedAt      interface{} `json:"updated_at,omitempty"`
	User           interface{} `json:"User,omitempty"`
}

type PhotoResponse struct {
	ID        int         `json:"id,omitempty"`
	Title     string      `json:"title,omitempty"`
	Caption   string      `json:"caption,omitempty"`
	PhotoURL  string      `json:"photo_url,omitempty"`
	UserID    int         `json:"user_id,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	User      interface{} `json:"User,omitempty"`
}

type CommentResponse struct {
	ID        int         `json:"id,omitempty"`
	Message   string      `json:"message,omitempty"`
	PhotoID   int         `json:"photo_id,omitempty"`
	UserID    int         `json:"user_id,omitempty"`
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
	User      interface{} `json:"User,omitempty"`
	Photo     interface{} `json:"Photo,omitempty"`
}
