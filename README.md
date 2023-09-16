# GO Simple CRUD
## Golang Simple CRUD using Golang and PostgreSQL

Simple CRUD using Golang and PostgreSQL. Course from myskill.id

### Environment

Check PostgreSQL config in .env file.
Database should be exist.
Table will created while server started.

```sql
CREATE TABLE IF NOT EXISTS public.books
(
    id integer NOT NULL DEFAULT nextval('books_id_seq'::regclass),
    title text COLLATE pg_catalog."default" NOT NULL,
    author text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default" NOT NULL,
    stock integer NOT NULL,
    CONSTRAINT books_pkey PRIMARY KEY (id)
)
```

### Start Server

```
go run main.go
```

server will run at http://localhost:8080

### Login

username : admin
password : password123