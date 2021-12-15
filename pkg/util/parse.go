package util

import "strings"

func ParseCsdnId(articleDetailUrl string) string {
	splitUrl := strings.Split(articleDetailUrl, "/")
	return splitUrl[len(splitUrl)-1]
}
