package sap_api_caller

type PurchasingInfoRecordReads struct {
	 ConnectionKey        string `json:"connection_key"`
	 Result               bool   `json:"result"`
	 RedisKey             string `json:"redis_key"`
	 Filepath             string `json:"filepath"`
	 APISchema            string `json:"api_schema"`
	 PurchasingInfoRecord string `json:"purchasing_info_record"`
	 Deleted              string `json:"deleted"`
}

type PurchasingInfoRecord           struct {
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
     IsRegularSupplier              string `json:"IsRegularSupplier"`
     AvailabilityStartDate          string `json:"AvailabilityStartDate"`
     AvailabilityEndDate            string `json:"AvailabilityEndDate"`
     Manufacturer                   string `json:"Manufacturer"`
     LastChangeDateTime             string `json:"LastChangeDateTime"`
     IsDeletedHeader                string `json:"IsDeleted_Header"`
     PurchasingOrganization         string `json:"PurchasingOrganization"`
     Plant                          string `json:"Plant"`
     PurchasingGroup                string `json:"PurchasingGroup"`
     MinimumPurchaseOrderQuantity   float64 `json:"MinimumPurchaseOrderQuantity"`
     StandardPurchaseOrderQuantity  float64 `json:"StandardPurchaseOrderQuantity"`
     MaterialPlannedDeliveryDurn    int    `json:"MaterialPlannedDeliveryDurn"`
     OverdelivTolrtdLmtRatioInPct   float64 `json:"OverdelivTolrtdLmtRatioInPct"`
     UnderdelivTolrtdLmtRatioInPct  float64 `json:"UnderdelivTolrtdLmtRatioInPct"`
     UnlimitedOverdeliveryIsAllowed string `json:"UnlimitedOverdeliveryIsAllowed"`
     LastReferencingPurchaseOrder   string `json:"LastReferencingPurchaseOrder"`
     LastReferencingPurOrderItem    int    `json:"LastReferencingPurOrderItem"`
     NetPriceAmount                 float64 `json:"NetPriceAmount"`
     MaterialPriceUnitQty           float64 `json:"MaterialPriceUnitQty"`
     PurchaseOrderPriceUnit         string `json:"PurchaseOrderPriceUnit"`
     PriceValidityEndDate           string `json:"PriceValidityEndDate"`
     InvoiceIsGoodsReceiptBased     string `json:"InvoiceIsGoodsReceiptBased"`
     TaxCode                        string `json:"TaxCode"`
     IncotermsClassification        string `json:"IncotermsClassification"`
     MaximumOrderQuantity           float64 `json:"MaximumOrderQuantity"`
     IsRelevantForAutomSrcg         string `json:"IsRelevantForAutomSrcg"`
     IsEvaluatedRcptSettlmtAllowed  string `json:"IsEvaluatedRcptSettlmtAllowed"`
     IsPurOrderAllwdForInbDeliv     string `json:"IsPurOrderAllwdForInbDeliv"`
     IsOrderAcknRqd                 string `json:"IsOrderAcknRqd"`
     IsMarkedForDeletionPOLevel     string `json:"IsMarkedForDeletion_PO_Level"`
}
 
type ConditionRecord            struct {
	 PurchasingInfoRecord       string `json:"PurchasingInfoRecord"`
     ConditionRecord            string `json:"ConditionRecord"`
     ConditionValidityEndDate   string `json:"ConditionValidityEndDate"`
     ConditionValidityStartDate string `json:"ConditionValidityStartDate"`
}

type ConditionSequentialNumber  struct {
	 PurchasingInfoRecord       string `json:"PurchasingInfoRecord"`
     ConditionRecord            string `json:"ConditionRecord"`
     ConditionSequentialNumber  int    `json:"ConditionSequentialNumber"`
     ConditionApplication       string `json:"ConditionApplication"`
     ConditionType              string `json:"ConditionType"`
     ConditionScaleQuantity     float64 `json:"ConditionScaleQuantity"`
     ConditionScaleQuantityUnit string `json:"ConditionScaleQuantityUnit"`
     ConditionScaleAmount       float64 `json:"ConditionScaleAmount"`
     ConditionRateValue         float64 `json:"ConditionRateValue"`
     ConditionQuantity          float64 `json:"ConditionQuantity"`
}
