//package main
//
//import (
//	"io"
//	"bufio"
//	"os"
//	"fmt"
//	"time"
//	"net/http"
//	"strings"
//)

//import (
//	"./kadai"
//	//"fmt"
//	"flag"
//	"log"
//	"fmt"
//)
//var n bool
//func init() {
//	// ポインタを指定して設定を予約
//	flag.BoolVar(&n, "n", false, "改行フラグ") }
//func main() {
//	fmt.Println("***********1-1**************")
//	kadai.Kadai_1_1()
//	fmt.Println("***********1-2**************")
//	kadai.Kadai_1_2()
//	fmt.Println("***********1-3**************")
//	kadai.Kadai_1_3()
//	fmt.Println("***********1-4**************")
//	kadai.Kadai_1_4()
//	fmt.Println("***********1-5**************")
//	kadai.Kadai_1_5()
//	fmt.Println("***********1-6**************")
//	kadai.Kadai_1_6(n)
//}

//func main() {
//	flag.Parse()
//	kadai.Syukudai_1()
//}


//func input(r io.Reader) <-chan string {
//	// TODO: チャネルを作る
//	done := make(chan string)
//	go func() {
//		s := bufio.NewScanner(r)
//		for s.Scan() {
//			// TODO: チャネルに読み込んだ文字列を送る
//			done <- s.Text()
//		}
//		// TODO: チャネルを閉じる
//		close(done)
//	}()
//	// TODO: チャネルを返す
//	return done
//}
//
//func main() {
//	fmt.Println("*************************************")
//	fmt.Println("*          GO言語単語テスト           *")
//	fmt.Println("* 　 　 　全三問/制限時間(15秒)        *")
//	fmt.Println("*          e→end/s→start            *")
//	fmt.Println("*                                   *")
//	fmt.Println("*************************************")
//	ch := input(os.Stdin)
//
//	var end_flag bool = false
//
//loop_fst:
//	for {
//		fmt.Print(">")
//		select {
//		case s := <-ch:
//
//			if(s == "e"){
//				end_flag = true
//			}
//
//			break loop_fst
//		}
//	}
//	var cnt int = 1
//	var time_cnt int = 0
//	var my_check = []bool{false, false, false, }
//	var my_ans = []bool{false, false, false, }
//	var question = []string{"dog", "cat", "mercari", }
//	if(end_flag == false) {
//		fmt.Println("Start!")
//	loop:
//		for {
//			if(my_check[cnt -1] == false) {
//				fmt.Println(fmt.Sprintf("第%d問(残り%d秒)", cnt, (15 - time_cnt)))
//				fmt.Println(question[cnt-1])
//				fmt.Print(">")
//				my_check[cnt -1] = true
//			}
//			select {
//			case s := <-ch:
//				if(s == question[cnt-1]){
//					fmt.Println("正解！")
//					my_ans[cnt-1] = true
//				}else {
//					fmt.Println("残念！")
//				}
//				cnt++
//				if cnt > 3{
//					if my_ans[0] && my_ans[1] && my_ans[2] {
//
//						MerucariCreate_main()
//						fmt.Println("complete!!")
//					}else{
//						fmt.Println("not complete!!")
//					}
//					break loop
//				}else{
//					break
//				}
//			case <-time.After(1 * time.Second):
//				time_cnt++
//				if(time_cnt >= 15) {
//					fmt.Println("timeout")
//					end_flag = true
//					break loop
//				}else{
//
//					break
//				}
//			}
//		}
//	}
//
//	if(end_flag == true){
//		fmt.Println("Good Bye.")
//	}
//}
//
//func MerucariCreate_main(){
//
//	fmt.Println("                                        .,,,.                                                       ")
//	fmt.Println("                                     `,,,,,,,,,                                                     ")
//	fmt.Println("                                    ,,,,,,,,,,,,.                                                   ")
//	fmt.Println("                                   ,,,,,,,,,,,,,,.                                                  ")
//	fmt.Println("                                  `,,,,,,,,,,,,,,,                                                  ")
//	fmt.Println("                                  ,,,,,,,,,,,,,,,,,                                                 ")
//	fmt.Println("                                  ,,,,,,,,,,,,,,,,,                                                 ")
//	fmt.Println("                                 .,,,,,,,,,,,,,,,,,`       ';:                                      ")
//	fmt.Println("                                 ,,,,,,,,,,,,,,,,,,.      .;;;;:                                    ")
//	fmt.Println("                                 ,,,,,,,,,,,,,,,,,,.      ;;;;;;;:                                  ")
//	fmt.Println("                                .,,,,,,,,,,,,,,,,,,`     ;;;;;;;;;;;                                ")
//	fmt.Println("                              .;;;,,,,,,,,,,,,,,,,,     ;;;;;;;;;;;;;;                              ")
//	fmt.Println("                            ,;;;;;,,,,,,,,,,,,,,,,,    `;;;;;;;;;;;;;;;;                            ")
//	fmt.Println("                          ,;;;;;;;;,,,,,,,,,,,,,,,     ;;;;;;;;;;;;;;;;;;;                          ")
//	fmt.Println("                        ,;;;;;;;;;;:,,,,,,,,,,,,,     ;;;;;;;;;;;;;;;;;;;;;;                        ")
//	fmt.Println("                      ,;;;;;;;;;;;;;:,,,,,,,,,,,     ,;;;;;;;;;;;;;;;;;;;;;;;;                      ")
//	fmt.Println("                    :;;;;;;;;;;;;;;;;;,,,,,,,,,,     ;;;,,,,,;;;;;;;;;;;;;;;;;;;                    ")
//	fmt.Println("                  :;;;;;;;;;;;;;;;;;;;;;;:,:;;;;    ;;,,,,,,,,,;;;;;;;;;;;;;;;;;;;                  ")
//	fmt.Println("                 ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;  ';,,,,,,,,,,,;;;;;;;;;;;;;;;;;;;                 ")
//	fmt.Println("                  ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;.;,,,,,,,,,,,,:;;;;;;;;;;;;;;;;;                  ")
//	fmt.Println("                   ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;';,,,,,,,,,,,,,;;;;;;;;;;;;;;;;,                  ")
//	fmt.Println("                   ,;;;;;;;;;;;;;;;;;;;;;;;;;;;;'''+,,,,,,,,,,,,,;;;;;;;;;;;;;;;;                   ")
//	fmt.Println("                    ;;;;;;;;;;;;;;;;;;;;;;;;;''''''',,,,,,,,,,,,,;;;;;;;;;;;;;;;                    ")
//	fmt.Println("                     ;;;;;;;;;;;;;;;;;;;;;;''''''''+,,,,,,,,,,,,,;;;;;;;;;;;;;;`                    ")
//	fmt.Println("                     `;;;;;;;;;;;;;;;;;;;''''''''''',,,,,,,,,,,,,;;;;;;;;;;;;;:                     ")
//	fmt.Println("                      :;;;;;;;;;;;;;;;;''''''''''''',,,,,,,,,,,,,;;;;;;;;;;;;;                      ")
//	fmt.Println("                       ;;;;;;;;;;;;;;'''''''''''''''',,,,,,,,,,,;;;;;;;;;;;;;                       ")
//	fmt.Println("                        ;;;;;;;;;;;''''''''''''''''''',,,,,,,,,'';;;;;;;;;;;.                       ")
//	fmt.Println("                        .;;;;;;;;''''''''''',,,,,;''''':,,,,,;''''';;;;;;;;;                        ")
//	fmt.Println("                         ';;;;;'''''''''''':,,,,,,+'''''''''''''''''';;;;;;                         ")
//	fmt.Println("                          ;;;'''''''''''''',,,,,,,,'''''''''''''''''''';;;                          ")
//	fmt.Println("                           '''''''''''''''',,,,,,,,'''''''''''''''''''''',                          ")
//	fmt.Println("                           ;;'''''''''''''',,,,,,,,''''''''''''''''''''';,                          ")
//	fmt.Println("                           ;;;;'''''''''''',,,,,,,'''''''''''''''''''';;;,                          ")
//	fmt.Println("                           ;;;;;;''''''''''',,,,,,'''''''''''''''''';;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;''''''''''',,+'''''''''''''''';;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;;'''''''''''''''''''''''''';;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;;;;'''''''''''''''''''''';;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;;;;;;'''''''''''''''''';;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;;;;;;;;'''''''''''''';;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;;;;;;;;;;'''''''''';;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;;;;;;;;;;;;'''''';;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;;;;;;;;;;;;;;'';;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;:.;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;    ,;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;:      ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;:  ;;   :;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;:  ;;;     .;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;:  ;;;:      ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;:  ;;;:  ;;`  ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;:  ;;;:  ;;;   ;;;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;:  ;;;:  ;;;.  ;;;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;: ;;;:  ;;;.  ';;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;:  ;;;.  ';;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;:  ;;;.  ';;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;;; ;;;.  ';;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;;;;;;;.  ';;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;;;;;;;.  ';;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                           ;;;;;;;;;;;;;;;;; ';;;;;;;;;;;;;;;;;;;;;;;;;;;,                          ")
//	fmt.Println("                            ';;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;.                           ")
//	fmt.Println("                              ';;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;.                             ")
//	fmt.Println("                                ';;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;.                               ")
//	fmt.Println("                                  ';;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;`                                 ")
//	fmt.Println("                                    ';;;;;;;;;;;;;;;;;;;;;;;;;;;`                                   ")
//	fmt.Println("                                      ;;;;;;;;;;;;;;;;;;;;;;;;`                                     ")
//	fmt.Println("                                        ;;;;;;;;;;;;;;;;;;;;`                                       ")
//	fmt.Println("                                          ;;;;;;;;;;;;;;;;`                                         ")
//	fmt.Println("                                            ;;;;;;;;;;;;                                            ")
//	fmt.Println("                                              ;;;;;;;;                                              ")
//	fmt.Println("                                                ;;;'                                                ")
//}
package main

import (
"encoding/json"
"log"
"net/http"
"strings"
	"fmt"
	"bytes"
)

func kaibun(s string) bool {
	if s == "" {
		return false
	}
    var i=0
	var j=0
	for /* TODO: 前と後ろから比較していく */ {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func handler(w http.ResponseWriter, r *http.Request) {

	// TODO: リクエストを取得する
	fmt.Fprintln(w, "hello", r.FormValue("msg"))
	var q = r.FormValue("msg")
	ss := strings.Split(q, ",")

	results := make([]bool, len(ss))
	for i := range ss {
		results[i] = kaibun(ss[i])
	}

	// TODO: JSONでレスポンスを返す
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(ss); err != nil {
		log.Fatal(err)
	}

}

func main() {
	// TODO: ハンドラを登録する
	http.HandleFunc("/", handler)
	// TODO: HTTPサーバを8080ポートで起動する
	http.ListenAndServe(":8080", nil)
}
