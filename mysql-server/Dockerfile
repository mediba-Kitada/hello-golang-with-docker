FROM mysql:5.6

# テーブル作成SQL
COPY ./ddl.sql /docker-entrypoint-initdb.d/

# テーブルにデータを投入するSQL
COPY ./dml.sql /docker-entrypoint-initdb.d/
