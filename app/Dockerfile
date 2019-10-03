FROM node:8-buster
MAINTAINER James Vasile "james@opentechstrategies.com"

## Local web server should serve static files directly, not pass those requests
## to our app
ENV STATIC_URL /static
ENV STATIC_PATH /app/static

# Create app directory
WORKDIR /usr/src/app

ENV GOPATH="/usr/local/src/golang"
RUN mkdir -p $GOPATH

# Golang depends
ENV TERM="xterm"
RUN apt-get update
RUN apt-get install -y golang-go golang-go.tools
RUN go get github.com/gin-gonic/gin

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY package*.json ./

RUN npm install

COPY . .

RUN node_modules/browserify/bin/cmd.js -o backend/public/bundle.js frontend/main.js
ENTRYPOINT ["npm"]
#CMD ["start"]
CMD ["run", "backend"]

