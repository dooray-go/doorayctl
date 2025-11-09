package result

import (
	model "github.com/dooray-go/dooray/openapi/model/messenger"
	"github.com/zbum/klo"
	"os"
)

func PrintMessengerResult(result *model.DirectSendResponse) error {
	prn, err := klo.PrinterFromFlag("",
		&klo.Specs{DefaultColumnSpec: "ID:{.result.id},SUCCESS:{.header.isSuccessful}"})

	if err != nil {
		return err
	}

	table, err := klo.NewSortingPrinter("{.id}", prn)
	if err != nil {
		panic(err)
	}
	err = table.Fprint(os.Stdout, result)
	if err != nil {
		return err
	}
	return nil
}
