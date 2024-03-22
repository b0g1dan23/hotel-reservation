# Hotel reservation backend

## Project outline

- users -> book room from an hotel
- admins -> going to check reservation/bookings
- Authentication and authorization -> JWT Tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> db management -> seeding, migration

## Resources

### MongoDB Driver

Documentation

```
https://www.mongodb.com/docs/drivers/go/current/quick-start/
```

Installing MongoDB Client

```
go get go.mongodb.org/mongo-driver/mongo
```

### GoFiber

Documentation

```
https://gofiber.io/
```

Installing GoFiber

```
go get github.com/gofiber/fiber/v2
```

## Docker

### Installing MongoDB as a Docker container

```
docker run -d -p 27017:27017 --name mongodb mongo:latest
```
