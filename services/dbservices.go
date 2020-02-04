package services

import (
	"database/sql"
	"fmt"
	"hb-server/models"
	"strings"

	_ "hb-server/github.com/go-sql-driver/mysql"
)

//GetProviderConfigs func
func GetProviderConfigs(db *sql.DB, publisherID string) *models.Config {

	publisher := GetPublisherInf(db, publisherID)

	adSlotInf := GetAdSlotInf(db, publisherID)
	fmt.Println("adSlotInf", len(adSlotInf))

	providers := GetAllProvidersDetails(db, publisher)
	fmt.Println("providers", len(providers))

	adslotProvidersInf := GetAllProvidersAdSlotInf(db, providers, publisherID)
	fmt.Println("adslotProvidersInf", len(adslotProvidersInf))

	config := GetConfigsData(publisher, adSlotInf, providers, adslotProvidersInf)

	return config
}

//GetAllProvidersAdSlotInf func
func GetAllProvidersAdSlotInf(db *sql.DB, Providers []*models.Provider, publisherID string) []*models.AdSlotProvider {
	adSlotProvidersInf := make([]*models.AdSlotProvider, 0)
	for _, provider := range Providers {
		adSlotProvidersInf = append(adSlotProvidersInf, GetProviderSlotData(db, provider.Providername, publisherID))
	}
	return adSlotProvidersInf
}

//GetProviderSlotData func
func GetProviderSlotData(db *sql.DB, providerName string, publisherID string) *models.AdSlotProvider {
	adSlotProvider := &models.AdSlotProvider{}
	rows, _ := db.Query("SELECT * FROM  adSlotProvider_"+providerName+" WHERE pubid = ?", publisherID)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&adSlotProvider.Pubid,
			&adSlotProvider.AdslotId,
			&adSlotProvider.Epc,
			&adSlotProvider.Ecc,
			&adSlotProvider.Floor_price,
			&adSlotProvider.Rev_share,
			&adSlotProvider.ProviderID,
		)
	}
	return adSlotProvider
}

//GetAllProvidersDetails func
func GetAllProvidersDetails(db *sql.DB, publisher *models.Publisher) []*models.Provider {
	providers := make([]*models.Provider, 0)
	providerIds := strings.Split(publisher.ProviderIds, ",")
	fmt.Println(providerIds)
	for _, provID := range providerIds {
		providers = append(providers, GetProviderInf(db, provID))
	}
	return providers
}

//GetProviderInf func
func GetProviderInf(db *sql.DB, providerID string) *models.Provider {
	provider := &models.Provider{}
	fmt.Println("ProviderId= ", providerID)

	rows, _ := db.Query("SELECT * FROM Provider WHERE providerID= ?", providerID)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&provider.ProviderID,
			&provider.Providername,
		)
	}
	return provider

}

//GetPublisherInf func
func GetPublisherInf(db *sql.DB, publisherID string) *models.Publisher {
	publisher := &models.Publisher{}
	rows, _ := db.Query("SELECT * FROM publisher WHERE pubid = ?", publisherID)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&publisher.Name,
			&publisher.IsActive,
			&publisher.Pubid,
			&publisher.ProviderIds,
		)
	}
	return publisher
}

//GetAdSlotInf func
func GetAdSlotInf(db *sql.DB, publisherID string) []*models.AdSlotsPlacement {
	adSlotInf := make([]*models.AdSlotsPlacement, 0)
	rows, _ := db.Query("SELECT * FROM AdSlotsPlacement WHERE pubid = ?", publisherID)
	defer rows.Close()
	for rows.Next() {
		adslot := &models.AdSlotsPlacement{}
		rows.Scan(
			&adslot.AdslotId,
			&adslot.Size,
			&adslot.Name,
			&adslot.Pubid,
			&adslot.ProviderIds,
		)
		adSlotInf = append(adSlotInf, adslot)
	}
	return adSlotInf
}
