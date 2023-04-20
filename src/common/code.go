package common

const (
	XE_UNEXPECTED_ERROR = "XE-99999"
	XE_CONFIG_ERROR     = "XE-00001"
	XE_LOADCONFIG_ERROR = "XE-00002"
	XE_COREV1_ERROR     = "XE-00003"
	XE_EXEC_ERROR       = "XE-00004"
)

var Code_description_definbtion map[string]string = map[string]string{
	XE_UNEXPECTED_ERROR: "未知错误",
	XE_CONFIG_ERROR:     "初始化配置异常",
	XE_LOADCONFIG_ERROR: "加载配置异常",
	XE_COREV1_ERROR:     "创建执行Api对象异常",
	XE_EXEC_ERROR:       "执行命令异常",
}
