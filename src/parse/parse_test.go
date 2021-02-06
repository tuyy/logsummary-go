package parse

import (
	"reflect"
	"testing"
)

func TestParseLineHappy(t *testing.T) {
	input := "[2021-02-01 07:30:20.356] host=test01nmip=10.80.19.2 name=tuyy msg=hello"

	_, err := ParseLog(input)
	if err != nil {
		t.Fatalf("invalid line err:%#v", err)
	}
}

func TestParseInvalidLine(t *testing.T) {
	input := "[2021-02-01 07:30:20.356] host=test01.nm ip=10.80.19.2 name=tuyy msg=hello"
	want := NewLogData()
	want.KeyValue = map[string]string{
		"host": "test01.nm",
		"ip": "10.80.19.2",
		"name": "tuyy",
		"msg": "hello",
	}

	logData, err := ParseLog(input)
	if err != nil {
		t.Fatalf("failed to parse line. err:%s\n", err)
	}
	if !reflect.DeepEqual(logData.KeyValue, want.KeyValue) {
		t.Fatalf("invalid result. parsed:%#v want:%#v", logData, want)
	}
}
