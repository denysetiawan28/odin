FROM golang:alpine as builder

#get git
RUN apk update && \
apk upgrade && \
apk add --no-cache git && \
apk add --no-cache tzdata

#add user for golang and maintainer
#RUN addgroup -S golang && adduser -S golang -G golang
#USER golang:golang
MAINTAINER titipaja.id

#working directory
ADD . /opt/odin
WORKDIR /opt/odin

#copy resource
COPY . .

#building
#RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o odin-server
#RUN chmod a+x tokyo-server-final

FROM scratch

COPY --from=builder /opt/odin/resources /opt/odin/resources
COPY --from=builder /opt/odin/odin-server /opt/odin/odin-server

WORKDIR /opt/odin

#expose network
EXPOSE 9080

#running
CMD [ "/opt/odin/odin-server" ]