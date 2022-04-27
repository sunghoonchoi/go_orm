package main

const DB_PATH = "orm.db"

type Member struct {
	ID       int
	UserName string
	Age      int
	TeamName string
	JoinDate string
}

const (
	TEST_MEMBER_ID_1        = 1
	TEST_MEMBER_USER_NAME_1 = "choisunghoon"
	TEST_MEMBER_AGE_1       = 43
	TEST_MEMBER_TEAM_NAME_1 = "blockchain"
	TEST_MEMBER_JOIN_DATE_1 = "2021.12.20"
)

const (
	TEST_MEMBER_ID_2        = 2
	TEST_MEMBER_USER_NAME_2 = "jack"
	TEST_MEMBER_AGE_2       = 27
	TEST_MEMBER_TEAM_NAME_2 = "blockchain"
	TEST_MEMBER_JOIN_DATE_2 = "2021.12.20"
)

const (
	TEST_MEMBER_ID_3        = 3
	TEST_MEMBER_USER_NAME_3 = "apdul"
	TEST_MEMBER_AGE_3       = 39
	TEST_MEMBER_TEAM_NAME_3 = "blockchain"
	TEST_MEMBER_JOIN_DATE_3 = "2021.12.20"
)

const (
	TEST_MEMBER_ID_4        = 4
	TEST_MEMBER_USER_NAME_4 = "jiyeon"
	TEST_MEMBER_AGE_4       = 25
	TEST_MEMBER_TEAM_NAME_4 = "backend"
	TEST_MEMBER_JOIN_DATE_4 = "2020.12.20"
)

const (
	TEST_MEMBER_ID_5        = 5
	TEST_MEMBER_USER_NAME_5 = "tom"
	TEST_MEMBER_AGE_5       = 35
	TEST_MEMBER_TEAM_NAME_5 = "backend"
	TEST_MEMBER_JOIN_DATE_5 = "2020.12.20"
)

type NFT struct {
	ID      int
	NftType string
	NftName string
	NftHash []byte
}

const (
	TEST_NFT_ID   = 1
	TEST_NFT_TYPE = "eos"
	TEST_NFT_NAME = "my first drawing nft"
	TEST_NFT_DATA = "myfirstnft myfirstnft myfirstnft myfirstnft myfirstnf tmyfirstnft myfirstnft"
)
