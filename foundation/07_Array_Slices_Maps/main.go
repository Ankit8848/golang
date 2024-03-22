package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("For Array in Go")

	// specify before using array

	var studentList [10]string

	// 1st way to declare

	studentList[0] = "ankit"
	studentList[4] = "aman"
	studentList[8] = "mental"

	fmt.Println("Student list is: ", studentList)
	fmt.Println("Student list is: ", len(studentList))

	// 2nd way to declare & initialize

	var friendList = [2]string{"ankit", "priya"}

	fmt.Println("friends is: ", friendList)

	// for slices =>  is a dynamically-sized, flexible view into the elements of an array.

	fmt.Println("For Slices in Go")
	// for slices, we dont have to specify the no.

	// 1st way to use slices
	var friendsList = []string{"ankit", "priya"}
	fmt.Println("Friend list is: ", len(friendsList))

	// here, we can't just add like arrays . we have to use append method

	// 1st comma syntax using
	friendsList = append(friendsList, "sup", "sona")
	fmt.Println(friendsList)

	// 2nd square and colon syntax using
	friendsList = append(friendsList[1:3])
	// last range is non-inclusive. it means it will not be included
	fmt.Println(friendsList)

	// 2nd way to use slices &
	marksExam := make([]int, 3)

	marksExam[0] = 77
	marksExam[1] = 55
	marksExam[2] = 99

	marksExam = append(marksExam, 44, 32, 24, 75)
	// so, append method will do allocation again

	fmt.Println(marksExam)

	sort.Ints(marksExam)
	// this is a package which helps a lot while sorting. but not in the array.
	fmt.Println(marksExam)
	fmt.Println(sort.IntsAreSorted(marksExam))

	// to remove a value from slices based on index

	var progLang = []string{"javascript", "python", "golang", "ruby", "java"}
	fmt.Println(progLang)
	var index int = 2
	progLang = append(progLang[:index], progLang[index+2:]...)
	fmt.Println(progLang)

	// for maps
	fmt.Println("For Maps in Go")

	progsLang := make(map[string]string)

	progsLang["JS"] = "Javascript"
	progsLang["RB"] = "Ruby"
	progsLang["PY"] = "Python"

	fmt.Println("List of all languages: ", progsLang)
	fmt.Println("JS short for: ", progsLang["JS"])
	// here to access we use key same as taken,if it is string, int, ...

	// to remove in maps

	delete(progsLang, "PY")
	fmt.Println("List of all languages: ", progsLang)

	// loops for maps in go

	// for _, value => also can be use. if you dont care about key and vice versa.
	for key, value := range progsLang {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
