package util

import (
	"github.com/zan8in/gologger"
	"os"
	"os/exec"
	"path/filepath"
)

func BuildLoaderFile(goFileName, exeFileName string) error {
	fileName := filepath.Base(exeFileName)
	cmd := exec.Command("go", "build", "-o", fileName, "-ldflags", "-w -s -H windowsgui", "-trimpath")
	// Cmd结构体的Dir字段来设置命令的工作目录
	workingDir := filepath.Dir(goFileName)
	cmd.Dir = workingDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 执行命令
	if err := cmd.Run(); err != nil {
		gologger.Error().Msgf("捆绑文件编译失败: %s\n", err)
		return err
	}
	gologger.Info().Msgf("捆绑文件编译成功: %s\n", exeFileName)
	return nil
}
