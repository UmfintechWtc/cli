package k8s

import (
	"bytes"
	"cli/src/common"
	"fmt"
	"os"

	xlogger "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

type K8sApiServer struct{}

func (K *K8sApiServer) initConfigInPod() (*rest.Config, *kubernetes.Clientset) {
	EnvConfig, EnvConfig_Err := rest.InClusterConfig()
	if EnvConfig_Err != nil {
		xlogger.Error(common.Rsperror(
			common.XE_CONFIG_ERROR,
			EnvConfig_Err.Error()))
		os.Exit(1)
	}
	ClientSet, ClientSet_Err := kubernetes.NewForConfig(EnvConfig)
	if ClientSet_Err != nil {
		xlogger.Error(common.Rsperror(
			common.XE_LOADCONFIG_ERROR,
			ClientSet_Err.Error()))
		os.Exit(2)
	}
	// kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	// config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	return EnvConfig, ClientSet
}

func (K *K8sApiServer) initCoreV1(podname, container, namespace string, command []string) remotecommand.Executor {
	config, clientset := K.initConfigInPod()
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

func (K *K8sApiServer) Run_exec(podname, container, namespace string, command []string) {
	var stdout, stderr bytes.Buffer
	exec := K.initCoreV1(podname, container, namespace, command)
	stream_err := exec.Stream(remotecommand.StreamOptions{
		Stdin:  nil,
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if stream_err != nil {
		xlogger.Error(common.Rsperror(
			common.XE_EXEC_ERROR,
			stream_err.Error()))
		os.Exit(2)

	}
	fmt.Println(stdout.String())
}
