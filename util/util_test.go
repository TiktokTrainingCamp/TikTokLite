package util

import (
	"fmt"
	"testing"
)

func TestUtil(t *testing.T) {
	//iv := []byte("12345678qwertyui")
	//key := []byte("12345678abcdefgh09876543alnkdjfh")
	//plainText := []byte("Hi,I'm lady_killer9afzegvrzevgCDSef")
	//cipherText, _ := AesEncrypt(plainText, key)
	//fmt.Printf("加密后:%s\n", cipherText)
	//decryptText, _ := AesDecrypt(cipherText, key)
	//fmt.Printf("解密后:%s\n", decryptText)
	//
	//// 测试iv大小错误
	//iv = []byte("lady_killer9")
	//cipherText, err := AesEncrypt(plainText, iv, key)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(0)
	//}
	//// 测试key大小错误
	//iv = []byte("12345678qwertyui")
	//key = []byte("12345678abcdefgh09876543abc")
	//cipherText, err = AesEncrypt(plainText, iv, key)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(0)
	//}
	var key = "ThisIsPrivateKey"
	token, _ := GenerateToken(1, key)
	fmt.Println(token)
	plainText, success := ParseToken(token, key)
	fmt.Println(plainText, success)
}
