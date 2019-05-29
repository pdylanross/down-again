package main

import (
	"net/http"
	"time"

	"github.com/pdylanross/down-again/pkg/daLog"
)

func main() {
	l, err := daLog.CreateLogger(&daLog.LoggerOptions{})

	if err != nil {
		panic(err.Error())
	}

	for true {
		pingSite("https://google.com", l)
		pingSite("https://microsoft.com", l)
		pingSite("https://amazon.com", l)
		pingSite("https://charter.com", l)

		time.Sleep(5 * time.Second)
	}
}

func pingSite(site string, l daLog.Logger) {
	start := time.Now()
	_, err := http.Get(site)
	elapsed := time.Since(start)

	if err != nil {
		l.LogFail(site)
	} else {
		l.LogSuccess(site, elapsed)
	}
}
