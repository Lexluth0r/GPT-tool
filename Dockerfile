FROM node:alpine as frontend

WORKDIR /app

COPY resource/package.json resource/package-lock.json /app/resource/

RUN set -x ; cd /app/resource \
&& npm install --registry=https://registry.npmmirror.com

COPY resource/ /app/resource/

RUN set -x ; cd /app/resource/ \
&& npm run build


FROM golang:1.17 as build

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app/

ADD go.mod .
ADD go.sum .
COPY . .
RUN go mod download -x
RUN go build -o chat cmd/chat.go


FROM centos:7
WORKDIR /usr/local/bin/

COPY --from=frontend /app/resource/dist/ /resource/dist/
COPY --from=build /app/chat .

CMD ["./chat"]
