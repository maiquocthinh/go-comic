package utils

import (
	"fmt"
	"regexp"
)

func DropboxShareLinkToDirectLink(sharedLink string) string {
	pattern := regexp.MustCompile(`\/scl\/fi\/([^\/]+\/[^\/]+)\?([^&]+)`)

	matches := pattern.FindStringSubmatch(sharedLink)
	if len(matches) != 3 {
		return "Invalid Dropbox shared link"
	}

	filePath := matches[1]
	params := matches[2]

	re := regexp.MustCompile(`&dl=0`)
	params = re.ReplaceAllString(params, "")

	directLink := fmt.Sprintf("https://dl.dropboxusercontent.com/s/%s?%s", filePath, params)
	return directLink
}
