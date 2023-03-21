# BOOKING CINEMA TIKET APP

- Project được viết bằng Go
- Sử dụng framework Gin, ORM: GORM
- Frontend project: https://github.com/neil-go-phan/seatmap-frontend
- Database docker container chạy trên port 2345
- Migration: [golang-migration](https://github.com/golang-migrate/migrate)

## HOW TO RUN
- Yêu cầu:
  - docker
  - nodejs
- Clone source code
  - Chạy frontend:
    - Mở terminal
    - CD vào folder frontend
    - Chạy các lệnh sau:
    - `npm install`
    - `npm run dev`
  - Chạy backend:
    - Mở terminal
    - CD vào folder backend
    - Chạy các lệnh sau (Linux và MacOS):
    - `make docker_prepare`
    - `make postgres`
    - `make server`


## ADMIN ACCOUNT
````json
{
{
    "username": "superadmin",
    "password": "goldenowl2023"
}
}

````

## Entity Relationship Diagrams