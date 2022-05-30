package assignments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type currentTime struct {
	CurrentTime string `json:"current_time"`
}

func GetTime(writer http.ResponseWriter, request *http.Request) {

	returnMap := make(map[string]string)
	timezones := request.URL.Query().Get("tz")

	if timezones == "" {
		loc, err := time.LoadLocation("UTC")
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprint(writer, "Invalid Timezone")
			return
		} else {
			nowTime := currentTime{
				time.Now().In(loc).String(),
			}
			returnMap["current_time"] = nowTime.CurrentTime
		}
	} else {

		splitTzs := strings.Split(timezones, ",")

		for index := 0; index < len(splitTzs); index++ {

			if splitTzs[index] != "" {
				loc, err := time.LoadLocation(splitTzs[index])
				if err != nil {
					writer.WriteHeader(http.StatusNotFound)
					fmt.Fprint(writer, "Invalid Timezone")
					return
				} else {
					nowTime := currentTime{
						time.Now().In(loc).String(),
					}
					returnMap[splitTzs[index]] = nowTime.CurrentTime
				}
			}
		}
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(returnMap)
}
