package main

import (
// FIXME: we should really parse the XML from Gomez.
//    "encoding/xml"
    "fmt"
    "github.com/bmizerany/pat"
    "github.com/scorredoira/email"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/smtp"
    "os"
    "strings"
)

// ecv check to see if everything is still running
func Ecv(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "SYSTEM_UP")
}

func GomezTest(w http.ResponseWriter, req *http.Request) {
    body, err  := ioutil.ReadAll(req.Body)
    // FIXME: we should log + move on rather than print to stdout and exit
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }
    fmt.Printf("%s", string(body))
}

// take the post, lookup list members, send email
func GomezAlert(w http.ResponseWriter, req *http.Request) {
    name          := req.URL.Query().Get(":name")
    body, err     := ioutil.ReadAll(req.Body)
    ringring      := "URL"

    // FIXME: remove once I can get POSTs from Gomez to work:
    email_subject := "alert from gomez"

    email_from    := "FROM"
    smtp_server   := "smtp.gmail.com:587"
    smtp_login    := "LOGIN"
    smtp_passwd   := "PASSWORD"
    smtp_host     := "smtp.gmail.com"


    // get the response and error for the "smart contact" url for the list
    // FIXME: we should define a list of folks to get alerted in case this
    // call fails - we can send out the regular alert plus a note about not
    // being able to talk to ringring for the oncall list.
    resp, err  := http.Get(ringring+name)

    // it looks like if you define something (for this example, err) and you do not use it in any
    // way, go will throw an error at compile time about it.
    // FIXME: we should log + move on rather than print to stdout and exit
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {

        // setting defer on this closes out the connection automatically
        // if we didn't set this we would have to do this manually.
        defer resp.Body.Close()
        email_addresses, err := ioutil.ReadAll(resp.Body)

        // FIXME: we need to parse the XML in the body from Gomez

        // FIXME: we should log + move on rather than print to stdout and exit
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }

        // send the email
        m := email.NewMessage(email_subject, string(body))
        m.From = email_from
        m.To   = strings.Split(string(email_addresses), "\n")
        err = email.Send(smtp_server, smtp.PlainAuth("", smtp_login, smtp_passwd, smtp_host), m)

    }
}

func main() {
    m := pat.New()
    m.Get("/ecv", http.HandlerFunc(Ecv))
    m.Post("/gomezalerts/:name", http.HandlerFunc(GomezAlert))
    m.Post("/gomeztest/:name", http.HandlerFunc(GomezTest))

    // Register this pat with the default serve mux so that other packages
    // may also be exported. (i.e. /debug/pprof/*)
    http.Handle("/", m)
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
