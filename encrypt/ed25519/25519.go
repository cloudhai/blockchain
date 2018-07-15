package ed25519

import (
	"github.com/tyler-smith/go-bip39"
	"github.com/golang/x/crypto/ed25519"
	"bytes"
	"fmt"
	"encoding/hex"
	"github.com/golang/x/crypto/ripemd160"
)

type EdKeys struct{
	Prikey ed25519.PrivateKey
	Pubkey ed25519.PublicKey
	Mnemonic string
	Password string
	Address string
}

func NewKeys(passwd string) *EdKeys{
	encrypt,_ :=bip39.NewEntropy(128);
	mnemonic,_ :=bip39.NewMnemonic(encrypt)
	seed := bip39.NewSeed(mnemonic,passwd)
	pub,pri,_ := ed25519.GenerateKey(bytes.NewReader(seed))
	return &EdKeys{Prikey:pri,Pubkey:pub,Password:passwd,Mnemonic:mnemonic}

}

func NewKeysBymnenonic(mnemonic,passwd string) *EdKeys{
	seed := bip39.NewSeed(mnemonic,passwd)
	fmt.Println(mnemonic)
	pub,pri,_ := ed25519.GenerateKey(bytes.NewReader(seed))
	fmt.Println(hex.EncodeToString(pub))
	return &EdKeys{Prikey:pri,Pubkey:pub,Password:passwd,Mnemonic:mnemonic}
}

func (self *EdKeys)GetAddress() string{
	if self.Address == ""{
		haser := ripemd160.New()
		haser.Write(self.Pubkey)
		bs := haser.Sum(nil)
		self.Address = hex.EncodeToString(bs)
	}
	return self.Address
}

func (self *EdKeys)Sign(data []byte)[]byte{
	return ed25519.Sign(self.Prikey,data)
}

func Verify(data,sig,pub []byte)bool{
	return ed25519.Verify(pub,data,sig)
}

