package user

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Thumbnail   string `json:"thumbnail"`
}

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	User *User `json:"user"`
}

func NewGetUserResponse() *GetUserResponse {
	return &GetUserResponse{
		User: &User{
			ID:          "1",
			Name:        "test",
			DisplayName: "test",
			Thumbnail:   "https://example.com",
		},
	}
}

func NewCreateUserResponse() *GetUserResponse {
	return &GetUserResponse{
		User: &User{
			ID:          "1",
			Name:        "test",
			DisplayName: "test",
			Thumbnail:   "https://example.com",
		},
	}
}
