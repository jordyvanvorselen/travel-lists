# travel-lists

## Local development

### Installing dependencies

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

We also need to install `sql-migrate` to apply our database migrations:

```sh
    go install github.com/rubenv/sql-migrate/...@latest
    asdf reshim golang
```

Now, we need to create a database and run the database migrations. We use `docker` and `docker-compose` for that. Install both of them and run the following commands to create your database:

```sh
    docker compose up

    sql-migrate up
```

So, because we don't want to write all our SQL ourselves, we will install `sqlboiler`:

```sh
    go install github.com/volatiletech/sqlboiler/v4@latest
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
```

To manage JS dependencies, we use [NPM](https://www.npmjs.com/) to manage them and [Vite](https://v2.vitejs.dev/) to build them into a single `web/assets/dist/index.min.js` file, that we can then import in our `Templ` templates.

To build the JavaScript dependencies, you can use:

```sh
    cd assets
    npm install
    npx vite build
```

To set up the local database connection, we need to set some env variables. Create a `.env` by copying the `env.sample` and setting the correct values:

```sh
    cp env.sample .env
```

To start the server which automatically recompiles the code on code changes, use [air](https://github.com/air-verse/air). Air will also build all our `JavaScript`, `CSS`, `sqlboiler` files and `Templ` templates and keep them up to date during development:

```sh
    go install github.com/air-verse/air@latest
    asdf reshim golang

    air
```

### Running the application

You can run the application and the needed dependencies using:

```sh
    docker compose up --detach
    
    sql-migrate up

    air
```
