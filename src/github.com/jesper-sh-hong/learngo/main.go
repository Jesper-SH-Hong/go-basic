package main

//go의 패키지들..
import (
	"fmt"
	"strings"

	"github.com/jesper-sh-hong/learngo/something"
)

//패키지명 main: 이 프로젝트를 컴파일 하고 사용(서버 시작)할 것이다.
//ex. 걍 라이브러리/오픈소스: main.go 필요 없음. (colly 보면 main.go 없음. 컴파일 필요 없으면.. )
//컴파일러가 main부터 찾음

// 둘다 int면 저렇게 한번만 힌트. 후자는 리턴타입. return 없으면 배우셈
func multiply(a, b int) int {
	return a * b
}

// 여러 값을 리턴하는 섹시한 Go..
func lenAndUpper(name string) (int, string) {
	defer fmt.Println("I'm done") //함수가 끝나고 실행됨. 몹시 유용. 이미지 업로드 후 파일 삭제, 파일 닫기, api 요청 등.
	return len(name), strings.ToUpper(name)
}

// 가변인자. ...로 표시
func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked return.
func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name) //새로 생성하지 않음. 대입이므로 := 아님
	uppercase = strings.ToUpper(name)
	return //return length, uppercase도 가능인데 이미 위에 선언했으니..
}

func main() {
	fmt.Println("Hi")
	//Go는 함수를 export 하려면 함수명을 대문자로
	//formatting 패키지의 Println 함수

	something.SayBye()
	//something.sayHello는 못 갖다 쓰지. export 안되서

	//variables & constants
	const name string = "jesper"

	//축약형(shorthand)은 func 안에서만 제공. 밖에서는 var로 선언해야
	// var name2 string = "jesp" 인데 걍 := 써도 고가 타입 추정
	name2 := "jesp"
	name2 = "kelly"
	fmt.Println(name)
	fmt.Println(name2)

	fmt.Println(multiply(2, 2))

	// totalLen, upperName := lenAndUpper("jesper")
	// fmt.Println(totalLen, upperName)

	totalLen, upperName := lenAndUpper2("jesper")
	fmt.Println(totalLen, upperName)

	//ignored value: _ 로 무시
	totalLen2, _ := lenAndUpper("jesper")
	fmt.Println(totalLen2)

	repeatMe("jesper", "cindy", "james", "kelly")

}
