package orm

import (
	"crypto/sha256"
	"testing"
)

type NFT struct {
	ID      int
	NftType string
	NftName string
	NftHash []byte
}

type Member struct {
	Name    string
	Birth   string
	Address string
	Created string
	Updated string
}

// ORM test 1
//db를 정상적으로 생성, 테스트 테이블 생성
func Test_MakeCreateQuery(t *testing.T) {

	// step1 : data setting
	var myNFT NFT
	var testNFTData = "myfirstnftmyfirstnftmyfirstnftmyfirstnftmyfirstnftmyfirstnftmyfirstnft"
	myNFT.NftName = "my first drawing nft"
	myNFT.NftType = "eos"

	hash := sha256.Sum256([]byte(testNFTData))
	copy(myNFT.NftHash, hash[0:])

	_, err := RegisterInOrmEngine(myNFT, "test.db", 1)
	if err != nil {
		t.Error("creat table fail")
	}
}
