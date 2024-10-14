# travel-lists

### Development setup

Install the required version of Go and other dependencies using [asdf](https://asdf-vm.com/).

```sh
    asdf plugin add golang https://github.com/asdf-community/asdf-golang
    asdf plugin add nodejs

    # Add the necessary configuration to your Shell configuration: https://github.com/asdf-community/asdf-golang?tab=readme-ov-file#goroot

    asdf install
```

After installing Go, you might want to install the Go language server [gopls](https://pkg.go.dev/golang.org/x/tools/gopls) for smart autocompletions:

```sh
    go install -v golang.org/x/tools/gopls@latest
    asdf reshim golang
```

If you want to automatically recompile the code on code changes, you can use [air](https://github.com/air-verse/air).

```sh
    go install github.com/air-verse/air@latest
    asdf reshim golang

    air
```
