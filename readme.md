# goption

Generator class fields as constructor options

## Install

	go get -v github.com/micln/goption

## Usage

```shell
$ goption

usage: goption --package=PACKAGE [<flags>]

Generator class fields as constructor options

Flags:
  -h, --help             Show context-sensitive help (also try --help-long and --help-man).
  -v, --verbose          show log
  -p, --package=PACKAGE  package full path
  -c, --classes=CLASSES  match classes name like "person,car".
  -w, --writefile        write generated code
```

## Demo

定义一个普通的 struct

```go
//go:generate goption -p . -c Person -w
type Person struct {
	Name      string
	Age       int
	Interests []string

	Parent   *Person
	Children [5]*Person

	Friends []*Person
}
```

- 注释 `//go:generate goption -p . -c Person -w` 可以帮助 IDE 进行自动生成

运行命令

    go:generate goption -p . -c Person -w

参数说明

- -p：要扫描的包路径
- -c：指定要生成的类名，选填，格式为 A,B,C。不填时为包内所有类生成代码
- -w：无此参数时，会把生成的代码输出到 stdout 自行后续处理；有此参数时会把生成的代码写到 `class_option_gen.go`

## todo

自定义 template 为 struct 生成各种模板代码。说白了就是把 Generator 给开放出来，我帮你解析好 struct，你直接写你要的代码格式。
