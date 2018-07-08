package encrypt

import (
	"math/big"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"encoding/hex"
	"bytes"
	"compress/gzip"
	"log"
)

type Keys struct{
	X []byte
	Y []byte
	D []byte
}

func NewKeys() *Keys{
	curve := elliptic.P256()
	pri,_ := ecdsa.GenerateKey(curve,rand.Reader)
	x := pri.PublicKey.X.Bytes()
	y := pri.PublicKey.Y.Bytes()
	d := pri.D.Bytes()
	return &Keys{x,y,d}
}

func (self *Keys)Pubkey() *ecdsa.PublicKey{
	pub := new(ecdsa.PublicKey)
	pub.Curve = elliptic.P256()
	var x *big.Int
	var y *big.Int
	x.SetBytes(self.X)
	y.SetBytes(self.Y)
	pub.X = x
	pub.Y = y
	return pub
}

func (self *Keys)Prikey() *ecdsa.PrivateKey{
	pri := new(ecdsa.PrivateKey)
	var d *big.Int
	d.UnmarshalText(self.D)
	pri.D = d
	pri.PublicKey = *self.Pubkey()
	return pri
}

func Sign(msg []byte,pri *ecdsa.PrivateKey)([]byte ,error){
	r,s,err := ecdsa.Sign(rand.Reader,pri,msg)
	if err != nil{
		log.Panic("sign err:",err)
	}
	rt,err := r.MarshalText()
	st,err := s.MarshalText()
	fmt.Printf("length of r :%d  lengthof s:%d\n",len(rt),len(st))
	fmt.Printf("rt:%s     st:%s\n",hex.EncodeToString(rt),hex.EncodeToString(st))
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	_,err = w.Write(append(append(rt,byte('+')),st...))
	w.Flush()
	return b.Bytes(),err
}

func Verify(msg,sign []byte,pubk *ecdsa.PublicKey)bool{
	gr, err := gzip.NewReader(bytes.NewBuffer(sign))
	defer gr.Close()
	if err != nil{
		log.Panic("gzip uncompre fail")
	}
	buf := make([]byte,1024)
	count, err := gr.Read(buf)

	rs := buf[:count]
	index := bytes.IndexByte(rs,'+')
	fmt.Printf("index of sing:%d\n",index)
	if index < 0{
		fmt.Println("index < 0")
		return false
	}
	fmt.Printf("r in verify:%s s:%s\n",hex.EncodeToString(rs[:index]),hex.EncodeToString(rs[index+1:]))
	var rint big.Int
	var sint big.Int
	rint.UnmarshalText(rs[:index])
	sint.UnmarshalText(rs[index+1:])
	result := ecdsa.Verify(pubk,msg,&rint,&sint)
	return result

	return false
}


func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}
