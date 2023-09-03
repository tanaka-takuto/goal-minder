gqlgen:
	go run github.com/99designs/gqlgen generate

generate:
	make gqlgen
	make dbApply
	make sqlBoilerGenerate

dbDryRun:
	mysqldef -uroot -ppassword -hGM-db -P3306 GM --dry-run --config=mysqldef.yml < infra/db/schema/ddl.sql

dbApply:
	mysqldef -uroot -ppassword -hGM-db -P3306 GM  --config=mysqldef.yml < infra/db/schema/ddl.sql
	$(MAKE) dbExport

dbExport:
	mysqldef -uroot -ppassword -hGM-db -P3306 GM --export  --config=mysqldef.yml > infra/db/schema/ddl.sql

sqlBoilerGenerate:
	sqlboiler mysql