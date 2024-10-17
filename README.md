# Online-Library-Backend
- 기존 [단대라이브러리](https://github.com/DKSH-WoongDo/Woongdo-API/) Express코드를 Go로 포팅하는 Repository

## 프로젝트 구성

- **`cmd/main.go`**: Entry Point
- **`internal/model`**: Prisma 스키마 및 Prisma Client 관련 코드
- **`internal/repository`**: DB 접근 관련 코드. 데이터 읽기&쓰기
- **`internal/service`**: 비즈니스 로직
- **`internal/controller`**: API Endpoint && HTTP 요청 처리
- **`pkg/crpyt`**: 내부 사용 암호화 라이브러리 (SHA256, JWT, Base64, ...)
- **`pkg/domain`**: 내부 사용 도메인
- **`pkg/utils`**: 내부 사용 라이브러리 (errResponseCreate, body validator, ...)

## 사용 라이브러리
- [fiber](https://github.com/gofiber/fiber) 
- [validator](https://github.com/go-playground/validator) 
- [prisma-client-go](https://github.com/steebchen/prisma-client-go) 