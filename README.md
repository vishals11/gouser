## gouser

User registration system using Golang

### Steps
#### A. Dependencies
- Golang
- Mysql Database

#### B. Setup
```
export PORT=<port>

export DB_HOST=<db host>
export DB_USERNAME=<db username>
export DB_PASSWORD=<db password>
export DB_PORT=<db port>
export DB_NAME=<database name>

export JWT_SIGNING_KEY=<key>
```

#### C. Sample Request
- Create User `POST http://localhost:8000/user`

```
{
    "name": "Vishal",
    "email": "vishal@gmail.com",
    "password": "abcd12345",
    "phone_no": "123456"
}
```

- Login User `PATCH http://localhost:8000/user`

```
{
    "email": "vishal@gmail.com",
    "password": "abcd12345"
}
```

- Update User `PUT http://localhost:8000/user`

Header: `Authorization` : `<token>`
```
{
    "email": "vishal@gmail.com",
    "phone_no": "123456"
}
```
