package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/drankou/go-vader/vader"
)

// Get preferred outbound ip of this machine
func getOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

func helloWorld(port string) {
	fmt.Println("--------TwitterTrendR: Backend server--------")
	fmt.Println()
	fmt.Println("  Launched on host: ", getOutboundIP())
	fmt.Println("  Launched on port: ", port)
	fmt.Println()
	fmt.Println("---------------------------------------------")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT not set switching to default port")
		port = "5000"
	}

	helloWorld(port)
	handleRequests(port)
}

func trendAnalysis(client *twitter.Client, trend twitter.Trend, wg *sync.WaitGroup) {
	results := search(client, trend.Name, 300)
	sentiment := sumScores(results)
	fmt.Printf("%28s\t-->\t Neg: %0.2f, Neu: %0.2f, Pos: %0.2f\n", trend.Name, sentiment["neg"], sentiment["neu"], sentiment["pos"])
	wg.Done()
}

func sumScores(tweets []twitter.Tweet) map[string]float64 {
	// vader
	sia := vader.SentimentIntensityAnalyzer{}
	err := sia.Init("data/vader_lexicon.txt", "data/emoji_utf8_lexicon.txt")
	if err != nil {
		log.Fatal(err)
	}

	results := map[string]float64{
		"compound": 0.0,
		"neg": 0.0,
		"neu": 0.0,
		"pos": 0.0,
	}
	n := float64(len(tweets))
	for _, tweet := range tweets {
		score := sia.PolarityScores(tweet.Text)
		results["compound"] += score["compound"] / n
		results["neg"] += score["neg"] / n
		results["neu"] += score["neu"] / n
		results["pos"] += score["pos"] / n
	}

	return results
}

func search(client *twitter.Client, query string, limit int) []twitter.Tweet {
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: query,
		Count: limit,
	})
	if err != nil {
		panic(err)
	}

	return search.Statuses
}

func getLocations(client *twitter.Client) {
	locs, _, _ := client.Trends.Available()
	for _, loc := range locs {
		fmt.Println(loc)
	}
}

func trendList(client *twitter.Client) []twitter.Trend {
	// rennes : 619163
	// global : 1 
	// paris : 615702
	// france : 23424819
	// us : 23424977
	trends, _, err := client.Trends.Place(23424977, nil)
	if err != nil {
		panic(err)
	}

	return trends[0].Trends
}

func streamList(client *twitter.Client) {
	params := &twitter.StreamSampleParams{
		Language: []string{"fr", "en"},
	}
	_, err := client.Streams.Sample(params)
	if err != nil {
		panic(err)
	}
}