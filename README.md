# fasthttpunit

### fasthttp单元测试框架

## [示例](/example)

> 1.运行单元测试

```shell
./example/path/unit.sh

./example/file/unit.sh
```

> 2.用例结果及覆盖率

![case](https://github.com/jhq0113/fasthttpunit/blob/main/imgs/result.png?raw=true)
![cover](https://github.com/jhq0113/fasthttpunit/blob/main/imgs/result1.png?raw=true)
![cover-html](https://github.com/jhq0113/fasthttpunit/blob/main/imgs/cover.png?raw=true)

## 使用

> 1.配置用例列表，用例支持yml、yaml、json扩展名的文件，配置用例目录后，框架会自动扫描目录下是用例文件

* yaml文件示例

```yaml
desc: '等于'
path: '/equal'
caseList:
  - desc: '1'
    params: 'a=e&c=1'
    expected: 'Hello World'

  - desc: '2'
    params: 'a=1&b=c'
    expected: 'Ok'
```

* json文件示例

```json
{
  "desc": "包含",
  "path": "/contains",
  "method": "POST",
  "caseList": [
    {
      "desc": "1",
      "params": "num=1&sign=1",
      "expected": "Hello",
      "expectedType": "contains"
    },
    {
      "desc": "2",
      "params": "num=2&sign=2",
      "expected": "hello",
      "expectedType": "contains"
    }
  ]
}
```

> 2.fasthttp程序代码

```go
package main

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Equal(ctx *fasthttp.RequestCtx) {
	_, _ = ctx.WriteString(`Hello World`)
}

func Contains(ctx *fasthttp.RequestCtx) {
	_, _ = ctx.WriteString(`24sdfq23rwasdfasdfHelloadfasdf23sadfasdfef2`)
}

func Pattern(ctx *fasthttp.RequestCtx) {
	_, _ = ctx.WriteString(`{"code": 200, "msg": "ok", "data":{}}`)
}

func loadRouter() *fasthttprouter.Router {
	r := fasthttprouter.New()
	r.GET("/equal", Equal)
	r.POST("/contains", Contains)
	r.GET("/pattern", Pattern)

	return r
}

func main() {
	r := loadRouter()

	server := fasthttp.Server{
		Handler: r.Handler,
	}

	err := server.ListenAndServe(":8080")
	if err != nil {
		log.Fatalf("server start error:%s\n", err.Error())
	}
}
```

> 3.单元测试代码

```go
package main

import (
	"fmt"
	"testing"

	"github.com/jhq0113/fasthttpunit"
)

func mockA() {
	fmt.Printf("mock:a\n")
}

func mockB() {
	fmt.Printf("mock:b\n")
}

func TestUnit(t *testing.T) {
	r := loadRouter()

	casePath := fasthttpunit.BinPath() + "/case"

	conf, err := fasthttpunit.LoadConf(casePath)
	if err != nil {
		t.Fatal(fasthttpunit.Red("load conf err: %s", err.Error()))
	}

	conf.Delay = 3

	u := fasthttpunit.NewUnitWithRouter(conf, t, r)
	u.Test(mockA, mockB)
}
```

> 4.单元测试执行脚本

```shell
#!/bin/bash
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)

cd $SCRIPTPATH

go test . -v -coverpkg=... -coverprofile=$SCRIPTPATH/unitout/app.out
go tool cover -func=$SCRIPTPATH/unitout/app.out -o $SCRIPTPATH/unitout/coverage.txt
go tool cover -html=$SCRIPTPATH/unitout/app.out -o $SCRIPTPATH/unitout/coverage.html
```
