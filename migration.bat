@echo off

set /p table=nome da tabela:

migrate create -ext sql -dir src/api/infra/database/migrations -seq create_%table%_table

@pause