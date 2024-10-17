# go-fiber-docker-example
## 프로젝트 구성

- **`cmd/main.go`**: Entry Point
- **`internal/model`**: Prisma 스키마 및 Prisma Client 관련 코드
- **`internal/repository`**: DB 접근 관련 코드. 데이터 읽기&쓰기
- **`internal/service`**: 비즈니스 로직
- **`internal/controller`**: API Endpoint && HTTP 요청 처리
- **`pkg/`**: 내부 사용 라이브러리