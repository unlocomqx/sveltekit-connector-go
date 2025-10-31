package users

func QueryUserInfo() (any, error) {
	return map[string]any{
		"id":       1,
		"username": "myuser",
	}, nil
}
