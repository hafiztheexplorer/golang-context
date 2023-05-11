package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// function context//////////////////////////////////////////////////////////////

func TestContext(t *testing.T) {

	Background := context.Background()
	fmt.Println(Background)

	Todo := context.TODO()
	fmt.Println(Todo)

}

// function context with value////////////////////////////////////////////////////

func TestContextWithValue(t *testing.T) {
	// hirarki tertinggi
	contextA := context.Background()
	// hirarki yang ikut dengan contextA
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")
	// hirarki yang ikut dengan contextB
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")
	// hirarki yang ikut dengan contextC
	contextF := context.WithValue(contextC, "f", "F")
	// hirarki yang ikut dengan contextF
	contextG := context.WithValue(contextF, "g", "G")

	//silahkan play fuction ini, dan lihat gimana susunan / hirarkinya
	fmt.Println("\n",
		"isi contextA = ", contextA, "\n",
		"isi contextB = ", contextB, "\n",
		"isi contextC = ", contextC, "\n",
		"isi contextD = ", contextD, "\n",
		"isi contextE = ", contextE, "\n",
		"isi contextF = ", contextF, "\n",
		"isi contextG = ", contextG,
	)

	// mencari key e, di mana dia tidak dalam satu parent context dengan contextG
	fmt.Println(contextG.Value("e"))

}

// function context with cancel/////////////////////////////////////////////////////
// contoh tanpa context cancel (yang dicomment)///

/* func FungsiPenghitungSederhana() chan int { */

// ctx context.Context adalah parameter context
func FungsiPenghitungSederhana(ctx context.Context) chan int {
	// membuat channel dengan nama destination, dan dia mengembalikan return value berupa int
	destination := make(chan int)

	// membuat goroutine denan blank function
	go func() {
		// kalau sudah selesai go routine akan diclose
		defer close(destination)

		counter := 1
		// perulangan yang tidak akan berhenti, karena menghitung terus tanpa limit
		/*
			for {
			 	destination <- counter
			 	counter++
			 }
		*/

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}

	}()
	return destination
}

// coba jalankan function ini , maka goroutine awal dan akhir akan berbeda
// ini yang dinamakan goroutine leaks, jadi goroutine masih tetap ada yang aktif
// mengirim data ke channel, padahal sudah diminta stop
func TestContextWithCancel(t *testing.T) {

	/*
		fmt.Println("Total GoRoutine yang terjadi =", runtime.NumGoroutine())
	*/

	fmt.Println("Total GoRoutine yang terjadi =", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	/*
		destination := FungsiPenghitungSederhana()
	*/
	destination := FungsiPenghitungSederhana(ctx)
	fmt.Println("Total GoRoutine yang terjadi =", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter =", n)
		if n == 10 {
			break
		}
	}
	cancel() // mengiirim signal cancel ke context
	time.Sleep(2 * time.Second)
	fmt.Println("Total GoRoutine yang terjadi =", runtime.NumGoroutine())
}
