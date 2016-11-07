FROM alpine
RUN mkdir -p /data /run/docker/plugins
COPY bin/plugin-no-remove /usr/bin/
CMD ["plugin-no-remove"]
