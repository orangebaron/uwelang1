package compiler

import "fmt"

const whitespace = " \n\t"

type codeStructure struct {
	beginning string
	end       string
	args      []*codeStructure
}

func parseCode(code string) {
	keywords := []string{"struct"}
	for loc := 0; loc < len(code); loc++ {
		//check for keywords
		for _, keyword := range keywords {
			if code[loc:len(keyword)+1] == keyword+" " {
				fmt.Println("keyword found yay", keyword)
			}
		}
	}
}
