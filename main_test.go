package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {

	t.Run("2", func(t *testing.T) {
		os.Args[1] = "./input.txt"
		main()
	})
	t.Run("1", func(t *testing.T) {
		main()
	})

}
