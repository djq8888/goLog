package goLog

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func removeOldestFile() {
	files, _ := ioutil.ReadDir("log")
	var infos []string
	for _, file := range files {
		if strings.Contains(file.Name(), "INFO") {
			infos = append(infos, file.Name())
		}
	}
	if len(infos) > maxFile {
		oldestFile := getOldestFile(infos)
		removeFile(oldestFile)
	}
}

func getOldestFile(files []string) string {
	var oldestFile string
	var oldestTime int64
	for _, file := range files {
		if strings.Contains(file, "INFO") {
			createTime := getCreateTime(file)
			if oldestTime == 0 {
				oldestTime = createTime
				oldestFile = file
			}
			if oldestTime > createTime {
				oldestTime = createTime
				oldestFile = file
			}
		}
	}
	return oldestFile
}

func getCreateTime(filename string) int64 {
	createTime, _ := time.Parse("2006-01-02_15_04_05", filename[4:23])
	timestamp := createTime.Unix()
	return timestamp
}

func removeFile(filename string) {
	os.Remove("log/"+filename)
}
