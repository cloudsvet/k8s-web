FROM debian
ADD ./web .
EXPOSE 3000
CMD ./web
