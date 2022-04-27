package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"orm"
)

func main() {

	//step 1 : define structure
	// type Member struct {
	// 	ID       int
	// 	UserName string
	// 	Age      int
	// 	TeamName string
	// 	JoinDate string
	// }
	var testMember1 Member
	testMember1.ID = TEST_MEMBER_ID_1
	testMember1.UserName = TEST_MEMBER_USER_NAME_1
	testMember1.Age = TEST_MEMBER_AGE_1
	testMember1.TeamName = TEST_MEMBER_TEAM_NAME_1
	testMember1.JoinDate = TEST_MEMBER_JOIN_DATE_1

	//step 2 : register it in orm engine
	// struct information is registered in orm engine
	fmt.Println("======================= create")
	ormEngine, err := orm.RegisterInOrmEngine(testMember1, DB_PATH, 1, 2) // Primary Keys are index 1 and 2 of the Member struct
	if err != nil {
		fmt.Println("regist in orm engine fail")
		log.Panic(err)
	}

	fmt.Println("======================= insert")
	err = orm.InsertToTable(ormEngine, &testMember1)
	if err != nil {
		fmt.Println("insert to table fail")
		log.Panic(err)
	}

	testMember2 := Member{TEST_MEMBER_ID_2, TEST_MEMBER_USER_NAME_2, TEST_MEMBER_AGE_2, TEST_MEMBER_TEAM_NAME_2, TEST_MEMBER_JOIN_DATE_2}
	testMember3 := Member{TEST_MEMBER_ID_3, TEST_MEMBER_USER_NAME_3, TEST_MEMBER_AGE_3, TEST_MEMBER_TEAM_NAME_3, TEST_MEMBER_JOIN_DATE_3}
	testMember4 := Member{TEST_MEMBER_ID_4, TEST_MEMBER_USER_NAME_4, TEST_MEMBER_AGE_4, TEST_MEMBER_TEAM_NAME_4, TEST_MEMBER_JOIN_DATE_4}
	testMember5 := Member{TEST_MEMBER_ID_5, TEST_MEMBER_USER_NAME_5, TEST_MEMBER_AGE_5, TEST_MEMBER_TEAM_NAME_5, TEST_MEMBER_JOIN_DATE_5}
	err = orm.InsertToTable(ormEngine, &testMember2)
	if err != nil {
		fmt.Println("insert to table fail")
		log.Panic(err)
	}
	err = orm.InsertToTable(ormEngine, &testMember3)
	if err != nil {
		fmt.Println("insert to table fail")
		log.Panic(err)
	}
	err = orm.InsertToTable(ormEngine, &testMember4)
	if err != nil {
		fmt.Println("insert to table fail")
		log.Panic(err)
	}
	err = orm.InsertToTable(ormEngine, &testMember5)
	if err != nil {
		fmt.Println("insert to table fail")
		log.Panic(err)
	}

	//step 3 : select from table
	// set first condition to select
	fmt.Println("======================= select")
	var selectCondition orm.SearchConditon
	selectCondition.ColumnName = orm.GetColumnNameAtIndex(ormEngine, 3) //team_name
	selectCondition.Value = "backend"
	selectCondition.Operator = orm.EQUAL
	// in the first condition LogicOp isn't needed

	// set second condition to select
	// in the second condition LogicOp (AND, OR) is needed
	var selectCondition2 orm.SearchConditon
	selectCondition2.LogicOp = orm.OR
	selectCondition2.ColumnName = orm.GetColumnNameAtIndex(ormEngine, 1) //user_name
	selectCondition2.Value = "choisunghoon"
	selectCondition2.Operator = orm.EQUAL

	// you can add addition condition to select
	// do select
	searchResults, err := orm.SearchRows(ormEngine, selectCondition, selectCondition2)
	if err != nil {
		log.Panic("search error")
	}
	// these are select results shape of array of map
	for i, v := range searchResults {
		fmt.Printf(" [SEARCH RESULT ] %d row search result =====\n", i)
		fmt.Println(v)
	}

	//step 4 : update row
	// set first condition to update
	fmt.Println("======================= update")
	var updateCondition orm.SearchConditon
	updateCondition.ColumnName = orm.GetColumnNameAtIndex(ormEngine, 3) //team_name
	updateCondition.Value = "blockchain"
	updateCondition.Operator = orm.EQUAL
	updateCondition.LogicOp = orm.OR

	// set target columns to update
	targetArray := make([]orm.TargetColumn, 0)
	var targetColumn orm.TargetColumn
	targetColumn.ColumnName = orm.GetColumnNameAtIndex(ormEngine, 3) //team_name
	targetColumn.Value = "backend"
	targetArray = append(targetArray, targetColumn)

	// execute update
	err = orm.UpdateRows(ormEngine, targetArray, updateCondition)
	if err != nil {
		fmt.Println("update rows fail")
		log.Panic(err)
	}

	//step 5 : delete condition
	// set first condition to delete row
	fmt.Println("======================= delete")
	var deleteCondition orm.SearchConditon
	deleteCondition.ColumnName = orm.GetColumnNameAtIndex(ormEngine, 3)
	deleteCondition.Value = "backend"
	deleteCondition.Operator = orm.EQUAL

	err = orm.DeleteRows(ormEngine, deleteCondition)
	if err != nil {
		fmt.Println("delete rows fail")
		log.Panic(err)
	}

	//step 6 : delete all
	fmt.Println("======================= delete all")
	err = orm.DeleteAll(ormEngine)
	if err != nil {
		fmt.Println("delete rows fail")
		log.Panic(err)
	}

	fmt.Println("======================= MEMBER TEST END")

	//step 7 : additional
	// struct NFT contains blob inside
	// check  orm engine treats this kind of data structure well
	// type NFT struct {
	// 	ID      int
	// 	NftType string
	// 	NftName string
	// 	NftHash []byte
	// }
	var myNFT NFT
	var testNFTData = TEST_NFT_DATA
	myNFT.ID = TEST_NFT_ID
	myNFT.NftName = TEST_NFT_NAME
	myNFT.NftType = TEST_NFT_TYPE
	hash := sha256.Sum256([]byte(testNFTData))
	copy(myNFT.NftHash, hash[0:])

	ormEngine2, err2 := orm.RegisterInOrmEngine(myNFT, DB_PATH, 1, 2)
	if err != nil {
		fmt.Println("regist in orm engine fail")
		log.Panic(err2)
	}

	fmt.Println("======================= additional")
	err = orm.InsertToTable(ormEngine2, &myNFT)
	if err != nil {
		fmt.Println("insert to table fail")
		log.Panic(err)
	}
}
