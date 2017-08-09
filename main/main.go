package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"os"

	cacher "github.com/chonla/inthemiddle/cacher"
	httper "github.com/chonla/inthemiddle/httper"
	logger "github.com/chonla/inthemiddle/logger"
	proxy "github.com/chonla/inthemiddle/proxy"
	goproxy "github.com/elazarl/goproxy"
	"github.com/fatih/color"
	"gopkg.in/elazarl/goproxy.v1/transport"
)

var (
	inPFunc  = color.New(color.FgYellow).SprintFunc()
	outPFunc = color.New(color.FgMagenta).SprintFunc()
)

type Args struct {
	Ip           string
	Port         string
	ExportFolder string
	Help         bool
	Record       bool
}

func main() {
	welcome()
	args := testArgs()
	startInTheMiddle(args)
}

func welcome() {
	logger.Info("*****************")
	logger.Info("* IN THE MIDDLE *")
	logger.Info("*****************")
}

func testArgs() (args Args) {
	flag.StringVar(&args.Ip, "ip", "0.0.0.0", "Listening IP address")
	flag.StringVar(&args.Port, "port", "8080", "Listening port")
	flag.StringVar(&args.ExportFolder, "export", "./fixtures", "Exporting folder")
	flag.BoolVar(&args.Record, "record", false, "Record mode (Record all activities)")
	flag.BoolVar(&args.Help, "?", false, "Show usage")
	flag.Parse()

	if args.Help {
		logger.Info("Usage of In the middle:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	return
}

func startInTheMiddle(args Args) {
	proxy.WaitForExitSignal()
	tr := transport.Transport{Proxy: transport.ProxyFromEnvironment}
	proxy.Start(proxy.Options{
		Ip:           args.Ip,
		Port:         args.Port,
		ExportFolder: args.ExportFolder,
		Record:       args.Record,
		OnRequest: func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			ctx.RoundTripper = goproxy.RoundTripperFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (resp *http.Response, err error) {
				ctx.UserData, resp, err = tr.DetailedRoundTrip(req)
				return
			})

			reqBody, err := httputil.DumpRequest(req, true)
			if err != nil {
				logger.Error(err)
				os.Exit(1)
			}

			r := httper.NewRequest(string(reqBody))

			logger.Info(inPFunc("--> ") + r.ToString())

			if !args.Record {
				resp, err := cacher.Find(req)
				if err == nil {
					logger.Debug("Cache HIT")
					return req, resp
				}
				logger.Debug("Cache MISSED")
			}

			return req, nil
		},
		OnResponse: func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			respBody, err := httputil.DumpResponse(resp, true)
			if err != nil {
				logger.Error(err)
				os.Exit(1)
			}

			r := httper.NewResponse(string(respBody))

			logger.Info(outPFunc("<-- ") + r.ToString())
			return resp
		},
	})
}
