FROM alpine
ADD auth-demo /auth-demo
ENTRYPOINT [ "/auth-demo" ]
