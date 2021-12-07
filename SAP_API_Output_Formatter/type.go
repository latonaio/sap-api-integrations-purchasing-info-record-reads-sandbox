package sap_api_output_formatter

type PurchasingInfoRecordReads struct {
	ConnectionKey        string `json:"connection_key"`
	Result               bool   `json:"result"`
	RedisKey             string `json:"redis_key"`
	Filepath             string `json:"filepath"`
	APISchema            string `json:"api_schema"`
	PurchasingInfoRecord string `json:"purchasing_info_record"`
	Deleted              bool   `json:"deleted"`
}

type General struct {
		PurchasingInfoRecord        string `json:"PurchasingInfoRecord"`
		Supplier                    string `json:"Supplier"`
		Material                    string `json:"Material"`
		MaterialGroup               string `json:"MaterialGroup"`
		PurgDocOrderQuantityUnit    string `json:"PurgDocOrderQuantityUnit"`
		SupplierMaterialNumber      string `json:"SupplierMaterialNumber"`
		SupplierRespSalesPersonName string `json:"SupplierRespSalesPersonName"`
		SupplierPhoneNumber         string `json:"SupplierPhoneNumber"`
		SupplierMaterialGroup       string `json:"SupplierMaterialGroup"`
		IsRegularSupplier           bool   `json:"IsRegularSupplier"`
		AvailabilityStartDate       string `json:"AvailabilityStartDate"`
		AvailabilityEndDate         string `json:"AvailabilityEndDate"`
		Manufacturer                string `json:"Manufacturer"`
		CreationDate                string `json:"CreationDate"`
		PurchasingInfoRecordDesc    string `json:"PurchasingInfoRecordDesc"`
		LastChangeDateTime          string `json:"LastChangeDateTime"`
		IsDeleted                   bool   `json:"IsDeleted"`
}

type PurchasingOrganizationPlant struct {
		PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
        PurchasingInfoRecordCategory   string `json:"PurchasingInfoRecordCategory"`
		PurchasingOrganization         string `json:"PurchasingOrganization"`
		Plant                          string `json:"Plant"`
		Supplier                       string `json:"Supplier"`
		Material                       string `json:"Material"`
		MaterialGroup                  string `json:"MaterialGroup"`
		PurgDocOrderQuantityUnit       string `json:"PurgDocOrderQuantityUnit"`
		PurchasingGroup                string `json:"PurchasingGroup"`
		MinimumPurchaseOrderQuantity   string `json:"MinimumPurchaseOrderQuantity"`
		StandardPurchaseOrderQuantity  string `json:"StandardPurchaseOrderQuantity"`
		MaterialPlannedDeliveryDurn    string `json:"MaterialPlannedDeliveryDurn"`
		OverdelivTolrtdLmtRatioInPct   string `json:"OverdelivTolrtdLmtRatioInPct"`
		UnderdelivTolrtdLmtRatioInPct  string `json:"UnderdelivTolrtdLmtRatioInPct"`
		UnlimitedOverdeliveryIsAllowed bool   `json:"UnlimitedOverdeliveryIsAllowed"`
		LastReferencingPurchaseOrder   string `json:"LastReferencingPurchaseOrder"`
		LastReferencingPurOrderItem    string `json:"LastReferencingPurOrderItem"`
		NetPriceAmount                 string `json:"NetPriceAmount"`
		MaterialPriceUnitQty           string `json:"MaterialPriceUnitQty"`
		PurchaseOrderPriceUnit         string `json:"PurchaseOrderPriceUnit"`
		PriceValidityEndDate           string `json:"PriceValidityEndDate"`
		InvoiceIsGoodsReceiptBased     bool   `json:"InvoiceIsGoodsReceiptBased"`
		TaxCode                        string `json:"TaxCode"`
		IncotermsClassification        string `json:"IncotermsClassification"`
		MaximumOrderQuantity           string `json:"MaximumOrderQuantity"`
		IsRelevantForAutomSrcg         string `json:"IsRelevantForAutomSrcg"`
		IsEvaluatedRcptSettlmtAllowed  bool   `json:"IsEvaluatedRcptSettlmtAllowed"`
		IsPurOrderAllwdForInbDeliv     bool   `json:"IsPurOrderAllwdForInbDeliv"`
		IsOrderAcknRqd                 bool   `json:"IsOrderAcknRqd"`
		IsMarkedForDeletion            bool   `json:"IsMarkedForDeletion"`
}

type PricingCondition struct {
        PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
        PurchasingInfoRecordCategory   string `json:"PurchasingInfoRecordCategory"`
        PurchasingOrganization         string `json:"PurchasingOrganization"`
        Supplier                       string `json:"Supplier"`
        Material                       string `json:"Material"`
	    MaterialGroup                  string `json:"MaterialGroup"`
	    Plant                          string `json:"Plant"`
	    ConditionRecord                string `json:"ConditionRecord"`
	    ConditionValidityEndDate       string `json:"ConditionValidityEndDate"`
	    ConditionValidityStartDate     string `json:"ConditionValidityStartDate"`
	    ConditionApplication           string `json:"ConditionApplication"`
	    ConditionType                  string `json:"ConditionType"`
}
