package main

import (
	"fmt"
	"os"

	"github.com/inancgumus/screen"
)

// --------------------------------------------------------------------------------------------------------------------

func add_node_frontend(root *trie_node) {
	// User facing function that takes input and calls add node.

	screen.Clear()
	screen.MoveTopLeft()

	// Take user input
	user_input := take_input_string("Enter a word to add:\n--------------------")
	add_word(root, user_input)
	continue_()
	menu(root)
}

// --------------------------------------------------------------------------------------------------------------------

func edit_node_frontend(root *trie_node) {
	// User facing function that takes input and calls edit node.

	screen.Clear()
	screen.MoveTopLeft()

	// Take user input
	user_input := take_input_string("Enter a node to edit:\n--------------------")
	new_names := take_input_string("New word: ")
	edit_node(root, new_names, user_input)
	continue_()
	menu(root)
}

// --------------------------------------------------------------------------------------------------------------------

func delete_word_frontend(root *trie_node) {
	// User facing function that takes input and calls delete node.

	screen.Clear()
	screen.MoveTopLeft()

	// Take user input
	user_input := take_input_string("Enter a word to delete:\n--------------------")
	delete_word(root, user_input)
	continue_()
	menu(root)
}

// --------------------------------------------------------------------------------------------------------------------

func continue_() {
	// Pauses the program until the user continues
	for {
		choice := take_input_int("\nPress 1 to continue")
		if choice == 1 {
			break
		} else {
			fmt.Println("Invalid input")
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------

func load_file_frontend(root *trie_node) {

	screen.Clear()
	screen.MoveTopLeft()

	user_input := take_input_string("Enter the path to a text file:\n--------------------")
	load_file(root, user_input)
	continue_()
	menu(root)
}

// --------------------------------------------------------------------------------------------------------------------

func menu(root *trie_node) {
	// Program main menu

	screen.Clear()
	screen.MoveTopLeft()

	// Display menu options
	fmt.Println("Trie\n-----")
	fmt.Println("1) Search for word")
	fmt.Println("2) Add word")
	fmt.Println("3) Edit word")
	fmt.Println("4) Delete word")
	fmt.Println("5) Print trie")
	fmt.Println("6) Load File")
	fmt.Println("7) Exit")
	choice := take_input_int("-------------------")
	// Check for valid inputs before calling the switch. Reload menu if invalid input given.
	if choice < 1 || choice > 7 {
		menu(root)
	} else {
		select_option(root, choice)
	}
}

// --------------------------------------------------------------------------------------------------------------------

func print_trie_frontend(root *trie_node) {
	// User facing function that prints the contents of the trie.

	screen.Clear()
	screen.MoveTopLeft()

	print_trie(root)
	continue_()
	menu(root)
}

// --------------------------------------------------------------------------------------------------------------------

func search_word_frontend(root *trie_node) {
	// User facing function that takes input and calls the search for word function.

	screen.Clear()
	screen.MoveTopLeft()

	// Take user input
	user_input := take_input_string("Enter a word to search for:\n--------------------")
	search_word(root, user_input)
	continue_()
	menu(root)
}

// --------------------------------------------------------------------------------------------------------------------
func select_option(root *trie_node, choice int) {

	// Converts user input into a function call.

	switch choice {
	case 1:
		search_word_frontend(root)
	case 2:
		add_node_frontend(root)
	case 3:
		edit_node_frontend(root)
	case 4:
		delete_word_frontend(root)
	case 5:
		print_trie_frontend(root)
	case 6:
		load_file_frontend(root)
	case 7:
		os.Exit(1)
	default:
		menu(root)
	}
}

// --------------------------------------------------------------------------------------------------------------------
