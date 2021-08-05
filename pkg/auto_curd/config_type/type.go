package config_type

type TableType struct {
	Name   string       `json:"name"`
	Mixin  []string     `json:"mixin"`
	Fields []*FieldType `json:"fields"`
}

type FieldType struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Comment     string `json:"comment"`
	Default     string `json:"default"`
	IsSensitive bool   `json:"is_sensitive"`
	IsOptional  bool   `json:"is_optional"`
	IsUnique    bool   `json:"is_unique"`
}

type Config struct {
	WorkPath      string             `json:"work_path"`
	Table         *TableType         `json:"table"`
	ProjectConfig *ProjectConfigType `json:"project_config"`
	Runtime       *RuntimeType       `json:"runtime"`
}

type ProjectConfigType struct {
	RouterPath     string `json:"router_path"`
	ModulePath     string `json:"module_path"`
	SchemaPath     string `json:"schema_path"`
	ServicePath    string `json:"service_path"`
	ControllerPath string `json:"controller_path"`
	IsForce        bool   `json:"is_force"`
}

type RuntimeType struct {
	ProjectPath     string `json:"project_path"`
	ProjectName     string `json:"project_name"`
	ControllerPath  string `json:"controller_path"`
	ControllerName  string `json:"controller_name"`
	ServicePath     string `json:"service_path"`
	ServiceName     string `json:"service_name"`
	ServiceTypePath string `json:"service_type_path"`
	ServiceTypeName string `json:"service_type_name"`
	SchemaPath      string `json:"schema_path"`
	SchemaName      string `json:"schema_name"`
	RouterPath      string `json:"router_path"`
}

type MiddlewareFn = func(config *Config) (*Config, error)
