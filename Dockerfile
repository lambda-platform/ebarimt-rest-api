FROM golang:latest

RUN apt-get update && \
    apt-get install -y libsqlite3-dev && \
    rm -rf /var/lib/apt/lists/*

ENV LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

COPY x64/libcrypto.so.1.0.0 /usr/lib
RUN ln -s /usr/lib/libcrypto.so.1.0.0 /usr/lib/libcrypto.so

COPY x64/libssl.so.1.0.0 /usr/lib
RUN ln -s /usr/lib/libssl.so.1.0.0 /usr/lib/libssl.so

COPY x64/libicudata.so.53.1 /usr/lib
RUN ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so.53 \
    && ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so.5 \
    && ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so

COPY x64/libicui18n.so.53.1 /usr/lib
RUN ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so.53 \
    && ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so.5 \
    && ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so

COPY x64/libicuuc.so.53.1 /usr/lib
RUN ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so.53 \
    && ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so.5 \
    && ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so

COPY x64/libQt5Core.so.5.4.1 /usr/lib
RUN ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so.5.4 \
    && ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so.5 \
    && ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so

COPY x64/libQt5Network.so.5.4.1 /usr/lib
RUN ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so.5.4 \
    && ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so.5 \
    && ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so

COPY x64/libQt5Script.so.5.4.1 /usr/lib
RUN ln -s /usr/lib/libQt5Script.so.5.4.1 /usr/lib/libQt5Script.so.5.4 \
    && ln -s /usr/lib/libQt5Script.so.5.4.1 /usr/lib/libQt5Script.so.5 \
    && ln -s /usr/lib/libQt5Script.so


# Add user
RUN useradd -ms /bin/bash ebarimtuser
USER ebarimtuser

# Set working directory
WORKDIR /home/ebarimtuser/app

# Copy the rest of the project files into the container at /app
COPY --chown=ebarimtuser:ebarimtuser . .

# Download dependencies
RUN go mod download

# Build the project
RUN go build -o main .

# Expose port 3000 for the application
EXPOSE 3000

# Run the application when the container starts
CMD ["./main"]
