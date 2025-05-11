package result

import (
	"github.com/dooray-go/dooray/openapi/model/account"
	"testing"
)

func TestPrintAccountResult(t *testing.T) {
	type args struct {
		result *account.GetMembersResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestPrintAccountResult",
			args: args{
				result: &account.GetMembersResponse{
					Result: []account.Member{
						{
							ID:                   "1",
							Name:                 "John Doe",
							ExternalEmailAddress: "johndo@test.com",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrintAccountResult(tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("PrintAccountResult() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
