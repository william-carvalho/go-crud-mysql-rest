 
FROM golang

COPY dist/go-crud-mysql-rest /bin/

EXPOSE 5001

ENTRYPOINT [ "/bin/go-crud-mysql-rest" ]