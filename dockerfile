FROM golang:1.14

WORKDIR /go/src/home-recipes
COPY . .

RUN go install -v ./cmd/homerecipes/.

CMD ["homerecipes"]