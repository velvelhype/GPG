package main

import "fmt"
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

func main() {
	fmt.Println("//////////////////////////////////////////");
	fmt.Println("valid maps");
	fmt.Println("//////////////////////////////////////////\n");

	fmt.Println("\n----success_cases-----");

	board := [][]string{
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
	output("empty", board)

	board = [][]string{
		{
			"1....",
			".....",
			".....",
			".....",
			".....",
		},
	}
	output("inv_number", board)

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

	fmt.Println("//////////////////////////////////////////");
	fmt.Println("random ---------");
	fmt.Println("//////////////////////////////////////////\n");	
}
