package report

import "testing"

func FuzzHumanSize(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(1024))
	f.Add(int64(1048576))
	f.Add(int64(1073741824))
	f.Add(int64(-1))
	f.Add(int64(9999999999999))

	f.Fuzz(func(t *testing.T, numBytes int64) {
		result := HumanSize(numBytes)
		if result == "" {
			t.Error("HumanSize returned empty string")
		}
	})
}

func FuzzFormatDigest(f *testing.F) {
	f.Add("sha256:abcdef1234567890abcdef1234567890")
	f.Add("")
	f.Add("sha256:short")
	f.Add("not-a-digest")
	f.Add("sha256:")

	f.Fuzz(func(t *testing.T, digest string) {
		// FormatDigest must not panic on any input
		_ = FormatDigest(digest)
	})
}
