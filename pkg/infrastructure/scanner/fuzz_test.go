package scanner

import (
	"encoding/json"
	"testing"
)

func FuzzFilterTags(f *testing.F) {
	// Seed corpus
	f.Add("latest")
	f.Add("3.12")
	f.Add("20.14-nonroot")
	f.Add("3.0.20240824")
	f.Add("arm64-debug")
	f.Add("v1.0.0-beta.1")
	f.Add("")
	f.Add("!@#$%^&*()")

	cfg := DefaultTagFilter()

	f.Fuzz(func(t *testing.T, tag string) {
		// FilterTags must not panic on any input
		result := FilterTags([]string{tag}, cfg)
		if len(result) > 1 {
			t.Errorf("FilterTags returned more tags than input: got %d", len(result))
		}
	})
}

func FuzzParseImagePatterns(f *testing.F) {
	f.Add("azurelinux/base/python")
	f.Add("docker.io/library/python:3-slim")
	f.Add("mcr.microsoft.com/dotnet/aspnet:8.0")
	f.Add("")
	f.Add("repo:with:multiple:colons")
	f.Add("a/b/c/d/e/f")

	f.Fuzz(func(t *testing.T, pattern string) {
		// ParseImagePatterns must not panic on any input
		repos, singles := ParseImagePatterns([]string{pattern})
		total := len(repos) + len(singles)
		if pattern != "" && total == 0 {
			t.Errorf("non-empty pattern %q produced no output", pattern)
		}
	})
}

func FuzzBuildFullImageName(f *testing.F) {
	f.Add("mcr.microsoft.com", "azurelinux/base/python", "3.12")
	f.Add("", "docker.io/library/node", "20-slim")
	f.Add("registry.example.com", "repo", "latest")

	f.Fuzz(func(t *testing.T, registry, repo, tag string) {
		// BuildFullImageName must not panic on any input
		result := BuildFullImageName(registry, repo, tag)
		if result == "" {
			t.Error("BuildFullImageName returned empty string")
		}
	})
}

func FuzzCleanVersion(f *testing.F) {
	f.Add("3.12.9")
	f.Add("3.12.9-1.azl3")
	f.Add("21.0.10+7-LTS")
	f.Add("")
	f.Add("not-a-version")

	f.Fuzz(func(t *testing.T, version string) {
		// cleanVersion must not panic on any input
		_ = cleanVersion(version)
	})
}

func FuzzParseSyftResult(f *testing.F) {
	f.Add(`{"artifacts":[]}`)
	f.Add(`{"artifacts":[{"name":"python","version":"3.12","type":"binary","metadata":{}}]}`)
	f.Add(`{}`)
	f.Add(``)
	f.Add(`{"artifacts": null}`)

	f.Fuzz(func(t *testing.T, jsonData string) {
		// parseSyftResult must not panic on any JSON input
		var output syftOutput
		if err := json.Unmarshal([]byte(jsonData), &output); err == nil {
			_ = parseSyftResult(&output)
		}
	})
}
