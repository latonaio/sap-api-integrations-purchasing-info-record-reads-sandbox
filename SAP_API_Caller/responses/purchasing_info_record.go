package responses

type PurchasingInfoRecord struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
			PurchasingInfoRecordCategory   string `json:"PurchasingInfoRecordCategory"`
			Supplier                       string `json:"Supplier"`
			Material                       string `json:"Material"`
			MaterialGroup                  string `json:"MaterialGroup"`
			PurgDocOrderQuantityUnit       string `json:"PurgDocOrderQuantityUnit"`
			SupplierMaterialNumber         string `json:"SupplierMaterialNumber"`
			SupplierRespSalesPersonName    string `json:"SupplierRespSalesPersonName"`
			SupplierPhoneNumber            string `json:"SupplierPhoneNumber"`
			SupplierMaterialGroup          string `json:"SupplierMaterialGroup"`
			IsRegularSupplier              bool   `json:"IsRegularSupplier"`
			AvailabilityStartDate          string `json:"AvailabilityStartDate"`
			AvailabilityEndDate            string `json:"AvailabilityEndDate"`
			Manufacturer                   string `json:"Manufacturer"`
			LastChangeDateTime             string `json:"LastChangeDateTime"`
			IsDeletedHeader                bool   `json:"IsDeleted_Header"`
			PurchasingOrganization         string `json:"PurchasingOrganization"`
			Plant                          string `json:"Plant"`
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
			IsRelevantForAutomSrcg         bool   `json:"IsRelevantForAutomSrcg"`
			IsEvaluatedRcptSettlmtAllowed  bool   `json:"IsEvaluatedRcptSettlmtAllowed"`
			IsPurOrderAllwdForInbDeliv     bool   `json:"IsPurOrderAllwdForInbDeliv"`
			IsOrderAcknRqd                 bool   `json:"IsOrderAcknRqd"`
			IsMarkedForDeletionPOLevel     bool   `json:"IsMarkedForDeletion_PO_Level"`
		} `json:"results"`
	} `json:"d"`
}
