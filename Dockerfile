FROM alpine
ADD skkm-srv /skkm-srv
ENTRYPOINT [ "/skkm-srv" ]
