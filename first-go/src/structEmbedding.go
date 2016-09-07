package main

import "fmt"

type Person struct{
	name string
	age int
}

func (p *Person) greeting(){ // Person 구조체에 메서드 연결
	fmt.Println("Hello~")
}

type Student1 struct{
	p Person // 이렇게 그냥 맴버변수로 가지고 있게 되면, has - a 관계가 된다, 임베딩 된 관계라고 하지는 않는다
	school string
	grade int
}

type Student2 struct{
	Person // 이렇게하면 임베딩 하는 것
	school string
	grade int
}


// 오버라이딩을 구현해봅시다
func (p *Student2) greeting(){
	fmt.Println("Hello Overriding~")
}

/*
	부모 구조체의 메서드 이름과 중복된다면 상속 과정의 맨 아래 메서드가 호출이 된다
	즉, 자식 구조체의 메서드가 부모 구조체의 메서드를 오버라이드 하는 효과가 생긴다
	부모 구조체의 메서드를 호출하고 싶다면, s.Person.greeding() 처럼 명시적으로 구조체 이름을 불러주어야 한다
*/

func main(){
	p := Person{"junsu", 30}
	p.greeting()

	s1 := Student1{p, "Soongsil", 1}
	s1.p.greeting() // 부르려면 이렇게 해야함
	// s1.greeting() // 여기서는 이렇게하면 에러남

	s2 := Student2{p, "Soongsil", 1}
	s2.Person.greeting() // 가능
	s2.greeting()		 // 가능

	// 오버라이딩 을 구현해봅시다
}