using mysql-client I want to connect

To connect to a MySQL server using the `mysql-client` from a Docker container, use the following command:

```
docker run --rm -it mysql:latest mysql -h <mysql_host> -P <port> -u <user> -p
```

Replace `<mysql_host>`, `<port>`, and `<user>` with the appropriate MySQL server address, port (usually 3306), and username. The `-p` flag will prompt for the password. For example, to connect to a MySQL instance running on `localhost` and mapped to port `13306` for the root user:

```
docker run --rm -it mysql:latest mysql -h 127.0.0.1 -P 13306 -u root -p
```

This will start a temporary container, run the MySQL client inside it, allow interactive use, and remove the container after exit. Make sure Docker networking and port mappings allow this connection. Use `mysql/mysql-server` as the image if you don't have `mysql:latest` locally[1][2].

Citations:
[1] Connect to mysql in a docker container from the host https://stackoverflow.com/questions/33001750/connect-to-mysql-in-a-docker-container-from-the-host
[2] 2.6.1 Basic Steps for MySQL Server Deployment with Docker https://dev.mysql.com/doc/mysql-linuxunix-excerpt/8.0/en/docker-mysql-getting-started.html
[3] Run and connect to MySQL in one step using docker run https://forums.docker.com/t/run-and-connect-to-mysql-in-one-step-using-docker-run/103237
[4] mysql - Official Image | Docker Hub https://hub.docker.com/_/mysql
[5] How to Set Up and Configure MySQL in Docker https://hevodata.com/learn/docker-mysql/
[6] Problem in connecting with mysql in docker - General https://forums.docker.com/t/problem-in-connecting-with-mysql-in-docker/32954
[7] How to install MySQL using Docker https://www.bytebase.com/reference/mysql/how-to/how-to-install-mysql-using-docker/
[8] How to use MySQL with Docker and Docker compose a ... https://geshan.com.np/blog/2022/02/mysql-docker-compose/
[9] docker-compose run --rm backend php yii migrate does not ... https://github.com/yiisoft/yii2-app-advanced/issues/458
[10] Connecting MySQL in Docker container via client ... https://discuss.openedx.org/t/connecting-mysql-in-docker-container-via-client-application-e-g-mysql-workbench/13786
[11] Docker container to connect with MySQL locally installed https://forums.docker.com/t/docker-container-to-connect-with-mysql-locally-installed/85000
[12] Manual MySQL upgrade https://docs.mailcow.email/troubleshooting/debug-mysql_upgrade/
[13] 1.11.1-beta10 (build: 6662) Mysql db migration schema but ... https://forums.docker.com/t/1-11-1-beta10-build-6662-mysql-db-migration-schema-but-no-data/11024
[14] Configure External Databases https://docs.dronahq.com/self-hosted-deployment/configure-external-databases/
[15] MySQL 8.4 Reference Manual :: Option Index https://dev.mysql.com/doc/refman/8.3/en/dynindex-option.html
[16] Docker compose for MySQL https://www.dotnetthailand.com/storage/mysql/docker-compose-for-mysql/
[17] MySQL 9.2 Reference Manual :: Option Index https://dev.mysql.com/doc/refman/9.2/en/dynindex-option.html
[18] Working with Databases https://phd.dmstr.io/en/docs/guide/tutorials/database.md
[19] Import MySQL Database - Kinsta® Docs https://kinsta.com/docs/wordpress-hosting/database-management/restore-import-mysql-database/
