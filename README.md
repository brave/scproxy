# scproxy

Redis backend for sccproxy that allows for read only with local caching and read write access over ssh forwarding.

This is somewhat of a hack since local caching would ideally be supported in sccache directly. It may make sense to use goma
here as well when the server side is open sourced (https://groups.google.com/a/chromium.org/forum/m/#!topic/chromium-dev/q7hSGr_JNzg).

## Getting Started

### Server

```
# Allow read/write and read only access for your public key
cat ~/.ssh/id_rsa.pub > sshd/authorized_keys.rw
cat ~/.ssh/id_rsa.pub > sshd/authorized_keys.ro

# Build and start server
docker-compose -f docker-compose.proxy.yml up --build
```

### Read only client with local redis caching
```
./start.sh  <server ip>
```

Currently this doesn't support encrypted ssh keys, and the above command assumes `~/.ssh/id_rsa`. Workaround for now is to create a second unencrypted ssh key that is only used for this and pass it as an argument to `./start.sh`.
```
ssh-keygen -f ~/.ssh/scproxy
./start.sh <server ip> scproxy
```

### Read write client
```
./start_rw.sh  <server ip>
```



### Prerequisites

* docker
* docker-compose

### Known issues

