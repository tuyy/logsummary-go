package parse

import (
	"encoding/json"
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
		StartTime: time.Date(2021, time.February, 1, 2, 0, 0, 0, time.Local),
		EndTime:   time.Date(2021, time.March, 8, 1, 0, 0, 0, time.Local),
		Total:     3,
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

func TestConvertToJson(t *testing.T) {
	summary := Summary{
		StartTime: time.Date(2009, time.January, 2, 0, 0, 0, 0, time.Local),
		EndTime:   time.Date(2021, time.February, 1, 0, 0, 0, 0, time.Local),
		Total:     1,
		KeyValue:  nil,
	}
	want := []byte(`{"start":"2009-01-02T00:00:00+09:00","end":"2021-02-01T00:00:00+09:00","Total":1}`)

	rz, err := json.Marshal(summary)
	if err != nil {
		t.Fatalf("failed to encode json. err:%#v want:%#v", err, want)
	}
	if !reflect.DeepEqual(rz, want) {
	    t.Fatalf("invalid result. rz:%#v want:%#v", rz, want)
	}
}
