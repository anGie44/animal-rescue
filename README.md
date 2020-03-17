# animal-rescue

To run the server on your system:

1. Create a `.env` file with your Auth0 Credentials (only includes vars referenced in `main.go`) [Sign up](https://auth0.com) for an account for free if you don't have one.
2. Update the `auth0-variables.js` file in `static/js` with your Auth0 Credentials.
3. Add `http://localhost:3000` to your Allowed Callback, and Allowed Logout URL's in your [Auth0 Management Dashboard](https://manage.auth0.com)
4. Make sure you have [dep](https://github.com/golang/dep) installed
5. Run `dep ensure` to install dependencies 
6. Run `go build` to create the binary (`animal-rescue`)
7. Run the binary : `./animal-rescue`

To run tests:

2. Run `dep ensure` to install dependencies
2. Run `go test ./...`
