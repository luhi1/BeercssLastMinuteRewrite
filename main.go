package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type EventInfo struct {
	EventName           string
	Points              int
	EventDescription    string
	EventDate           string
	RoomNumber          int
	AdvisorNames        string
	Location            string
	LocationDescription string
	Sport               string
	SportDescription    string
	Attendance          []StudentAttendance
}

type StudentAttendance struct {
	StudentName   string
	StudentNumber int
}

type Winners struct {
	RandomNinthWinner    string
	RandomTenthWinner    string
	RandomEleventhWinner string
	RandomTwelvthWinner  string
	NinthWinners         []string
	TenthWinners         []string
	EleventhWinners      []string
	TwelvthWinners       []string
}

func main() {
	//ADD THIS IN TO OTHER FILE
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/teacherEvents", loginHandler)
	http.HandleFunc("/winners", logHandler)
	http.HandleFunc("/teacherCreateEvent", logiHandler)
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//ADD EVENTIMAGE INTO HERE ONE DAY
	//ADD ATTENDANCE INTO THE STRUCT AS AN ARRAY OF ANOTHER STRUCT TO STORE ALL THE ATTENDANCE ATTRIBUTES.
	//THE CHECK MARK WILL = THE STUDENT NUMBER ATTRIBUTE OF THAT STRUCT
	//HAVE TO HAVE A WAY TO DELETE SHIT WHEN THEY GET UNCHECKED
	var Events []EventInfo
	for i := 0; i < 3; i++ {
		Events = append(Events, EventInfo{
			EventName:           "asdfasdf",
			Points:              0,
			EventDescription:    "asdf",
			EventDate:           "2017-06-01",
			RoomNumber:          0,
			AdvisorNames:        "asdf",
			Location:            "asdf",
			LocationDescription: "asdf",
			Sport:               "asdf",
			SportDescription:    "asdf",
			Attendance: []StudentAttendance{
				{
					StudentName:   "",
					StudentNumber: 2,
				},
				{
					StudentName:   "1",
					StudentNumber: 1,
				},
			},
		})
	}
	tplExec(w, "teacherEvents.gohtml", Events)
}

func logiHandler(w http.ResponseWriter, r *http.Request) {
	//ADD EVENTIMAGE INTO HERE ONE DAY
	//ADD ATTENDANCE INTO THE STRUCT AS AN ARRAY OF ANOTHER STRUCT TO STORE ALL THE ATTENDANCE ATTRIBUTES.
	//THE CHECK MARK WILL = THE STUDENT NUMBER ATTRIBUTE OF THAT STRUCT
	//HAVE TO HAVE A WAY TO DELETE SHIT WHEN THEY GET UNCHECKED
	tplExec(w, "teacherCreateEvent.gohtml", nil)
}
func logHandler(w http.ResponseWriter, r *http.Request) {
	//ADD EVENTIMAGE INTO HERE ONE DAY
	//ADD ATTENDANCE INTO THE STRUCT AS AN ARRAY OF ANOTHER STRUCT TO STORE ALL THE ATTENDANCE ATTRIBUTES.
	//THE CHECK MARK WILL = THE STUDENT NUMBER ATTRIBUTE OF THAT STRUCT
	//HAVE TO HAVE A WAY TO DELETE SHIT WHEN THEY GET UNCHECKED
	joe := Winners{
		RandomNinthWinner:    "a",
		RandomTenthWinner:    "b",
		RandomEleventhWinner: "c",
		RandomTwelvthWinner:  "d",
		NinthWinners:         []string{"asdf", "adafsd"},
		TenthWinners:         []string{"asdf"},
		EleventhWinners:      []string{"asdf"},
		TwelvthWinners:       []string{"asdf"},
	}
	tplExec(w, "winners.gohtml", joe)
}

func tplExec(w http.ResponseWriter, filename string, information any) error {
	temp := template.Must(template.ParseFiles(filename))

	err := temp.Execute(w, information)
	//@TODO: REMOVE
	if err != nil {
		return err
	}
	return nil
}
