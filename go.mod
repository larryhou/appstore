module appstore

go 1.24

replace github.com/larryhou/appstoreconnect => ./src/github.com/larryhou/appstoreconnect

require github.com/larryhou/appstoreconnect v0.0.0-00010101000000-000000000000

require github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
