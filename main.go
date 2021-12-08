package main

import (
	sap_api_caller "sap-api-integrations-purchasing-info-record-reads/SAP_API_Caller"
	"sap-api-integrations-purchasing-info-record-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Purchasing_Info_Record_material-group-sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetPurchasingInfoRecord(
		inoutSDC.PurchasingInfoRecord.PurchasingInfoRecord,
		inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.PurchasingInfoRecordCategory,
		inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.Supplier,
		inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.Material,
		inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.PurchasingOrganization,
		inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.Plant,
		inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.MaterialGroup,
		inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.PricingCondition.ConditionType,
	)
}
