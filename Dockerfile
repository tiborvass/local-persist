FROM alpine

COPY bin/local-persist /usr/bin/
CMD ["local-persist"]
