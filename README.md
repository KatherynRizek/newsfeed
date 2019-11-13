# markdown-generator

## Build the app
* Build the app with `docker build -t go-docker-dev .`
* Go into the docker container with `docker run --rm -it -p 8080:8080 -v $(pwd):/go/src/app \ markdown-generator bash`
* When in the container, run the app with `src/main.go`

## Dependencies
* New app dependencies can be added in the `glide.yaml` file
* To update dependcies after added in `glide.yml`, run `glide udpate` within the container

## Credits
* The Dockerfile was used from https://blog.hasura.io/the-ultimate-guide-to-writing-dockerfiles-for-go-web-apps-336efad7012c/
* The app is being built from this tutorial: https://codegangsta.gitbooks.io/building-web-apps-with-go/content/index.html
