package types

type ServerConfig struct {
	// 模块名称
	ModelName string `json:"model_name"`
	// 模块描述
	ModelComment string `json:"model_comment"`
	// 项目路径配置
	ProjectConfig ProjectConfig `json:"project_config"`
	// 生成配置
	GenerateConfig GenerateConfig `json:"generate_config"`
	// 数据表定义
	TableConfig TableConfig `json:"table_config"`
	// 缓存临时数据
	TmpData TmpData `json:"tmp_data"`
}

type ProjectConfig struct {
	// 项目绝对路径，可以自动录入
	WorkPath string `json:"work_path"`
	// 包路径，可以自动录入
	ModulePath string `json:"module_path"`
	// 路由文件存放位置
	RouterPath string `json:"router_path"`
	// 路由前缀
	RoutingPrefix string `json:"routing_prefix"`
	// Schema存放位置
	SchemaPath string `json:"schema_path"`
	// Service存放位置
	ServicePath string `json:"service_path"`
	// Controller存放位置
	ControllerPath string `json:"controller_path"`
}

type GenerateConfig struct {
	// 是否覆盖
	IsForce bool `json:"is_force"`
	// 是否存在更新操作
	IsHasUpdate bool `json:"is_has_update"`
	// 是否存在删除操作
	IsHasDelete bool `json:"is_has_delete"`
	// 是否存在查看操作
	IsHasShow bool `json:"is_has_show"`
	// 是否需要鉴权
	IsHasJwtAuth bool `json:"is_has_jwt_auth"`
	// 是否需要权限认证
	IsHasRbac bool `json:"is_has_rbac"`
	// 是否注入到系统菜单
	IsAddMenu bool `json:"is_add_menu"`
	// 父级菜单ID
	ParentMenuId int `json:"parent_menu_id"`
}

type TableConfig struct {
	// 是否开启软删除
	IsSoftDel bool `json:"is_soft_del"`
	// mixin混入  例如 AuditMixin  会自动添加 创建时间 删除时间 更新时间
	Mixin  []string     `json:"mixin"`
	Fields []FieldsItem `json:"fields"`
}

type FieldsItem struct {
	FieldType string `json:"field_type"`
	// 字段名称
	FieldName string `json:"field_name"`
	// 字段描述
	FieldComment string `json:"field_comment"`
	// 字段默认值
	FieldDefault string `json:"field_default"`
	// 是否为敏感字段
	IsSensitive bool `json:"is_sensitive"`
	// 是否可以为空
	IsOptional bool `json:"is_optional"`
	// 是否唯一
	IsUnique bool `json:"is_unique"`
	// 是否存在于插入表单
	IsInsert bool `json:"is_insert"`
	// 是否存在于编辑表单
	IsEdit bool `json:"is_edit"`
	// 是否列表展示字段
	IsListShow bool `json:"is_list_show"`
	// 是否列表默认展示字段
	IsListDefaultShow bool `json:"is_list_default_show"`
	// 是否为搜索字段
	IsSearch bool `json:"is_search"`
	// 查询方式
	SearchType string `json:"search_type"`
	// 展示方式
	ShowType string `json:"show_type"`
	// 展示配置说明
	ShowConfig string `json:"show_config"`
	// 字典类型
	DictType string `json:"dict_type"`
}

type TmpData struct {
	// 自动生成程序目录
	WorkPath string `json:"work_path"`
	// 下划线命名
	ModelNameCase string `json:"model_name_case"`
	// 全小写命名
	ModelNameLower string `json:"model_name_lower"`
	// 控制器路径
	ControllerPath string `json:"controller_path"`
	// 控制器名称
	ControllerName string `json:"controller_name"`
	// 服务路径
	ServicePath string `json:"service_path"`
	// 服务名称
	ServiceName string `json:"service_name"`
	// 服务类型路径
	ServiceTypePath string `json:"service_type_path"`
	// 服务类型名称
	ServiceTypeName string `json:"service_type_name"`
	// Schema路径
	SchemaPath string `json:"schema_path"`
	// Schema名称
	SchemaName string `json:"schema_name"`
	// 所需注入路由路径
	RouterPath string `json:"router_path"`
}

//注入方法
type MiddlewareFn = func(config *ServerConfig) (*ServerConfig, error)
