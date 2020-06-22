package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/avct/uasurfer"
)

func main() {
	var count int
	ua := &uasurfer.UserAgent{}
	stats := stats{
		BrowserIDs: make(map[uasurfer.BrowserID]int),
		OSIDs:      make(map[uasurfer.OSID]int),
		DeviceIDs:  make(map[uasurfer.DeviceID]int),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		count++
		ua.Reset()
		uasurfer.ParseUserAgent(scanner.Text(), ua)
		stats.BrowserIDs[ua.Browser.ID]++
		stats.OSIDs[ua.OS.Name]++
		stats.DeviceIDs[ua.DeviceID]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Printf("Read %d useragents\n", count)
	fmt.Println()
	stats.Summary(count, os.Stdout)
}

type stats struct {
	OSIDs      map[uasurfer.OSID]int
	BrowserIDs map[uasurfer.BrowserID]int
	DeviceIDs  map[uasurfer.DeviceID]int
}

func (s *stats) Summary(total int, dest io.Writer) {
	browserCounts := make([]stringCount, 0, len(s.BrowserIDs))
	for k, v := range s.BrowserIDs {
		browserCounts = append(browserCounts, stringCount{name: k.String(), count: v})
	}
	sort.Slice(browserCounts, func(i, j int) bool { return browserCounts[j].count < browserCounts[i].count }) // by count reversed
	fmt.Fprintf(dest, "Browsers\n")
	err := writeTable(browserCounts, total, dest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "writing summary: %v", err)
		return
	}

	fmt.Fprintln(dest)
	osCounts := make([]stringCount, 0, len(s.OSIDs))
	for k, v := range s.OSIDs {
		osCounts = append(osCounts, stringCount{name: k.String(), count: v})
	}
	sort.Slice(osCounts, func(i, j int) bool { return osCounts[j].count < osCounts[i].count }) // by count reversed
	fmt.Fprintf(dest, "Operating Systems\n")
	err = writeTable(osCounts, total, dest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "writing summary: %v", err)
		return
	}

	fmt.Fprintln(dest)
	deviceCounts := make([]stringCount, 0, len(s.DeviceIDs))
	for k, v := range s.DeviceIDs {
		deviceCounts = append(deviceCounts, stringCount{name: k.String(), count: v})
	}
	sort.Slice(deviceCounts, func(i, j int) bool { return deviceCounts[j].count < deviceCounts[i].count }) // by count reversed
	fmt.Fprintf(dest, "Device Types\n")
	err = writeTable(deviceCounts, total, dest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "writing summary: %v", err)
		return
	}
}

func writeTable(counts []stringCount, total int, dest io.Writer) error {
	tw := tabwriter.NewWriter(dest, 10, 1, 2, ' ', 0)
	for i := range counts {
		fmt.Fprintf(tw, "%s\t%d (%.2f%%)\n", counts[i].name, counts[i].count, percent(counts[i].count, total))
	}
	return tw.Flush()
}

type stringCount struct {
	name  string
	count int
}

func percent(num, den int) float64 {
	return float64(num) / float64(den) * 100.0
}
