package main

import (
	"regexp"
	"strings"
)

const TRACKER_URL string = "https://tracker.yandex.ru/"

var taskIdsRegex, _ = regexp.Compile("(\\w+-\\d+)")

func findTasksIds(msg string) []string {
	tasksIds := taskIdsRegex.FindAllString(msg, -1)
	return tasksIds
}

func thereIsTasks(tasksIds []string) bool {
	return len(tasksIds) > 0
}

func getUrl(taskId string) string {
	return TRACKER_URL + taskId
}

func getTextForTasks(tasksIds []string) string {
	urls := Map(tasksIds, getUrl)
	return strings.Join(urls, "\n")
}
