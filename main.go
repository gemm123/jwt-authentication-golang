package main

import "jwt-authentication-golang/database"

func main() {
	database.Connect("host=localhost user=postgres password=gemmq123456 dbname=jwt_demo port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	database.Migrate()
}
