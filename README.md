# In The Middle

Stateful mock server

## To record activities

Start In The Middle

```sh
go run main/main.go -record
```

Set browser's proxy to In The Middle's listening IP and port and start the activities.

## To replay activities

```sh
go run main/main.go
```

Replaying recorded activities, In-The-Middle will load activities from stub.json from ```export folder```.

## Command line options

* ```-record``` to start record mode. Default is ```false```.
* ```-ip <ip>``` to specify listening IP address. Default is ```0.0.0.0```.
* ```-port <ip>``` to specify listening port. Default is ```8080```.
* ```-export <path>``` to specify exporting directory. Default is ```./fixtures```.
* ```-?``` to show options screen.

## Dependencies

Use ```go get ./...``` to install all dependencies.

* ```github.com/fatih/color```
* ```gopkg.in/elazarl/goproxy.v1```
* ```github.com/kr/pretty```
* ```gopkg.in/xmlpath.v1```
* ```github.com/jmoiron/jsonq```

## Examples

See ```examples``` directory for example.

## Get Started

1. Start in-the-middle in RECORD mode with default setting.
2. Use ```curl --proxy http://localhost:8080 http://anyHTTPdomain``` to record request and response. You should see request and response in in-the-middle console log.
3. Stop in-the-middle by Ctrl-C. File stub.json should be created in ./fixtures folder.
4. Start in-the-middle in REPLAY mode with default setting.
5. Use ```curl --proxy http://localhost:8080 http://anyHTTPdomain``` to get cached response. If request hits cache, you should see "Cache HIT" message in in-the-middle console. Otherwise, you should see "Cache MISSED" message.
6. You can modify cached content by editing files in ./fixtures folder.

## Github

https://github.com/chonla/in-the-middle
