package service

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rotisserie/eris"
	"github.com/valyala/fasthttp"
)

var EthereumService = ethereumService{}

type ethereumService struct{}

func (ethereumService) CreateAddress(fctx *fasthttp.RequestCtx) (string, error) {

	// 1. 生成私鑰
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", eris.Wrap(err, "生成私鑰失敗")
	}

	// 2. 從私鑰獲取公鑰
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", eris.Wrap(err, "轉換公鑰失敗")
	}

	// 3. 從公鑰獲取地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("私钥:", hex.EncodeToString(privateKey.D.Bytes()))
	return strings.ToLower(address.Hex()), nil
}
