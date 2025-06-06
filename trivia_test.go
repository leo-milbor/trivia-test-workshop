package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"testing"
)

func Test_GoldenMaster(t *testing.T) {
	for i := range 1_000 {
		t.Run(fmt.Sprint("golden master ", i), func(t *testing.T) {
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("failed to create pipe: %w", err)
			}
			stdoutBackup := os.Stdout
			os.Stdout = w

			notAWinner := false

			game := NewGame()

			game.Add("Chet")
			game.Add("Pat")
			game.Add("Sue")

			rd := rand.New(rand.NewSource(int64(i)))

			for {
				game.Roll(rd.Intn(5) + 1)

				if rd.Intn(9) == 7 {
					notAWinner = game.WrongAnswer()
				} else {
					notAWinner = game.WasCorrectlyAnswered()

				}

				if !notAWinner {
					break
				}
			}

			b := bytes.Buffer{}
			err = w.Close()
			io.Copy(&b, r)
			os.Stdout = stdoutBackup
			got := b.String()
			// saveGoldenMaster(t, got, i)
			compareGoldenMaster(t, got, i)
		})
	}
}

func saveGoldenMaster(t *testing.T, want string, i int) {
	file, err := os.Create(fmt.Sprint("testdata/golden_master_", i))
	if err != nil {
		t.Fatalf("create golden master: %v", err)
	}
	if _, err := file.WriteString(want); err != nil {
		t.Fatalf("write golden master: %v", err)
	}
}

func compareGoldenMaster(t *testing.T, got string, i int) {
	file, err := os.ReadFile(fmt.Sprint("testdata/golden_master_", i))
	if err != nil {
		t.Fatalf("read golden master: %v", err)
	}
	want := string(file)
	if want != got {
		t.Fatalf("compare golden master:\nWANT:\n%s\n\nGOT:\n%s", want, got)
	}
}
