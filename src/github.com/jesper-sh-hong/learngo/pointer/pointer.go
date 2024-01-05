package pointer

import "fmt"

func Run() {
	a := 2
	// b := a //deep copy. 값 복사.
	b := &a
	a = 10
	fmt.Println(a, b)
	fmt.Println(a, &a, b)
	fmt.Println(*b)

	*b = 20  //포인터를 통해서 얘가 참조중인 a값 변경
	fmt.Println(a)



	//어레이.는 얼마나 큰 지 알려줘야 함. <-> slice. 크기 x
	names := [5]string{"nico", "lynn", "dal"}
	names[3] = "alala"
	names[4] = "alala"
	// names[5] = "alala" //idx out of bound.. error

	//slice. append로 
	names2 := []string{"nico", "lynn", "dal"}
	names2 = append(names2, "alala") //append는 새로운 slice를 만들어서 리턴. names2는 안 바뀜.. js,파이썬과 차이
	fmt.Println(names2)


	//map. key-value. 딕셔너리 같음
	nico := map[string]string{"name": "nico", "age": "12"} //[키타입]밸류타입
	fmt.Println(nico)
}