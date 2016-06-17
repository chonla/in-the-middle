package cacher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
    "errors"

	logger "inthemiddle/logger"
	matcher "inthemiddle/matcher"
)

var (
	toFolder string
	cache    = Cache{}
)

func SetExportFolder(path string) {
	toFolder = path
}

func Find(req *http.Request) (*http.Response, error) {
    for k, v := range cache {
        if v.Match(req) {
            return cache[k].GetResponse(), nil
        }
    }
    return nil, errors.New("Cache missed.")
}

func Load(filename string) error {
    data, err := ioutil.ReadFile(toFolder + "/" + filename)
    if (err != nil) {
        return err
    }
    err = json.Unmarshal(data, &cache)
    if (err != nil) {
        logger.Debug("Unable to load " + filename)
        logger.Debug(err)
        cache = Cache{}
        return err
    }
    for k, _ := range cache {
		cache[k].currentState = "__default"
	}

    logger.Debug(toFolder + "/" + filename + " has been loaded.")

	matcher.Initialize()

    return nil
}

func Store(req *http.Request, resp *http.Response) {
	url := req.URL.String()
	respBody, _ := httputil.DumpResponse(resp, true)
	cache = append(cache, CacheItem{
		Key: url,
		Matcher: RequestMatcher{
			Pattern: url,
			Type:    "plain",
		},
		//req:  req,
		//resp: resp,
        currentState: "__default",
		ResponseStates: map[string]Responsible{
			"__default": Responsible{
				Return: string(respBody),
			},
		},
	})
}

func Flush() {
	for k, v := range cache {
		safeName := createSafeName(v.Key)
		body := v.ResponseStates["__default"].Return
		dumpCache(safeName, body)
		cache[k].Key = safeName
		cache[k].ResponseStates["__default"] = Responsible{
			Return: "file://" + toFolder + "/" + safeName,
		}
	}
	jsonString, _ := json.MarshalIndent(cache, "", "    ")
	err := ioutil.WriteFile(toFolder+"/stub.json", jsonString, 0755)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
