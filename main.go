package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	version     = "0.1.0"
	defaultPort = ":9337"
)

var buf bytes.Buffer
var logger = log.New(&buf, "INFO:", log.Lshortfile)
var logInfo = func(info string) {
	logger.Output(2, info)
}

// Health represents a healthcheck status and ranged durations
type Health struct {
	max    string
	min    string
	status bool
}

func (h *Health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/health endpoint called")

	var n = 0

	imax, _ := strconv.Atoi(h.max)
	imin, _ := strconv.Atoi(h.min)

	if imax > imin && imin >= 0 {
		n = rand.Intn(imax-imin+1) + imin
		h.status = true
	}

	fmt.Printf("Sleeping %d seconds...\n", n)
	time.Sleep(time.Duration(n) * time.Second)

	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]interface{}{
		"healthy":    h.status,
		"remoteAddr": getRemoteAddr(r),
		"healthMax:": imax,
		"healthMin:": imin,
		"n":          n,
	})
	w.Write(resp)
}

func getRemoteAddr(r *http.Request) string {
	xforwarded := r.Header.Get("X-FORWARDED-FOR")
	if xforwarded != "" {
		return xforwarded
	}
	return r.RemoteAddr
}

func main() {
	rand.Seed(time.Now().UnixNano())
	max := os.Getenv("MAXTIME")
	min := os.Getenv("MINTIME")

	logInfo(fmt.Sprintf("MAXTIME: %s MINTIME: %s", max, min))
	fmt.Print(&buf)

	http.Handle("/health", &Health{max, min, false})
	log.Fatal(http.ListenAndServe(defaultPort, nil))

	logInfo("Done")
	fmt.Print(&buf)
}
