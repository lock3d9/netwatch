package main

import (
	"fmt"
	"net"
	"time"
)

func startping(address string, count int, verbose bool) {
	successful := 0
	failed := 0
	var totalTime time.Duration
	fmt.Println("ping... ", address)

	for i := 1; i <= count; i++ {
		start := time.Now()
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		elapsed := time.Since(start)
		if err != nil {
			failed++
		} else {
			successful++
			totalTime += elapsed
		}

		if err != nil {
			failed++
			fmt.Printf("[%d] failed: %v\n", i, err)

		} else {
			successful++
			totalTime += elapsed

			fmt.Printf("[%d] connected time: %.2fms\n",
				i, elapsed.Seconds()*1000)
			if verbose == true {
				fmt.Printf(" %s -> %s\n",
					conn.LocalAddr(),
					conn.RemoteAddr(),
				)
			}
			conn.Close()
		}

		if i < count {
			time.Sleep(time.Second)
		}
		if successful > 0 {
			average := totalTime / time.Duration(successful)
			fmt.Println("Average:", average)
		}
	}
	fmt.Println("\nNetwatch Statistics\n")
	fmt.Printf(
		" Successful: %d\n Failed: %d\n Total: %d\n ",
		successful,
		failed,
		successful+failed,
	)
	if successful > 0 {
		avg := totalTime / time.Duration(successful)
		fmt.Printf("\nAverage: %.2fms\n", avg.Seconds()*1000)
	}
	total := successful + failed
	if total > 0 {
		failurerate := float64(failed) / float64(successful+failed) * 100
		fmt.Printf("Failure rate: %.2f%%\n", failurerate)
	}
}
