package services

import (
	"fmt"
	"hb-server/models"
	"hb-server/util"
)

//GetConfigsData func
func GetConfigsData(Publisher *models.Publisher, AdslotsInf []*models.AdSlotsPlacement, Providers []*models.Provider, AdslotProviderInf []*models.AdSlotProvider) *models.Config {

	adslots := GetAdSlots(AdslotsInf)
	providers := GetProviders(Providers)
	providersMap := GetProvidersMap(AdslotProviderInf)

	config := &models.Config{}
	config.Adslots = adslots
	config.Providers = providers
	config.ProvidersMap = providersMap

	return config
}

//GetProvidersMap func
func GetProvidersMap(AdslotProviderInf []*models.AdSlotProvider) map[int64]models.ProvidersMapInf {

	groupedAdSlotInf := util.GroupByAdSlotIDgo(AdslotProviderInf)

	providersMapList := make(map[int64]models.ProvidersMapInf)

	for adslotID, adslotProvider := range groupedAdSlotInf {

		//	providersInfMap := make(models.ProvidersMapInf)
		mapArray := make(map[string]*models.AdslotPlacementDetails)

		for _, provider := range adslotProvider {

			adslotPlacementDetails := &models.AdslotPlacementDetails{
				provider.ProviderID,
				provider.Rev_share,
				provider.Epc,
				provider.Floor_price,
				provider.Ecc,
			}

			mapArray[provider.ProviderID] = adslotPlacementDetails

			//	providersInfArr = append(providersInfArr, mapArray)
		}
		providersMapList[adslotID] = mapArray
	}
	fmt.Println(len(groupedAdSlotInf))

	return providersMapList
}

//GetProviders func
func GetProviders(Providers []*models.Provider) models.ProvidersInfMap {

	providerList := make(models.ProvidersInfMap)

	for _, currProvider := range Providers {
		provider := &models.Provider{
			currProvider.ProviderID,
			currProvider.Providername,
		}
		providerList[currProvider.ProviderID] = provider
	}
	return providerList
}

//GetAdSlots func
func GetAdSlots(AdslotsInf []*models.AdSlotsPlacement) models.AdslotListMap {

	adSlotListMap := make(models.AdslotListMap)

	for _, currAdslot := range AdslotsInf {
		adslot := &models.Adslot{
			currAdslot.AdslotId,
			currAdslot.Size,
			currAdslot.Name,
			currAdslot.ProviderIds,
		}

		adSlotListMap[currAdslot.AdslotId] = adslot
	}
	return adSlotListMap
}
