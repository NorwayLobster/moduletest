package moduletest

import (
	"github.com/NorwayLobster/gomodone"
	"github.com/fyf2173/moduledependencytest"
)

func Hello() string {
	return "Hello World"
}

func Hello2() string {
	return moduledependencytest.Hello()
}

func Hello1() string {
	return gomodone.SayHi("World")
}

func Proverb() string {
	return "proverb"
}
