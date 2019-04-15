package main

import "fmt"

func main() {
	s := []string{"this", "th", "is", "famous",
		"Word", "break", "b", "r", "e", "a", "k",
		"br", "bre", "brea", "ak", "problem"}
	word := "Wordbreakproblem"
	wordBreak(s, word, "")
}

func wordBreak(s []string, word string, output string) {
	if len(word) == 0 {
		fmt.Println("output", output)
		return
	}
	for i := 1; i <= len(word); i++ {
		prefix := word[0:i]
		if contains(prefix, s) {
			wordBreak(s, word[i:], output+" "+prefix)
		}
	}
}

func contains(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

// public static void wordBreak(List<String> dict, String str, String out)
// 	{
// 		// if we have reached the end of the String,
// 		// print the output String

// 		if (str.length() == 0)
// 		{
// 			System.out.println(out);
// 			return;
// 		}

// 		for (int i = 1; i <= str.length(); i++)
// 		{
// 			// consider all prefixes of current String
// 			String prefix = str.substring(0, i);

// 			// if the prefix is present in the dictionary, add prefix to the
// 			// output String and recurse for remaining String

// 			if (dict.contains(prefix)) {
// 				wordBreak(dict, str.substring(i), out + " " + prefix);
// 			}
// 		}
// 	}

// 	// main function
// 	public static void main(String[] args)
// 	{
// 		// List of Strings to represent dictionary
// 		List<String> dict = Arrays.asList("this", "th", "is", "famous",
// 									"Word", "break", "b", "r", "e", "a", "k",
// 									"br", "bre", "brea", "ak", "problem");

// 		// input String
// 		String str = "Wordbreakproblem";

// 		wordBreak(dict, str, "");
// 	}
