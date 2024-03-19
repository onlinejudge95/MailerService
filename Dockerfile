FROM golang:1.20 as build

WORKDIR /code

COPY go.mod go.sum ./

COPY . .

RUN go build -tags jsoniter -o mailer

FROM public.ecr.aws/lambda/provided:al2023

WORKDIR /code

COPY --from=build /code/mailer ./mailer

ENTRYPOINT [ "./mailer" ]
