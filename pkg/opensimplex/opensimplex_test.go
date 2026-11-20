//nolint:varnamelen // allow short variable names due to test readability
package opensimplex_test

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/KEINOS/go-noise/pkg/opensimplex"
	"github.com/stretchr/testify/require"
)

func TestSamplesMatch(t *testing.T) {
	samples := loadSamples(t)

	// Fused multiply-add can round differently from the reference samples.
	// The largest observed absolute difference on darwin/arm64 was 3.71e-13.
	const sampleTolerance = 1e-12

	n := opensimplex.New(0)

	for _, s := range samples {
		var expected, actual float64

		switch len(s) {
		case 3:
			expected = s[2]
			actual = n.Eval64(s[0], s[1])
		case 4:
			expected = s[3]
			actual = n.Eval64(s[0], s[1], s[2])
		case 5:
			continue
		default:
			t.Fatalf("Unexpected size sample: %d", len(s))
		}

		require.InDelta(t, expected, actual, sampleTolerance,
			"%dD sample at %v", len(s)-1, s[:len(s)-1])
	}
}

// ----------------------------------------------------------------------------
//  Helper Functions
// ----------------------------------------------------------------------------

func loadSamples(t *testing.T) [][]float64 {
	t.Helper()

	var samples [][]float64

	f, err := os.Open("../../testdata/opensimplex.json")
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := f.Close(); err != nil {
			t.Errorf("close sample file: %v", err)
		}
	})

	dec := json.NewDecoder(f)

	for {
		var sample []float64

		err := dec.Decode(&sample)
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatal(err)
		}

		samples = append(samples, sample)
	}

	return samples
}
