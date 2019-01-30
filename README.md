# Golang web server 小練習

## Create:
[POST] 
```
http://localhost:8080/create
```
#### Parameter:
name(string, required)
age(string, required)
gender(int, required)

## Read:
[GET] 
```
http://localhost:8080/read/{userId}
http://localhost:8080/read
```

## Update:
[POST] 
```
http://localhost:8080/update
```
#### Parameter:
user_id(string, required)
name(string, optional)
age(string, optional)
gender(int, optional)

## Delete:
[GET]
```
http://localhost:8080/delete/{userId}
```
