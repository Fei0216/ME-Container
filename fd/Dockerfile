FROM gbaeke/gocv-4.0.0-build as build       
ENV GOPATH /go
ENV PATH $GOPATH/src:$GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH
RUN go get -u -d gocv.io/x/gocv       
      
                
FROM gbaeke/gocv-4.0.0-run       
      
COPY fd.go /go/fd.go
COPY haarcascade_frontalface_alt.xml /go/haarcascade_frontalface_alt.xml
COPY test.jpeg /go/test.jpeg

CMD ["go","run", "fd.go"]

EXPOSE 8180