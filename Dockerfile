# Use a base image that has timezone configuration
FROM golang:1.20

# Set environment variable TZ to Asia/Bangkok
ENV TZ=Asia/Bangkok

# Install tzdata package for setting timezones
RUN apt-get update && apt-get install -y tzdata

# Configure the timezone
RUN ln -fs /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && dpkg-reconfigure -f noninteractive tzdata

# Your application build instructions here
WORKDIR /app
COPY . .
RUN go build -o main .

# Run your application
CMD ["./main"]