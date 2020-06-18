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

Token is valid for only 24 Hours

After new login old token will be invalidated

Only non empty fields will be updated in the database
```
{
    "email": "vishal@gmail.com",
    "password": "ABCD12345",
    "phone_no": "123456"
}
```
