package result

import (
	model "github.com/dooray-go/dooray/openapi/model/project"
	"testing"
)

func TestPrintPostsResult(t *testing.T) {
	type args struct {
		result *model.GetPostsResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestPrintPostsResult",
			args: args{
				result: &model.GetPostsResponse{
					Result: []model.PostInfo{
						{
							ID:            "1",
							Subject:       "Test Post 1",
							TaskNumber:    "PROJECT-1",
							Closed:        false,
							CreatedAt:     "2023-10-01T00:00:00+09:00",
							WorkflowClass: "working",
							Number:        1,
							Priority:      "high",
							DueDate:       "2023-10-15T00:00:00+09:00",
						},
						{
							ID:            "2",
							Subject:       "Test Post 2",
							TaskNumber:    "PROJECT-2",
							Closed:        true,
							CreatedAt:     "2023-10-02T00:00:00+09:00",
							WorkflowClass: "closed",
							Number:        2,
							Priority:      "medium",
							DueDate:       "2023-10-20T00:00:00+09:00",
						},
					},
					TotalCount: 2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrintPostsResult(tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("PrintPostsResult() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}