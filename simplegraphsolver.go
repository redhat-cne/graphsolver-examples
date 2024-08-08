package main

import (
	"github.com/openshift/ptp-operator/test/pkg/client"
	lib "github.com/redhat-best-practices-for-k8s/graphsolver-lib"
	l2lib "github.com/redhat-best-practices-for-k8s/l2discovery-lib"
)

// Runs Solver to find optimal configurations
func main() {
	const (
		// problem/scenario name
		findOCProblemName = "OC"

		// unique id for each tag, e.g. solution role
		tagSlave       = 0
		tagGrandmaster = 1
	)

	// create an OC client
	client.Client = client.New("")

	// Initialize l2 library
	l2lib.GlobalL2DiscoveryConfig.SetL2Client(client.Client, client.Client.Config)

	// Collect L2 info
	config, err := l2lib.GlobalL2DiscoveryConfig.GetL2DiscoveryConfig(true)
	if err != nil {
		return
	}

	// initialize L2 config in solver
	lib.GlobalConfig.SetL2Config(config)

	// Initializing problems
	lib.GlobalConfig.InitProblem(
		findOCProblemName,
		[][][]int{
			{{int(lib.StepNil), 0, 0}},         // step1
			{{int(lib.StepSameLan2), 2, 0, 1}}, // step2
		},
		[]int{tagSlave: 0, tagGrandmaster: 1},
	)

	// Run solver for problem
	lib.GlobalConfig.Run(findOCProblemName)

	// print first solution
	lib.GlobalConfig.PrintAllSolutions()
}
