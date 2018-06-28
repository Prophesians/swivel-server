package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type SwivelServerHandler struct {
}

func (ssh *SwivelServerHandler) IsAlive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")
	}
}

type article struct {
	Header string
	Summary string
	Link string
}

type response struct {
	Topics []article
}

func (ssh *SwivelServerHandler) GetArticles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		article1 := article{
			`'Your Time On Facebook' may soon be monitored`,
			 `Facebook has reportedly introduced a new feature, tentatively titled "Your Time On Facebook", that would provide information about how much time users spent on the platform each day in a week, along with the average time spent on the site per day.`,
			 `https://www.financialexpress.com/industry/technology/your-time-on-facebook-may-soon-be-monitored/1217160/`,
		}

		article2 := article{
			`'Hacker bypasses iOS passcode and it's surprisingly easy`,
			 `Passcodes have pretty much become the standard security measure of choice for most iPhone users. Even in the presence of more advanced biometric solutions, like Face ID, the sheer convenience and approachability of a four, six or even longer digit number, makes it the ideal fallback security measure. The way it works on iOS is simple, yet efficient - you get a total of 10 attempts to enter the code. Fail all of them and the data will get automatically wiped, for security. The number of input attempts is tracked by a hardware module, called the Secure Enclave, making it pretty impossible to actually disable the limit or circumvent it directly. As an extra any brute-force measure, each consecutive pin entry has a slightly longer processing time.`,
			 `https://www.gsmarena.com/hacker_bypasses_ios_passcode_and_its_surprisingly_easy-news-31832.php`,
		}

		article3 := article{
			`Ali Zafar files defamation suit against Meesha Shafi`,
			 `'Shafi had received a legal notice sent by Zafar’s counsel earlier as well, which asked her to issue an apology on Twitter, failing which he would file a PKR 1 billion-defamation case against her.In a social media post, the 36-year-old had alleged that Zafar harassed her on multiple occasions and that “my conscience does not allow me to be silent anymore.”Zafar, a Pakistani singer-songwriter, has worked in several Bollywood films including ‘Tere Bin Laden,’ ‘Mere Brother Ki Dulhan,’ ‘Chashme Baddoor,’ and ‘Dear Zindagi.'(ANI)'`,
			 `https://www.siasat.com/news/ali-zafar-files-defamation-suit-against-meesha-shafi-1372560/`,
		}

		article4 := article{
			`'Your Time On Facebook' may soon be monitored`,
			 `Facebook has reportedly introduced a new feature, tentatively titled "Your Time On Facebook", that would provide information about how much time users spent on the platform each day in a week, along with the average time spent on the site per day.`,
			 `https://www.financialexpress.com/industry/technology/your-time-on-facebook-may-soon-be-monitored/1217160/`,
		}

		article5 := article{
			`'Your Time On Facebook' may soon be monitored`,
			 `Facebook has reportedly introduced a new feature, tentatively titled "Your Time On Facebook", that would provide information about how much time users spent on the platform each day in a week, along with the average time spent on the site per day.`,
			 `https://www.financialexpress.com/industry/technology/your-time-on-facebook-may-soon-be-monitored/1217160/`,
		}

		articles := []article{article1, article2, article3, article4, article5}

		js, err := json.Marshal(response{articles})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

//func (ssh *SwivelServerHandler) SaveTags(context *Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		err := models.SaveTags(r)
//	}
//}
