package moduletest

// import "github.com/fyf2173/moduledependencytest"
import "github.com/NorwayLobster/gomodone"

func Hello() string {
	return "Hello World"
}

// func Hello() string {
// return moduledependencytest.Hello()
// }

func Hello1() string {
	return gomodone.SayHi("World")
}
