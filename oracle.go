// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"uppg1"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)
	go func() {
		for {
			go prophecy(<-questions, answers)
		}
	}()
	go func() {
		for {
			go prophecy("", answers)
			time.Sleep(time.Duration(15+rand.Intn(45)) * time.Second)
		}
	}()
	go func() {
		for {
			fmt.Println(<-answers)
		}
	}()
	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)

	if question != "" {
		// Find the longest word.
		longestWord := ""
		words := strings.Fields(strings.Trim(question, ".?!")) // Fields extracts the words into a slice.
		for _, w := range words {
			if len(w) > len(longestWord) {
				longestWord = w
			}
		}

		// Cook up some pointless nonsense.
		nonsense := []string{
			"It is certain",
			"It is decidedly so",
			"Without a doubt",
			"Yes – definitely",
			"You may rely on it",
			"As I see it, yes",
			"Most likely",
			"Outlook good",
			"Yes",
			"Signs point to yes",
			"Reply hazy, try again",
			"Ask again later",
			"Better not tell you now",
			"Cannot predict now",
			"Concentrate and ask again",
			"Don't count on it",
			"My reply is no",
			"My sources say no",
			"Outlook not so good",
			"Very doubtful"}
		answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
	} else {
		nonsense := []string{
			"The moon is dark.",
			"The sun is bright.",
		}
		answer <- nonsense[rand.Intn(len(nonsense))]
	}
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
