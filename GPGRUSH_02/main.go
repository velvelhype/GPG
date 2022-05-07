package main

import "fmt"
import "math/rand"
import "piscine"

func output(str string, board [][]string) {
	fmt.Print(str);
	fmt.Println("=========\n");
	for h := 0; h < len(board); h++{ 
		for i := 0; i < len(board[h]); i++{ 
			for j := 0; j < len(board[h][i]); j++{
				fmt.Print(string(board[h][i][j]))
				// fmt.Print(" ")
			}
		fmt.Println("")
	}
		fmt.Println("\nanswer--\n")
		piscine.FOVSecuring(board[h])
		// fmt.Println("dummy!!\n")
		fmt.Println("")
	}
}

func randFit()string {
	if rand.Intn(7) == 1 {
		return (string)('0' + rand.Intn(8) + 1)
	} else {
		return (string)('.')
	}
}

func main() {
	fmt.Println("//////////////////////////////////////////");
	fmt.Println("random maps");
	fmt.Println("//////////////////////////////////////////\n");
	board := [][]string{
		{
			".....",
			".....",
			".....",
			".....",
			".....",
		},
	}
	for i := 0; i < 100; i++{
		for i := 0; i < 5; i++{
			board[0][i] = randFit();
			board[0][i] += randFit();
			board[0][i] += randFit();
			board[0][i] += randFit();
			board[0][i] += randFit();
		}
	output("rand_test", board)
	}
	return

	fmt.Println("//////////////////////////////////////////");
	fmt.Println("valid maps");
	fmt.Println("//////////////////////////////////////////\n");

	fmt.Println("\n----success_cases-----");

	board = [][]string{
		{
			"...2.",
			"..6.4",
			"5...6",
			"7.6..",
			".3...",
		},
	}
	output("pdf_v1", board)

	board = [][]string{
		{
			".....",
			"8.8..",
			".7.7.",
			"..8.5",
			".....",
		},
	}
	output("pdf_v2", board)

	board = [][]string{
		{
			"37...",
			"..8..",
			".....",
			"..8..",
			"...69",
		},
	}
	output("pdf_v3", board)


	board = [][]string{
		{
			"9....",
			".....",
			".....",
			".....",
			".....",
		},
	}
	output("pdf_v3", board)

	fmt.Println("\n----fail_cases-----");

	board = [][]string{
		{
			".....",
			".....",
			"..33.",
			"..3..",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			".....",
			".....",
			".....",
			".....",
			".....",
		},
	}
	output("v1", board)

	board = [][]string{
		{
			".....",
			".....",
			".....",
			".....",
			".....",
		},
	}

	fmt.Println("//////////////////////////////////////////");
	fmt.Println("invalid maps---------");
	fmt.Println("//////////////////////////////////////////\n");

	board = [][]string{
		{
			"...",
		},
	}
	output("pdf_inv1", board)

	board = [][]string{
		{
			"...A.",
			"..6.4",
			"5....6",
			"7.6..",
			".3...",
		},
	}
	output("pdf_inv2", board)

	board = [][]string{
		{
			"BBBBB",
			"BBBBB",
			"BB9BB",
			"BBBBB",
			"BBBBB",
		},
	}
	output("pdf_inv2", board)

	board = [][]string{
		{
			".....",
			".....",
			"B....",
			".....",
			".....",
		},
	}
	output("B_contaminated", board)
	return;

	fmt.Println("//////////////////////////////////////////");
	fmt.Println("random ---------");
	fmt.Println("//////////////////////////////////////////\n");

	board = [][]string{
		{
			"...1.",
			"..6.4",
			"5..3.",
			"7.6..",
			".3...",
		},
	}
	output("random", board)

	board = [][]string{
		{
			".....",
			".....",
			".....",
			".....",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			"...2.",
			"..6.4",
			"5....",
			"7.6..",
			".3...",
		},
	}
	output("random", board)

	board = [][]string{
		{
			"...1.",
			"..6.4",
			"5....6",
			"7.6..",
			".3...",
		},
	}
	output("random", board)

	board = [][]string{
		{
			"...3.",
			"..6.4",
			"5....6",
			"7.6..",
			".3...",
		},
	}
	output("random", board)

	board = [][]string{
		{
			"...6.",
			"..6.4",
			"5....6",
			"7.6..",
			".3...",
		},
	}
	output("random", board)

	board = [][]string{
		{
			"...1.",
			"..6.4",
			"5....",
			"7.6..",
			".3...",
		},
	}
	output("random", board)

	board = [][]string{
		{
			".....",
			".....",
			"..3..",
			".....",
			".....",
		},
	}
	output("empty", board)
	board = [][]string{
		{
			".....",
			".....",
			"..3..",
			"...5.",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			".....",
			".22..",
			".....",
			".....",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			".....",
			".2...",
			".....",
			".....",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			".....",
			".3.3.",
			".....",
			".4.3.",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			".....",
			".33..",
			".....",
			".....",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			".....",
			".88..",
			".....",
			".....",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			".....",
			".4.4.",
			".4.4.",
			".....",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			"5....",
			".5...",
			"..5..",
			"...5.",
			"....5",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			"6....",
			".....",
			".6...",
			"...6.",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			"7....",
			".7...",
			"..7..",
			"...7.",
			"....7",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			".3...",
			"..3..",
			"..3..",
			"..3..",
			"..3..",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			"..3..",
			"..3..",
			"..3..",
			"..3..",
			"..3..",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			"..5..",
			"..5..",
			"..5..",
			"..5..",
			"..5..",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			"...3.",
			".5.3.",
			".....",
			".2...",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			"...4.",
			".5...",
			".....",
			".37..",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			"3...5",
			".....",
			".....",
			".....",
			".....",
		},
	}
	output("empty", board)

	board = [][]string{
		{
			"99999",
			"99999",
			"99999",
			"99999",
			"99999",
		},
	}
	output("empty", board)
}
