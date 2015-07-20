package main

import(
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "html/template"
)

type Page struct {
    Title string
    Body []byte
}

func (p *Page) save() error {
    filename := "files/" + p.Title + ".txt"

    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := "files/" + title + ".txt"
    body, err := ioutil.ReadFile(filename)

    if err != nil {
        return nil, err
    }

    return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles("views/" + tmpl + ".html")
    t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)

    renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)

    if err != nil {
        p = &Page{Title: title}
    }

    renderTemplate(w, "edit", p)
}

// func saveHandler(w http.ResponseWriter, *r http.Request) {
//
// }

func main() {
    dirname := "files"
    d, _ := os.Open(dirname)

    files, err := d.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    for _, file := range files {
        if file.Mode().IsRegular() {
            fmt.Println(file.Name(), file.Size(), "bytes")
        }
    }

    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    // http.HandleFunc("/save/", saveHandler)
    http.ListenAndServe(":8080", nil)
}

// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hi there, welcome to the GoLang to-do app")
// }
//
// func main() {
//     http.HandleFunc("/", handler)
//     http.ListenAndServe(":8080", nil)
// }
