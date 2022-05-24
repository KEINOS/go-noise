package opensimplex_test

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/KEINOS/go-noise/pkg/opensimplex"
	"github.com/stretchr/testify/require"
)

var samples [][]float64

func loadSamples(t *testing.T) [][]float64 {
	t.Helper()

	if len(samples) > 0 {
		return samples
	}

	f, err := os.Open("../../testdata/opensimplex.json")
	require.NoError(t, err)

	defer f.Close()

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

func TestSamplesMatch(t *testing.T) {
	samples := loadSamples(t)

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

		if expected != actual {
			t.Fatalf("Expected %v, got %v for %dD sample at %v",
				expected, actual, len(s)-1, s[:len(s)-1])
		}
	}
}
