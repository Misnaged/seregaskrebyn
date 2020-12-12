   package main


   import (
       "fmt"
       "os"
       "log"
      "github.com/badoux/goscraper"



   )
      var (
            outfile, _ = os.Create("cheeba.txt")  /// это я так логи записываю в файл отдельный, по - крестьянски 
  l      = log.New(outfile, "", 0)
          )
   const  myUA string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36" // const ваще по приколу написал. можно и var. Это типа User Agent


   func main() {
urls := []string{
"https://контактыч.com/id1",
"https://контактыч.com/id1",
"https://контактыч.com/id1",
"https://контактыч.com/id1",
"https://контактыч.com/id1",
"https://контактыч.com/id1",
"https://контактыч.com/id1",
"https://контактыч.com/id1",
"https://контактыч.com/id1",
"https://контактыч.com/id1",
"https://контактыч.com/id1",
"https://контактыч.com/id1",
"https://контактыч.com/id1",
"https://контактыч.com/id1",}// адресочки - адрески!}
 

   ch := make(chan string)  

       for _, url := range urls {
           go fetch(url, ch) // start a goroutine собственно
       }
       for range urls {
           l.Println(<-ch) // receive from channel ch 
       }
   }

   func fetch(url string, ch chan<- string) {
s, err := goscraper.Scrape(url, 5, 5000, myUA, "http://167.172.109.12:46329")
      if err != nil {
        fmt.Println(err)  /// В любой непонятной ситуации делай проверку на наличие ошибок!
        return  
  }   /// даже если их нет это ничего не значит ибо = см. выше ^^^^^^

      ch <-  s.Preview.Title  /// вот я ебал мозг себе с этим говном просто жесть сколько времени. 
///~~~~~~___  fmt.Printf("Title : %s\n", s.Preview.Title) ___~~~~~//// вот так было и       постоянно прилетало "multiple-value fmt.Printf() in single-value context"  я чуть не свихнулся

   }