package main

import (
	"flag"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/slmcmahon/go-azdo"
	slmcommon "github.com/slmcmahon/go-common"
)

func main() {
	var (
		patFlag     string
		orgFlag     string
		projectFlag string
		libFlag     int
	)

	flag.StringVar(&patFlag, "pat", "", "Personal Access Token")
	flag.StringVar(&orgFlag, "org", "", "Azure Devops Organization")
	flag.StringVar(&projectFlag, "project", "", "Azure DevOps Project")
	flag.IntVar(&libFlag, "lib", 0, "Azure Devops Library ID")

	flag.Parse()

	pat, err := slmcommon.CheckEnvOrFlag(patFlag, "AZDO_PAT")
	if err != nil {
		log.Fatal(err)
	}
	org, err := slmcommon.CheckEnvOrFlag(orgFlag, "AZDO_ORG")
	if err != nil {
		log.Fatal(err)
	}
	project, err := slmcommon.CheckEnvOrFlag(projectFlag, "AZDO_PROJECT")
	if err != nil {
		log.Fatal(err)
	}

	if libFlag == 0 {
		log.Fatal("No library value was specified")
	}

	ops := azdo.NewAZDOOperations(pat, org, project)

	libData, err := ops.GetVariableLibraries(libFlag)
	if err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Value"})

	for varName, variable := range libData[0].Variables {
		table.Append([]string{varName, variable.Value})
	}
	table.Render()
}
