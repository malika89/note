package golang

import (
	"encoding/json"
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
	"time"
)


type ts struct {
	C   string
	K   string
	T   interface{}
	Max int
	Cn  string
	Min int
	Tls bool
	Success bool
	Result interface{}
}


func ExampleJson(in *ts) {

	t1 := time.Now()

	var res interface{}
	for i := 0; i < 100000; i++ {
		// encode
		b, _ := json.Marshal(in)
		// decode
		var out = ts{}
		_ = json.Unmarshal(b, &out)
		if res ==nil{
			res = out
		}
	}
	t2 := time.Now()
	fmt.Println("Json 消耗时间：", t2.Sub(t1), "秒")
	fmt.Println(res)
}

func ExampleMsgpack(in *ts) {

	t1 := time.Now()
	var res interface{}
	for i := 0; i < 100000; i++ {
		// encode
		b, _ := msgpack.Marshal(in)
		// decode
		var out = ts{}
		_ = msgpack.Unmarshal(b, &out)
		if res ==nil{
			res = out
		}
	}
	t2 := time.Now()
	fmt.Println("msgpack 消耗时间：", t2.Sub(t1), "秒")
	fmt.Println(res)
}

func main(){
	var in = &ts{
		C:   "LOCK",
		K:   "31uEbMgunupShBVTewXjtqbBv5MndwfXhb",
		T:   1000,
		Max: 200,
		Cn:  "中文",
		Min: 10,
		Tls:true,
		Success:false,
		Result: map[string]interface{}{"total":1,"data":""},
	}

	ExampleMsgpack(in)
	ExampleJson(in)

}
