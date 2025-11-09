package result

import (
	model "github.com/dooray-go/dooray/openapi/model/project"
	"testing"
)

func TestPrintProjectsResult(t *testing.T) {
	type args struct {
		result *model.GetProjectsResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestPrintProjectsResult",
			args: args{
				result: &model.GetProjectsResponse{
					Result: []model.ProjectInfo{
						{
							ID:          "1",
							Code:        "techcenter",
							Description: "기술센터 업무용 프로젝트 입니다.",
							State:       "active",
							Scope:       "public",
							Type:        "project",
						},
						{
							ID:          "2",
							Code:        "engineering",
							Description: "Engineering team project",
							State:       "active",
							Scope:       "private",
							Type:        "project",
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
			if err := PrintProjectsResult(tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("PrintProjectsResult() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}