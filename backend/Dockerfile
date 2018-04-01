# Use build container
FROM golang:1.10 as devserver
# Workdir is path in your docker image from where all your commands will be executed
WORKDIR /go/src/github.com/budougumi0617/react-golang/backend
# Add all from your project inside workdir of docker image
ADD . /go/src/github.com/budougumi0617/react-golang/backend
# Then run your script to install dependencies and build application
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure
# Expose port for access to your app outside of container
EXPOSE 8080

# Starting server
CMD [ "go", "run", "./main.go" ]

# How to use this Dockerfile by docker command
# docker build -f ./Dockerfile -t apiserver --no-cache=true .
# docker run --rm --name exec-apiserver -p 8888:8080 -it apiserver:latest
# curl http://localhost:8888
