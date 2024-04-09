package pkg

func BuildResponse_(key string, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"key":     key,
		"message": message,
		"data":    data,
	}
}
