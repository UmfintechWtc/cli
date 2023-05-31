package k8s

import (
	"bytes"
	"cli/src/common"
	"fmt"
	"os"

	xlogger "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func initCoreV1(initK8sConn *K8sApiServer, podname, containerName, namespace string, command []string) remotecommand.Executor {
	kubernetesReq := initK8sConn.K8sClient.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podname).
		Namespace(namespace).
		SubResource("exec")
	kubernetesReq.VersionedParams(
		&v1.PodExecOptions{
			Container: containerName,
			Command:   command,
			Stdin:     false,
			Stdout:    true,
			Stderr:    true,
			TTY:       false,
		},
		scheme.ParameterCodec,
	)
	executorApi, executor_err := remotecommand.NewSPDYExecutor(initK8sConn.K8sConfig, "POST", kubernetesReq.URL())
	if executor_err != nil {
		xlogger.Error(common.Rsperror(
			common.XE_COREV1_ERROR,
			executor_err.Error()))
		os.Exit(2)
	}
	return executorApi
}

func (K *K8sApiServer) Run_exec(initK8sConn *K8sApiServer, podName, containerName, namespace string, command []string) {
	var stdout, stderr bytes.Buffer
	exec := initCoreV1(initK8sConn, podName, containerName, namespace, command)
	stream_err := exec.Stream(remotecommand.StreamOptions{
		Stdin:  nil,
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if stream_err != nil {
		fmt_message := fmt.Sprintf("%s[%s]:%s\n", podName, containerName, command)
		xlogger.Error(common.Rsperror(
			common.XE_EXEC_ERROR,
			fmt_message+stream_err.Error()))
		os.Exit(2)
	}
	xlogger.Infof("namespace:[%s] -> Pod:[%s] -> container:[%s] \n命令输出: \n%s", namespace, podName, containerName, stdout.String())
}
