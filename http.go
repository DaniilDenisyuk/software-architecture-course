package main
import(
    "encoding/json"
    "net/http"
    "time"
)

func main() {
    http.Handle("/time", http.HandlerFunc(GetTime))
    http.ListenAndServe(":8795", nil)	
}

func GetTime(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(time.RFC3339)
    json, err := json.Marshal("{ time: "+string(tm)+"}")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if r.Method == http.MethodGet {
        w.Header().Set("Content-Type", "application/json")
        w.Write(json)
    }
}	


