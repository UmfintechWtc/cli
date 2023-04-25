package k8s

import (
	"bytes"
	"cli/src/common"
	"fmt"
	"os"
	"path/filepath"

	xlogger "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
)

type K8sApiServer struct{}

func (K *K8sApiServer) initConfigInPod(mode string) (*rest.Config, *kubernetes.Clientset) {
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
		return EnvConfig, ClientSet
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
		return FileConfig, ClientSet
	} else {
		xlogger.Error(common.Rsperror(
			common.XE_UNEXPECTED_ERROR,
			"未指明exec_pod当前执行环境类型"))
		os.Exit(2)
	}
	return nil, nil
}

func (K *K8sApiServer) initCoreV1(podname, container, namespace, mode string, command []string) remotecommand.Executor {
	config, clientset := K.initConfigInPod(mode)
	kubernetesReq := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podname).
		Namespace(namespace).
		SubResource("exec").Param("container", container)
	kubernetesReq.VersionedParams(
		&v1.PodExecOptions{
			Command: command,
			Stdin:   false,
			Stdout:  true,
			Stderr:  true,
			TTY:     false,
		},
		scheme.ParameterCodec,
	)
	executorApi, executor_err := remotecommand.NewSPDYExecutor(config, "POST", kubernetesReq.URL())
	if executor_err != nil {
		xlogger.Error(common.Rsperror(
			common.XE_COREV1_ERROR,
			executor_err.Error()))
		os.Exit(2)
	}
	return executorApi
}

func (K *K8sApiServer) Run_exec(podname, container, namespace, mode string, command []string) {
	var stdout, stderr bytes.Buffer
	exec := K.initCoreV1(podname, container, namespace, mode, command)
	stream_err := exec.Stream(remotecommand.StreamOptions{
		Stdin:  nil,
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if stream_err != nil {
		fmt_message := fmt.Sprintf("%s[%s]:%s\n", podname, container, command)
		xlogger.Error(common.Rsperror(
			common.XE_EXEC_ERROR,
			fmt_message+stream_err.Error()))
		os.Exit(2)
	}
	xlogger.Info("命令输出: ", stdout.String())
}
