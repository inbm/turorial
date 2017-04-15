package main

import (
	"calc"
	"fmt"
	_ "math"
	_ "mynet"
	"net"
	"os"

	"bufio"
	_ "file"
	_ "math"
	"math/rand"
	_ "runtime"
	_ "sync"
	"time"
	_ "unsafe"

	_ "runtime"

	_ "math"

	"reflect"
	_ "runtime"
	. "shape"
	_ "sync"
	_ "sync/atomic"
)

type Data struct {
	tag    string
	buffer []int
}

func (d *Data) setData() {

}

func _p(a ...interface{}) {
	fmt.Println(a...)
}

func _t(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}

func _v(v interface{}) reflect.Value {
	return reflect.ValueOf(v)
}

func h(args []reflect.Value) []reflect.Value {
	_p("hello")
	return nil
}

func main() {

	var hello func()

	fn := reflect.ValueOf(&hello).Elem()
	v := reflect.MakeFunc(fn.Type(), h)

	fn.Set(v)

	hello()

	//runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var a *int = new(int)
	*a = 3

	_p(_t(a))

	_p("*->", _v(a))

	_p("v->", _v(a).Elem())

	var b interface{}
	b = 10

	_p("*->", _v(b))
	//_p(_v(a).Elem().Int())

	s1 := Shape(&Retangle{10, 10})
	c1 := Circle{100, 1000, 39}
	_p(s1)

	name, ok := reflect.TypeOf(c1).FieldByName("X")
	_p(name, ok)
	_p(reflect.TypeOf(s1).Method(0))

	/*runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {       // 고루틴 2,000개 생성
			atomic.AddInt32(&data, 1)// 1씩 더함
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {       // 고루틴 1,000개 생성
			atomic.AddInt32(&data, -1)
			wg.Done()
		}()
	}

	wg.Wait()
	_p(data)*/

	/*var r1 Retangle
	r2 := new(Retangle)
	r3 := Retangle{}

	_p(r1, r2, r3)




	var type_assertion interface{} = "hello"

	_type_assertion := type_assertion.(string)
	_p("_type_assertion: ", _type_assertion)

	var shape = Shape(&Retangle{})


	var rectangle = shape.(*Retangle)

	_p("shape: ", shape, " rectangle: ", rectangle)*/

	/*runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	wg := new(sync.WaitGroup) // 대기 그룹 생성

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		wg.Add(1)            // 반복할 때마다 wg.Add 함수로 1씩 추가
		go func(n int) {     // 고루틴 10개 생성

			fmt.Println(n)
			wg.Done()    // 고루틴이 끝났다는 것을 알려줌
		}(i)
	}

	wg.Wait() // 모든 고루틴이 끝날 때까지 기다림
	*/
	//time.Sleep(5 * time.Second)

	/*

		pool := sync.Pool{
			New: func() interface{} {
				data := new(Data)
				data.tag = "new"
				data.buffer = make([]int, 10)
				return data
			},
		}

		for i:=0; i<10; i++{
			go func() {
				//data := pool.Get().(Data)
				data := pool.Get().(*Data)

				for index:=range data.buffer{
					data.buffer[index] = rand.Intn(100)
				}
				_p(data)
				data.tag = "used"
				pool.Put(data)
			}()
		}


		for i := 0; i < 10; i++ {
			go func() {                               // 고루틴 10개 생성
				data := pool.Get().(*Data)        // 풀에서 *Data 타입으로 데이터를 가져옴
				n := 0
				for index := range data.buffer {
					data.buffer[index] = n    // 슬라이스에 짝수 저장
					n += 2
				}
				fmt.Println(data)                 // data 내용 출력
				data.tag = "used"                 // 객체가 사용되었다는 태그 설정
				pool.Put(data)                    // 풀에 객체 보관
			}()
		}

	*/

	/*
		shapes := []Shape{&Circle{0, 0, 5}, &Ellipse{Circle{0, 0, 3}, 10}, &Retangle{7,7}}

		for _, shape := range shapes{
			_p(shape.Area())
		}
	*/

	/*

		runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

		var mutex = new(sync.Mutex)    // 뮤텍스 생성
		var cond = sync.NewCond(mutex) // 뮤텍스를 이용하여 조건 변수 생성

		c := make(chan bool, 3) // 비동기 채널 생성

		for i := 0; i < 3; i++ {
			go func(n int) {                        // 고루틴 3개 생성
				mutex.Lock()                    // 뮤텍스 잠금, cond.Wait() 보호 시작
				c <- true                       // 채널 c에 true를 보냄
				_p("wait begin : ", n)
				cond.Wait()                     // 조건 변수 대기
				_p("wait end : ", n)
				mutex.Unlock()                  // 뮤텍스 잠금 해제, cond.Wait() 보호 종료

			}(i)
		}

		for i := 0; i < 3; i++ {
			<-c // 채널에서 값을 꺼냄, 고루틴 3개가 모두 실행될 때까지 기다림
		}


			mutex.Lock()                // 뮤텍스 잠금, cond.Signal() 보호 시작
			_p("signal")
			cond.Broadcast()               // 대기하고 있는 고루틴을 하나씩 깨움
			mutex.Unlock()              // 뮤텍스 잠금 해제, cond.Signal() 보고 종료

	*/

	/*

		runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

		var data int = 0
		var rwMutex = new(sync.RWMutex)

		go func() {                                       // 값을 쓰는 고루틴
			for i := 0; i < 3; i++ {
				rwMutex.Lock()
				data += 1                         // data에 값 쓰기
				_p("write   : ", data)   // data 값을 출력
				time.Sleep(10 * time.Millisecond) // 10 밀리초 대기
				rwMutex.Unlock()
			}
		}()

		go func() {                                       // 값을 읽는 고루틴
			for i := 0; i < 3; i++ {
				rwMutex.RLock()
				_p("read 1 : ", data)    // data 값을 출력(읽기)
				time.Sleep(1 * time.Second)       // 1초 대기
				rwMutex.RUnlock()
			}
		}()

		go func() {                                       // 값을 읽는 고루틴
			for i := 0; i < 3; i++ {
				rwMutex.RLock()

				time.Sleep(1 * time.Second)       // 2초 대기
				_p("read 2 : ", data)    // data 값을 출력(읽기)
				rwMutex.RUnlock()
			}
		}()


		rwMutex.RLock()
		time.Sleep(3 * time.Second)              // 10초 동안 프로그램 실행
		_p("read 3 : ", data)
		rwMutex.RUnlock()

	*/

	/*
		runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

		var data = []int{} // int형 슬라이스 생성
		var mutex = new(sync.Mutex)

		go func() {                             // 고루틴에서
			for i := 0; i < 1000; i++ {     // 1000번 반복하면서
				mutex.Lock()
				data = append(data, 1)  // data 슬라이스에 1을 추가
				mutex.Unlock()

				runtime.Gosched()       // 다른 고루틴이 CPU를 사용할 수 있도록 양보
			}
		}()

		go func() {                             // 고루틴에서
			for i := 0; i < 1000; i++ {     // 1000번 반복하면서
				mutex.Lock()
				data = append(data, 2)  // data 슬라이스에 1을 추가
				mutex.Unlock()

				runtime.Gosched()       // 다른 고루틴이 CPU를 사용할 수 있도록 양보
			}
		}()

		time.Sleep(2 * time.Second)      // 2초 대기

		_p(data)
		_p(len(data))           // data 슬라이스의 길이 출력
	*/

	/*

		c1 := make(chan int)
		c2 := make(chan string)

		go func(){
			for{
				c1<-10
				time.Sleep(100*time.Millisecond)
			}
		}()

		go func(){
			for{
				c2<-"hello"
				time.Sleep(500*time.Millisecond)
			}
		}()

		go func(){
			for{
				select {
				case i := <-c1:
					_p("c1:", i)
				case s := <-c2:
					_p("c2:", s)
				}
			}
		}()

		time.Sleep(10*time.Second)

	*/

	//suspend();

	/*

		n := []int{1, 2, 3, 4, 5}
		r := sumRange(n...) // ...를 사용하여 가변인자에
		// 슬라이스를 바로 넘겨줌

		_p(r) // 15


		_p(file.Read("./src/main.go"))
	*/

	/*terrestrialPlanet := map[string]map[string]float32{
		"Mercury": map[string]float32{
			"meanRadius":    2439.7,
			"mass":          3.3022E+23,
			"orbitalPeriod": 87.969,
		},
		"Venus": map[string]float32{
			"meanRadius":    6051.8,
			"mass":          4.8676E+24,
			"orbitalPeriod": 224.70069,
		},
		"Earth": map[string]float32{
			"meanRadius":    6371.0,
			"mass":          5.97219E+24,
			"orbitalPeriod": 365.25641,
		},
		"Mars": map[string]float32{
			"meanRadius":    3389.5,
			"mass":          6.4185E+23,
			"orbitalPeriod": 686.9600,
		},
	}

	_p(terrestrialPlanet["Mars"]["mass"]) // 6.4185E+23*/
	/*

		for i := 1; i <= 100; i++ {
			switch {                         // case에 조건식을 지정했으므로 판단할 변수는 생략
			case i%3 == 0 && i%5 == 0:       // 3의 배수이면서 5의 배수일 때
				_p("FizzBuzz")  // FizzBuzz 출력
			case i%3 == 0:                   // 3의 배수일 때
				_p("Fizz")      // Fizz 출력
			case i%5 == 0:                   // 5의 배수일 때
				_p("Buzz")      // Buzz 출력
			default:                         // 아무 조건에도 해당하지 않을 때
				_p(i)           // 숫자 출력
			}
		}
	*/

	//
	//go producer(c)
	//go consumer(c)

	/*

		go func() {
			c <- 1
		}()

		a, ok := <-c       // 채널이 닫혔는지 확인
		_p(a, ok) // 1 true: 채널은 열려 있고 1을 꺼냄

		close(c)           // 채널을 닫음
		a, ok = <-c        // 채널이 닫혔는지 확인
		_p(a, ok)
	*/

	/*
		go func(){
			for i:=0; i<5; i++{
				c <- i


			}
			close(c)
		}()


		for i:=range c{
			//a, ok := <-c
			_p(i)
			//_p(i, a, ok)
		}*/

	//channel basic
	/*
		channel := make(chan int)
		go sum(1, 2, channel)

		n := <-channel

		_p(n)
	*/

	// sync channel

	/*
		done := make(chan bool)

		count :=3

		go func() {
			for i := 0; i < count; i++ {
				done <- true
				_p("gorutine:", i)
				time.Sleep(1 * time.Second)
			}
		}()



		for i := 0; i < count; i++ {
			<-done
			_p("main:", i)
		}
	*/

	/*

		runtime.GOMAXPROCS(1)

		hello := "hello"

		for i:=0; i<3; i++{
			go func(n int){
				_p(hello, n)
			}(i)
		}

	*/

	//suspend()

	/*

		var student Student

		_p(unsafe.Sizeof(student))

		student.school = "han"
	*/

	/*
		closure("success", func(s string){
			println(s)
		})

		serve()
	*/
}

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex = map[string]Vertex{
	"lee":  Vertex{1.0, 2.0},
	"good": Vertex{3.0, 4.0},
}

var pow = []int{}

func closure(s string, f func(string)) {
	println("start " + s)
	f("connect")
	_p(calc.Sum(5, 6))
}

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	school string
	grade  int
}

func (p *Person) greeting() {
	_p("hello")
}

func hello(n int) {

	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	_p(n)

}

func suspend() {
	_p("wait:")
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		_p(err)
	} else {
		_p(string(sentence))
	}
}

func sum(a int, b int, c chan int) {
	c <- a + b
}

func producer(c chan<- int) { // 보내기 전용 채널
	for i := 0; i < 5; i++ {
		c <- i
	}

	c <- 100 // 채널에 값을 보냄

	//_p(<-c) // 채널에서 값을 꺼내면 컴파일 에러
}

func consumer(c <-chan int) { // 받기 전용 채널
	for i := range c {
		_p(i)
	}

	_p("by consumer", <-c) // 채널에 값을 꺼냄

	//c <- 1        // 채널에 값을 보내면 컴파일 에러
}

func SumAndDiff(a int, b int) (int, int) {
	return a + b, a - b // 리턴할 값 또는 변수를 차례대로 나열
}

func sumRange(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}

	return total
}

func serve() {
	service := "0.0.0.0:8889"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// go routine 실행
		go handleclient(conn)
	}

}

func handleclient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		println(n)
		if err != nil {
			return
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}
