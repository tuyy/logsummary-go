package parse

import "time"

type Summary struct {
	startTime time.Time
	endTime   time.Time
	total     int
	KeyValue  map[string]map[string]int
}

func MakeSummary(logDataList []LogData) Summary {
	if len(logDataList) == 0 {
		return Summary{}
	}

	result := Summary{
		startTime: logDataList[0].LoggedTime,
		endTime:   logDataList[0].LoggedTime,
		total:     len(logDataList),
		KeyValue:  make(map[string]map[string]int),
	}

	for _, logData := range logDataList {
		if logData.LoggedTime.Before(result.startTime) {
			result.startTime = logData.LoggedTime
		}
		if logData.LoggedTime.After(result.endTime) {
			result.endTime = logData.LoggedTime
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
