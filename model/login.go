package model

type LoginInfo struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    LoginData `json:"data"`
}

type LoginData struct {
	PermissionIDList []string `json:"permissionIdList"`
	Token            string   `json:"token"`
	UserInfo         UserInfo `json:"userInfo"`
	IsLicenseEnd     int      `json:"isLicenseEnd"`
	LicenseEndTime   string   `json:"licenseEndTime"`
}

type UserInfo struct {
	UserID         string `json:"userId"`
	Username       string `json:"username"`
	RealName       string `json:"realName"`
	RoleID         string `json:"roleId"`
	RoleName       string `json:"roleName"`
	DepartmentID   string `json:"departmentId"`
	DepartmentName string `json:"departmentName"`
	AuthStartTime  string `json:"authStartTime"`
	AuthEndTime    string `json:"authEndTime"`
}
