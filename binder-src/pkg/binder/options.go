package binder

import (
	"errors"
	_ "fmt"
	"github.com/zan8in/goflags"
	"github.com/zan8in/gologger"
)

// 1.定义参数结构体
type (
	Options struct {
		Payload string
		File    string
		Output  string
	}
)

// 2.解析参数
func ParseOptions() *Options {
	options := &Options{}
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`binder`)

	flagSet.CreateGroup("input", "Input",
		flagSet.StringVarP(&options.Payload, "payload", "p", "", "input a payload.exe"),
		flagSet.StringVarP(&options.File, "file", "f", "", "input a normal file"),
		flagSet.StringVarP(&options.Output, "output", "o", "result", "save output path"),
	)

	flagSet.Parse()
	// 解析参数是否有值
	err := options.validateOptions()
	if err != nil {
		gologger.Fatal().Msgf("Program exiting: %s\n", err)
	}
	return options
}

var (
	errNoFile1 = errors.New("no input payload.exe file")
	errNoFile2 = errors.New("no input normal file")
	errNoInput = errors.New("no input,please binder.exe -h")
)

// Options结构体实现的方法
func (options *Options) validateOptions() (err error) {
	if options.File == "" && options.Payload == "" {
		return errNoInput
	} else if options.Payload == "" {
		return errNoFile1
	} else if options.File == "" {
		return errNoFile2
	}
	return nil
}
