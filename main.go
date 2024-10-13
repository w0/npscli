package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/w0/npscli/nps"
)

const APIKEY = ""

func printTable(data nps.Activities) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	for i := 0; i < data.Total; i++ {

		if i%4 == 0 {
			fmt.Fprintln(w)
		}
		fmt.Fprintf(w, "%d: %s\t", i, data.Data[i].Name)
	}

	fmt.Fprintln(w)
	w.Flush()
}

func printParks(parks []nps.Park) {
	parkIndexes := randomIndexs(len(parks))

	fmt.Println("Consider visiting one of these ranomly selected parks!")
	for i, index := range parkIndexes {
		fmt.Printf("%d: %s, located in the state(s) of %s: \n\t\t%s\n\n",
			i+1,
			parks[index].FullName,
			parks[index].States,
			parks[index].Url)
	}

}

func randomIndexs(max int) []int {
	var indexes []int

	indexes = append(indexes, rand.IntN(max))
	indexes = append(indexes, rand.IntN(max))
	indexes = append(indexes, rand.IntN(max))

	return indexes
}

func main() {
	categories, err := nps.GetCategories(APIKEY)

	if err != nil {
		fmt.Println(err)
		return
	}

	printTable(categories)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Select a category between 0 - %d: ", categories.Total-1)
	text, _ := reader.ReadString('\n')

	selection, _ := strconv.Atoi(strings.Split(text, "\n")[0])

	if selection < 0 || selection >= categories.Total {
		fmt.Println("Bad. No. We dont access indexes outside of the available range!")
		return
	}

	fmt.Println()
	fmt.Println("Looking up parks that support the activity:", categories.Data[selection].Name)
	parks, err := nps.GetParkByCategory(categories.Data[selection].Id, APIKEY)
	if err != nil {
		fmt.Println(err)
		return
	}
	printParks(parks.Data[0].Parks)
}
