FROM mariadb:10.5

RUN apt update ; \
    apt install -y openssl ; \
    apt install ; \
    rm -rf /var/lib/apt/lists/* \
    ;

# COPY initial_content.sql.enc /initial_content.sql.enc

# RUN --mount=type=secret,id=domain_secret,target=/run/domain_secret.txt \
#     cat /run/domain_secret.txt ; \
#     openssl enc \
#     -aes-256-cbc \
#     -d \
#     -salt \
#     -in /initial_content.sql.enc \
#     -out /initial_content.sql \
#     -pass file:/run/domain_secret.txt \
#     && \
#     rm /initial_content.sql.enc ; \
#     mv /initial_content.sql /docker-entrypoint-initdb.d/ ; \
#     chown -R mysql:mysql /docker-entrypoint-initdb.d ; \
#     head -n 30 /docker-entrypoint-initdb.d/initial_content.sql

EXPOSE 3306
