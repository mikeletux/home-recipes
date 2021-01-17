FROM golang:1.14

WORKDIR /go/src/home-recipes
COPY . .

RUN go install -v ./cmd/homerecipes/. \
    && mkdir -p /var/home-recipes

CMD ["homerecipes"]