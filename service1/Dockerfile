FROM busybox:latest
RUN mkdir /app
ADD Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
ADD service1 /app/service1
ADD config /app/config
RUN chmod -R 777 /app/service1
WORKDIR /app
CMD ["/app/service1"]