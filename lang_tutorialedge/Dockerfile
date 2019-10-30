# Base image specification
FROM golang:1.12.0-alpine3.9

# Create an /app folder
# RUN mkdir /app

# Copy current directory to app ?
COPY . /go/src/github.com/al-un/tutos-go/lang_tutorialedge

# Working directory?
WORKDIR /go/src/github.com/al-un/tutos-go/lang_tutorialedge

# # Where am I?
# RUN pwd
# # Where is it?
# RUN echo $GOPATH
# # what's
# RUN ls -l
# Build: main.go must serve an HTTP stuff
RUN go build -o main-simple-web-server .

# Run!
CMD ["/go/src/github.com/al-un/tutos-go/lang_tutorialedge/main-simple-web-server"]