package main

import (
	sap_api_caller "sap-api-integrations-purchasing-info-record-reads/SAP_API_Caller"
	"sap-api-integrations-purchasing-info-record-reads/file_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := file_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Purchasing_Info_Record_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetPurchasingInfoRecord(
		inoutSDC.Supplier.Material,
		inoutSDC.Supplier.Material.PurchasingOrganization.Plant,
		inoutSDC.Supplier.MaterialGroup,
		inoutSDC.Supplier.MaterialGroup.PurchasingOrganization.Plant,
		inoutSDC.PurchasingInfoRecord,
		inoutSDC.PurchasingInfoRecord.ConditionType,
	)
}