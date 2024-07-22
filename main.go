package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	jokeStore := []string{"Wife told me to take the spider out instead of killing it... We had some drinks, cool guy, wants to be a web developer.",
		"Why did the scarecrow win an award? Because he was outstanding in his field.",
		"Is the pool safe for diving? It deep ends.",
		"The shovel was a ground-breaking invention.",
		"I made a belt out of watches once... It was a waist of time."}
	n := len(jokeStore)
	fmt.Fprintf(w, jokeStore[rand.Intn(n)])
}

func handlerApi(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	req.Header.Add("Accept", "text/plain")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprint(w, string(body))

}

func madHandler(w http.ResponseWriter, r *http.Request) {
	names := map[string]string{
		"Josh":        "he",
		"June":        "she",
		"Emily":       "she",
		"Michael":     "he",
		"Sarah":       "she",
		"David":       "he",
		"Anna":        "she",
		"John":        "he",
		"Laura":       "she",
		"Daniel":      "he",
		"Sophia":      "she",
		"James":       "he",
		"Emma":        "she",
		"Matthew":     "he",
		"Olivia":      "she",
		"Anthony":     "he",
		"Grace":       "she",
		"Andrew":      "he",
		"Hannah":      "she",
		"Christopher": "he",
	}
	occupations := []string{
		"Software Engineer",
		"Data Scientist",
		"Graphic Designer",
		"Product Manager",
		"Teacher",
		"Nurse",
		"Electrician",
		"Chef",
		"Mechanical Engineer",
		"Doctor",
		"Artist",
		"Writer",
		"Architect",
		"Marketing Specialist",
		"Sales Manager",
		"Pharmacist",
	}
	computingDevices := []string{
		"a Laptop",
		"a Desktop Computer",
		"a Tablet",
		"a Smartphone",
		"a Smartwatch",
	}
	bodyParts := []string{
		"Head",
		"Arm",
		"Leg",
		"Hand",
		"Foot",
		"Eye",
		"Ear",
		"Nose",
		"Mouth",
		"Torso",
		"Shoulder",
		"Knee",
		"Elbow",
		"Wrist",
		"Ankle",
	}
	moodWords := []string{
		"Happy",
		"Sad",
		"Excited",
		"Angry",
		"Bored",
		"Nervous",
		"Content",
		"Joyful",
		"Frustrated",
		"Anxious",
		"Relaxed",
		"Surprised",
		"Depressed",
		"Hopeful",
		"Irritated",
	}
	actionWords := []string{
		"Run",
		"Walk",
		"Jump",
		"Swim",
		"Dance",
		"Sing",
		"Write",
		"Read",
		"Cook",
		"Drive",
		"Climb",
		"Draw",
		"Talk",
		"Listen",
		"Sleep",
	}
	var keys []string
	for key := range names {
		keys = append(keys, key)
	}
	name := keys[rand.Intn(len(keys))]
	occupation := occupations[rand.Intn(len(occupations))]
	computing_device := computingDevices[rand.Intn(len(computingDevices))]
	body_part := bodyParts[rand.Intn(len(bodyParts))]
	mood_word := moodWords[rand.Intn(len(moodWords))]
	action_word := actionWords[rand.Intn(len(actionWords))]

	strArr := []string{
		name,
		" is a ",
		strconv.Itoa(20 + rand.Intn(60)),
		"-year old ",
		occupation,
		" who has been struggling with a lot of job-related stress. He/she decided to try a new application to relieve stress, which runs on ",
		computing_device,
		" to help improve his/her mood. The application senses his/her mood through a device he/she wears on his/her ",
		body_part,
		". When the device senses that he/she is ",
		mood_word,
		" it responds ",
		action_word,
		"."}

	fmt.Fprint(w, strings.Join(strArr, ""))

}

func main() {

	http.HandleFunc("/joke", handler)
	http.HandleFunc("/madlib", madHandler)
	http.HandleFunc("/joke/api", handlerApi)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
