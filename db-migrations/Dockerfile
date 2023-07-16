FROM migrate/migrate
WORKDIR /go/src/github.com/ow0sh/fullstacktodoapp
COPY migrations ./migrations
COPY migrate.sh ./migrate.sh
RUN chmod +x migrate.sh
ENTRYPOINT [ "./migrate.sh" ]