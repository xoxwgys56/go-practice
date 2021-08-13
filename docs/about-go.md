# About go lang

## start from `main`
`C`언어와 마찬가지로 `go`도 `main`함수에서 시작한다는 특징을 갖고 있다. 만약 `main`함수가 없다면 프로그램은 시작되지 않는다.

## [package “main” and func “main”](https://stackoverflow.com/questions/42333488/package-main-and-func-main)

```go
package main
```

```go
func main()
```

무엇이 다른걸까?

먼저 함수 `main`에 대해 살펴보자. `go`의 entry point는 `main` 함수다. 이것에 대한 설명은 찾기 어려웠는데 [go spec Program_execution](https://golang.org/ref/spec#Program_execution)을 보면, 설명 되어있다.

>완전한 프로그램은 홀로 연결되어 있도록 만들어져 있고, 다른 곳에서 `import`되지 않은 `main package`여야 한다.
>
>`main package`는 반드시 `main`이라고 명시된 함수를 선언해야하고, 어떠한 `argument`도 갖지 말아야 하며, 어떠한 값도 반환하지 않아야 한다.

```go
func main() {...}
```

프로그램은 `main package`를 초기화하고 `main` 함수를 호출하는 것으로 시작된다.
이 함수 호출 결과가 반환되면, 프로그램이 종료된다. 이것은 다른 (`non-main`) `goroutine`이 완료될 때까지 기다리지 않는다.

위 내용을 보면 알 수 있듯이 `main package`외에서 사용된 `main` 에는 특별한 의미를 두지 않지만, `package main`에서 사용된 `main`은 `reserved keyword`라는 것을 알 수 있다.

그렇기때문에 만약 `main` 함수를 `non-main package`에서 선언하는 것은 괜찮다. 다만 그 함수는 그저 이름만 `main`인 함수가 된다. (되도록 선언하지 않는게 좋을 것이다.)

