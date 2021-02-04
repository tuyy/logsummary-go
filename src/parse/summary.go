package parse

import (
	"time"
)

type Summary struct {
	StartTime time.Time                 `json:"start"`
	EndTime   time.Time                 `json:"end"`
	Total     int                       `json:"Total"`
	KeyValue  map[string]map[string]int `json:",omitempty"`
}

func MakeSummary(logDataList []LogData) Summary {
	if len(logDataList) == 0 {
		return Summary{}
	}

	result := Summary{
		StartTime: logDataList[0].LoggedTime,
		EndTime:   logDataList[0].LoggedTime,
		Total:     len(logDataList),
		KeyValue:  make(map[string]map[string]int),
	}

	for _, logData := range logDataList {
		if logData.LoggedTime.Before(result.StartTime) {
			result.StartTime = logData.LoggedTime
		}
		if logData.LoggedTime.After(result.EndTime) {
			result.EndTime = logData.LoggedTime
		}

		for key, val := range logData.KeyValue {
			if v, ok := result.KeyValue[key]; ok {
				v[val]++
			} else {
				result.KeyValue[key] = make(map[string]int)
				result.KeyValue[key][val] = 1
			}
		}
	}
	return result
}
