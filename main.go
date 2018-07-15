package main

import (
	//"gitee.com/johng/gkvdb/gkvdb"
	//"log"
	//"time"
	//"fmt"
	"blockchain/encrypt/ed25519"
)


func main(){
	//db,err := gkvdb.New("./db")
	//defer db.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//t1:= time.Now()
	////for i:= 0;i<128000;i++{
	////	var val [1024*8]byte
	////	rand.Read(val[:])
	////	db.Set([]byte("cloud"+strconv.Itoa(i)),val[:])
	////}
	//v := db.Get([]byte("cloud2340"))
	//fmt.Println(len(v))
	//t2 := time.Now().Sub(t1)
	//fmt.Println(t2.Seconds())
	ed25519.NewKeysBymnenonic("throw bamboo tell magic can boat book speak shoot powder surprise clip")
}
