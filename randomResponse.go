package main 

import(
    "fmt"
    "math/rand"
    "time"
    "regexp"
    "strings"
)

func main(){

    //asking eliza questions 
    fmt.Println("I’m not sure what you’re trying to say. Could you explain it to me?")
    fmt.Println(ElizaResponse("I’m not sure what you’re trying to say. Could you explain it to me?"))
    fmt.Println("Father was a teacher.")
    fmt.Println(ElizaResponse("Father was a teacher."))
    fmt.Println("I’m looking forward to the weekend.")
    fmt.Println(ElizaResponse("I’m looking forward to the weekend."))
    fmt.Println("I was my father’s favourite.")
    fmt.Println(ElizaResponse("I was my father’s favourite."))
    fmt.Println("My grandfather was French!")
    fmt.Println(ElizaResponse("My grandfather was French!"))
    fmt.Println("I am happy.")
    fmt.Println(ElizaResponse("I am happy."))
    fmt.Println("I am not happy with your responses.")
    fmt.Println(ElizaResponse("I am not happy with your responses."))
    fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
    fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
    fmt.Println("I am supposed to just take what you’re saying at face value?")
    fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))
    fmt.Println("how are you eliza?")
    fmt.Println(ElizaResponse("how are you eliza?"))

    

    rand.Seed(time.Now().UTC().UnixNano())    
}
//returs an answerr to user's question
func ElizaResponse(input string )string{

    //creating random response
    eliza := [] string{
        "I’m not sure what you’re trying to say. Could you explain it to me?",
        "How does that make you feel?",
        "Why do you say that?",
    }

    //Recognise “father”
    matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`,input)
    //if father is true then return why dont you.... 
    if matched == true{
         return "why don't you tell me more about your father?"
    }

    //part 3
    // re := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)
    //“i am “, “I AM “, “I’m “, “Im “, “i’m “
    re := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)
 
 	if match := re.MatchString(input); match {
 
 		return Reflect(input)
 	}
    //part 6
    matc, _ := regexp.MatchString(`how are you(.*)`,input)
    //if father is true then return why dont you.... 
    if matc == true{
         return "I am good what about you?"
    }



    return eliza[rand.Intn(len(eliza))]
}

//https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
func Reflect(input string) string {
	// Split the input on word boundaries.
    
	tokens := strings.Split(input, " ")
	
	// List the reflections.
	reflections := [][]string{
		{"your", "my"},
 		{"you’re", "i am"},
 		{"I", "you"},
 		{"you", "I"},
 		{"me", "you"},
        {"am", "are"},
	}
	
	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}
	
	// Put the tokens back together.
    joinString := strings.Join(tokens, " ")
    return fmt.Sprintf("How do you know %s?", joinString)
}