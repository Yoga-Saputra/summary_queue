FROM golang:latest

# Update OS and install lib
RUN apt-get update \
    && apt-get install -y git

# Set working directory
WORKDIR /app

# Download and install go air live reload (development purpose)
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

# Command which applies when container from this image runs
EXPOSE 8090
CMD air