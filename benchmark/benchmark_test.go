package benchmark

import (
	"fmt"
	"go.etcd.io/etcd/tools/benchmark/cmd"
	"os"
	"testing"
)


func TestBenchmark(t *testing.T){
	os.Args = []string{
		"--endpoints=http://127.0.0.1:2379", //目标机器
		"--target-leader",//目标服务器
		"--conns=1",//连接数
		"--clients=1",//链接客户端
		"put",
		"--key-size=8", //键的大小(以字节为单位)
		"--sequential-keys",
		"--total=10000",//键的总数
		"--val-size=256",//值得大小(以字节为单位)
	}
	//TODO HANK 这里
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
