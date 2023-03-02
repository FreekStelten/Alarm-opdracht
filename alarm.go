package main

import (
    "fmt" //pakket om uitvoer naar de console af te drukken
    "strconv" //pakket om de gebruikersinvoer van een tekenreeks naar een geheel getal te converteren.	
)

func main() {
    // Ask the user for a number
    var input string
    fmt.Print("voer een geldig nummer in: ") //vraagt de gebruiker om hier input in te zetten	
    fmt.Scanln(&input) //hier wordt de input opgehaald en bewaard voor de volgende code


    // Convert the number string to an integer
    num, err := strconv.Atoi(input)
    if err != nil { //strconv.atoi controlleert de input of het een geldig nummer is, anders print die de foutmelding uit.
        fmt.Println("Dit is geen geldig nummer.")
        return  //hiermee stopt die de code.
    }

    // Print the word with an increasing number
    for i := 1; i <= num; i++ {  //hiermee telt die steeds op en zet een loop op tot het NUM (ingevoerde getal) is bereikt.
        fmt.Printf("Alarm %d!\n", i) // hiermee print die de optelling uit. het start bij 1 en gaat per 1 omhoog. /n betekend dat 
		// er steeds een nieuwe regel moet beginnen
    }
}