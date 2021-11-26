package file_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo            string      `json:"document_no"`
		DeliverTo             string      `json:"deliver_to"`
		Quantity              float64     `json:"quantity"`
		PickedQuantity        float64     `json:"picked_quantity"`
		Price                 float64     `json:"price"`
	    Batch                 string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             float64     `json:"quantity"`
		CompletedQuantity    float64     `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 float64     `json:"quantity"`
			CompletedQuantity        float64     `json:"completed_quantity"`
			ErroredQuantity          float64     `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity float64     `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                     string `json:"api_schema"`
	MaterialCode                  string `json:"material_code"`
	Plant                         string `json:"plant/supplier"`
	Stock                         float64`json:"stock"`
	DocumentType                  string `json:"document_type"`
	DocumentNo                    string `json:"document_no"`
	PlannedDate                   string `json:"planned_date"`
	ValidatedDate                 string `json:"validated_date"`
	Deleted                       string `json:"deleted"`
}

type SDC struct {
	ConnectionKey        string `json:"connection_key"`
	Result               bool   `json:"result"`
	RedisKey             string `json:"redis_key"`
	Filepath             string `json:"filepath"`
	PurchasingInfoRecord struct {
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
		ConditionRecord                struct {
			ConditionRecord            string `json:"ConditionRecord"`
			ConditionValidityEndDate   string `json:"ConditionValidityEndDate"`
			ConditionValidityStartDate string `json:"ConditionValidityStartDate"`
			ConditionSequentialNumber  struct {
				ConditionSequentialNumber  int    `json:"ConditionSequentialNumber"`
				ConditionApplication       string `json:"ConditionApplication"`
				ConditionType              string `json:"ConditionType"`
				ConditionScaleQuantity     float64 `json:"ConditionScaleQuantity"`
				ConditionScaleQuantityUnit string `json:"ConditionScaleQuantityUnit"`
				ConditionScaleAmount       float64 `json:"ConditionScaleAmount"`
				ConditionRateValue         float64 `json:"ConditionRateValue"`
				ConditionQuantity          float64 `json:"ConditionQuantity"`
			} `json:"ConditionSequentialNumber"`
		} `json:"ConditionRecord"`
	} `json:"PurchasingInfoRecord"`
	APISchema            string `json:"api_schema"`
	PurchasingInfoRecord string `json:"purchasing_info_record"`
	Deleted              string `json:"deleted"`
}
