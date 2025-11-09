package result

import (
	model "github.com/dooray-go/dooray/openapi/model/project"
	"github.com/zbum/klo"
	"os"
)

func PrintProjectsResult(result *model.GetProjectsResponse) error {
	specs := &klo.Specs{
		DefaultColumnSpec: "ID:{.id},CODE:{.code},TYPE:{.type},SCOPE:{.scope},STATE:{.state}",
		WideColumnSpec:    "CODE:{.code},TYPE:{.type},SCOPE:{.scope},DESCRIPTION:{.description}",
	}
	prn, err := klo.PrinterFromFlag("", specs)

	if err != nil {
		return err
	}

	table, err := klo.NewSortingPrinter("{.id}", prn)
	if err != nil {
		panic(err)
	}
	err = table.Fprint(os.Stdout, result.Result)
	if err != nil {
		return err
	}
	return nil
}
