package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

var SelectOption = Type("SelectOption", func() {
	Description("Select Options")
	Field(1, "value", String, "Value", func() {
		Default("")
	})
	Field(2, "label", String, "Label", func() {
		Default("")
	})
})

var OrderListFiltersResultData = Type("OrderListFiltersResultData", func() {
	Description("OrderListFiltersResultData")
	Field(1, "keywords_type_list", ArrayOf(SelectOption), "Keywords Type List", func() {
		Description("List of keywords types")
	})
	Field(2, "platform_list", ArrayOf(SelectOption), "Platform List", func() {
		Description("List of user's platforms")
	})
	Field(3, "store_list", ArrayOf(SelectOption), "Store List", func() {
		Description("List of user's stores")
	})
	Field(4, "warehouse_list", ArrayOf(SelectOption), "Warehouse List", func() {
		Description("List of user's warehouses")
	})
	Field(5, "country_list", ArrayOf(SelectOption), "Country List", func() {
		Description("List of user's countries")
	})
})

var OrderListFiltersResult = Type("OrderListFilters", func() {
	Description("Filters for listing orders")
	Extend(BaseResponse)
	Field(1, "data", OrderListFiltersResultData, "data")
})

var OrderQueryPayload = Type("OrderQueryPayload", func() {
	Description("Order Query Params")
	Field(1, "id", ArrayOf(String), "Order ID", func() {
		Description("Order ID")
		Example([]string{"1", "2"})
	})
	Field(2, "platform_order_no", String, "Platform Order No", func() {
		Description("Platform Order No")
	})
	Field(3, "listing_sku", String, "Listing SKU", func() {
		Description("Listing SKU")
	})
	Field(4, "sku", String, "SKU", func() {
		Description("SKU")
	})
	Field(5, "nss_tracking_number", String, "NSS Tracking Number", func() {
		Description("NSS Tracking Number")
	})
	Field(6, "shipping_name", String, "Shipping Name", func() {
		Description("Shipping Name")
	})
	Field(7, "platform", String, "Platform", func() {
		Example("shopify")
	})
	Field(8, "status", ArrayOf(String), "Status", func() {
		Example([]string{"1", "2"})
	})
	Field(9, "store_id", String, "Store Id", func() {
		Example("1")
	})
	Field(10, "warehouse_id", String, "Warehouse Id", func() {
		Example("1")
	})
	Field(11, "country_code", String, "Country Code (二字码)", func() {
		Example("CN")
	})
	Field(12, "created_at_start", String, "Created Start Date", func() {
		Example("2022-01-01")
	})
	Field(13, "created_at_end", String, "Created End Date", func() {
		Example("2022-01-02")
	})
	Field(14, "ship_date_start", String, "Ship Start Date", func() {
		Example("2022-01-01")
	})
	Field(15, "ship_date_end", String, "Ship End Date", func() {
		Example("2022-01-02")
	})
	Field(16, "offline_order", String, "Is Offline Order", func() {
		Example("1")
	})
	Field(17, "page", String, "Page", func() {
		Example("1")
	})
	Field(18, "page_size", String, "Page Size", func() {
		Example("10")
	})
	Extend(AuthToken)
})

var OrderQueryParams = func() {
	Param("id")
	Param("platform_order_no")
	Param("listing_sku")
	Param("sku")
	Param("nss_tracking_number")
	Param("shipping_name")
	Param("platform")
	Param("status")
	Param("store_id")
	Param("warehouse_id")
	Param("country_code")
	Param("created_at_start")
	Param("created_at_end")
	Param("ship_date_start")
	Param("ship_date_end")
	Param("offline_order")
	Param("page")
	Param("page_size")
}

var OrderListItemShipInfo = Type("OrderListItemShipInfo", func() {
	Description("Order List Item Ship Info")
	Field(1, "shipping_name", String, "Shipping Name", func() {
		Example("John")
	})
	Field(2, "shipping_country", String, "Shipping Country", func() {
		Example("US")
	})
	Field(3, "zip_code", String, "Shipping Phone", func() {
		Example("+1 123 456 789")
	})
})

var ListOrderItem = Type("OrderItem", func() {
	Description("List Order Item")
	Field(1, "sku", String, "SKU", func() {
		Example("NSS-BLUE-2021")
	})
	Field(2, "qty", Int, "Qty", func() {
		Example(1)
	})
})

var ChannelOption = Type("ChannelOption", func() {
	Description("Channel Option")
	Field(1, "id", Int64, "Channel Cost Id", func() {
		Example(1)
	})
	Field(2, "channel_name", String, "Channel Name", func() {
		Example("GLOBAL-A-B&N")
	})
	Field(2, "shipping_cost", Float64, "Shipping Cost", func() {
		Example(36.73)
	})
	Field(4, "channel_type_name", String, "Channel Type Name", func() {
		Example("Amazon")
	})
	Field(5, "min_normal_days", Int, "Min Normal Days", func() {
		Example(1)
	})
	Field(6, "max_normal_days", Int, "Max Normal Days", func() {
		Example(3)
	})
	Field(7, "charge_weight", Int64, "Charge Weight", func() {
		Example(12)
	})
	Field(8, "fuel_fee", Float64, "Fuel Fee", func() {
		Example(12.45)
	})
	Field(9, "misc_fee", Float64, "Misc Fee", func() {
		Example(12.45)
	})
	Field(10, "processing_fee", Float64, "Processing Fee", func() {
		Example(12.45)
	})
	Field(11, "total_fee", Float64, "Total Fee", func() {
		Example(12.45)
	})
})

var HoldReason = Type("HoldReason", func() {
	Description("Hold Reason")
	Field(1, "type", String, "Hold Reason Type", func() {
		Example("inventory")
	})
	Field(2, "reason", String, "Hold Reason", func() {
		Example("insufficient stock.")
	})
})

var ListItem = Type("List Item", func() {
	Description("Order List Item")
	Field(1, "id", Int32, "Order Id", func() {
		Example(1)
	})
	Field(2, "platform_order_id", String, "Platform Order id", func() {
		Example("12345")
	})
	Field(3, "platform_order_no", String, "Platform Order No", func() {
		Example("12345")
	})
	Field(4, "platform", String, "Platform", func() {
		Example("shopify")
	})
	Field(5, "store_id", Int32, "Store Id", func() {
		Example(1)
	})
	Field(6, "store_name", String, "Store Name", func() {
		Example("NSS Demo Store")
	})
	Field(7, "shipping_info", OrderListItemShipInfo, "Shipping Info")
	Field(8, "channel_id", Int32, "Channel Id", func() {
		Example(10)
	})
	Field(9, "channel_name", String, "Channel Name", func() {
		Example("Fedex")
	})
	Field(10, "channel_type", Int, "Channel Type", func() {
		Enum(1, 2, 3)
		Example(1)
	})
	Field(11, "channel_type_name", String, "Channel Type Name", func() {
		Example("Fastest")
	})
	Field(12, "delivery_cost", Float64, "Delivery Cost", func() {
		Example(12.45)
	})
	Field(13, "nss_tracking_number", String, "NSS Tracking Number", func() {
		Example("123456789")
	})
	Field(14, "items", ArrayOf(ListOrderItem), "Items")
	Field(15, "warehouse_id", Int32, "Warehouse Id", func() {
		Example(1)
	})
	Field(16, "warehouse_name", String, "Warehouse Name", func() {
		Example("China Warehouse")
	})
	Field(17, "status", Int, "Status", func() {
		Example(1)
	})
	Field(18, "status_name", String, "Status Name", func() {
		Example("Pending")
	})
	Field(19, "created_at", String, "Created At", func() {
		Example("2020-01-01 00:00:00")
	})
	Field(20, "ship_date", String, "Ship Date", func() {
		Example("2020-01-01 00:00:00")
	})
	Field(21, "channel_options", ArrayOf(ChannelOption), "Channel Options")
	Field(22, "hold_reasons", ArrayOf(HoldReason), "Hold Reasons")
})

var GetOrderListData = Type("GetOrderListData", func() {
	Field(1, "list", ArrayOf(ListItem), "List")
	Field(2, "mate", MetaData, "meta")
})

var GetOrderListResult = Type("GetOrderListResult", func() {
	Description("Get Order List Result")
	Extend(BaseResponse)
	Field(1, "data", GetOrderListData, "Data")
})

var OrderCountData = Type("OrderCountData", func() {
	Description("Order Counts")
	Field(1, "total", Int, "Total", func() {
		Example(10)
	})
	Field(2, "ready_to_ship", Int, "Ready To Ship", func() {
		Example(3)
	})
	Field(3, "pending", Int, "Pending", func() {
		Example(3)
	})
	Field(4, "shipped", Int, "Shipped", func() {
		Example(2)
	})
	Field(5, "cancelled", Int, "Cancelled", func() {
		Example(1)
	})
	Field(6, "exception", Int, "Exception", func() {
		Example(1)
	})
})

var OrderCountResult = Type("OrderCountResult", func() {
	Description("Order Counts Result")
	Extend(BaseResponse)
	Field(1, "data", OrderCountData, "Order Counts")
})

var UploadOrdersPayload = Type("UploadOrdersPayload", func() {
	Description("Upload Orders Payload")
	Field(1, "file", Bytes, "file", func() {
		Example("order.xlsx")
		MaxLength(864000)
	})
	Field(2, "file_name", String, "file name", func() {
		Example("xxx")
		MaxLength(50)
	})
	Extend(AuthToken)
	Required("file", "file_name")
})

var UploadOrdersData = Type("UploadOrdersData", func() {
	Description("Upload Orders Data")
	Field(1, "total_count", Int, "Total Count", func() {
		Example(100)
	})
	Field(2, "success_count", Int, "Success Count", func() {
		Example(955)
	})
	Field(3, "fail_count", Int, "Fail Count", func() {
		Example(45)
	})
	Field(4, "result_file", String, "Result File", func() {
		Example("orders_result.excel (base64)")
	})
})

var UploadOrdersResult = Type("UploadOrdersResult", func() {
	Description("Upload Orders Result")
	Extend(BaseResponse)
	Field(1, "data", UploadOrdersData, "Data")
})

var ExportOrderResult = Type("ExportOrderResult", func() {
	Description("Export Order Result")
	Attribute("length", Int64, "Length is the downloaded content length in bytes.", func() {
		Example(4 * 1024 * 1024)
	})
	Required("length")
})
