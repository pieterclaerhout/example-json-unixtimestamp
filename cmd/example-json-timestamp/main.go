package main

import (
	"encoding/json"
	"time"

	"github.com/pieterclaerhout/go-log"
)

const jsonResponse = `{
    "status": "ok",
    "last_check": 1572428388
}`

type Response struct {
	Status    string `json:"status"`
	LastCheck int64  `json:"last_check"`
}

type ImprovedResponse struct {
	Status    string `json:"status"`
	LastCheck *Time  `json:"last_check"`
}

func main() {

	var r Response
	if err := json.Unmarshal([]byte(jsonResponse), &r); err != nil {
		log.Error(err)
		return
	}

	log.InfoDump(r, "r")

	timestamp := time.Unix(r.LastCheck, 0)
	log.InfoDump(timestamp.String(), "timestamp")

	var rImproved ImprovedResponse
	if err := json.Unmarshal([]byte(jsonResponse), &rImproved); err != nil {
		log.Error(err)
		return
	}

	log.InfoDump(rImproved, "rImproved")
	log.InfoDump(rImproved.LastCheck.String(), "rImproved.LastCheck")

	jsonBytes, err := json.MarshalIndent(rImproved, "", "  ")
	if err != nil {
		log.Error(err)
		return
	}

	log.Info(string(jsonBytes))

}
