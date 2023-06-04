FROM debian:latest

COPY . .

ENTRYPOINT [ "/bin/exec" ]