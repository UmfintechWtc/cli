package command

import "cli/src/client/k8s"

type PodCommand struct {
	PodName   string
	Container string
	NameSpace string
	Mode      string
}

func (P *PodCommand) Exec(args []string) {
	k8sApi := k8s.K8sApiServer{}
	k8sApi.Run_exec(P.PodName, P.Container, P.NameSpace, P.Mode, args)
}
