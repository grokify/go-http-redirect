package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func parseRedirectStatus(status string, defaultStatus int) int {
	status = strings.TrimSpace(status)
	if len(status) == 0 {
		return defaultStatus
	}
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		return defaultStatus
	}
	if statusInt == http.StatusMultipleChoices ||
		statusInt == http.StatusMovedPermanently ||
		statusInt == http.StatusFound ||
		statusInt == http.StatusSeeOther ||
		statusInt == http.StatusNotModified ||
		statusInt == http.StatusUseProxy ||
		statusInt == http.StatusTemporaryRedirect ||
		statusInt == http.StatusPermanentRedirect {
		return statusInt
	}
	return defaultStatus
}

func main() {
	http.Handle("/", http.RedirectHandler(
		os.Getenv("REDIRECT_URL"),
		parseRedirectStatus(os.Getenv("HTTP_STATUS"), 302)))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
