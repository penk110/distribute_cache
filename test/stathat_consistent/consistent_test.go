package stathat_consistent

import (
	"fmt"
	"testing"
)

func TestNewConsistent(t *testing.T) {

	_consistent := NewConsistent()

	_consistent.Add("c_1")
	_consistent.Add("c_2")

	var n = "c_3"
	elm, err := _consistent.Get(n)
	if err != nil {
		t.Errorf("get element failed, key: %s err: %s", n, err.Error())
		return
	}

	t.Logf("get element success, elm: %s", elm)

}

func TestNewConsistent2(t *testing.T) {
	cons := NewConsistent()
	cons.Add("cacheA")
	cons.Add("cacheB")
	cons.Add("cacheC")

	server1, err := cons.Get("user_1")
	server2, err := cons.Get("user_2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("server1:", server1) //输出 server1: cacheC
	fmt.Println("server2:", server2) //输出 server2: cacheA

	fmt.Println()

	//user_1在cacheA上，把cacheA删掉后看下效果
	cons.Remove("cacheA")
	server1, err = cons.Get("user_1")
	server2, err = cons.Get("user_2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("server1:", server1) //输出 server1: cacheC,和删除之前一样，在同一个server上
	fmt.Println("server2:", server2) //输出 server2: cacheB,换到另一个server了
}
