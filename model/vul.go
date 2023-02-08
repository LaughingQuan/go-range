package model

type VulInfo struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    VulData `json:"data"`
}

type Records struct {
	VulResultID      string      `json:"vulResultId"`
	SecurityLevelID  int         `json:"securityLevelId"`
	VulName          string      `json:"vulName"`
	VulSerial        string      `json:"vulSerial"`
	VulURL           string `json:"vulUrl"`
	VulMidTypeID     string      `json:"vulMidTypeId"`
	IastParam        string      `json:"iastParam"`
	DetectEngineID   int         `json:"detectEngineId"`
	Status           int         `json:"status"`
	IastActiveVerify interface{} `json:"iastActiveVerify"`
	CreateDate       string      `json:"createDate"`
	FirstCreateDate  string      `json:"firstCreateDate"`
	MergeCount       interface{} `json:"mergeCount"`
	AssignUserID     interface{} `json:"assignUserId"`
	MergeVulCount    int         `json:"mergeVulCount"`
}

type VulData struct {
	Records          []Records     `json:"records"`
	Total            int           `json:"total"`
	Size             int           `json:"size"`
	Current          int           `json:"current"`
	Orders           []interface{} `json:"orders"`
	OptimizeCountSQL bool          `json:"optimizeCountSql"`
	HitCount         bool          `json:"hitCount"`
	CountID          interface{}   `json:"countId"`
	MaxLimit         interface{}   `json:"maxLimit"`
	SearchCount      bool          `json:"searchCount"`
	Pages            int           `json:"pages"`
}

type ConfigVul struct {
	Data []ConfigVulData `json:"data"`
}

type ConfigVulData struct {
	VulURL  string `json:"vulUrl"`
	VulName string `json:"vulName"`
}
