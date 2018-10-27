package tracker

import (
	"regexp"
	"strings"

	"github.com/rndd/discord-bot/utils"
)

const TRACKER_URL string = "https://tracker.yandex.ru/"

var taskIdsRegex, _ = regexp.Compile(utils.AROUND_ID_RE + "([a-zA-Z]{2,}-\\d+)" + utils.AROUND_ID_RE)

func FindTasksIds(msg string) []string {
	var tasksIds []string
	foundGroup := taskIdsRegex.FindAllStringSubmatch(msg, -1)

	for _, found := range foundGroup {
		tasksIds = append(tasksIds, found[1])
	}

	return tasksIds
}

func ThereIsTasks(tasksIds []string) bool {
	return len(tasksIds) > 0
}

func getUrl(taskId string) string {
	return TRACKER_URL + taskId
}

func GetTextForTasks(tasksIds []string) string {
	urls := utils.Map(tasksIds, getUrl)
	return strings.Join(urls, "\n")
}
