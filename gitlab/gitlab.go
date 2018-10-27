package gitlab

import (
	"regexp"
	"strings"

	"github.com/rndd/discord-bot/utils"
)

const MR_URL = "https://gitlab.com/smartup-school/learning-system-web/merge_requests/"

var mrRe, _ = regexp.Compile(utils.AROUND_ID_RE + "(?:\\!)(\\d+)" + utils.AROUND_ID_RE)

func FindMRIds(msg string) []string {
	var mrIds []string
	foundGroup := mrRe.FindAllStringSubmatch(msg, -1)

	for _, found := range foundGroup {
		mrIds = append(mrIds, found[1])
	}

	return mrIds
}

func getUrl(id string) string {
	return MR_URL + id
}

func GetTextForMr(mrIds []string) string {
	urls := utils.Map(mrIds, getUrl)
	return strings.Join(urls, "\n")
}
