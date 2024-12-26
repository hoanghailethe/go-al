package cmd

import "fmt"

func isValid(s string) bool {
    stack := make(Stack, 0)
    for _, v := range(s) {
        // fmt.Print(i)
        switch os := string(v); os {
            case "(": stack=stack.Push(")")
            case "[": stack=stack.Push("]")
            case "{": stack=stack.Push("}")
            default: 
                stack2, value := stack.Pop()
                stack = stack2
                // fmt.Println(value)
                if value != os {
                    return false
                }
                // fmt.Println(stack)
        }
    }
    return len(stack)==0   
}


type Stack []string
 
func (s Stack) Push(str string) Stack {
    return append(s, str)
}

func (s Stack) Pop() (Stack, string) {
    if (len(s) == 0) {return []string{}, ""}
    return s[:len(s)-1], string(s[len(s)-1])
}


// GPT:
func isValid(s string) bool {
    stack := make([]rune, 0) // Use []rune for better performance with characters
    for _, char := range s {
        switch char {
        case '(': 
            stack = append(stack, ')') // Push the expected closing parenthesis
        case '[': 
            stack = append(stack, ']')
        case '{': 
            stack = append(stack, '}')
        default:
            // Pop the top element from the stack
            if len(stack) == 0 || stack[len(stack)-1] != char {
                return false
            }
            stack = stack[:len(stack)-1] // Remove the top element
        }
    }
    return len(stack) == 0 // Valid if stack is empty
}

