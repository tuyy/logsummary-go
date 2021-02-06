package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"logsummary-go/src/parse"
	"os"
)

const outputFileName = "result.txt"

func main() {
	buf := ReadFileAndGetNewBuffer()

	logDataList := GetLogDataList(buf)

	summary := MakeSummaryString(logDataList)

	WriteFileWithSummary(summary)
}

func WriteFileWithSummary(summary string) {
	f, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(f, summary)
}

func MakeSummaryString(logDataList []parse.LogData) string{
	summary := parse.MakeSummary(logDataList)

	contents, err := json.MarshalIndent(summary, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(contents)
}

func GetLogDataList(buf *bytes.Buffer) []parse.LogData {
	var logDataList []parse.LogData

	scan := bufio.NewScanner(buf)
	for scan.Scan() {
		line := scan.Text()
		logData, err := parse.ParseLog(line)
		if err != nil {
			log.Fatal(err, line)
		}
		logDataList = append(logDataList, *logData)
	}
	if err := scan.Err(); err != nil {
		panic(err)
	}

	return logDataList
}

func ReadFileAndGetNewBuffer() *bytes.Buffer {
	// Note: 파일로 읽어도되지만.. 임의로 text로 작성함
	b := []byte(`[2021-02-01 07:30:20.356] host=test02.nm ip=10.80.20.8 name=tuyy msg=hello
[2021-02-01 01:10:18.336] host=test01.nm ip=10.80.19.3 name=tuyy1 msg=hello
[2021-02-01 02:32:24.558] host=test03.nm ip=10.80.19.4 name=tuyy2 msg=hello
[2021-02-01 03:40:31.751] host=test01.nm ip=10.80.19.5 name=tuyy3 msg=hello
[2021-02-01 04:35:22.312] host=test04.nm ip=10.80.19.6 name=tuyy4 msg=hello
[2021-02-01 05:40:10.152] host=test03.nm ip=10.80.19.7 name=tuyy5 msg=hello`)

	return bytes.NewBuffer(b)
}
