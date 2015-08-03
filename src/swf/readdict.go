package swf

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//敏感词库
var link map[string]interface{}

//读取字段
func ReadSwfDict(config ...string) {
	var fi *os.File
	if len(config) == 0 {
		fi, _ = os.Open("./baidu_fsw.dict")
	} else {
		fi, _ = os.Open(config[0])
	}
	defer fi.Close()
	bs, _ := ioutil.ReadAll(fi)
	fswwords := strings.Split(string(bs), "\r\n")
	log.Println(len(fswwords))
	//要组装关键字
	link = make(map[string]interface{})
	var nowMap *map[string]interface{}
	for _, key := range fswwords {
		nowMap = &link
		for i := 0; i < len(key); i++ {
			kc := key[i : i+1]
			if v, ok := (*nowMap)[kc]; ok {
				nowMap, _ = v.(*map[string]interface{})
			} else {
				newMap := map[string]interface{}{}
				newMap["YN"] = "N"
				(*nowMap)[kc] = &newMap
				nowMap = &newMap
			}

			if i == len(key)-1 {
				(*nowMap)["YN"] = "Y"
			}

		}
	}
}
