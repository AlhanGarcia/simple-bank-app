FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download deps
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the code
COPY *.go ./
COPY handlers/*.go ./handlers/
COPY models/*.go ./models/
COPY database/*.go ./database/

# Build
RUN go build -o /simple-bank-app

# Expose port 8080
EXPOSE 8080

# Run the application
CMD [ "/simple-bank-app" ]