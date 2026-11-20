package main

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveImgGIFOverwritesShorterFile(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "output.gif")
	frame := image.NewPaletted(image.Rect(0, 0, 1, 1), color.Palette{color.Black, color.White})
	animation := &gif.GIF{
		Image: []*image.Paletted{frame, frame},
		Delay: []int{0, 0},
	}

	var longGIF bytes.Buffer
	require.NoError(t, gif.EncodeAll(&longGIF, animation))
	require.NoError(t, SaveImgGIF(animation, path))

	got, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, longGIF.Bytes(), got)

	animation.Image = animation.Image[:1]
	animation.Delay = animation.Delay[:1]

	var shortGIF bytes.Buffer
	require.NoError(t, gif.EncodeAll(&shortGIF, animation))
	require.Greater(t, longGIF.Len(), shortGIF.Len())
	require.NoError(t, SaveImgGIF(animation, path))

	got, err = os.ReadFile(path)
	require.NoError(t, err)
	require.Len(t, got, shortGIF.Len())
	require.Equal(t, shortGIF.Bytes(), got)
}
