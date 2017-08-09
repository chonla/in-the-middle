package proxy

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	goproxy "github.com/elazarl/goproxy"

	logger "github.com/chonla/inthemiddle/logger"

	cacher "github.com/chonla/inthemiddle/cacher"
)

type Options struct {
	Ip           string
	Port         string
	ExportFolder string
	Record       bool
	OnRequest    func(*http.Request, *goproxy.ProxyCtx) (*http.Request, *http.Response)
	OnResponse   func(*http.Response, *goproxy.ProxyCtx) *http.Response
}

var (
	proxy   *goproxy.ProxyHttpServer
	options Options
)

func Start(op Options) {
	options = op

	proxy = goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	proxy.OnRequest().DoFunc(onRequestHandler)
	proxy.OnResponse().DoFunc(onResponseHandler)

	cacher.SetExportFolder(options.ExportFolder)

	if !options.Record {
		cacher.Load("stub.json")
		logger.Info("In the middle has been started in REPLAY mode. Press ^C to terminate In the middle.")
	} else {
		logger.Info("In the middle has been started in RECORD mode. Press ^C to terminate In the middle.")
	}

	logger.Info("Current settings")
	logger.Info(options)

	http.ListenAndServe(options.Ip+":"+options.Port, proxy)
}

func onRequestHandler(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	if options.OnRequest != nil {
		return options.OnRequest(req, ctx)
	}
	return req, nil
}

func onResponseHandler(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	if options.Record {
		if resp.Header.Get("X-Cacher") != "In-The-Middle" {
			cacher.Store(ctx.Req, resp)
		}
	}
	if options.OnResponse != nil {
		return options.OnResponse(resp, ctx)
	}
	return resp
}

func WaitForExitSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c

		if !options.Record {
			logger.Info("In the middle has been terminated from REPLAY mode.")
		} else {
			logger.Info("Flush cache to file.")
			cacher.Flush()
			logger.Info("In the middle has been terminated from RECORD mode. See stub.json in export folder for exported session.")
		}
		os.Exit(1)
	}()
}
