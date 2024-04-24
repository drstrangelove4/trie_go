package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var count int = 0

// --------------------------------------------------------------------------------------------------------------------

func add_word(root *trie_node, word string) {
	// Takes a string, turns it into a hash of that string and then adds it to the trie.

	// Create an array of ints that hash of the word. Exits on failure (invalid runes and numbers).
	hash := hash_word(word)
	if hash == nil {
		return
	}

	// Set a pointer to the working node to allow for easier manipulation.
	current_node := root

	// For each hash in the hash array add a child and preceeding child nodes to those child nodes for each hash in the
	// hash arrry.
	for i := 0; i < len(hash); i++ {
		if current_node.children[hash[i]] == nil {
			current_node.children[hash[i]] = new_node(false, false)
		}
		if i == len(hash)-1 {
			current_node.terminal = true
			return
		}
		current_node = current_node.children[hash[i]]
	}
	fmt.Println(word, " has been added.")
}

// --------------------------------------------------------------------------------------------------------------------

func check_children(node *trie_node) bool {
	// A helper function that looks for children in the passed node. Simply scans for non null values in the child node
	// array that each trie_node holds.

	// Nothing found.
	if node == nil {
		return false
	}

	// If there is a non nil entry then a child is detected.
	for i := 0; i < LETTERS; i++ {
		if node.children[i] != nil {
			return true
		}
	}

	// No child nodes have been found.
	return false
}

// --------------------------------------------------------------------------------------------------------------------

func convert_hash(hash_array []int) string {
	// Takes an array of integers and converts them into runes, building them into a string.

	// A numberic value we want to add to each word to convert back from a hash value to a rune.
	a_rune := int(rune('a'))

	// Return string
	var return_string string

	for i := 0; i < len(hash_array); i++ {
		return_string += string(rune(a_rune + hash_array[i]))
	}
	return return_string
}

// --------------------------------------------------------------------------------------------------------------------

func delete_word(root *trie_node, word string) {
	// Clears termination status from the last node and then attempts to remove child nodes that have no other child
	// nodes from the trie. This function starts from the end of the passed word and works backwards to prevent
	// removing nodes that are used by other words in the trie.

	// Create an array of hashes(ints) for the passed word.
	hash := hash_word(word)
	if hash == nil {
		fmt.Println("Cannot delete word")
		return
	}
	length := len(hash)

	// Find all nodes releated to the word.
	var nodes []*trie_node
	current_node := root
	for i := 0; i < length; i++ {
		if current_node.children[hash[i]] == nil {
			fmt.Println(word, "is not in this trie")
			return
		}
		nodes = append(nodes, current_node)
		current_node = current_node.children[hash[i]]
	}

	// Remove terminal status from last node.
	nodes[length-1].terminal = false

	// Remove childless nodes from the trie.
	for i := length - 1; i >= 0; i-- {
		status := check_children(nodes[i].children[hash[i]])
		if !status {
			nodes[i].children[hash[i]] = nil
		}
	}

	fmt.Println("Removed", word, "from the trie.")
}

// --------------------------------------------------------------------------------------------------------------------

func edit_node(root *trie_node, new_word string, old_word string) {
	// This 'edits' the node. It does this by removing the word passed (if it's in the trie) and adding back in the word
	// The user wants. Doing it this way allows us to reuse previously built functions.

	// Delete the old word from the trie.
	delete_word(root, old_word)

	// Add the new word passed to the trie.
	add_word(root, new_word)

	fmt.Println("Successfully changed", old_word, "to", new_word)

}

// --------------------------------------------------------------------------------------------------------------------

func hash_rune(char rune) (int, error) {
	// Takes an alphabetical rune and converts it to an integer value between 0-25.

	// Get the integer value of the rune 'a'.
	a_rune := rune('a')

	// Convert the letter to a number between 0 and 25. This will give us an index value for the current rune. Indexes
	// outside of 0 to 25 throws an error (the user passed a non alphabetic rune to the function).
	result := int(char) - int(a_rune)
	if result < 0 || result > 25 {
		return result, errors.New("index out of range")
	} else {
		return result, nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

func hash_word(word string) []int {
	// Takes a string and convert it into an array of hashes.

	// Converts the word to lowercase. This allows creation of an array of hashes for all possible alphabetical strings.
	word = strings.ToLower(word)

	// Get the hashes for a word. Return on an index that is out of range to prevent adding partial words.
	var hashes []int
	for i := 0; i < len(word); i++ {

		// Get the value of the current rune.
		current_rune := rune(word[i])
		current_hash, err := hash_rune(current_rune)

		// Error handling.
		if err != nil {
			fmt.Println("there was an error hashing the word:", err)
			return nil

			// Add the hash to the slice.
		} else {
			hashes = append(hashes, current_hash)
		}
	}
	return hashes
}

// --------------------------------------------------------------------------------------------------------------------

func load_file(root *trie_node, location string) {
	// Takes a location from the user and attempts to load words from a file into the trie.

	// Check for valid file extension.
	if location[len(location)-4:] != ".txt" {
		fmt.Println("Invalid file format passed to program. Please enter a '.txt' file")
		return
	}

	// Attempt to load the file at the passed location. Abort procedure if there is an error.
	file, err := os.Open(location)
	if err != nil {
		fmt.Println("There was an error loading the file:", err)
		return
	}
	// Defer the file close until it is no longer used.
	defer file.Close()

	// Parse the file for words. Scans file for words (file splits on whitespace). Abort procedure if there is an error.
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		add_word(root, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("There was an error parsing that file:", err)
		return
	}

	fmt.Println("File has been loaded in the trie.")
}

// --------------------------------------------------------------------------------------------------------------------

func new_node(root_status bool, terminal_status bool) *trie_node {
	// Creates and returns a pointer to a new trie node.

	new_node := &trie_node{}
	new_node.is_root = root_status
	new_node.terminal = terminal_status
	return new_node
}

// --------------------------------------------------------------------------------------------------------------------

func print_trie(root *trie_node) {
	// Initialises the print array function. Checks for an empty trie and passes an empty array to the recursive function
	// used to print the trie.

	// Deal with empty trie
	if root == nil {
		fmt.Println("The trie is empty")
		return
	}
	// Call the recursive function that will print the trie.
	var start_array []int
	print_trie_rec(root, start_array)

	// Report the amount of words stored in the trie, reset the counter to prevent double counts on future calls.
	fmt.Println("\nThere are", count, "words in the trie")
	count = 0
}

func print_trie_rec(root *trie_node, hash_index []int) {
	// A recursive function which interates over the array and finds populated child nodes, stores hash values
	// and calls the function to convert them into a string on terminal status.

	for i := 0; i < LETTERS; i++ {

		// Skip empty values
		if root.children[i] != nil {
			// Copy the array to prevent values leaking into the next word.
			array_copy := hash_index

			// Add a hash to the array
			array_copy = append(array_copy, i)

			// Print the array if we encounter a terminal bool = true
			if root.terminal {
				fmt.Println(convert_hash(array_copy))
				count++
			}

			// Continue down the trie
			print_trie_rec(root.children[i], array_copy)
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------

func search_word(root *trie_node, search string) {
	// Searches the trie for a string.

	// Convert the string to an array of ints.
	hash := hash_word(search)
	if hash == nil {
		return
	}

	// Test to see if there is a value in the node.
	current_node := root
	for i := 0; i < len(hash); i++ {
		// Unable to find a trie entry that fully matches the hash.
		if current_node.children[hash[i]] == nil {
			fmt.Println(search, "not found in this trie.")
			return
		}

		// Finds a terminal node that matches the hash array.
		if i == len(hash)-1 && current_node.terminal {
			fmt.Println("Trie contains", search)
			return

			// Terminal node not found at the end of the hash array.
		} else if i == len(hash)-1 && !current_node.terminal {
			fmt.Println(search, "not found in this trie.")
		}
		current_node = current_node.children[hash[i]]
	}
}

// --------------------------------------------------------------------------------------------------------------------

func take_input_int(prompt string) int {
	// Prompts the user for numerical input until a valid number is entered. Returns that number to the calling function.

	for {
		fmt.Println(prompt)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		converted_value, err := strconv.Atoi(strings.TrimSpace(input))
		if err == nil {
			return converted_value
		} else {
			fmt.Println("Invalid input")
			continue
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------

func take_input_string(prompt string) string {
	// Uses a reader to take input from stdin and returns the string with preceeding and trailing whitespace removed.

	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// --------------------------------------------------------------------------------------------------------------------
