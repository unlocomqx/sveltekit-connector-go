package routes

func queryUserInfo() (any, error) {
	return map[string]any{
		"id":       1,
		"username": "myuser",
	}, nil
}
