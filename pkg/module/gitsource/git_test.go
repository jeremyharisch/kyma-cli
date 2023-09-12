package gitsource

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"testing"
)

func Test_fetchRepoURLFromRemotes(t *testing.T) {
	type args struct {
		gitRemotes []*git.Remote
		remoteName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Should return correct GitHub URL",
			args: args{
				gitRemotes: createTestRemotes(),
				remoteName: "upstream",
			},
			want:    "github.com/kyma-test/test",
			wantErr: false,
		},
		{name: "Should return return error due remote not existing",
			args: args{
				gitRemotes: createTestRemotes(),
				remoteName: "non-existent",
			},
			want:    "",
			wantErr: true,
		},
		{name: "Should return return error due invalid URL",
			args: args{
				gitRemotes: createTestRemotes(),
				remoteName: "invalidURL",
			},
			want:    "",
			wantErr: true,
		},
		{name: "Should return return error due to empty remotes",
			args: args{
				gitRemotes: []*git.Remote{},
				remoteName: "upstream",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchRepoURLFromRemotes(tt.args.gitRemotes, tt.args.remoteName)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchRepoURLFromRemotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fetchRepoURLFromRemotes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func createTestRemotes() []*git.Remote {
	return []*git.Remote{
		git.NewRemote(nil, &config.RemoteConfig{
			Name: "upstream",
			URLs: []string{"github.com/kyma-test/test"},
		}),
		git.NewRemote(nil, &config.RemoteConfig{
			Name: "origin",
			URLs: []string{"github.com/user-test/test"},
		}),
		git.NewRemote(nil, &config.RemoteConfig{
			Name: "invalidURL",
			URLs: []string{"\t"},
		}),
	}

}
