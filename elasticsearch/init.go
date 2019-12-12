package elasticsearch

// import (
// 	"fmt"
// 	"os"

// 	"gopkg.in/olivere/elastic.v2"
// )

// type Tweet struct {
// 	User    string
// 	Message string
// }

// func EsInit() {
// 	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(os.Getenv("ES_CONF")))
// 	if err != nil {
// 		fmt.Println("connect es error", err)
// 		return
// 	}

// 	fmt.Println("conn es succ")

// 	for i := 0; i < 20; i++ {
// 		tweet := Tweet{User: "olivere", Message: "Take Five"}
// 		_, err = client.Index().
// 			Index("twitter").
// 			Type("tweet").
// 			Id(fmt.Sprintf("%d", i)).
// 			BodyJson(tweet).
// 			Do()
// 		if err != nil {
// 			// Handle error
// 			panic(err)
// 			return
// 		}
// 	}

// 	fmt.Println("insert succ")
// }
