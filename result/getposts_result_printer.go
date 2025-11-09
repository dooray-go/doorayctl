package result

import (
	model "github.com/dooray-go/dooray/openapi/model/project"
	"github.com/zbum/klo"
	"os"
)

func PrintPostsResult(result *model.GetPostsResponse) error {
	specs := &klo.Specs{
		DefaultColumnSpec: "NUMBER:{.number},TASK_NUMBER:{.taskNumber},SUBJECT:{.subject},WORKFLOW:{.workflowClass},CLOSED:{.closed},CREATED_AT:{.createdAt}",
		WideColumnSpec:    "TASK_NUMBER:{.taskNumber},SUBJECT:{.subject},WORKFLOW:{.workflowClass},PRIORITY:{.priority},DUE_DATE:{.dueDate}",
	}
	prn, err := klo.PrinterFromFlag("", specs)

	if err != nil {
		return err
	}

	table, err := klo.NewSortingPrinter("{.number}", prn)
	if err != nil {
		panic(err)
	}
	err = table.Fprint(os.Stdout, result.Result)
	if err != nil {
		return err
	}
	return nil
}