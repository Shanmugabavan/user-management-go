package controller

func JsonError(message string) string {
	return `{"message": "` + message + `"}`
}
