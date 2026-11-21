# Basic Usage Example Sources

## 2D

This example plots one-dimensional noise as a 2D line graph. It uses Go 1.22 or later. The example module declares Gonum Plot and uses the go-noise source from this checkout. The custom noise plot gives the same value each time you evaluate the same seed and coordinates.

Run from the repository root:

```sh
(cd _example/2d && go run -mod=readonly .)
```

The command writes `2d_perlin.png`, `2d_opensimplex.png`, and `2d_pseudorandom.png` in `_example/2d`. It overwrites any existing files with these names.

## 3D

Example of 2 dimmension image with time lapse as the 3rd dimmension.

- Run: `cd ./_example/3d && go run .`

## Advanced Configuration

By default the `alpha` and `beta` values of Perling noise are set to `2.0` and the iteration to `3`.

To change them, overwrite the `Smoothness` and `Scale` field values of the object as follows:

```go
alpha := 1.5
beta := 1.5

n, err := noise.New(noise.Perlin, seed)
n.Smoothness = alpha
n.Scale = beta
```
