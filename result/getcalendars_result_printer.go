package result

import (
	model "github.com/dooray-go/dooray/openapi/model/calendar"
	"github.com/thediveo/klo"
	"os"
)

func PrintCalendarsResult(result *model.GetCalendarsResponse) error {
	specs := &klo.Specs{
		DefaultColumnSpec: "ID:{.id},NAME:{.name},TYPE:{.type},CREATED_AT:{.createdAt},OWNER:{.ownerOrganizationMemberId}",
		WideColumnSpec:    "NAME:{.name},CREATED_AT:{.createdAt}",
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
