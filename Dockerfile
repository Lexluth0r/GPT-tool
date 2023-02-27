FROM node:19.6.1-alpine3.17 as frontend

WORKDIR /app/resource

COPY resource/package.json resource/package-lock.json /app/resource/

ADD resource/package.json .
ADD resource/yarn.lock .

RUN yarn install

COPY resource/ /app/resource

RUN yarn run build

FROM golang:1.17 as backend

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /usr/share

ADD go.mod .
ADD go.sum .
COPY . .
RUN go mod download -x
RUN go build -o chat cmd/chat.go

FROM centos:7
WORKDIR /usr/local/bin/

COPY --from=frontend /app/resource/dist/ resource/dist/
COPY --from=backend /usr/share/chat .


CMD ["./chat"]


