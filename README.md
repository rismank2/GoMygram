# GoMygram
Final Project Scalable Web Service with Golang - DTS Kominfo
library/package yang digunakan
1. github.com/lib/pq
2. github.com/gorilla/mux
3. github.com/gorm.io/gorm 
4. github.com/gorm.io/driver/mysql
5. github.com/dgrijalva/jwt-go
6. golang.org/x/crypto

#POSTMAN REST API"

#Menambhakan user
Method Post : /users/register
{
  "age" : 23,
  "email" : "risman@gmail.com",
  "password" : "risman",
  "username" : "risman"
}
![1](https://user-images.githubusercontent.com/64664885/196326006-a66c9004-0683-41a7-a9b8-1b2d872d39bc.PNG)

#User login
Method Post : /users/login
{
  "email" : "risman@gmail.com",
  "password" : "risman"
}
![2](https://user-images.githubusercontent.com/64664885/196326289-8e3a3f38-9818-40fc-a8cc-9f5d1c6d0f0e.PNG)

