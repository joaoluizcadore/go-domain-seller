package utils_test

import (
	"testing"

	"github.com/joaoluizcadore/domain-seller/internal/utils"
)

func TestCleanDomain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		domain string
		want   string
	}{
		{
			name:   "domain with www",
			domain: "www.google.com",
			want:   "google.com",
		},
		{
			name:   "domain without www",
			domain: "google.com",
			want:   "google.com",
		},
		{
			name:   "domain with www and path",
			domain: "www.google.com.br/path",
			want:   "google.com.br",
		},
		{
			name:   "domain without www and path",
			domain: "google.com.br/path",
			want:   "google.com.br",
		},
		{
			name:   "domain with http",
			domain: "http://google.com.br/path",
			want:   "google.com.br",
		},
		{
			name:   "domain with https",
			domain: "https://google.com.br/path",
			want:   "google.com.br",
		},
		{
			name:   "domain with http and www",
			domain: "http://www.google.com.br/path",
			want:   "google.com.br",
		},
		{
			name:   "domain with https and www",
			domain: "https://www.google.com.br/path",
			want:   "google.com.br",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.CleanDomain(tt.domain); got != tt.want {
				t.Errorf("CleanDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
