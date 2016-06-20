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

Dependencies

* ```github.com/fatih/color```
* ```gopkg.in/elazarl/goproxy.v1```
* ```github.com/kr/pretty```
* ```gopkg.in/xmlpath.v1```
