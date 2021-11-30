package responses

type PricingCondition struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PurchasingInfoRecord       string `json:"PurchasingInfoRecord"`
			ConditionRecord            string `json:"ConditionRecord"`
			ConditionValidityEndDate   string `json:"ConditionValidityEndDate"`
			ConditionValidityStartDate string `json:"ConditionValidityStartDate"`
			ConditionSequentialNumber  string `json:"ConditionSequentialNumber"`
			ConditionApplication       string `json:"ConditionApplication"`
			ConditionType              string `json:"ConditionType"`
			ConditionScaleQuantity     string `json:"ConditionScaleQuantity"`
			ConditionScaleQuantityUnit string `json:"ConditionScaleQuantityUnit"`
			ConditionScaleAmount       string `json:"ConditionScaleAmount"`
			ConditionRateValue         string `json:"ConditionRateValue"`
			ConditionQuantity          string `json:"ConditionQuantity"`
		} `json:"results"`
	} `json:"d"`
}
