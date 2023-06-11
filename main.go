package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	type placeholder [5]string
	fmt.Println("hii")

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digit := [...]placeholder{
		zero,
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine,
		colon,
	}

	for {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		tim := [...]placeholder{
			digit[hour/10],
			digit[hour%10],
			colon,
			digit[min/10],
			digit[min%10],
			colon,
			digit[sec/10],
			digit[sec%10],
		}

		for index := range colon {
			for t := range tim {
				fmt.Printf(" %v ", tim[t][index])
			}
			fmt.Println(" ")

		}

		println(" ")
		time.Sleep(time.Second)
	}

}
