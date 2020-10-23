FROM busybox:latest
RUN mkdir /app
ADD Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
ADD service2 /app/service2
ADD config /app/config
RUN chmod -R 777 /app/service2
WORKDIR /app
CMD ["/app/service2"]