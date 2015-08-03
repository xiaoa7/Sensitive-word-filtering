package swf

import (
	"strings"
)

const (
	REPLTXT = "*************************************"
)

//过滤,验证是否存在敏感词
func Match(txt string) bool {
	for i := 0; i < len(txt); i++ {
		nowMap := &link
		length := 0   // 匹配标识数默认为0
		flag := false // 敏感词结束标识位：用于敏感词只有1位的情况
		for j := i; j < len(txt); j++ {
			word := txt[j : j+1]
			nowMap, _ = (*nowMap)[word].(*map[string]interface{})
			if nowMap != nil {
				length = length + 1
				tag, _ := (*nowMap)["YN"].(string)
				if "Y" == tag {
					flag = true
					return true
				}
			} else {
				break
			}
		}

		if length < 2 || !flag {
			length = 0
		}
		if length > 0 {
			//log.Println(txt[i : i+length])
			i = i + length - 1
		}
	}
	return false
}

//替换
func Repl(txt string) string {
	keywords := []string{}
	for i := 0; i < len(txt); i++ {
		nowMap := &link
		length := 0   // 匹配标识数默认为0
		flag := false // 敏感词结束标识位：用于敏感词只有1位的情况
		for j := i; j < len(txt); j++ {
			word := txt[j : j+1]
			nowMap, _ = (*nowMap)[word].(*map[string]interface{})
			if nowMap != nil {
				length = length + 1
				tag, _ := (*nowMap)["YN"].(string)
				if "Y" == tag {
					flag = true
				}
			} else {
				break
			}
		}

		if length < 2 || !flag {
			length = 0
		}
		if length > 0 {
			keywords = append(keywords, txt[i:i+length])
			i = i + length - 1
		}
	}
	for _, keyword := range keywords {
		txt = strings.Replace(txt, keyword, REPLTXT[:len(keyword)], -1)
	}
	return txt
}
