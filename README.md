# mountebank-proxy-to-github

Using Go to make requests to the Github API through a mountebank proxy which allows for "caching" to improve testing performance and feedback loop.

# setup

## last run with the following

- node: `14.16.0`
- npm: `7.6.0`
- go: `1.16`

1. install nodejs packages via `npm install` (note it runs patch-package as well)
2. adjust Github PAT in `main.go` if needed, this needs the permission for `repo: public_repo`.

# setup and running

1. start mountebank proxy server with `npm run mb:configed`, leave this running
2. visit mountebank web-ui at `localhost:2525` and check the imposters page (link in the top nav)
3. run the go script in a separate terminal using `go run main.go --proxy 9999` to return the mountebank readme
4. see the results in the web-ui at `http://localhost:2525/imposters/9999`
