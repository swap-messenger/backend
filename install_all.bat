cd src/api/chat
go install
cd ..
cd user
go install
cd ..
cd methods
go install
cd ..
go install
cd ..
go install
cd messages
go install
cd ..
cd message_engine
go install
cd ..
cd ..
cd models
go install
cd ..
cd db_work
go install
cd ..
git rm --cached app.db
go run src/main.go
