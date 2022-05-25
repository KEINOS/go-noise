# CONTRIBUTING

## Pull Request

**Any PR for improvement is welcome!**

We will merge it as soon as it passes the CIs and not a prank-kind implementation. ;-)

- PR Branch: `main`
    - It is recommended to do a "[Draft-PR](https://github.blog/2019-02-14-introducing-draft-pull-requests/)" before the actual implementation if the fix is big. However, feel free to discard it as well!
- CI/CD: [Github Actions](./.github/workflows)
    - `go test ./...`
    - `golangci-lint run`
    - `golint ./...`
    - Code coverage check: 100% of coverage.

## Bug reports

- [Issues](https://github.com/KEINOS/go-noise/issues)
- If possible, please attach a simple reproduction code sample of the error. PRs for the fixes are much appreciated. 🙏
