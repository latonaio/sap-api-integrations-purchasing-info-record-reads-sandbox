package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
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
		
func (c *SAPAPICaller) AsyncGetPurchasingInfoRecord(Supplier, Material, PurchasingOrganization, Plant, MaterialGroup, PurchasingInfoRecord, ConditionType string) {
	wg := &sync.WaitGroup{}

	wg.Add(6)
	go func() {
		c.Material(Supplier, Material)
		wg.Done()
	}()
	
	go func() {
		c.MaterialPlant(Supplier, Material, PurchasingOrganization, Plant)
		wg.Done()
	}()
	
	go func() {
		c.MaterialGroup(Supplier, MaterialGroup)
		wg.Done()
	}()
	
	go func() {
		c.MaterialGroupPlant(Supplier, MaterialGroup, PurchasingOrganization, Plant)
		wg.Done()
	}()
	
	go func() {
		c.PurchasingInfoRecord(PurchasingInfoRecord)
		wg.Done()
	}()
	
	go func() {
	    c.PricingCondition(PurchasingInfoRecord, ConditionType)
	    wg.Done()
	}()
	
	wg.Wait()
}

func (c *SAPAPICaller) Material(Supplier, Material string) {
	res, err := c.callPurchasingInfoRecordSrvAPIRequirementMaterial("/A_PurchasingInfoRecord", Supplier, Material)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) MaterialPlant(Supplier, Material, PurchasingOrganization, Plant string) {
	res, err := c.callPurchasingInfoRecordSrvAPIRequirementMaterialPlant("/A_PurchasingInfoRecord('{PurchasingInfoRecord}')/to_PurgInfoRecdOrgPlantData", Supplier, Material, PurchasingOrganization, Plant)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) MaterialGroup(Supplier, MaterialGroup string) {
	res, err := c.callPurchasingInfoRecordSrvAPIRequirementMaterialGroup("/A_PurchasingInfoRecord", Supplier, MaterialGroup)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

func (c *SAPAPICaller) MaterialGroupPlant(Supplier, MaterialGroup, PurchasingOrganization, Plant string) {
	res, err := c.callPurchasingInfoRecordSrvAPIRequirementMaterialGroupPlant("/A_PurchasingInfoRecord('{PurchasingInfoRecord}')/to_PurgInfoRecdOrgPlantData", Supplier, MaterialGroup, PurchasingOrganization, Plant)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)
	
func (c *SAPAPICaller) PurchasingInfoRecord(PurchasingInfoRecord string) {
	res, err := c.callPurchasingInfoRecordSrvAPIRequirementPurchasingInfoRecord("/A_PurchasingInfoRecord", PurchasingInfoRecord)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) PurchasingInfoRecordPricingCondition(PurchasingInfoRecord, ConditionType string) {
	res, err := c.callPurchasingInfoRecordSrvAPIRequirementPricingCondition("/A_PurInfoRecdPrcgCndn", PurchasingInfoRecord, ConditionType)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementMaterial(api, Supplier, Material string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Supplier, Material")
	params.Add("$filter", fmt.Sprintf("Supplier eq '%s' and Material eq '%s'", Supplier, Material))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementMaterialPlant(api, Supplier, Material, PurchasingOrganization, Plant string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Supplier, Material, PurchasingOrganization, Plant")
	params.Add("$filter", fmt.Sprintf("Supplier eq '%s' and Material eq '%s' and PurchasingOrganization eq '%s' and Plant eq '%s'", Supplier, Material, PurchasingOrganization, Plant))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementMaterialGroup(api, Supplier, MaterialGroup string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Supplier, MaterialGroup")
	params.Add("$filter", fmt.Sprintf("Supplier eq '%s' and MaterialGroup eq '%s'", Supplier, MaterialGroup))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementMaterialGroupPlant(api, Supplier, MaterialGroup, PurchasingOrganization, Plant string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Supplier, MaterialGroup, PurchasingOrganization, Plant")
	params.Add("$filter", fmt.Sprintf("Supplier eq '%s' and MaterialGroup eq '%s' and PurchasingOrganization eq '%s' and Plant eq '%s'", Supplier, MaterialGroup, PurchasingOrganization, Plant))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementPurchasingInfoRecord(api, PurchasingInfoRecord string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "PurchasingInfoRecord")
	params.Add("$filter", fmt.Sprintf("PurchasingInfoRecord  eq '%s'", PurchasingInfoRecord))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementPricingCondition(api, PurchasingInfoRecord, ConditionType string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "PurchasingInfoRecord, ConditionType")
	params.Add("$filter", fmt.Sprintf("PurchasingInfoRecord  eq '%s' and ConditionType eq '%s'", PurchasingInfoRecord, ConditionType))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}