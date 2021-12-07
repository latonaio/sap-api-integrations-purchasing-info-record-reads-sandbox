package responses

type PricingCondition struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
            PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
            PurchasingInfoRecordCategory   string `json:"PurchasingInfoRecordCategory"`
            PurchasingOrganization         string `json:"PurchasingOrganization"`
            Supplier                       string `json:"Supplier"`
			Material                       string `json:"Material"`
			MaterialGroup                  string `json:"MaterialGroup"`
	        Plant                          string `json:"Plant"`
			ConditionRecord                string `json:"ConditionRecord"`
			ConditionValidityEndDate       string `json:"ConditionValidityEndDate"`
			ConditionValidityStartDate     string `json:"ConditionValidityStartDate"`
			ConditionApplication           string `json:"ConditionApplication"`
			ConditionType                  string `json:"ConditionType"`
		} `json:"results"`
	} `json:"d"`
}
