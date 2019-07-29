package etcdtest

import (
	"hank.com/etcd-3.3.12-annotated/etcdmain"
	"os"
	"testing"
)

func TestEtcdMain(t *testing.T){
	os.Args = []string{"etcd-3.3.12-test","--name","infraO",
		"--initial-advertise-peer-urls","http://127.0.0.1:2380",
		"--listen-peer-urls","http://127.0.0.1:2380",
		"--listen-client-urls","http://127.0.0.1:2379",
		"--advertise-client-urls","http://127.0.0.1:2379",
		"--initial-cluster-token","etcd-cluster-1",
		"--initial-cluster","infraO=http://127.0.0.1:2380,infral=http://127.0.0.1:2382,infra2=http://127.0.0.1:2384",
		"--initial-cluster-state","new"}
	etcdmain.Main()//服务端主入口
}
