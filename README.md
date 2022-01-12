```text
官方安装教程：
https://github.com/facebook/rocksdb/blob/main/INSTALL.md


mac默认安装路径：
/usr/local/Cellar/rocksdb/6.22.1


配置gorocksdb，path/to/rocksdb修改为本机rocksdb所有在路径：
CGO_CFLAGS="-I/path/to/rocksdb/include" \ CGO_LDFLAGS="-L/path/to/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \


CGO_CFLAGS="-I/usr/local/Cellar/rocksdb/6.22.1/include" \ 
CGO_LDFLAGS="-L/usr/local/Cellar/rocksdb/6.22.1 -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" 

https://zhuanlan.zhihu.com/p/129049724
https://learnku.com/docs/go-szgbf/1.0/principle-of-consistent-hash-algorithm/8813

```

##### 参考

1.go cgo https://pkg.go.dev/cmd/cgo

2.一致性哈希 https://www.stathat.com/c/consistent    https://www.jianshu.com/p/5198b869374a

3.rocksdb官网 http://rocksdb.org/

4.rocksdb开源库 https://github.com/tecbot/gorocksdb

5.CSDN https://blog.csdn.net/qq_40697071/article/details/103791892

6.gossip协议



