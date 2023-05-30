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

func (K *K8sApiServer) Run_exec(containerName, namespace, mode string, podIndex int, command []string) {
	var stdout, stderr bytes.Buffer
	initK8sConn := K.initConfigInPod(mode)
	targetPodList := GetContainerBelongToAllPod(initK8sConn, namespace, containerName)
	if !common.CheckIndexIsExceedListLen(podIndex, targetPodList) {
		xlogger.Fatalf("执行的第%d个Pod超过匹配结果,请确认:[%s]", podIndex, targetPodList)
	}
	exec := initCoreV1(initK8sConn, targetPodList[podIndex], containerName, namespace, command)
	stream_err := exec.Stream(remotecommand.StreamOptions{
		Stdin:  nil,
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if stream_err != nil {
		fmt_message := fmt.Sprintf("%s[%s]:%s\n", targetPodList[podIndex], containerName, command)
		xlogger.Error(common.Rsperror(
			common.XE_EXEC_ERROR,
			fmt_message+stream_err.Error()))
		os.Exit(2)
	}
	xlogger.Infof("容器[%s]匹配到多个Pod:%s\n执行的第%d个Pod[%s]，命令输出: \n%s", containerName, targetPodList, podIndex+1, targetPodList[podIndex], stdout.String())
}
