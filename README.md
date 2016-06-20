# In The Middle

Stateful mock server

To record activities

```sh
go run main/main.go -record
```

To replay activities

```sh
go run main/main.go
```

Replaying recorded activities, In-The-Middle will load activities from stub.json from ```export folder```.

Command line options

* ```-record``` to start record mode. Default is ```false```.
* ```-ip <ip>``` to specify listening IP address. Default is ```0.0.0.0```.
* ```-port <ip>``` to specify listening port. Default is ```8080```.
* ```-export <path>``` to specify exporting directory. Default is ```./fixtures```.
* ```-?``` to show options screen.

Dependencies

* ```github.com/fatih/color```
* ```gopkg.in/elazarl/goproxy.v1```
* ```github.com/kr/pretty```
* ```gopkg.in/xmlpath.v1```
