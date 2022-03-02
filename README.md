# sample-go-web-app

### cURL

```sh
curl -i -X POST -H "Content-Type: application/json" -d '{"content":"It is my first go web application","author":"Song Yu Chen"}' http://127.0.0.1:8080/post/

curl -i -X GET http://127.0.0.1:8080/post/1

curl -i -X PUT -H "Content-Type: application/json" -d '{"content":"It is my first go web app"}' http://127.0.0.1:8080/post/1

curl -i -X DELETE http://127.0.0.1:8080/post/1
```

