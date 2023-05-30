package command

import "cli/src/client/k8s"

type PodCommand struct {
	Container string
	NameSpace string
	Mode      string
	Index     int
}

func (P *PodCommand) Exec(args []string) {
	k8sApi := k8s.K8sApiServer{}
	k8sApi.Run_exec(P.Container, P.NameSpace, P.Mode, P.Index, args)
}
