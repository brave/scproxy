## Example
```
$ docker-compose up &
[1] 1111
$ Creating network "scproxy_public" with the default driver
Creating network "scproxy_private" with the default driver
Creating scproxy_scproxy_1      ... done
Creating scproxy_redis_remote_1 ... done
Creating scproxy_redis_local_1  ... done
Attaching to scproxy_redis_remote_1, scproxy_scproxy_1, scproxy_redis_local_1
redis_remote_1  | 1:C 20 May 04:26:44.843 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
redis_remote_1  | 1:C 20 May 04:26:44.844 # Redis version=4.0.9, bits=64, commit=00000000, modified=0, pid=1, just started
redis_remote_1  | 1:C 20 May 04:26:44.844 # Configuration loaded
redis_remote_1  |                 _._                                                  
redis_remote_1  |            _.-``__ ''-._                                             
redis_remote_1  |       _.-``    `.  `_.  ''-._           Redis 4.0.9 (00000000/0) 64 bit
redis_remote_1  |   .-`` .-```.  ```\/    _.,_ ''-._                                   
redis_remote_1  |  (    '      ,       .-`  | `,    )     Running in standalone mode
redis_remote_1  |  |`-._`-...-` __...-.``-._|'` _.-'|     Port: 6379
redis_remote_1  |  |    `-._   `._    /     _.-'    |     PID: 1
redis_remote_1  |   `-._    `-._  `-./  _.-'    _.-'                                   
redis_remote_1  |  |`-._`-._    `-.__.-'    _.-'_.-'|                                  
redis_remote_1  |  |    `-._`-._        _.-'_.-'    |           http://redis.io        
redis_remote_1  |   `-._    `-._`-.__.-'_.-'    _.-'                                   
redis_remote_1  |  |`-._`-._    `-.__.-'    _.-'_.-'|                                  
redis_remote_1  |  |    `-._`-._        _.-'_.-'    |                                  
redis_remote_1  |   `-._    `-._`-.__.-'_.-'    _.-'                                   
redis_remote_1  |       `-._    `-.__.-'    _.-'                                       
redis_remote_1  |           `-._        _.-'                                           
redis_remote_1  |               `-.__.-'                                               
redis_remote_1  | 
redis_remote_1  | 1:M 20 May 04:26:44.845 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
redis_remote_1  | 1:M 20 May 04:26:44.845 # Server initialized
redis_local_1   | 1:C 20 May 04:26:45.837 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
redis_remote_1  | 1:M 20 May 04:26:44.845 # WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled.
redis_remote_1  | 1:M 20 May 04:26:44.845 * DB loaded from disk: 0.000 seconds
redis_remote_1  | 1:M 20 May 04:26:44.845 * Ready to accept connections
redis_local_1   | 1:C 20 May 04:26:45.837 # Redis version=4.0.9, bits=64, commit=00000000, modified=0, pid=1, just started
redis_local_1   | 1:C 20 May 04:26:45.837 # Configuration loaded
redis_local_1   |                 _._                                                  
redis_local_1   |            _.-``__ ''-._                                             
redis_local_1   |       _.-``    `.  `_.  ''-._           Redis 4.0.9 (00000000/0) 64 bit
redis_local_1   |   .-`` .-```.  ```\/    _.,_ ''-._                                   
redis_local_1   |  (    '      ,       .-`  | `,    )     Running in standalone mode
redis_local_1   |  |`-._`-...-` __...-.``-._|'` _.-'|     Port: 6379
redis_local_1   |  |    `-._   `._    /     _.-'    |     PID: 1
redis_local_1   |   `-._    `-._  `-./  _.-'    _.-'                                   
redis_local_1   |  |`-._`-._    `-.__.-'    _.-'_.-'|                                  
redis_local_1   |  |    `-._`-._        _.-'_.-'    |           http://redis.io        
redis_local_1   |   `-._    `-._`-.__.-'_.-'    _.-'                                   
redis_local_1   |  |`-._`-._    `-.__.-'    _.-'_.-'|                                  
redis_local_1   |  |    `-._`-._        _.-'_.-'    |                                  
redis_local_1   |   `-._    `-._`-.__.-'_.-'    _.-'                                   
redis_local_1   |       `-._    `-.__.-'    _.-'                                       
redis_local_1   |           `-._        _.-'                                           
redis_local_1   |               `-.__.-'                                               
redis_local_1   | 
redis_local_1   | 1:M 20 May 04:26:45.839 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
redis_local_1   | 1:M 20 May 04:26:45.839 # Server initialized
redis_local_1   | 1:M 20 May 04:26:45.840 # WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled.
redis_local_1   | 1:M 20 May 04:26:45.840 * DB loaded from disk: 0.000 seconds
redis_local_1   | 1:M 20 May 04:26:45.840 * Ready to accept connections
scproxy_1       | PONG <nil>
scproxy_1       | PONG <nil>
scproxy_1       | 2018/05/20 04:26:48 started server at :6379

$ redis-cli get redis_remote
scproxy_1       | 2018/05/20 04:27:09 [DEBUG] 'redis_remote' not found in local cache
"1"
$ redis-cli get redis_local
"2"
$ redis-cli get doesnt_exist
scproxy_1       | 2018/05/20 04:27:25 [DEBUG] 'doesnt_exist' not found in local cache
(nil)
scproxy_1       | 2018/05/20 04:27:25 [ERROR] redis: nil
$ redis-cli set doesnt_exist jk
OK
$ redis-cli get doesnt_exist # retrieve key from local cache
"jk"
```

