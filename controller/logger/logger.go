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

			providersReponses := make(models.ProvidersBidResponse)
			err = decoder.Decode(&providersReponses)

			if LogProviderResponses(db, providersReponses) {
				fmt.Println("log inserted for auction result")
			}

		} else if actionString == "logAuctionParticipant" {

			auctionResult := make(models.AuctionResult)
			err = decoder.Decode(&auctionResult)

			LogAuctionParticipantList(db, auctionResult)

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

func LogAuctionParticipantList(db *sql.DB, participantList models.AuctionResult) bool {
	for _, slotValues := range participantList {
		for _, providers := range slotValues {
			for _, auctionResponse := range providers {
				if !InsertAuctionParticipantLog(db, auctionResponse) {
					return false
				}
			}
		}
	}
	return true
}

func InsertAuctionParticipantLog(db *sql.DB, auctionResponse models.AuctionResponse) bool {

	stmt, err := db.Prepare("INSERT INTO AuctionParticipantLog VALUES (?,?,?,?,?,?,?,?,?,?,?)")
	_, err = stmt.Exec(auctionResponse.Pubid, auctionResponse.AuctionID, auctionResponse.Auction_placementID, auctionResponse.BidPrice, auctionResponse.ID, auctionResponse.Ecc,
		auctionResponse.Epc, auctionResponse.Size, auctionResponse.RevShare, auctionResponse.SharedBid, auctionResponse.Status)

	if err != nil {
		panic(err)
		return false
	}
	return true
}

func LogProviderResponses(db *sql.DB, providersResponses models.ProvidersBidResponse) bool {
	for _, slotValues := range providersResponses {
		for _, adslotBidInf := range slotValues {
			if !InsertProviderResponseLog(db, adslotBidInf) {
				return false
			}
		}
	}
	return true
}

//InsertProviderResponseLog func
func InsertProviderResponseLog(db *sql.DB, response models.BidResponse) bool {
	stmt, err := db.Prepare("INSERT INTO ProviderResponseLog VALUES (?,?,?,?,?,?,?,?)")
	_, err = stmt.Exec(response.Pubid, response.BidPrice, response.ID, response.Ecc, response.Epc, response.Size, response.RevShare, response.Status)

	if err != nil {
		panic(err)
		return false
	}
	return true
}

/*
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
*/
