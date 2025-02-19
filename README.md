# chzzk-sdk-golang

`chzzk-sdk-golang`은 Go 언어로 작성된 SDK로, `chzzk-api` 패키지를 통해 특정 API와 상호 작용할 수 있도록 설계되었습니다.

## 설치

Go 모듈을 사용하여 `chzzk-sdk-golang`을 설치하려면, 프로젝트의 `go.mod` 파일에 다음과 같이 추가하세요:

```go
require (
    github.com/ArsCaelum/chzzk-sdk-golang latest
)
```

그런 다음, 다음 명령어를 실행하여 모듈을 다운로드합니다:

```bash
go mod tidy
```

## 사용법

`chzzk-api` 패키지를 임포트한 후, 제공되는 기능을 활용하여 API와 상호 작용할 수 있습니다.

```go
package main

import (
    "github.com/ArsCaelum/chzzk-sdk-golang/chzzk-api"
)

func main() {
    // chzzk-api 패키지의 기능을 사용하여 로직 구현
}
```

자세한 사용 방법과 예제는 `chzzk-api` 패키지의 소스 코드를 참고하시기 바랍니다.

## 기여

버그 제보, 기능 추가 요청 또는 풀 리퀘스트는 언제나 환영합니다.
기여하시려면 이 저장소를 포크하고 변경 사항을 반영한 후 풀 리퀘스트를 생성해주세요.

## 라이선스

이 프로젝트는 MIT 라이선스 하에 배포됩니다. 자세한 내용은 [LICENSE](https://github.com/ArsCaelum/chzzk-sdk-golang/blob/main/LICENSE) 파일을 참고하세요.

더 자세한 정보와 업데이트는 [GitHub 저장소](https://github.com/ArsCaelum/chzzk-sdk-golang)를 방문하여 확인하시기 바랍니다.

