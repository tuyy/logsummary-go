package parse

import (
	"reflect"
	"testing"
	"time"
)

func TestSummaryHappy(t *testing.T) {
	input := []LogData{
		{time.Date(2021, time.February, 1, 3, 0, 0, 0, time.Local),
			map[string]string{"name": "tu", "age": "31", "result": "SUCCESS"}},
		{time.Date(2021, time.March, 8, 1, 0, 0, 0, time.Local),
			map[string]string{"name": "yy", "age": "31", "result": "SUCCESS"}},
		{time.Date(2021, time.February, 1, 2, 0, 0, 0, time.Local),
			map[string]string{"name": "tu", "age": "31", "result": "INVALID"}},
	}
	var want = Summary{
		startTime: time.Date(2021, time.February, 1, 2, 0, 0, 0, time.Local),
		endTime:   time.Date(2021, time.March, 8, 1, 0, 0, 0, time.Local),
		total:     3,
		KeyValue: map[string]map[string]int{
			"name": {
				"tu": 2,
				"yy": 1,
			},
			"age": {
				"31": 3,
			},
			"result": {
				"SUCCESS": 2,
				"INVALID": 1,
			},
		},
	}
	rz := MakeSummary(input)
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. rz:%#v want:%#v", rz, want)
	}

}
