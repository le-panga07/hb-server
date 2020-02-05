package logger

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hb-server/models"
	"net/http"
)

//Log get config maps
func Log(db *sql.DB, actionString string) http.HandlerFunc {

	fmt.Println(actionString)

	fn := func(res http.ResponseWriter, req *http.Request) {

		fmt.Println("auctionResult got")

		decoder := json.NewDecoder(req.Body)

		var err interface{}

		if actionString == "logProviderResponse" {
			auctionResult := make(models.AuctionResult)
			err = decoder.Decode(&auctionResult)
			if LogAuctionResult(db, auctionResult) {
				fmt.Println("log inserted for auction result")
			}
		} else if actionString == "logAuctionParticipant" {
			var participantLog []models.ParticipantLog
			err = decoder.Decode(&participantLog)
			LogParticipantList(db, participantLog)
		}

		if err != nil {
			panic(err)
		}
		//fmt.Println(auctionResult)
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(&struct{ status string }{"OK"})

	}
	return http.HandlerFunc(fn)
}

func LogParticipantList(db *sql.DB, participantList []models.ParticipantLog) {
	for _, participant := range participantList {
		InsertParticipantLog(db, participant)
	}
}

func InsertParticipantLog(db *sql.DB, dbparticipant models.ParticipantLog) bool {
	stmt, err := db.Prepare("INSERT INTO AuctionParticipantLog VALUES (?,?,?)")
	_, err = stmt.Exec(dbparticipant.ProviderId, dbparticipant.ECC, dbparticipant.EPC)

	if err != nil {
		panic(err)
		return false
	}
	return true
}

func LogAuctionResult(db *sql.DB, auctionResult models.AuctionResult) bool {
	for _, slotValues := range auctionResult {
		for _, arrayBidRes := range slotValues {
			for index, bidRersult := range arrayBidRes {
				isWinner := false
				if index == 0 {
					isWinner = true
				}
				if !InsertAuctionLog(db, bidRersult, isWinner) {
					return false
				}
			}
		}
	}
	return true
}

//InsertLog func
func InsertAuctionLog(db *sql.DB, bidResult models.BidResult, isWinner bool) bool {

	stmt, err := db.Prepare("INSERT INTO ProviderResponseLog VALUES (?, ?,?,?,?,?,?)")
	_, err = stmt.Exec(bidResult.BidPrice, bidResult.ProviderID, bidResult.Adcode[0:10], bidResult.Ecc, bidResult.Epc, bidResult.Size, isWinner)

	if err != nil {
		panic(err)
		return false
	}
	return true
}
