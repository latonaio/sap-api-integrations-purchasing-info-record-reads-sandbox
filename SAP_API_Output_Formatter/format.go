package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchasing-info-record-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToGeneral(raw []byte, l *logger.Logger) (*General, error) {
	pm := &responses.General{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to General. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &General{
		PurchasingInfoRecord:        data.PurchasingInfoRecord,
		Supplier:                    data.Supplier,
		Material:                    data.Material,
		MaterialGroup:               data.MaterialGroup,
		PurgDocOrderQuantityUnit:    data.PurgDocOrderQuantityUnit,
		SupplierMaterialNumber:      data.SupplierMaterialNumber,
		SupplierRespSalesPersonName: data.SupplierRespSalesPersonName,
		SupplierPhoneNumber:         data.SupplierPhoneNumber,
		SupplierMaterialGroup:       data.SupplierMaterialGroup,
		IsRegularSupplier:           data.IsRegularSupplier,
		AvailabilityStartDate:       data.AvailabilityStartDate,
		AvailabilityEndDate:         data.AvailabilityEndDate,
		Manufacturer:                data.Manufacturer,
		CreationDate:                data.CreationDate,
		PurchasingInfoRecordDesc:    data.PurchasingInfoRecordDesc,
		LastChangeDateTime:          data.LastChangeDateTime,
		IsDeleted:                   data.IsDeleted,
	}, nil
}

func ConvertToPurchasingOrganizationPlant(raw []byte, l *logger.Logger) (*PurchasingOrganizationPlant, error) {
	pm := &responses.PurchasingOrganizationPlant{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PurchasingOrganizationPlant. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &PurchasingOrganizationPlant{
		PurchasingInfoRecord:           data.PurchasingInfoRecord,
        PurchasingInfoRecordCategory:   data.PurchasingInfoRecordCategory,
		PurchasingOrganization:         data.PurchasingOrganization,
		Plant:                          data.Plant,
		Supplier:                       data.Supplier,
		Material:                       data.Material,
		MaterialGroup:                  data.MaterialGroup,
		PurgDocOrderQuantityUnit:       data.PurgDocOrderQuantityUnit,
		PurchasingGroup:                data.PurchasingGroup,
		MinimumPurchaseOrderQuantity:   data.MinimumPurchaseOrderQuantity,
		StandardPurchaseOrderQuantity:  data.StandardPurchaseOrderQuantity,
		MaterialPlannedDeliveryDurn:    data.MaterialPlannedDeliveryDurn,
		OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
		UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
		UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
		LastReferencingPurchaseOrder:   data.LastReferencingPurchaseOrder,
		LastReferencingPurOrderItem:    data.LastReferencingPurOrderItem,
		NetPriceAmount:                 data.NetPriceAmount,
		MaterialPriceUnitQty:           data.MaterialPriceUnitQty,
		PurchaseOrderPriceUnit:         data.PurchaseOrderPriceUnit,
		PriceValidityEndDate:           data.PriceValidityEndDate,
		InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
		TaxCode:                        data.TaxCode,
		IncotermsClassification:        data.IncotermsClassification,
		MaximumOrderQuantity:           data.MaximumOrderQuantity,
		IsRelevantForAutomSrcg:         data.IsRelevantForAutomSrcg,
		IsEvaluatedRcptSettlmtAllowed:  data.IsEvaluatedRcptSettlmtAllowed,
		IsPurOrderAllwdForInbDeliv:     data.IsPurOrderAllwdForInbDeliv,
		IsOrderAcknRqd:                 data.IsOrderAcknRqd,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}, nil
}

func ConvertToPricingCondition(raw []byte, l *logger.Logger) (*PricingCondition, error) {
	pm := &responses.PricingCondition{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PricingCondition. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &PricingCondition{
        PurchasingInfoRecord:           data.PurchasingInfoRecord,
        PurchasingInfoRecordCategory:   data.PurchasingInfoRecordCategory,
        PurchasingOrganization:         data.PurchasingOrganization,
        Supplier:                       data.Supplier,
        Material:                       data.Material,
	    MaterialGroup:                  data.MaterialGroup,
	    Plant:                          data.Plant,
	    ConditionRecord:                data.ConditionRecord,
	    ConditionValidityEndDate:       data.ConditionValidityEndDate,
	    ConditionValidityStartDate:     data.ConditionValidityStartDate,
	    ConditionApplication:           data.ConditionApplication,
	    ConditionType:                  data.ConditionType,
	}, nil
}
