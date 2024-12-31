# Week 5

---

## Docker를 활용한 컨테이너화

---

### 1. 실행환경 구축
> Go는 단일 바이너리 파일을 build를 통해 생성함. 
> 다른 라이브러리나 환경 설정 없이 실행 가능.

### 2. 멀티 스테이지 빌드
> 1. Go의 go.mod, go.sum 파일을 복사하여 첫번째 빌드.
> 2. 최종 배포용 컨테이너 구성 (OS)
> 3. 오픈소스 air를 이용하여 파일이 변경될 때 마다 자동 빌드 및 실행. 그러기 위해 .air.toml 파일을 생성.

- 단, 현업에서는 테스트 코드를 작성하여 테스트를 진행하고, air를 쓰고 curl로 확인하는 것은 드뭄.

### 3. Docker compose를 이용한 멀티 컨테이너 구성
> 1. docker-compose.yml 파일로 여러 컨테이너를 한번에 실행 가능.
> 1.1 docker compose build --no-cache 사용.

### 4. Makefile
> 명령어를 간편하게 한번에 실행할 수 있도록 하는 파일.
> .PHONY를 사용하여 명령어를 실행할 수 있도록 함.
> 문법에 맞게 작성.

### 5. Github Actions
> CI/CD를 위한 Github Actions를 사용하여 빌드 및 테스트를 자동화.
> 기존 PR 기반이면 실제 머지가 되었을 때 빌드 및 테스트가 진행되지만, Github Actions를 사용하면 PR이 올라올 때부터 빌드 및 테스트가 진행됨.
> .github/workflows를 루트 디렉토리에 test.yml, .octocov.yml 파일을 생성하여 작성.
> .octocov.yml 파일은 테스트 커버리지를 확인하기 위한 파일. 테스트 비율, 테스트 시간 등을 확인할 수 있음.

#### 5.1 Golangci를 활용한 정적 분석
> workflows/golangci.yml와 .golangci.yml은 golangci-lint를 사용하여 정적 분석을 진행.
> 코드 품질 검토, 불필요한 코드 제거, 코드 품질 향상을 위해 사용.
> .golangci.yml의 enable의 unused값을 넣어주면 사용하지 않는 코드도 활용 가능.

## 2. http 서버를 약한 결합 구성으로 변경
> Config/config.go 파일을 생성하여 환경 변수를 사용하여 설정을 변경.
> port와 env를 설정하여 사용.
```go 
type Config struct {
    Port int
    Env string
}
```
> main.go 파일에서 Config를 사용하여 환경 변수를 설정.
> run 함수도 변경

## 3. 추가 구현
> 1. Server type 구조체를 생성하여 서버를 구성.
> 2. NewServer 함수를 생성하여 서버를 생성.
> 3. health check를 위한 핸들러 생성.