## 미들웨어를 만드는 법

### 미들웨어 패턴

- 시그니처란? 그냥 type hinting과 같은 느낌
- handler를 받고, handler를 넘김.

### 동적으로 받는 값을 이용한 미들웨어 패턴
#### Closure 이용 필요: 
- XXX 타입값(DB 연결 및 로거 등 다양한 Context), 환경 변수, 실행 시에 받는 인자를 넣기 위하여 클로저를 이용.

### 복원 미들웨어
- panic 발생 시 복원하는 미들웨어. 요청 단위로 복원을 하고, 이건 독립된 고루틴에서 발생하기 때문에, 특정 panic이 일어나도 서버가 죽지 않는다.
- 하지만 panic이 발생했을 때 오류 응답이나 구조를 통일시키기 위해 복원 미들웨어를 사용한다.

### 로그 미들웨어
- 요청 처리 시작 시간, 처리 시간, status code, http 메서드 및 path, 쿼리 파라미터 및 헤더 정보, 요청 바디, 응답 바디 등을 로그로 남긴다.
- Go의 HTTP request body는 스트림 데이터 구조로 한 번 밖에 읽을 수 없다. (Body만 그럼) 
따라서 요청 바디를 로깅하기 위해서는 요청 바디를 읽기 전에 별도 버퍼에 복사하는 등 추가 처리가 필요하다.
- io.NopCloser 함수로 복사한 body를 *bytes.Buffer 타입으로 변환한 다음 Close 메서드를 호출하여 복사한 body를 자동으로 닫아준다.
- zap은 빠르고 강력한 로그 라이브러리이다.
- 복사를 하기에 이 미들웨어 처리 부분에서 요청을 받아서 끝내야 하고, 단순 계산 시에 메모리가 2배가 들기 때문에 이미지 등의 BLOB과 같은 큰 데이터를 처리할 때는 주의해야 한다.

### status code 및 response body를 저장하는 미들웨어
- http.Hanlder 타입의 시그니처로 response를 나타내는 http.ResponseWriter 인터페이스는 읽기 관련 메서드를 가지고 있지 않다. 따라서 response body를 저장하기 위해서는 ResponseWriter 인터페이스를 구현한 구조체를 만들어야 한다.
- type rwWrapper struct {
    rw http.ResponseWriter
    mw io.Writer
    status int
}
- 이걸 그냥 interface로 만들어서 사용하면 됨.
- 추가적으로, request나 response body 부분에 사용자 개인 정보가 들어갈 수 있기 때문에, 이 부분은 주의해야 한다.

### context.Context 타입값에 정보를 부여하는 미들웨어
- *http.Request 타입값의 WithContext 메서드나 Clone 메서드를 이용하여 context.Context 타입값에 정보를 부여할 수 있다. 
- request 당 유지되는 타입





