package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-purchasing-info-record-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetPurchasingInfoRecord(purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant, materialGroup, conditionType string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(purchasingInfoRecord)
				wg.Done()
			}()
		case "Material":
			func() {
				c.Material(purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant)
				wg.Done()
			}()
		case "MaterialGroup":
			func() {
				c.MaterialGroup(purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant)
				wg.Done()
			}()

		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) General(purchasingInfoRecord string) {
	generalData, err := c.callPurchasingInfoRecordSrvAPIRequirementGeneral("A_PurchasingInfoRecord", purchasingInfoRecord)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(generalData)

	orgPlantData, err := c.callToPurgInfoRecdOrgPlantData(generalData[0].ToPurgInfoRecdOrgPlantData)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(orgPlantData)
	validityData, err := c.callToPurInfoRecdPrcgCndnValidity(orgPlantData.ToPurInfoRecdPrcgCndnValidity)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(validityData)

	cndnData, err := c.callToPurInfoRecdPrcgCndn(validityData.ToPurInfoRecdPrcgCndn)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(cndnData)
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementGeneral(api, purchasingInfoRecord string) ([]sap_api_output_formatter.General, error) {
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithGeneral(req, purchasingInfoRecord)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurgInfoRecdOrgPlantData(url string) (*sap_api_output_formatter.ToPurgInfoRecdOrgPlantData, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurgInfoRecdOrgPlantData(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurInfoRecdPrcgCndnValidity(url string) (*sap_api_output_formatter.ToPurInfoRecdPrcgCndnValidity, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurInfoRecdPrcgCndnValidity(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurInfoRecdPrcgCndn(url string) (*sap_api_output_formatter.ToPurInfoRecdPrcgCndn, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurInfoRecdPrcgCndn(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Material(purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant string) {
	materialData, err := c.callPurchasingInfoRecordSrvAPIRequirementMaterial("A_PurgInfoRecdOrgPlantData", purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(materialData)

	validityData, err := c.callToPurInfoRecdPrcgCndnValidity(materialData[0].ToPurInfoRecdPrcgCndnValidity)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(validityData)

	cndnData, err := c.callToPurInfoRecdPrcgCndn(validityData.ToPurInfoRecdPrcgCndn)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(cndnData)
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementMaterial(api, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant string) ([]sap_api_output_formatter.PurchasingOrganizationPlant, error) {
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithMaterial(req, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPurchasingOrganizationPlant(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurInfoRecdPrcgCndnValidity2(url string) (*sap_api_output_formatter.ToPurInfoRecdPrcgCndnValidity, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurInfoRecdPrcgCndnValidity(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurInfoRecdPrcgCndn2(url string) (*sap_api_output_formatter.ToPurInfoRecdPrcgCndn, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurInfoRecdPrcgCndn(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MaterialGroup(purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant string) {
	materialGroupData, err := c.callPurchasingInfoRecordSrvAPIRequirementMaterialGroup("A_PurgInfoRecdOrgPlantData", purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(materialGroupData)

	validityData, err := c.callToPurInfoRecdPrcgCndnValidity2(materialGroupData[0].ToPurInfoRecdPrcgCndnValidity)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(validityData)

	cndnData, err := c.callToPurInfoRecdPrcgCndn2(validityData.ToPurInfoRecdPrcgCndn)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(cndnData)
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementMaterialGroup(api, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant string) ([]sap_api_output_formatter.PurchasingOrganizationPlant, error) {
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithMaterialGroup(req, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPurchasingOrganizationPlant(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurInfoRecdPrcgCndnValidity3(url string) (*sap_api_output_formatter.ToPurInfoRecdPrcgCndnValidity, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurInfoRecdPrcgCndnValidity(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurInfoRecdPrcgCndn3(url string) (*sap_api_output_formatter.ToPurInfoRecdPrcgCndn, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurInfoRecdPrcgCndn(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithGeneral(req *http.Request, purchasingInfoRecord string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchasingInfoRecord eq '%s'", purchasingInfoRecord))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithMaterial(req *http.Request, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchasingInfoRecord ne '' and PurchasingInfoRecordCategory ne '' and Supplier eq '%s' and Material eq '%s' and PurchasingOrganization eq '%s' and Plant eq '%s'", supplier, material, purchasingOrganization, plant))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithMaterialGroup(req *http.Request, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchasingInfoRecord ne '' and PurchasingInfoRecordCategory ne '' and Supplier eq '%s' and MaterialGroup eq '%s' and PurchasingOrganization eq '%s' and Plant eq '%s'", supplier, materialGroup, purchasingOrganization, plant))
	req.URL.RawQuery = params.Encode()
}

// func (c *SAPAPICaller) getQueryWithPricingConditionMaterial(req *http.Request, purchasingInfoRecord, PurchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant, conditionType string) {
//	params := req.URL.Query()
//	params.Add("$filter", fmt.Sprintf("PurchasingInfoRecord ne '' and PurchasingInfoRecordCategory ne '' and Supplier eq '%s' and Material eq '%s' and PurchasingOrganization='%s' and Plant eq '%s' and conditionType eq '%s'", purchasingInfoRecord, PurchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant, conditionType))
//	req.URL.RawQuery = params.Encode()
// }

// func (c *SAPAPICaller) getQueryWithPricingConditionMaterialGroup(req *http.Request, purchasingInfoRecord, PurchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant, conditionType string) {
//	params := req.URL.Query()
//	params.Add("$filter", fmt.Sprintf("PurchasingInfoRecord ne '' and PurchasingInfoRecordCategory ne '' and Supplier eq '%s' and MaterialGroup eq '%s' and PurchasingOrganization='%s' and Plant eq '%s' and conditionType eq '%s'", purchasingInfoRecord, PurchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant, conditionType))
//	req.URL.RawQuery = params.Encode()
// }
