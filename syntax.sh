#create db migrations
migrate create -ext sql -dir db/migrations nama_file_migration

#run
#mysql://user:password@tcp(host:port)/nama_database
migrate -database "mysql://root:@tcp(localhost:3306)/go_pzn_clone" -path db/migrations 20230414210021_create_table_courses.down

#set expire jwt
claims["exp"] = time.Now().Add(time.Minute * 15).Unix()


users > courses = many to many
categories > courses = many to many