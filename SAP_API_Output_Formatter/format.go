package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchasing-info-record-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToPurchasingInfoRecord(raw []byte, l *logger.Logger) *PurchasingInfoRecord {
	pm := &responses.PurchasingInfoRecord{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &PurchasingInfoRecord{
		PurchasingInfoRecord           data.PurchasingInfoRecord,
		PurchasingInfoRecordCategory   data.PurchasingInfoRecordCategory,
		Supplier                       data.Supplier,
		Material                       data.Material,
		MaterialGroup                  data.MaterialGroup,
		PurgDocOrderQuantityUnit       data.PurgDocOrderQuantityUnit,
		SupplierMaterialNumber         data.SupplierMaterialNumber,
		SupplierRespSalesPersonName    data.SupplierRespSalesPersonName,
		SupplierPhoneNumber            data.SupplierPhoneNumber,
		SupplierMaterialGroup          data.SupplierMaterialGroup,
		IsRegularSupplier              data.IsRegularSupplier,
		AvailabilityStartDate          data.AvailabilityStartDate,
		AvailabilityEndDate            data.AvailabilityEndDate,
		Manufacturer                   data.Manufacturer,
		LastChangeDateTime             data.LastChangeDateTime,
		IsDeletedHeader                data.IsDeleted_Header,
		PurchasingOrganization         data.PurchasingOrganization,
		Plant                          data.Plant,
		PurchasingGroup                data.PurchasingGroup,
		MinimumPurchaseOrderQuantity   data.MinimumPurchaseOrderQuantity,
		StandardPurchaseOrderQuantity  data.StandardPurchaseOrderQuantity,
		MaterialPlannedDeliveryDurn    data.MaterialPlannedDeliveryDurn,
		OverdelivTolrtdLmtRatioInPct   data.OverdelivTolrtdLmtRatioInPct,
		UnderdelivTolrtdLmtRatioInPct  data.UnderdelivTolrtdLmtRatioInPct,
		UnlimitedOverdeliveryIsAllowed data.UnlimitedOverdeliveryIsAllowed,
		LastReferencingPurchaseOrder   data.LastReferencingPurchaseOrder,
		LastReferencingPurOrderItem    data.LastReferencingPurOrderItem,
		NetPriceAmount                 data.NetPriceAmount,
		MaterialPriceUnitQty           data.MaterialPriceUnitQty,
		PurchaseOrderPriceUnit         data.PurchaseOrderPriceUnit,
		PriceValidityEndDate           data.PriceValidityEndDate,
		InvoiceIsGoodsReceiptBased     data.InvoiceIsGoodsReceiptBased,
		TaxCode                        data.TaxCode,
		IncotermsClassification        data.IncotermsClassification,
		MaximumOrderQuantity           data.MaximumOrderQuantity,
		IsRelevantForAutomSrcg         data.IsRelevantForAutomSrcg,
		IsEvaluatedRcptSettlmtAllowed  data.IsEvaluatedRcptSettlmtAllowed,
		IsPurOrderAllwdForInbDeliv     data.IsPurOrderAllwdForInbDeliv,
		IsOrderAcknRqd                 data.IsOrderAcknRqd,
		IsMarkedForDeletionPOLevel     data.IsMarkedForDeletion_PO_Level,
		ConditionRecord                data.ConditionRecord,
		ConditionValidityEndDate       data.ConditionValidityEndDate,
		ConditionValidityStartDate     data.ConditionValidityStartDate,
		ConditionSequentialNumber      data.ConditionSequentialNumber,
		ConditionApplication           data.ConditionApplication,
		ConditionType                  data.ConditionType,
		ConditionScaleQuantity         data.ConditionScaleQuantity,
		ConditionScaleQuantityUnit     data.ConditionScaleQuantityUnit,
		ConditionScaleAmount           data.ConditionScaleAmount,
		ConditionRateValue             data.ConditionRateValue,
		ConditionQuantity              data.ConditionQuantity,
	}
}
