package main

import (
        "net/http"
        "io/ioutil"
        "time"
        "log"
        "os"
        "math/rand"
        "strconv"
        "fmt"
        "github.com/RayofLightz/tardis/core"
)

type BruteTry struct {
        username string
        password string
}

func NewBruteTry(user string, password string) BruteTry{
        var tmp BruteTry
        tmp.username = user
        tmp.password = password
        return tmp
}

func (brute *BruteTry) logEntry(){
        log.Println("Login attempt for user " + brute.username + " with password " + brute.password)
}

func homeFunc(writer http.ResponseWriter, req *http.Request){
        //Set the header to trick bots into thinking that this is a real webcam
        writer.Header().Set("Server", "SQ-WEBCAM")
        f, err := ioutil.ReadFile("index.html")
        if err != nil{
                log.Println(err)
        }
        writer.Write(f)
}

func asciiArt(){
        //Because all fun tools have ascii art ;) 
        policeBox := `
            ___
    _______(_@_)_______
    | POLICE      BOX |
    |_________________|
     | _____ | _____ |
     | |###| | |###| |
     | |###| | |###| |   
     | _____ | _____ |   
     | || || | || || |
     | ||_|| | ||_|| |  
     | _____ |$_____ |  
     | || || | || || |  
     | ||_|| | ||_|| | 
     | _____ | _____ |
     | || || | || || |   
     | ||_|| | ||_|| |         
     |       |       |        
     *****************
`
    fmt.Println(policeBox)
    fmt.Println("Tardis v1")
}

func main(){
    //This is all fairly routine golang http code
    //Check out https://golang.org/pkg/net/http/
    //For documentation
    cnf, err := core.LoadConfig()
    if err != nil{
            log.Println(err)
    }
    f, err := os.OpenFile(cnf.LogPath, os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0444)
    log.SetOutput(f)
    http.HandleFunc("/", homeFunc)
    http.HandleFunc("/css.css", func (w http.ResponseWriter, req *http.Request){http.ServeFile(w, req, "css.css")})
    fileHandle := http.FileServer(http.Dir("public/jpg"))
    http.Handle("/jpg/", http.StripPrefix("/jpg/", fileHandle))
    http.HandleFunc("/home.htm", func (writer http.ResponseWriter, req *http.Request){
            writer.Header().Set("Server", "SQ-WEBCAM")
            if req.Method == "POST"{
                err := req.ParseForm()
                if err != nil{
                    log.Println(err)
                }
                //check for the values
                creds := NewBruteTry(req.PostForm.Get("username"), req.PostForm.Get("password"))
                creds.logEntry()
                if cnf.EnableFakeSuccessPage == true{
                        //Random seed for fake sucess generator
                        rand.Seed(time.Now().UnixNano())
                        randNum := rand.Intn(5)
                        fmt.Println(randNum)
                        if randNum == 4{
                                log.Println("Serving fake success page")
                                content, err := ioutil.ReadFile(cnf.FakeSuccessPage)
                                if err != nil{
                                        log.Println(err)
                                }else{
                                        writer.Write(content)
                                }
                        }else{
                                writer.Write([]byte(`<html><body onload="alert('Username or password is invalid')"></body></html>`))
                        }
                } else{
                        writer.Write([]byte(`<html><body onload="alert('Username or password is invalid')"></body></html>`))
                }
            }
    })
    asciiArt()
    addr := ":" + strconv.Itoa(cnf.Port)
    log.Fatal(http.ListenAndServe(addr, nil))
}
