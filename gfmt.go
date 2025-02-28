package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func main() {
	fmt.Println(Green, "(G)it(F)ormate(M)essage(T)emplate", Reset)

	allowedTypes := []string{"feat", "fix", "docs", "style", "test", "refactor", "perf", "build", "ops", "chore", "revert", "reapply"}

	reader := bufio.NewReader(os.Stdin)

	var typeInput, scopeInput, messageInput, ticketInput string

	for {
		fmt.Print("Enter type: ")
		typeInput, _ = reader.ReadString('\n')
		typeInput = strings.TrimSpace(typeInput)

		isAllowed := false
		for _, t := range allowedTypes {
			if typeInput == t {
				isAllowed = true
				break
			}
		}

		if isAllowed {
			break
		} else {
			fmt.Println(Red, "Error: Invalid type. Allowed types are:", allowedTypes, Reset)
		}
	}

	for {
		fmt.Print("Enter scope (max 5 chars): ")
		scopeInput, _ = reader.ReadString('\n')
		scopeInput = strings.TrimSpace(scopeInput)

		if len(scopeInput) <= 5 {
			break
		} else {
			fmt.Println(Red, "Error: Scope must be 5 characters or less.", Reset)
		}
	}

	for {
		fmt.Print("Enter message (max 40 chars): ")
		messageInput, _ = reader.ReadString('\n')
		messageInput = strings.TrimSpace(messageInput)

		if len(messageInput) <= 40 {
			break
		} else {
			fmt.Println(Red, "Error: Message must be 40 characters or less.", Reset)
		}
	}

	for {
		fmt.Print("Enter Jira ticket number (format KLUE-2025): ")
		ticketInput, _ = reader.ReadString('\n')
		ticketInput = strings.TrimSpace(ticketInput)

		matched, _ := regexp.MatchString(`^[A-Z]{2,4}-\d{1,5}$`, ticketInput)
		if matched {
			break
		} else {
			fmt.Println(Red, "Error: Invalid Jira ticket number format. Expected format is KLUE-2025.", Reset)
		}
	}

	fmt.Printf("%s%s(%s): %s, %s%s\n", Yellow, typeInput, scopeInput, messageInput, ticketInput, Reset)

}
