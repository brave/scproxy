FROM redis:4
EXPOSE 6379
COPY redis.conf /usr/local/etc/redis/redis.conf
COPY data/dump.rdb /data/dump.rdb
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
