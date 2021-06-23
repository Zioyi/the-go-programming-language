package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track1 struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks1 = []*Track1{
	{"Go", "Delilah", "From the Roots Up", 2012, length1("3m38s")},
	{"Go", "Moby", "Moby", 1992, length1("3m37s")},
	{"Go", "Alicia Keys", "As I Am", 2007, length1("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length1("4m24s")},
}

func length1(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type myStableSort struct {
	clickedColumn []string
	s             []*Track1
}

func (x myStableSort) Len() int { return len(x.s) }
func (x myStableSort) Less(i, j int) bool {
	for _, column := range x.clickedColumn {
		switch column {
		case "title", "Title":
			if x.s[i].Title != x.s[j].Title {
				return x.s[i].Title < x.s[j].Title
			}
		case "artist", "Artist":
			if x.s[i].Artist != x.s[j].Artist {
				return x.s[i].Artist < x.s[j].Artist
			}
		default:
			panic("invalid column")
		}
	}
	return false
}
func (x myStableSort) Swap(i, j int) { x.s[i], x.s[j] = x.s[j], x.s[i] }

func printTracks1(tracks []*Track1) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func main() {
	m := myStableSort{[]string{"title"}, tracks1}
	sort.Sort(m)
	printTracks1(tracks1)

	m = myStableSort{[]string{"title", "artist"}, tracks1}
	sort.Sort(m)
	printTracks1(tracks1)
}
