FROM alpine

RUN apk update && apk add flite
ADD kubehostname /kubehostname
ENTRYPOINT ["/kubehostname"]