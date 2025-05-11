package result

import (
	"github.com/dooray-go/dooray/openapi/model"
	messengermodel "github.com/dooray-go/dooray/openapi/model/messenger"
	"testing"
)

func TestPrintMessengerResult(t *testing.T) {
	type args struct {
		result *messengermodel.DirectSendResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestPrintMessengerResult",
			args: args{
				result: &messengermodel.DirectSendResponse{
					Header: model.ResponseHeader{
						IsSuccessful: true,
					},
					Result: messengermodel.DirectSendResult{
						ID: 1,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrintMessengerResult(tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("PrintMessengerResult() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
