package stringx

import url2 "net/url"

func PasswordObfuscatedURL(url string) string {
	parsed, err := url2.Parse(url)
	if err != nil {
		return url
	}
	if parsed.User != nil {
		parsed.User = url2.UserPassword(parsed.User.Username(), "xxxxxxx")
	}
	return parsed.String()
}
