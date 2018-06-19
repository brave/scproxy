FROM ubuntu:18.04

WORKDIR /app

RUN apt-get update && apt-get install -y ssh autossh
COPY ./run.sh .


EXPOSE 6379
CMD ["/app/run.sh"]
