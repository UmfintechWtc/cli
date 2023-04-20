package command

import "cli/src/client/k8s"

type CliCommand struct {
	PodName   string
	Container string
	NameSpace string
}

func (C *CliCommand) Exec(args []string) {
	// fmt.Println(C.Container)
	// fmt.Println(C.NameSpace)
	// fmt.Println(C.PodName)
	// fmt.Println(args)
	k8sApi := k8s.K8sApiServer{}
	k8sApi.Run_exec(C.PodName, C.Container, C.NameSpace, args)
}
