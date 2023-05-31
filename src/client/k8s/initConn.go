package k8s

import (
	"cli/src/common"
	"os"
	"path/filepath"

	xlogger "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sApiServer struct {
	K8sConfig *rest.Config
	K8sClient *kubernetes.Clientset
}

func (K *K8sApiServer) InitConfigInPod(mode string) *K8sApiServer {
	if mode == common.SUB_CMD_CLI_POD_TYPE {
		// Pod内加载sa默认信息及环境变量
		EnvConfig, EnvConfig_Err := rest.InClusterConfig()
		if EnvConfig_Err != nil {
			xlogger.Error(common.Rsperror(
				common.XE_CONFIG_ERROR,
				common.SUB_CMD_CLI_POD_TYPE+EnvConfig_Err.Error()))
			os.Exit(1)
		}
		ClientSet, ClientSet_Err := kubernetes.NewForConfig(EnvConfig)
		if ClientSet_Err != nil {
			xlogger.Error(common.Rsperror(
				common.XE_LOADCONFIG_ERROR,
				common.SUB_CMD_CLI_POD_TYPE+ClientSet_Err.Error()))
			os.Exit(2)
		}
		return &K8sApiServer{K8sClient: ClientSet, K8sConfig: EnvConfig}
		// return EnvConfig, ClientSet
	} else if mode == common.SUB_CMD_CLI_HOST_TYPE {
		// 宿主机加载kubeconfig
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		FileConfig, FileConfig_err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if FileConfig_err != nil {
			xlogger.Error(common.Rsperror(
				common.XE_CONFIG_ERROR,
				common.SUB_CMD_CLI_HOST_TYPE+FileConfig_err.Error()))
			os.Exit(1)
		}
		ClientSet, ClientSet_Err := kubernetes.NewForConfig(FileConfig)
		if ClientSet_Err != nil {
			xlogger.Error(common.Rsperror(
				common.XE_LOADCONFIG_ERROR,
				common.SUB_CMD_CLI_HOST_TYPE+ClientSet_Err.Error()))
			os.Exit(2)
		}
		return &K8sApiServer{K8sClient: ClientSet, K8sConfig: FileConfig}

		// return FileConfig, ClientSet
	} else {
		xlogger.Error(common.Rsperror(
			common.XE_UNEXPECTED_ERROR,
			"未指明exec_pod当前执行环境类型"))
		os.Exit(2)
	}
	return nil
}
