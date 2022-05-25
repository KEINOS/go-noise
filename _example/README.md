# Basic Usage Example Sources

## 2D

Example of 2D noise generation.

- Run: `cd ./_example/2d && go run .`

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
