package main

import (
        "fmt"
        "net/http"
        "log"
        "io/ioutil"
        "crypto/rand"
        "github.com/sendgrid/sendgrid-go"
        "os/exec"
        "strings"
)

//encode decode router
func CodeRouter (w http.ResponseWriter, req *http.Request) {
        //Get Email Values
        to := req.FormValue("from")
        subject := req.FormValue("subject")
        subject = "Re: "+subject
        body:= req.FormValue("text")

        //Get Uploaded File
         fname= getUpload("attachment1")
        
        if strings.Contains(subject, "encode") {
           fmt.Println("encode accepted")
           encode(to,subject,body,fname)
        }

        if strings.Contains(subject, "decode") {
           fmt.Println("decode accepted")
           decode(to,subject,body,fname
        }
}

func getUpload(formName) string {
        //returns actual file name
        file, handler, err := req.FormFile(formName)
        if err != nil {
                fmt.Println(err)
        }
        data, err := ioutil.ReadAll(file)
        if err != nil {
                fmt.Println(err)
        }
        err = ioutil.WriteFile(handler.Filename, data, 0777)
        if err != nil {
                fmt.Println(err)
        }

return handler.Filename
}

func encode(to string, subject string, body string, fileName string) {
        msg := embedFile(fileName, body)
        emailBack(fileName, subject, to, msg)

}

func decode (to string, subject string, body string, fileName string) {
        msg := extractMsg(fileName)
        emailBack(fileName, subject, to, msg)
}

func extractMsg(fileName string) string {
        fname := randString(10)
        out, err1 := exec.Command("steghide", "extract", "-sf", fileName, "-p", "sendgrid", "-xf", fname).Output()
        fmt.Printf(string(out))
        if err1 != nil {
                log.Fatal(err1)
        }


        body, err := ioutil.ReadFile(fname)
        if err != nil { panic(err) }

        out2, err2 := exec.Command("rm", "-f", fname).Output()
        fmt.Printf(string(out2))
        if err2 != nil {
                log.Fatal(err2)
        }


        return string(body)

}

func embedFile(carrierFile string, embedText string) string {

        fname := randString(10)
        err := ioutil.WriteFile(fname, []byte(embedText), 0644)
        if err != nil {
                log.Fatal(err)
        }


        out, err1 := exec.Command("steghide", "embed", "-cf", carrierFile, "-ef", fname, "-p", "sendgrid").Output()
        fmt.Printf(string(out))
        if err1 != nil {
                log.Fatal(err1)
        }


        out2, err2 := exec.Command("rm", "-f", fname).Output()
        fmt.Printf(string(out2))
        if err2 != nil {
                log.Fatal(err2)
        }

        body:= "The attached file is now encoded with the body of your previous email. To see hidden text, send another email with the attached file below and set the subject to \"Decode\""

        return body
}

func emailBack(fileName string, subject string, to string, body string) {

        sg := sendgrid.NewSendGridClient("username", "password")
        message := sendgrid.NewMail()
        message.AddTo(to)
        message.AddSubject(subject)
        message.AddText(body)
        message.AddFrom("your address")
        message.AddAttachment(fileName)
        if r := sg.Send(message); r == nil {
           fmt.Println("Email sent!")
        } else {
           fmt.Println(r)
        }

        //deletes image locally after emailing it
        out2, err2 := exec.Command("rm", "-f", fileName).Output()
        fmt.Printf(string(out2))
        if err2 != nil {
                log.Fatal(err2)
        }

}

func randString(n int) string {
    const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var bytes = make([]byte, n)
    rand.Read(bytes)
    for i, b := range bytes {
        bytes[i] = alphanum[b % byte(len(alphanum))]
    }
    return string(bytes)
}

func main() {
        http.HandleFunc("/upload", CodeRouter)
        err := http.ListenAndServe(":3000", nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }

