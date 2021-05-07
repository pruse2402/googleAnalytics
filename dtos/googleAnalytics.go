package dtos

import "cloud.google.com/go/bigquery"

type GA_GoogleAnalytics struct {
	EventDate                  string              `json:"EventDate,nullable"`
	EventTimestamp             int64               `json:"EventTimestamp,nullable"`
	EventName                  string              `json:"EventName,nullable"`
	EventParams                *[]GA_EventParams   `json:"EventParams,nullable"`
	EventPreviousTimestamp     int64               `json:"EventPreviousTimestamp,nullable"`
	EventValueInUsd            float64             `json:"EventValueInUsd,nullable"`
	EventBundleSequenceID      int64               `json:"EventBundleSequenceID,nullable"`
	EventServerTimestampOffset int64               `json:"EventServerTimestampOffset,nullable"`
	UserID                     string              `json:"UserID,nullable"`
	UserPseudoID               string              `json:"UserPseudoID,nullable"`
	UserProperties             []GA_UserProperties `json:"UserProperties,nullable"`
	UserFirstTouchTimestamp    int64               `json:"UserFirstTouchTimestamp,nullable"`
	UserLtv                    *GA_UserLtv         `json:"UserLtv,nullable"`
	Device                     *GA_Device          `json:"Device,nullable"`
	Geo                        *GA_Geo             `json:"Geo,nullable"`
	AppInfo                    *GA_AppInfo         `json:"AppInfo,nullable"`
	TrafficSource              *GA_TrafficSource   `json:"TrafficSource,nullable"`
	StreamID                   string              `json:"StreamID,nullable"`
	Platform                   string              `json:"Platform,nullable"`
	EventDimensions            *GA_EventDimensions `json:"EventDimensions,nullable"`
	Ecommerce                  *GA_Ecommerce       `json:"Ecommerce,nullable"`
	Items                      []GA_Items          `json:"Items,nullable"`
}

type GA_Ecommerce struct {
	TotalItemQuantity    int64   `json:"TotalItemQuantity,nullable"`
	PurchaseRevenueInUsd float64 `json:"PurchaseRevenueInUsd,nullable"`
	PurchaseRevenue      float64 `json:"PurchaseRevenue,nullable"`
	RefundValueInUsd     float64 `json:"RefundValueInUsd,nullable"`
	RefundValue          float64 `json:"RefundValue,nullable"`
	ShippingValueInUsd   float64 `json:"ShippingValueInUsd,nullable"`
	ShippingValue        float64 `json:"ShippingValue,nullable"`
	TaxValueInUsd        float64 `json:"TaxValueInUsd,nullable"`
	TaxValue             float64 `json:"TaxValue,nullable"`
	UniqueItems          int64   `json:"UniqueItems,nullable"`
	TransactionID        string  `json:"TransactionID,nullable"`
}

type GA_EventDimensions struct {
	Hostname string `json:"Hostname,nullable"`
}

type GA_EventParams struct {
	Key   string   `json:"Key,nullable"`
	Value GA_Value `json:"Value,nullable"`
}

type GA_Value struct {
	StringValue string  `json:"StringValue,nullable"`
	IntValue    int64   `json:"IntValue,nullable"`
	FloatValue  float32 `json:"FloatValue,nullable"`
	DoubleValue float32 `json:"DoubleValue,nullable"`
}

type GA_UserProperties struct {
	Key   string                 `json:"Key,nullable"`
	Value GA_UserPropertiesValue `json:"Value,nullable"`
}

type GA_UserPropertiesValue struct {
	StringValue        string  `json:"StringValue,nullable"`
	IntValue           int64   `json:"IntValue,nullable"`
	FloatValue         float64 `json:"FloatValue,nullable"`
	DoubleValue        float64 `json:"DoubleValue,nullable"`
	SetTimestampMicros int64   `json:"SetTimestampMicros,nullable"`
}

type GA_UserLtv struct {
	Revenue  bigquery.NullFloat64 `json:"Revenue,nullable"`
	Currency string               `json:"Currency,nullable"`
}

type GA_Device struct {
	Category               string      `json:"Category,nullable"`
	MobileBrandName        string      `json:"MobileBrandName,nullable"`
	MobileModelName        string      `json:"MobileModelName,nullable"`
	MobileMarketingName    string      `json:"MobileMarketingName,nullable"`
	MobileOsHardwareModel  string      `json:"MobileOsHardwareModel,nullable"`
	OperatingSystem        string      `json:"OperatingSystem,nullable"`
	OperatingSystemVersion string      `json:"OperatingSystemVersion,nullable"`
	VendorID               string      `json:"VendorID,nullable"`
	AdvertisingID          string      `json:"AdvertisingID,nullable"`
	Language               string      `json:"Language,nullable"`
	IsLimitedAdTracking    string      `json:"IsLimitedAdTracking,nullable"`
	TimeZoneOffsetSeconds  int64       `json:"TimeZoneOffsetSeconds,nullable"`
	Browser                string      `json:"Browser,nullable"`
	BrowserVersion         string      `json:"BrowserVersion,nullable"`
	WebInfo                *GA_WebInfo `json:"WebInfo,nullable"`
}

type GA_WebInfo struct {
	Browser        string `json:"Browser,nullable"`
	BrowserVersion string `json:"BrowserVersion,nullable"`
	Hostname       string `json:"Hostname,nullable"`
}

type GA_Geo struct {
	Continent    string `json:"Continent,nullable"`
	Country      string `json:"Country,nullable"`
	Region       string `json:"Region,nullable"`
	City         string `json:"City,nullable"`
	SubContinent string `json:"SubContinent,nullable"`
	Metro        string `json:"Metro,nullable"`
}

type GA_AppInfo struct {
	ID            string `json:"ID,nullable"`
	Version       string `json:"Version,nullable"`
	InstallStore  string `json:"InstallStore,nullable"`
	FirebaseAppId string `json:"FirebaseAppId,nullable"`
	InstallSource string `json:"InstallSource,nullable"`
}

type GA_TrafficSource struct {
	Name   string `json:"Name,nullable"`
	Medium string `json:"Medium,nullable"`
	Source string `json:"Source,nullable"`
}

type GA_Items struct {
	ItemID           string  `json:"ItemID,nullable"`
	ItemName         string  `json:"ItemName,nullable"`
	ItemBrand        string  `json:"ItemBrand,nullable"`
	ItemVariant      string  `json:"ItemVariant,nullable"`
	ItemCategory     string  `json:"ItemCategory,nullable"`
	ItemCategory2    string  `json:"ItemCategory2,nullable"`
	ItemCategory3    string  `json:"ItemCategory3,nullable"`
	ItemCategory4    string  `json:"ItemCategory4,nullable"`
	ItemCategory5    string  `json:"ItemCategory5,nullable"`
	PriceInUsd       float64 `json:"PriceInUsd,nullable"`
	Price            float64 `json:"Price,nullable"`
	Quantity         int64   `json:"Quantity,nullable"`
	ItemRevenueInUsd float64 `json:"ItemRevenueInUsd,nullable"`
	ItemRevenue      float64 `json:"ItemRevenue,nullable"`
	ItemRefundInUsd  float64 `json:"ItemRefundInUsd,nullable"`
	ItemRefund       float64 `json:"ItemRefund,nullable"`
	Coupon           string  `json:"Coupon,nullable"`
	Affiliation      string  `json:"Affiliation,nullable"`
	LocationId       string  `json:"LocationId,nullable"`
	ItemListID       string  `json:"ItemListID,nullable"`
	ItemListName     string  `json:"ItemListName,nullable"`
	ItemListIndex    string  `json:"ItemListIndex,nullable"`
	PromotionID      string  `json:"PromotionID,nullable"`
	PromotionName    string  `json:"PromotionName,nullable"`
	CreativeName     string  `json:"CreativeName,nullable"`
	CreativeSlot     string  `json:"CreativeSlot,nullable"`
}
