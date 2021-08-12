package utils

func UrlIn(url string) (res bool) {
	taskUrl := []string{
		"/clusters/tasks/check", "/dns/tasks/check", "/messages/badge",
	}
	for _, v := range taskUrl {
		if v == url {
			return true
		}
	}
	return false
}
