package dtos

import "cloud.google.com/go/bigquery"

type GoogleAnalytics struct {
	EventDate                  bigquery.NullString  `bigquery:"event_date,nullable"`
	EventTimestamp             bigquery.NullInt64   `bigquery:"event_timestamp,nullable"`
	EventName                  bigquery.NullString  `bigquery:"event_name,nullable"`
	EventParams                []EventParams        `bigquery:"event_params,nullable"`
	EventPreviousTimestamp     bigquery.NullInt64   `bigquery:"event_previous_timestamp,nullable"`
	EventValueInUsd            bigquery.NullFloat64 `bigquery:"event_value_in_usd,nullable"`
	EventBundleSequenceID      bigquery.NullInt64   `bigquery:"event_bundle_sequence_id,nullable"`
	EventServerTimestampOffset bigquery.NullInt64   `bigquery:"event_server_timestamp_offset,nullable"`
	UserID                     bigquery.NullString  `bigquery:"user_id,nullable"`
	UserPseudoID               bigquery.NullString  `bigquery:"user_pseudo_id,nullable"`
	UserProperties             []UserProperties     `bigquery:"user_properties,nullable"`
	UserFirstTouchTimestamp    bigquery.NullInt64   `bigquery:"user_first_touch_timestamp,nullable"`
	UserLtv                    *UserLtv             `bigquery:"user_ltv,nullable"`
	Device                     *Device              `bigquery:"device,nullable"`
	Geo                        *Geo                 `bigquery:"geo,nullable"`
	AppInfo                    *AppInfo             `bigquery:"app_info,nullable"`
	TrafficSource              *TrafficSource       `bigquery:"traffic_source,nullable"`
	StreamID                   bigquery.NullString  `bigquery:"stream_id,nullable"`
	Platform                   bigquery.NullString  `bigquery:"platform,nullable"`
	EventDimensions            *EventDimensions     `bigquery:"event_dimensions,nullable"`
	Ecommerce                  *Ecommerce           `bigquery:"ecommerce,nullable"`
	Items                      []Items              `bigquery:"items,nullable"`
}

type Ecommerce struct {
	TotalItemQuantity    bigquery.NullInt64   `bigquery:"total_item_quantity,nullable"`
	PurchaseRevenueInUsd bigquery.NullFloat64 `bigquery:"purchase_revenue_in_usd,nullable"`
	PurchaseRevenue      bigquery.NullFloat64 `bigquery:"purchase_revenue,nullable"`
	RefundValueInUsd     bigquery.NullFloat64 `bigquery:"refund_value_in_usd,nullable"`
	RefundValue          bigquery.NullFloat64 `bigquery:"refund_value,nullable"`
	ShippingValueInUsd   bigquery.NullFloat64 `bigquery:"shipping_value_in_usd,nullable"`
	ShippingValue        bigquery.NullFloat64 `bigquery:"shipping_value,nullable"`
	TaxValueInUsd        bigquery.NullFloat64 `bigquery:"tax_value_in_usd,nullable"`
	TaxValue             bigquery.NullFloat64 `bigquery:"tax_value,nullable"`
	UniqueItems          bigquery.NullInt64   `bigquery:"unique_items,nullable"`
	TransactionID        bigquery.NullString  `bigquery:"transaction_id,nullable"`
}

type EventDimensions struct {
	Hostname bigquery.NullString `bigquery:"hostname,nullable"`
}

type EventParams struct {
	Key   bigquery.NullString `bigquery:"key,nullable"`
	Value Value               `bigquery:"value,nullable"`
}

type Value struct {
	StringValue bigquery.NullString  `bigquery:"string_value,nullable"`
	IntValue    bigquery.NullInt64   `bigquery:"int_value,nullable"`
	FloatValue  bigquery.NullFloat64 `bigquery:"float_value,nullable"`
	DoubleValue bigquery.NullFloat64 `bigquery:"double_value,nullable"`
}

type UserProperties struct {
	Key   bigquery.NullString `bigquery:"key,nullable"`
	Value UserPropertiesValue `bigquery:"value,nullable"`
}

type UserPropertiesValue struct {
	StringValue        bigquery.NullString  `bigquery:"string_value,nullable"`
	IntValue           bigquery.NullInt64   `bigquery:"int_value,nullable"`
	FloatValue         bigquery.NullFloat64 `bigquery:"float_value,nullable"`
	DoubleValue        bigquery.NullFloat64 `bigquery:"double_value,nullable"`
	SetTimestampMicros bigquery.NullInt64   `bigquery:"set_timestamp_micros,nullable"`
}

type UserLtv struct {
	Revenue  bigquery.NullFloat64 `bigquery:"revenue,nullable"`
	Currency bigquery.NullString  `bigquery:"currency,nullable"`
}

type Device struct {
	Category               bigquery.NullString `bigquery:"category,nullable"`
	MobileBrandName        bigquery.NullString `bigquery:"mobile_brand_name,nullable"`
	MobileModelName        bigquery.NullString `bigquery:"mobile_model_name,nullable"`
	MobileMarketingName    bigquery.NullString `bigquery:"mobile_marketing_name,nullable"`
	MobileOsHardwareModel  bigquery.NullString `bigquery:"mobile_os_hardware_model,nullable"`
	OperatingSystem        bigquery.NullString `bigquery:"operating_system,nullable"`
	OperatingSystemVersion bigquery.NullString `bigquery:"operating_system_version,nullable"`
	VendorID               bigquery.NullString `bigquery:"vendor_id,nullable"`
	AdvertisingID          bigquery.NullString `bigquery:"advertising_id,nullable"`
	Language               bigquery.NullString `bigquery:"language,nullable"`
	IsLimitedAdTracking    bigquery.NullString `bigquery:"is_limited_ad_tracking,nullable"`
	TimeZoneOffsetSeconds  bigquery.NullInt64  `bigquery:"time_zone_offset_seconds,nullable"`
	Browser                bigquery.NullString `bigquery:"browser,nullable"`
	BrowserVersion         bigquery.NullString `bigquery:"browser_version,nullable"`
	WebInfo                *WebInfo            `bigquery:"web_info,nullable"`
}

type WebInfo struct {
	Browser        bigquery.NullString `bigquery:"browser,nullable"`
	BrowserVersion bigquery.NullString `bigquery:"browser_version,nullable"`
	Hostname       bigquery.NullString `bigquery:"hostname,nullable"`
}

type Geo struct {
	Continent    bigquery.NullString `bigquery:"continent,nullable"`
	Country      bigquery.NullString `bigquery:"country,nullable"`
	Region       bigquery.NullString `bigquery:"region,nullable"`
	City         bigquery.NullString `bigquery:"city,nullable"`
	SubContinent bigquery.NullString `bigquery:"sub_continent,nullable"`
	Metro        bigquery.NullString `bigquery:"metro,nullable"`
}

type AppInfo struct {
	ID            bigquery.NullString `bigquery:"id,nullable"`
	Version       bigquery.NullString `bigquery:"version,nullable"`
	InstallStore  bigquery.NullString `bigquery:"install_store,nullable"`
	FirebaseAppId bigquery.NullString `bigquery:"firebase_app_id,nullable"`
	InstallSource bigquery.NullString `bigquery:"install_source,nullable"`
}

type TrafficSource struct {
	Name   bigquery.NullString `bigquery:"name,nullable"`
	Medium bigquery.NullString `bigquery:"medium,nullable"`
	Source bigquery.NullString `bigquery:"source,nullable"`
}

type Items struct {
	ItemID           bigquery.NullString  `bigquery:"item_id,nullable"`
	ItemName         bigquery.NullString  `bigquery:"item_name,nullable"`
	ItemBrand        bigquery.NullString  `bigquery:"item_brand,nullable"`
	ItemVariant      bigquery.NullString  `bigquery:"item_variant,nullable"`
	ItemCategory     bigquery.NullString  `bigquery:"item_category,nullable"`
	ItemCategory2    bigquery.NullString  `bigquery:"item_category2,nullable"`
	ItemCategory3    bigquery.NullString  `bigquery:"item_category3,nullable"`
	ItemCategory4    bigquery.NullString  `bigquery:"item_category4,nullable"`
	ItemCategory5    bigquery.NullString  `bigquery:"item_category5,nullable"`
	PriceInUsd       bigquery.NullFloat64 `bigquery:"price_in_usd,nullable"`
	Price            bigquery.NullFloat64 `bigquery:"price,nullable"`
	Quantity         bigquery.NullInt64   `bigquery:"quantity,nullable"`
	ItemRevenueInUsd bigquery.NullFloat64 `bigquery:"item_revenue_in_usd,nullable"`
	ItemRevenue      bigquery.NullFloat64 `bigquery:"item_revenue,nullable"`
	ItemRefundInUsd  bigquery.NullFloat64 `bigquery:"item_refund_in_usd,nullable"`
	ItemRefund       bigquery.NullFloat64 `bigquery:"item_refund,nullable"`
	Coupon           bigquery.NullString  `bigquery:"coupon,nullable"`
	Affiliation      bigquery.NullString  `bigquery:"affiliation,nullable"`
	LocationId       bigquery.NullString  `bigquery:"location_id,nullable"`
	ItemListID       bigquery.NullString  `bigquery:"item_list_id,nullable"`
	ItemListName     bigquery.NullString  `bigquery:"item_list_name,nullable"`
	ItemListIndex    bigquery.NullString  `bigquery:"item_list_index,nullable"`
	PromotionID      bigquery.NullString  `bigquery:"promotion_id,nullable"`
	PromotionName    bigquery.NullString  `bigquery:"promotion_name,nullable"`
	CreativeName     bigquery.NullString  `bigquery:"creative_name,nullable"`
	CreativeSlot     bigquery.NullString  `bigquery:"creative_slot,nullable"`
}
