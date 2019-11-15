# markdown-generator

## Build the app
* Make sure there are `glide.yaml` and `glide.lock` files present at the same level as the Dockerfile; they can be empty
* Build the app with `docker build -t <name> .`
* Go into the docker container with `docker run --rm -it -p 8080:8080 -v $(pwd):/go/src/app \` and then `<name> bash`
* When in the container, run the app with `go run <file>`. For example, if you would like to run the `main.go` file under the `src` directory, use the command `go run src/main.go`

## Dependencies
* New app dependencies can be added in the `glide.yaml` file
* To update dependcies after added in `glide.yml`, run `glide udpate` within the container

## Credits
* The Dockerfile was used from https://blog.hasura.io/the-ultimate-guide-to-writing-dockerfiles-for-go-web-apps-336efad7012c/
* The app is being built from this tutorial: https://freshman.tech/web-development-with-go/
