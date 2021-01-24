FROM golang:1.14

WORKDIR /go/src/home-recipes
COPY . .

RUN mkdir -p /go/home-recipes/storage \
    && go build -o /go/home-recipes/home-recipes /go/src/home-recipes/cmd/homerecipes/.

EXPOSE 8080/tcp

ENTRYPOINT ["/go/home-recipes/home-recipes"]