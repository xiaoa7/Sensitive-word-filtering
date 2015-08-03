# Sensitive-word-filtering 
敏感词过滤
直接从Java版DFA算法实现，转移过来，并未作修改，词库直接从百度搜索的一个暂用。
可替换成其他词库，也可自行添加词条。


在项目init方法中调用ReadSwfDict加载词库

在检查点调用Match方验证，或调用Repl方法替换

具体使用见swf_test.go