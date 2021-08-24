# simple-web-store

Run postgres database, ensure the data folder is empty
```shell
cd db
docker-compose up -d
```

Run Backend
```bash
go run main.go
```

For testing apis postman collection is available Simple Web Store.postman_collection.json

Run React App
```bash
cd store-ui
yarn start
```
