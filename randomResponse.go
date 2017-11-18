package main 

import(
    "fmt"
    "math/rand"
    "time"
    "regexp"
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


    return eliza[rand.Intn(len(eliza))]
}

