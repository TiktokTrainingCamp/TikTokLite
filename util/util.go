package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// GeneratePreview 生成预览封面
func GeneratePreview(videoName string) (string, error) {
	path := "public/"
	suffix := ".jpg"
	input := path + videoName
	pictureName := strings.Split(videoName, ".")[0] + suffix
	output := path + pictureName
	cmdArguments := []string{"-i", input, "-t", "1", "-r", "1", "-q:v", "2", "-f", "image2", output}

	cmd := exec.Command("ffmpeg", cmdArguments...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		//log.Fatal(err)
		return "", err
	}
	fmt.Printf("command output: %q", out.String())
	return pictureName, nil
}

//func AesEncrypt(plainText, iv, key []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	if len(iv) != block.BlockSize() {
//		return nil, err
//	}
//	// create a CTR interface
//	stream := cipher.NewCTR(block, iv)
//	cipherText := make([]byte, len(plainText))
//	// encrypt or decrypt
//	stream.XORKeyStream(cipherText, plainText)
//	return cipherText, nil
//}

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 移除
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	if unPadding < 0 || unPadding > length {
		return nil, errors.New("加密字符串错误！")
	}
	return data[:(length - unPadding)], nil
}

// aesEncrypt 加密
func aesEncrypt(data []byte, key []byte) ([]byte, error) {
	// NewCipher creates and returns a new cipher.Block. The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

// aesDecrypt 解密
func aesDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

// GenerateToken 生成token
func GenerateToken(userId int, keyString string) (string, bool) {
	timeStamp := time.Now().Unix()
	data := []byte(fmt.Sprint(userId) + "." + fmt.Sprint(timeStamp))
	key := []byte(keyString)
	crypt, err := aesEncrypt(data, key)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	token := base64.StdEncoding.EncodeToString(crypt)
	return token, true
}

// ParseToken 解析token
func ParseToken(token string, keyString string) (string, bool) {
	key := []byte(keyString)
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		fmt.Println(err)
		return err.Error(), false
	}
	plain, err := aesDecrypt(data, key)
	if err != nil {
		fmt.Println(err)
		return err.Error(), false
	}
	plainText := fmt.Sprintf("%s", plain)
	return plainText, strings.Contains(plainText, ".")
}
