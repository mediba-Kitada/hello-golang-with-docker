FROM mysql:5.6

# テーブル作成SQL
COPY ./mant_sv_api.sql /docker-entrypoint-initdb.d/

# テーブルにデータを投入するSQL
COPY ./total_qa.sql /docker-entrypoint-initdb.d/
