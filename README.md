# scproxy

Redis backend for sccproxy that allows for read only with local caching and read write access over ssh forwarding.

This is somewhat of a hack since local caching would ideally be supported in sccache directly. It may make sense to use goma
here as well when the server side is open sourced (https://groups.google.com/a/chromium.org/forum/m/#!topic/chromium-dev/q7hSGr_JNzg).

## Getting Started

### Server

Add read and read/write public keys to sshd/authorized_keys.ro and sshd/authorized_keys.rw then start with docker-compose.

```
docker-compose -f docker-compose.proxy.yml 
```

### Read only client with local redis caching
```
./start.sh  <server ip>
```

### Read write client
```
./start_rw.sh  <server ip>
```

### Prerequisites

* docker
* docker-compose
