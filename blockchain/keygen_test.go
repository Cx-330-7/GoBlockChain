package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
)

// 生成一对基于椭圆曲线的公钥和私钥
// 返回值：*ecdsa.PrivateKey：生成的私钥；*ecdsa.PublicKey：对应的公钥，是从私钥中导出的。
func generateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	curve := elliptic.P256() // 选取了 secp256r1 圆锥曲线，实现了 elliptic.Curve 接口的曲线对象
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println("Error generating key pair:", err)
		return nil, nil
	}
	return privateKey, &privateKey.PublicKey
}

// 对文档进行签名
// privateKey *ecdsa.PrivateKey：用于签名的私钥。
// document string：需要签名的文档内容。
// 返回值：
// string：文档的 SHA-256 哈希值，以十六进制字符串表示。
// string：签名值，以十六进制字符串表示。
/**
. 为什么不能直接传递 hash
hash 是 [32]byte 类型的定长数组，而 ecdsa.Sign 需要的是 []byte 类型。
hash[:] 是一个切片操作，用于将定长数组 [32]byte 转换为切片 []byte。
[:] 是 Go 的切片语法，表示从数组中取出整个数据范围。
转换后，底层数据仍然是原数组，但类型变成了切片
*/
func signDocument(privateKey *ecdsa.PrivateKey, document string) (string, string) {
	// 计算文档的哈希值
	hash := sha256.Sum256([]byte(document))
	// 使用私钥对哈希值签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:]) //最后的签名内容就是 (r,s)
	if err != nil {
		fmt.Println("Error signing document:", err)
		return "", ""
	}

	// 使用 fmt.Sprintf 格式化字符串，将 r 和 s 的 16 进制形式拼接在一起，赋值给变量 signature
	signature := fmt.Sprintf("%s%s", r.Text(16), s.Text(16))
	return hex.EncodeToString(hash[:]), signature
}

func verifySignature(publicKey *ecdsa.PublicKey, hashedDoc, signature string) bool {
	hash, err := hex.DecodeString(hashedDoc)
	if err != nil {
		fmt.Println("Error decoding hashed document:", err)
		return false
	}
	// 将签名解码为 r 和 s
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		fmt.Println("Error decoding signature:", err)
		return false
	}
	r := new(big.Int).SetBytes(signatureBytes[:len(signatureBytes)/2])
	s := new(big.Int).SetBytes(signatureBytes[len(signatureBytes)/2:])
	// 验证签名
	return ecdsa.Verify(publicKey, hash, r, s)
}

func TestKeyGen(t *testing.T) {
	// 生成密钥对
	privateKey, publicKey := generateKeyPair()

	// 打印私钥和公钥
	fmt.Printf("Private Key: %s\n", hex.EncodeToString(privateKey.D.Bytes()))
	fmt.Printf("Public Key: %s\n", hex.EncodeToString(publicKey.X.Bytes()))

	// 待签名文档
	document := "zhuanzhang20yuan"

	// 签名文档
	hashedDoc, signature := signDocument(privateKey, document)
	fmt.Printf("Hashed Document: %s\n", hashedDoc)
	fmt.Printf("Signature: %s\n", signature)

	// 验证签名
	isValid := verifySignature(publicKey, hashedDoc, signature)
	fmt.Printf("Signature valid: %v\n", isValid)
}
