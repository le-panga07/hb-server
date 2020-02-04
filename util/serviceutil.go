package util

import (
	"fmt"
	"hb-server/models"
)

//GroupByAdSlotIDgo func
func GroupByAdSlotIDgo(AdslotProviderInf []*models.AdSlotProvider) models.IntToStuctArrayMap {

	fmt.Println("Group len start ", len(AdslotProviderInf))

	adslotGroupMap := make(models.IntToStuctArrayMap)
	for _, adslotprovider := range AdslotProviderInf {

		if _, containsKey := adslotGroupMap[adslotprovider.AdslotId]; !containsKey {
			adslotGroupMap[adslotprovider.AdslotId] = make([]*models.AdSlotProvider, 0)
		}
		adslotGroupMap[adslotprovider.AdslotId] = append(adslotGroupMap[adslotprovider.AdslotId], adslotprovider)
	}

	return adslotGroupMap
}

//append(adslotGroupMap[adslotprovider.AdslotId], adslotprovider)
