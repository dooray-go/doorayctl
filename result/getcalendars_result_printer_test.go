package result

import (
	model "github.com/dooray-go/dooray/openapi/model/calendar"
	"testing"
)

func TestPrintCalendarsResult(t *testing.T) {
	type args struct {
		result *model.GetCalendarsResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestPrintCalendarsResult",
			args: args{
				result: &model.GetCalendarsResponse{
					Result: []model.Calendar{
						{
							ID:                        "1",
							Name:                      "Test Calendar",
							Type:                      "private",
							CreatedAt:                 "2023-10-01T00:00:00+09:00",
							OwnerOrganizationMemberID: "123123213",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrintCalendarsResult(tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("PrintCalendarsResult() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
