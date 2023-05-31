package command

import (
	"cli/src/client/k8s"
	"cli/src/common"
	"net"
	"strings"

	xlogger "github.com/sirupsen/logrus"
)

type PodCommand struct {
	QueryCondition string
	ContainerName  string
	NameSpace      string
	Mode           string
	Index          int
}

func (P *PodCommand) Exec(args []string) {
	podName := ""
	containerName := ""
	k8sApi := k8s.K8sApiServer{}
	initK8sConn := k8sApi.InitConfigInPod(P.Mode)
	if net.ParseIP(P.QueryCondition) == nil {
		podNameList := k8s.GetContainerBelongToAllPod(initK8sConn, P.NameSpace, P.QueryCondition)
		if !common.CheckIndexIsExceedListLen(P.Index, podNameList) || P.Index == 0 {
			xlogger.Fatalf("index: [%d] out of range, choose one of: [%s]", P.Index, common.ListSpecialFmt(podNameList))
		}
		podName = podNameList[P.Index-1]
		xlogger.Infof("specified for pod %s, match all pod:%s", podName, podNameList)
		containerName = P.QueryCondition

	} else {
		containerNameList := []string{}
		podName, containerNameList = k8s.GetContainerIPBelongToAllPod(initK8sConn, P.NameSpace, P.QueryCondition)
		if P.ContainerName == "Defaulted container" {
			containerName = containerNameList[0]
			xlogger.Infof("Defaulted container \"%s\" out of: %s", containerName, strings.Join(containerNameList, ", "))
		} else {
			containerName = P.ContainerName
		}

	}
	k8sApi.Run_exec(initK8sConn, podName, containerName, P.NameSpace, args)
}
