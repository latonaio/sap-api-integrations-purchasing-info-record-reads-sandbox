package main

import (
	sap_api_caller "sap-api-integrations-purchasing-info-record-reads/SAP_API_Caller"
	"sap-api-integrations-purchasing-info-record-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Purchasing_Info_Record_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetPurchasingInfoRecord(
        inoutSDC.PurchasingInfoRecord.Supplier,
        inoutSDC.PurchasingInfoRecord.Material,
        inoutSDC.PurchasingInfoRecord.PurchasingOrganization,
        inoutSDC.PurchasingInfoRecord.Plant,
        inoutSDC.PurchasingInfoRecord.MaterialGroup,
        inoutSDC.PurchasingInfoRecord.PurchasingInfoRecord,
        inoutSDC.PurchasingInfoRecord.PricingCondition.ConditionSequentialNumber.ConditionType,
	)
}
