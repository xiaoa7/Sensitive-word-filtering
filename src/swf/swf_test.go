package swf

import (
	"log"
	"testing"
)

func TestMatch(t *testing.T) {
	//加载敏感词库，仅需要初始化一次
	ReadSwfDict("./baidu_fsw.dict")
	//验证是否有敏感词
	ret := Match("奥运会举办的挺好的！")
	log.Println(ret)
	//替换敏感词为**号
	ret2 := Repl("奥运会举办的挺好的！")
	log.Println(ret2)
}
