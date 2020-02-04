package models

type Adslot struct {
	AdslotID    int64  `json:"id"`
	Size        string `json:"size"`
	Adslotname  string `json:"adslotname"`
	ProviderIds string `json:"providersid"`
}

type AdslotListMap map[int64]*Adslot
type ProvidersInfMap map[string]*Provider

type Provider struct {
	ProviderID   string `json:"id"`
	Providername string `json:"providername"`
}

type ProvidersList struct {
	Providerkey string    `json:"providerkey"`
	Provider    *Provider `json:"provider"`
}

type AdslotPlacementDetails struct {
	Id             string  `json:"id"` // providerID
	Rev_share      float32 `json:"revshare"`
	Extplacementid int64   `json:"epc"`
	FloorPrice     float32 `json:"bidprice"`
	Extpublisherid string  `json:"ecc"`
}

type ProvidersMapInf map[string]*AdslotPlacementDetails

type Config struct {
	Adslots      AdslotListMap             `json:"adslots"`
	Providers    ProvidersInfMap           `json:"providers"`
	ProvidersMap map[int64]ProvidersMapInf `json:"providersmap"` // adslotid:int64 as key in map
}

// db models

type Publisher struct {
	Name        string
	IsActive    bool
	Pubid       string
	ProviderIds string
}

type AdSlotsPlacement struct {
	AdslotId    int64
	Size        string
	Name        string
	Pubid       string
	ProviderIds string
}

type AdSlotProvider struct {
	Pubid       string
	AdslotId    int64
	Epc         int64
	Ecc         string
	Floor_price float32
	Rev_share   float32
	ProviderID  string
}

type IntToStuctArrayMap map[int64][]*AdSlotProvider

type BidResult struct {
	BidPrice   float32 `json:"bidPrice"`
	Adcode     string  `json:"adcode"`
	ProviderID string  `json:"providerid"`
	Ecc        string  `json:"ecc"`
	Epc        int64   `json:"epc"`
	Size       string  `json:"size"`
}

type AuctionResult map[string]map[string][]BidResult
