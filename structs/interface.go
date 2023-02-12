package structs

type ConfigFile struct {
	Profile string `json:"profile"`
	Host    string `json:"host"`
}

type RoleMapping struct {
	Backend_roles []string `json:"backend_roles"`
	Host          []string `json:"hosts"`
	User          []string `json:"users"`
}

type Role struct {
	ClusterPermissions []string             `json:"cluster_permissions"`
	IndexPermissions   []Index_Permissions  `json:"index_permissions"`
	TenantPermissions  []Tenant_Permissions `json:"tenant_permissions"`
}

type Index_Permissions struct {
	IndexPatterns  []string `json:"index_patterns"`
	DLS            string   `json:"dls"`
	FLS            []string `json:"fls"`
	MaskedFields   []string `json:"masked_fields"`
	AllowedActions []string `json:"allowed_actions"`
}

type Tenant_Permissions struct {
	TenantPatterns []string `json:"tenant_patterns"`
	AllowedActions []string `json:"allowed_actions"`
}
