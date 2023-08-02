# Simple usage with a mounted data directory:
# > docker build -t adv .
#
# > docker run -it -p 26657:26657 -v ~/.advd:/root/.advd adv advd init testing
# > docker run -it -p 26657:26657 -v ~/.advd:/root/.advd adv advd start
FROM golang:alpine AS build-env

# Install minimum necessary dependencies,
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/ewagmig/adv

# Add source files
COPY . .

# build adv
RUN make install


# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/adv /usr/bin/adv
COPY --from=build-env /go/bin/billboardcli /usr/bin/billboardcli

# Run advd by default, omit entrypoint to ease using container with advcli
CMD ["adv"]
