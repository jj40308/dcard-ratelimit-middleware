# dcard-ratelimit-middleware
實作一個 Server 並滿足以下需求:
- 每個 IP 每分鐘只能接受 60 個 requests
- 在 response 顯示目前 requests 的量，超過限制的話顯示 "Error"，例如一分鐘內第 30 個 request 顯示 30，第 61 個 request 顯示 Error
- 可以使用任意資料庫，也可自行設計 in-memory 資料結構，並在文件中說明理由
- 請附上測試
- 請不要使用 rate limit Library

## 專案架構
### 使用環境及工具
- 使用 Gin 作為開發框架
- Redis 作為紀錄 ratelimit 參照的資料庫
- Dockerize API Server, 並用 docker-compose 方便將 API Server 跟 Redis 跑起來
- 利用 vegeta 作為 integration test 的工具，模擬單一 IP request 超過 ratelimit 的狀況

### 如何運行專案
```bash
// Build API Server image
docker build -t dcard-ratelimit-middleware -f Dockerfile .
docker-compose up -d
```

API Server 預期會跑在 localhost:8080
```bash
curl localhost:8080/v1/get-remaining-requests
```

### 如何執行測試
Unit test
```bash
make test
```
Integration test
```bash
docker-compose up -d
make integration-test
```
