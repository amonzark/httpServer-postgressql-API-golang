# httpServer-postgressql-API-golang
Simple HTTP Server CRUD to postgressql using Golang Language Program

Hi my name's Dwiyan Ramadhan AKA amonzark. In this github repo I just recently learn how to create an API using gin framework (https://github.com/gin-gonic/gin).
This program is really simple no advance code here. You can try it as long you following the instruction here. If you have any concerns feel free to comment.

Big Thanks.

##Requirement you need to install
There are several items that you need to install first before start running the code.
1. docker (https://docs.docker.com/engine/install/)
2. postgresql (use docker)
3. pgadmin4 (use docker)
4. postman (https://www.postman.com/downloads/)

for checking the request and response using http in postman you can use these
```bash
//GET List student
curl --location 'http://localhost:8085/students' \
--header 'Content-Type: application/json' \
--data ''

//POST Create student
curl --location 'http://localhost:8085/students/create' \
--header 'Content-Type: application/json' \
--data '{"ID":"F004", "name":"Grogu Stein", "age":5, "grade":9}'

//GET ID Student
curl --location 'http://localhost:8085/students/get/F004' \
--header 'Content-Type: application/json' \
--data ''

//GET Age Student
curl --location --request GET 'http://localhost:8085/students/20' \
--header 'Content-Type: application/json' \
--data ''

//POST Update Student
curl --location 'http://localhost:8085/students/update/F004' \
--header 'Content-Type: application/json' \
--data '{"age":30, "ID":"F004"}'

//POST Delete Student
curl --location --request POST 'http://localhost:8085/students/delete/F004' \
--header 'Content-Type: application/json' \
--data ''
```
