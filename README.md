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

We need to install `templ` globally to be able to generate our templates:

```sh
    go install github.com/a-h/templ/cmd/templ@latest
    asdf reshim golang
```

Because we don't want to manage JS dependencies by hand, we use [NPM](https://www.npmjs.com/) to manage them and [Vite](https://v2.vitejs.dev/) to build them into a single `web/assets/dist/index.min.js` file, that we can then import in our `Templ` templates.

To build the JavaScript dependencies, you can use:

```sh
    cd assets
    npm install
    npx vite build
```

To start the server which automatically recompiles the code on code changes, use [air](https://github.com/air-verse/air). Air will also build all our `JavaScript`, `CSS` and `Templ` templates and keep them up to date during development:

```sh
    go install github.com/air-verse/air@latest
    asdf reshim golang

    air
```
