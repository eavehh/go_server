package main

import (
    "fmt"
    "os"
    "net/http"
//    "time"
)

 //обработчик HTTP-запросов
 /*
    w http.ResponseWriter — интерфейс для записи ответа на HTTP-запрос.
    r *http.Request — указатель на структуру запроса, содержащую информацию о HTTP-запросе, который был сделан.
*/

/*

func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Serving %s\n", r.URL.Path)
    fmt.Printf("Served: %s\n", r.Host)
}

/*
    Это еще одна функция-обработчик, называемая timeHandler, которая возвра-
    щает текущее время в формате HTML. Все вызовы fmt.Fprintf() отправляют
    данные обратно HTTP-клиенту, тогда как выходные данные fmt.Printf() вы-
    водятся на терминале, на котором работает веб-сервер.
*/

/*
func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)

	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}
*/

func myFunc(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Im a server")       //отправляет текстовый ответ клиенту который делает запрос 
    fmt.Println("Обслужено: ", r.Host)  //показывает значение хоста с которого был сделан запрос 
}



func main() {
    PORT := ":8001"
    arguments := os.Args

    if len(arguments) != 1{
        PORT = arguments[1]
    }
    fmt.Println("Using port number", PORT)

//    http.HandleFunc("/time", timeHandler)
//    http.HandleFunc("/", myHandler)
    http.HandleFunc("/my", myFunc)

    err := http.ListenAndServe(PORT, nil)
    if err != nil{
        fmt.Println(err)
        return
    }
}
