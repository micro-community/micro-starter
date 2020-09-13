FROM alpine
ADD auth /auth
ENTRYPOINT [ "/auth" ]
