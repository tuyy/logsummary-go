package parse

import (
	"errors"
	"time"
)

const logLayout = "2006-01-02 15:04:05.999"

type LogData struct {
	LoggedTime time.Time
	KeyValue   map[string]string
}

func NewLogData() *LogData {
	result := &LogData{}
	result.KeyValue = make(map[string]string)
	return result
}

func ParseLog(line string) (*LogData, error) {
	result := NewLogData()

	dtStr, nextIdx := getDateTimeStrAndNextIdx(line)
	if nextIdx == -1 {
		return nil, errors.New("invalid datetime in log")
	}

	result.LoggedTime = getLocalTime(dtStr)
	result.KeyValue = parseKeyValue(line, nextIdx)

	return result, nil
}

func parseKeyValue(line string, nextIdx int) map[string]string {
	result := make(map[string]string)
	key, val := "", ""
	for i := nextIdx; i < len(line); i++ {
		switch line[i] {
		case '=':
			key, val = val, ""
		case ' ':
			if key != "" {
				result[key] = val
				key, val = "", ""
			}
		default:
			val += string(line[i])
		}
	}
	if key != "" {
		result[key] = val
		key, val = "", ""
	}
	return result
}

func getLocalTime(dtStr string) time.Time {
	dt, err := time.ParseInLocation(logLayout, dtStr, time.Local)
	if err != nil {
		panic(err)
	}
	return dt
}

func getDateTimeStrAndNextIdx(line string) (string, int) {
	var stack []string

	result := ""
	idx := 0
	for _, c := range line {
		idx++

		switch {
		case c == '[':
			stack = append(stack, string(c))
		case c == ']':
			if stack[len(stack)-1] == "[" {
				return result, idx
			}
		default:
			result += string(c)
		}
	}
	return "", -1
}
