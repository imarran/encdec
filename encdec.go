package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const debugMode int = 0 // setting to a 1 will print debug info to the console

func generateHashTable(NoOfSets int) []string {
	var strSet = []string{" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz\t\r\n"}
	strTmp := ""
	strChar := ""
	strPos := 0

	for i := 1; i <= NoOfSets; i++ {
		strTmp = strSet[0]
		strNewSet := ""
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		//fmt.Print("Started with ")
		for len(strTmp) > 0 { // while there is still length...
			//fmt.Println(strTmp)
			strPos = r.Intn(len(strTmp)) //get a random character position
			strChar = strTmp[strPos : strPos+1]
			strNewSet = strNewSet + strChar                  //fetch the caracter and append it to a new string
			strTmp = strings.Replace(strTmp, strChar, "", 1) //remove the used character
			//fmt.Print(strPos)
			//fmt.Print(" removed " + strChar + " ")
		}
		//fmt.Println()
		strSet = append(strSet, strNewSet)
	}
	return strSet
}

func encryptStr(strIn string, hashIn []string) string {
	strOutput := ""
	keyIndex := 0
	strPos := 0
	hashKeys := len(hashIn)
	keyStr := ""

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	startIndex := r.Intn(hashKeys) //get a random starting place in the hash table
	keyIndex = startIndex

	if debugMode == 1 {
		fmt.Print("len(hashIn): ")
		fmt.Println(len(hashIn))
		fmt.Print("keyIndex: ")
		fmt.Println(keyIndex)
	}

	for strPos < len(strIn) { //iterate character by character
		//find the current character's position in the base string (index 0)
		inPos := strings.Index(hashIn[0], strIn[strPos:strPos+1])
		if debugMode == 1 {
			fmt.Print("inPos in loop: ")
			fmt.Println(inPos)
		}

		//assign and append the matching character in the keyString
		if debugMode == 1 {
			fmt.Print("keyIndex in loop: ")
			fmt.Println(keyIndex)
		}
		keyStr = hashIn[keyIndex]
		if debugMode == 1 {
			fmt.Println("keyStr in loop: " + keyStr)
		}
		strOutput += keyStr[inPos : inPos+1]

		//move to the next key
		keyIndex += 1
		if keyIndex >= len(hashIn) {
			keyIndex = 1
		}

		//move to the next character
		strPos += 1
	}

	if startIndex < 9 {
		strOutput += "0"
	}
	strOutput += fmt.Sprintf("%v", startIndex)
	return strOutput
}

func decryptStr(strIn string, hashIn []string) string {
	strOutput := ""
	keyIndex := 0
	strPos := 0
	keyStr := ""
	tmpStr := ""

	//get the starting place in the hash table from last 2 chars in strIn
	tmpStr = strIn[len(strIn)-2 : len(strIn)]
	startIndex, err := strconv.Atoi(tmpStr)
	if err != nil {
		return "Your input string is not valid for this decryption!"
	}
	if debugMode == 1 {
		fmt.Println("Found start key index: " + tmpStr)
	}
	strIn = strings.TrimSuffix(strIn, tmpStr)
	if debugMode == 1 {
		fmt.Println("Altered strIn: " + strIn)
	}
	keyIndex = startIndex

	if debugMode == 1 {
		fmt.Print("len(hashIn): ")
		fmt.Println(len(hashIn))
		fmt.Print("keyIndex: ")
		fmt.Println(keyIndex)
	}

	for strPos < len(strIn) { //iterate character by character
		//find the current character's position in the current key string
		inPos := strings.Index(hashIn[keyIndex], strIn[strPos:strPos+1])
		if debugMode == 1 {
			fmt.Print("inPos in loop: ")
			fmt.Println(inPos)
		}

		//assign and append the matching character in the base string (hashIn[0])
		keyStr = hashIn[0]
		strOutput += keyStr[inPos : inPos+1]

		//move to the next key
		keyIndex += 1
		if keyIndex >= len(hashIn) {
			keyIndex = 1
		}

		//move to the next character
		strPos += 1
	}

	return strOutput
}

func main() {
	originalStr := "the original string...before we mess it up :)"

	var hashTable = []string{}
	hashTable = generateHashTable(9) // you can now see hashTable[0] .. [n]
	if debugMode == 1 {
		for _, value := range hashTable {
			fmt.Println(value)
		}
		fmt.Println("\n\n")
		fmt.Println("Original String: " + originalStr)
	}

	// Encrypt the string
	fmt.Print("Starting Encryption...")
	encryptedStr := encryptStr(originalStr, hashTable)
	fmt.Println("...Encryption Complete")
	fmt.Println("->" + encryptedStr + "<-")

	// Decrypt the string
	fmt.Print("Starting Decryption...")
	decryptedStr := decryptStr(encryptedStr, hashTable)
	fmt.Println("...Decryption Complete")
	fmt.Println("->" + decryptedStr + "<-")
}
