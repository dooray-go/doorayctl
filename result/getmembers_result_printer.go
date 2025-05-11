package result

import (
	model "github.com/dooray-go/dooray/openapi/model/account"
	"github.com/thediveo/klo"
	"os"
)

func PrintAccountResult(result *model.GetMembersResponse) error {
	prn, err := klo.PrinterFromFlag("",
		&klo.Specs{DefaultColumnSpec: "ID:{.id},NAME:{.name},EXTERNAL_EMAIL:{.externalEmailAddress}"})

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
