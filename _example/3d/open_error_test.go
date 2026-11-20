package main

import (
	"image"
	"image/color"
	"image/gif"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveImgGIFPreservesOpenError(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "missing", "output.gif")
	frame := image.NewPaletted(image.Rect(0, 0, 1, 1), color.Palette{color.Black, color.White})
	animation := &gif.GIF{
		Image: []*image.Paletted{frame},
		Delay: []int{0},
	}

	err := SaveImgGIF(animation, path)
	require.ErrorIs(t, err, os.ErrNotExist)

	var pathErr *os.PathError
	require.ErrorAs(t, err, &pathErr)
	require.Equal(t, "open", pathErr.Op)
	require.Equal(t, path, pathErr.Path)
}
