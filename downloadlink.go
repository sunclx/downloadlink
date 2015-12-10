package downloadlink

import (
	"encoding/base64"
	"net/url"
)

func decode64(s string) (string, err) {
	data, err := base64.StdEncoding.DecodeString(s)
	return string(data), err
}
func encode64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func escape(s string) string {
	return url.QueryEscape(s)
}
func unescape(s string) (string,err) {
	return result, err := url.QueryUnescape(s)
}

//ThunderDecode convert thunder url to true url
func ThunderDecode(url string) (string, error) {
	url = url[len("thunder://"):]

	data, err := base64.StdEncoding.DecodeString(url)
	if err != nil {
		return "", err
	}
	url = string(data)
	
	url, err = unescape(url)
	if err != nil {
		return "", err
	}
	return url[2 : len(url)-2], nil
}

//ThunderEncode convert true url to thunder url
func ThunderEncode(url string) string {
	return "thunder://" + encode64("AA"+escape(url)+"ZZ")
}
