package models

import "testing"

func TestImageInfo_GetFullRepoURL(t *testing.T) {
	tests := []struct {
		name   string
		fields ImageInfo
		want   string
	}{
		{
			name: "Standard case with ip/port",
			fields: ImageInfo{
				Name:       "nginx",
				Tag:        "2.13.2",
				Repository: "192.168.1.10:5000",
			},
			want: "192.168.1.10:5000/nginx:2.13.2",
		},
		{
			name: "Standard case with domain",
			fields: ImageInfo{
				Name:       "nginx",
				Tag:        "latest",
				Repository: "docker.io",
			},
			want: "docker.io/nginx:latest",
		},
		{
			name: "Standard case with missing name",
			fields: ImageInfo{
				Name:       "",
				Tag:        "latest",
				Repository: "docker.io",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := &ImageInfo{
				Name:       tt.fields.Name,
				Tag:        tt.fields.Tag,
				Repository: tt.fields.Repository,
			}
			if got := info.GetFullRepoURL(); got != tt.want {
				t.Errorf("ImageInfo.GetFullRepoURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
