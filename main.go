package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.RedirectHandler(
		os.Getenv("REDIRECT_URL"),
		parseRedirectStatus(os.Getenv("HTTP_STATUS"), 302)))

	timeout := time.Second

	svr := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		IdleTimeout:       timeout,
		ReadHeaderTimeout: timeout,
		ReadTimeout:       timeout,
		WriteTimeout:      timeout,
		MaxHeaderBytes:    1 << 20,
	}

	err := svr.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
