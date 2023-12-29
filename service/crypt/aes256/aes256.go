package aes256

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
)

// Encrypt は指定された文字列を32バイトのキー (AES-256)を使って暗号化します。
func Encrypt(source string, key string) (result string, err error) {
	// AES暗号化
	encrypted, err := encryptAES([]byte(key), []byte(source))
	if err != nil {
		return "", err
	}

	// URL非対応文字列への変換
	base64Encoded := base64.StdEncoding.EncodeToString(encrypted)
	urlEncoded := url.QueryEscape(base64Encoded)

	return urlEncoded, nil
}

func Decrypt(source string, key string) (result string, err error) {

	// URL非対応文字列の変換処理
	urlDecoded, err := url.QueryUnescape(source)
	if err != nil {
		return "", err
	}

	// 暗号化文字列の取得
	base64Decoded, err := base64.StdEncoding.DecodeString(urlDecoded)
	if err != nil {
		return "", err
	}

	// 復号化
	decrypted, err := decryptAES([]byte(key), base64Decoded)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

// PKCS#7でパディングを行う関数
func pkcs7Pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS#7でパディングを削除する関数
func pkcs7Unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, fmt.Errorf("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

// AES暗号化を行う関数
func encryptAES(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext = pkcs7Pad(plaintext, aes.BlockSize)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// IVを暗号文の先頭に追加
	return append(iv, ciphertext[aes.BlockSize:]...), nil
}

// AES復号化を行う関数
func decryptAES(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return pkcs7Unpad(ciphertext)
}
