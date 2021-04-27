package dtos

import "cloud.google.com/go/bigquery"

type GoogleAnalytics struct {
	EventDate                  bigquery.NullString  `json:"event_date,nullable"`
	EventTimestamp             bigquery.NullInt64   `json:"event_timestamp,nullable"`
	EventName                  bigquery.NullString  `json:"event_name,nullable"`
	EventParams                []EventParams        `json:"event_params,nullable"`
	EventPreviousTimestamp     bigquery.NullInt64   `json:"event_previous_timestamp,nullable"`
	EventValueInUsd            bigquery.NullFloat64 `json:"event_value_in_usd,nullable"`
	EventBundleSequenceID      bigquery.NullInt64   `json:"event_bundle_sequence_id,nullable"`
	EventServerTimestampOffset bigquery.NullInt64   `json:"event_server_timestamp,nullable"`
	UserID                     bigquery.NullString  `json:"user_id,nullable"`
	UserPseudoID               bigquery.NullString  `json:"user_pseudo_id,nullable"`
	UserProperties             []UserProperties     `json:"user_properties,nullable"`
	UserFirstTouchTimestamp    bigquery.NullInt64   `json:"user_first_touch_timestamp,nullable"`
	// UserLtv                    UserLtv              `bigquery:"user_ltv,nullable"`
	// Device Device `bigquery:"device,nullable"`
	Geo           Geo                 `json:"geo,nullable"`
	AppInfo       AppInfo             `json:"appInfo,nullable"`
	TrafficSource TrafficSource       `json:"trafficSource,nullable"`
	StreamID      bigquery.NullString `json:"stream_id,nullable"`
	Platform      bigquery.NullString `json:"platform,nullable"`
	// EventDimensions EventDimensions     `bigquery:"event_dimensions,nullable"`
	// Ecommerce Ecommerce `bigquery:"ecommerce,nullable"`
	Items []Items `json:"items,nullable"`
}

type Ecommerce struct {
	TotalItemQuantity    bigquery.NullInt64   `json:"total_item_quantity,nullable"`
	PurchaseRevenueInUsd bigquery.NullFloat64 `json:"purchase_revenue_in_usd,nullable"`
	PurchaseRevenue      bigquery.NullFloat64 `json:"purchase_revenue,nullable"`
	RefundValueInUsd     bigquery.NullFloat64 `json:"refund_value_in_usd,nullable"`
	RefundValue          bigquery.NullFloat64 `json:"refund_value,nullable"`
	ShippingValueInUsd   bigquery.NullFloat64 `json:"shipping_value_in_usd,nullable"`
	ShippingValue        bigquery.NullFloat64 `json:"shipping_value,nullable"`
	TaxValueInUsd        bigquery.NullFloat64 `json:"tax_value_in_usd,nullable"`
	TaxValue             bigquery.NullFloat64 `json:"tax_value,nullable"`
	UniqueItems          bigquery.NullInt64   `json:"unique_items,nullable"`
	TransactionID        bigquery.NullString  `json:"transaction_id,nullable"`
}

type EventDimensions struct {
	Hostname bigquery.NullString `json:"hostname,nullable"`
}

type EventParams struct {
	Key   bigquery.NullString `json:"key,nullable"`
	Value Value               `json:"value,nullable"`
}

type Value struct {
	StringValue bigquery.NullString  `json:"string_value,nullable"`
	IntValue    bigquery.NullInt64   `json:"int_value,nullable"`
	FloatValue  bigquery.NullFloat64 `json:"float_value,nullable"`
	DoubleValue bigquery.NullFloat64 `json:"double_value,nullable"`
}

type UserProperties struct {
	Key   bigquery.NullString `json:"key,nullable"`
	Value UserPropertiesValue `json:"value,nullable"`
}

type UserPropertiesValue struct {
	StringValue        bigquery.NullString  `json:"string_value,nullable"`
	IntValue           bigquery.NullInt64   `json:"int_value,nullable"`
	FloatValue         bigquery.NullFloat64 `json:"float_value,nullable"`
	DoubleValue        bigquery.NullFloat64 `json:"double_value,nullable"`
	SetTimestampMicros bigquery.NullInt64   `json:"set_timestamp_micros,nullable"`
}

type UserLtv struct {
	Revenue  bigquery.NullFloat64 `json:"revenue,nullable"`
	Currency bigquery.NullString  `json:"currency,nullable"`
}

type Device struct {
	Category               bigquery.NullString `json:"category,nullable"`
	MobileBrandName        bigquery.NullString `json:"mobile_brand_name,nullable"`
	MobileModelName        bigquery.NullString `json:"mobile_model_name,nullable"`
	MobileMarketingName    bigquery.NullString `json:"mobile_marketing_name,nullable"`
	MobileOsHardwareModel  bigquery.NullString `json:"mobile_os_hardware_model,nullable"`
	OperatingSystem        bigquery.NullString `json:"operating_system,nullable"`
	OperatingSystemVersion bigquery.NullString `json:"operating_system_version,nullable"`
	VendorID               bigquery.NullString `json:"vendor_id,nullable"`
	AdvertisingID          bigquery.NullString `json:"advertising_id,nullable"`
	Language               bigquery.NullString `json:"language,nullable"`
	IsLimitedAdTracking    bigquery.NullString `json:"is_limited_ad_tracking,nullable"`
	TimeZoneOffsetSeconds  bigquery.NullInt64  `json:"time_zone_offset_seconds,nullable"`
	Browser                bigquery.NullString `json:"browser,nullable"`
	BrowserVersion         bigquery.NullString `json:"browser_version,nullable"`
	WebInfo                WebInfo             `json:"web_info,nullable"`
}

type WebInfo struct {
	Browser        bigquery.NullString `json:"browser,nullable"`
	BrowserVersion bigquery.NullString `json:"browser_version,nullable"`
	Hostname       bigquery.NullString `json:"hostname,nullable"`
}

type Geo struct {
	Continent    bigquery.NullString `json:"continent,nullable"`
	Country      bigquery.NullString `json:"country,nullable"`
	Region       bigquery.NullString `json:"region,nullable"`
	City         bigquery.NullString `json:"city,nullable"`
	SubContinent bigquery.NullString `json:"sub_continent,nullable"`
	Metro        bigquery.NullString `json:"metro,nullable"`
}

type AppInfo struct {
	ID            bigquery.NullString `json:"id,nullable"`
	Version       bigquery.NullString `json:"version,nullable"`
	InstallStore  bigquery.NullString `json:"install_store,nullable"`
	FirebaseAppId bigquery.NullString `json:"firebase_app_id,nullable"`
	InstallSource bigquery.NullString `json:"install_source,nullable"`
}

type TrafficSource struct {
	Name   bigquery.NullString `json:"name,nullable"`
	Medium bigquery.NullString `json:"medium,nullable"`
	Source bigquery.NullString `json:"source,nullable"`
}

type Items struct {
	ItemID           bigquery.NullString  `json:"item_id,nullable"`
	ItemName         bigquery.NullString  `json:"item_name,nullable"`
	ItemBrand        bigquery.NullString  `json:"item_brand,nullable"`
	ItemVariant      bigquery.NullString  `json:"item_variant,nullable"`
	ItemCategory     bigquery.NullString  `json:"item_category,nullable"`
	ItemCategory2    bigquery.NullString  `json:"item_category2,nullable"`
	ItemCategory3    bigquery.NullString  `json:"item_category3,nullable"`
	ItemCategory4    bigquery.NullString  `json:"item_category4,nullable"`
	ItemCategory5    bigquery.NullString  `json:"item_category5,nullable"`
	PriceInUsd       bigquery.NullFloat64 `json:"price_in_usd,nullable"`
	Price            bigquery.NullFloat64 `json:"price,nullable"`
	Quantity         bigquery.NullInt64   `json:"quantity,nullable"`
	ItemRevenueInUsd bigquery.NullFloat64 `json:"item_revenue_in_usd,nullable"`
	ItemRevenue      bigquery.NullFloat64 `json:"item_revenue,nullable"`
	ItemRefundInUsd  bigquery.NullFloat64 `json:"item_refund_in_usd,nullable"`
	ItemRefund       bigquery.NullFloat64 `json:"item_refund,nullable"`
	Coupon           bigquery.NullString  `json:"coupon,nullable"`
	Affiliation      bigquery.NullString  `json:"affiliation,nullable"`
	LocationId       bigquery.NullString  `json:"location_id,nullable"`
	ItemListID       bigquery.NullString  `json:"item_list_id,nullable"`
	ItemListName     bigquery.NullString  `json:"item_list_name,nullable"`
	ItemListIndex    bigquery.NullString  `json:"item_list_index,nullable"`
	PromotionID      bigquery.NullString  `json:"promotion_id,nullable"`
	PromotionName    bigquery.NullString  `json:"promotion_name,nullable"`
	CreativeName     bigquery.NullString  `json:"creative_name,nullable"`
	CreativeSlot     bigquery.NullString  `json:"creative_slot,nullable"`
}
