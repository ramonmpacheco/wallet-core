# PLEASE IGNORE MY COMMENTS BELOW, I DO THAT IN ORDE TO LEARN

# First step
FROM golang:1.19 as builder
# Need be here in order to avoid: 
# '"std" is not an importable package'
WORKDIR /go/src
# Copy all files from the host 
# into the docker image
COPY ../../ ./
# Build main.go file with the name 'app' into / folder
RUN go build -o /app ./cmd/walletcore
# Second step
FROM golang:1.19
# Install lib for kafka
RUN apt-get update && apt-get install -y librdkafka-dev
# Copy from the first step 
COPY --from=builder /app /app
# Open port 3003, because chi
EXPOSE 3003
# Execute the app folder content
CMD ["/app"]