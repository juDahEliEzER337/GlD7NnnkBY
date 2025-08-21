// 代码生成时间: 2025-08-21 23:55:19
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "io"
    "log"
)

// PasswordTool 提供密码加密和解密的功能
type PasswordTool struct {
    key []byte
}

// NewPasswordTool 创建一个新的 PasswordTool 实例
func NewPasswordTool(key []byte) *PasswordTool {
    return &PasswordTool{key: key}
}

// Encrypt 加密密码
func (p *PasswordTool) Encrypt(plaintext string) (string, error) {
    // 确保密钥长度符合AES-256的要求
    if len(p.key) != 32 {
        return "", errors.New("密钥长度必须是32字节")
    }

    block, err := aes.NewCipher(p.key)
    if err != nil {
        return "", err
    }

    // 随机生成一个偏移量
    var nonce [12]byte
    if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
        return "", err
    }

    // 填充plaintext使其长度是块大小的倍数
    plaintext = PKCS7Padding(plaintext, aes.BlockSize)

    // 加密
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    copy(ciphertext[:12], nonce[:])
    block.Encrypt(ciphertext[12:], []byte(plaintext))

    // 将nonce和密文合并并进行base64编码
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 解密密码
func (p *PasswordTool) Decrypt(ciphertext string) (string, error) {
    // 将base64编码的字符串解码
    decoded, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }

    if len(decoded) < aes.BlockSize+12 {
        return "", errors.New("解码后的密文长度不足")
    }

    // 确保密钥长度符合AES-256的要求
    if len(p.key) != 32 {
        return "", errors.New("密钥长度必须是32字节")
    }

    block, err := aes.NewCipher(p.key)
    if err != nil {
        return "", err
    }

    // 从密文中提取nonce
    var nonce [12]byte
    copy(nonce[:], decoded[:12])

    // 解密
    ciphertext = decoded[12:]
    if len(ciphertext)%aes.BlockSize != 0 {
        return "", errors.New("密文长度不是块大小的整数倍")
    }
    plaintext := make([]byte, len(ciphertext))
    block.Decrypt(plaintext, ciphertext)

    // 移除填充
    plaintext = PKCS7UnPadding(plaintext)

    return string(plaintext), nil
}

// PKCS7Padding 填充函数
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

// PKCS7UnPadding 去填充函数
func PKCS7UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

func main() {
    // 示例用法
    key := []byte("this-is-a-very-secret-key-123")
    passwordTool := NewPasswordTool(key)

    plaintext := "my-secret-password"
    encrypted, err := passwordTool.Encrypt(plaintext)
    if err != nil {
        log.Fatalf("加密失败: %v", err)
    }
    log.Printf("加密后的密码: %s", encrypted)

    decrypted, err := passwordTool.Decrypt(encrypted)
    if err != nil {
        log.Fatalf("解密失败: %v", err)
    }
    log.Printf("解密后的密码: %s", decrypted)
}