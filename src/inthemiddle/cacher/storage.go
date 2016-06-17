package cacher

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

    logger "inthemiddle/logger"
)

var maxSlot = 1000

func createSafeName(url string) string {
	re, _ := regexp.Compile(`[^A-Za-z0-9_\-\.]`)
	safeName := re.ReplaceAllString(url, "_")

	re, _ = regexp.Compile(`_+$`)
	safeName = re.ReplaceAllString(safeName, "")

	re, _ = regexp.Compile(`^_+`)
	safeName = re.ReplaceAllString(safeName, "")

	safeName = safeName + ".txt"

	safeName = getSafeSlot(safeName)

	return safeName
}

func getSafeSlot(fileName string) string {
	slot := 0

	for slot < maxSlot {
		tmpName := fmt.Sprintf("%d_%s", slot, fileName)
		_, err := os.Stat(toFolder + "/" + tmpName)
		if err != nil {
			return tmpName
		}
		slot = slot + 1
	}
	return fileName
}

func dumpCache(fileName string, content string) {
	finfo, err := os.Stat(toFolder)
	if err != nil || !finfo.IsDir() {
		err = os.MkdirAll(toFolder, 0755)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
	}

	b := []byte(content)
	err = ioutil.WriteFile(toFolder+"/"+fileName, b, 0755)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
