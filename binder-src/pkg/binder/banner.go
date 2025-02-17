package binder

import (
	_ "fmt"
	"github.com/zan8in/gologger"
)

func ShowBanner(VERSION string, CONTENT string) {
	gologger.Print().Msgf("\n|||\tB I N D E R\t|||\t%s\n大白哥免杀圈子内部工具\tauthor:大白哥\n%s\n\n", VERSION, CONTENT)
}
