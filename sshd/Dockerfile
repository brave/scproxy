FROM ubuntu:18.04

ARG env=default

RUN apt-get update && apt-get install -y openssh-server
RUN mkdir /var/run/sshd
RUN useradd -s /bin/false -m user 

COPY sshd_config /etc/ssh/sshd_config
COPY --chown=user:root authorized_keys.${env} /home/user/.ssh/authorized_keys

EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]