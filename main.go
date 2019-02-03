package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/ammario/ipisp"
	"github.com/bhendo/awsipranges"
	"github.com/jedib0t/go-pretty/table"
)

func main() {
	client, err := ipisp.NewDNSClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	ipAddress := flag.String("ip", "", "IP Address to search")

	flag.Parse()

	resp, err := client.LookupIP(net.ParseIP(*ipAddress))
	if err != nil {
		log.Fatalf("Error looking up %s: %v", *ipAddress, err)
	}

	a, err := awsipranges.New(http.DefaultClient)
	if err != nil {
		fmt.Println(err)
	}

	type results struct {
		Network  string
		Region   string
		Services []string
	}

	matches := &results{}

	for _, prefix := range a.Prefixes {
		if prefix.IP_Prefix == resp.Range.String() {
			matches.Network = prefix.IP_Prefix
			matches.Region = prefix.Region
			matches.Services = append(matches.Services, prefix.Service)
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"", ""})

	t.AppendRows([]table.Row{
		{"Name", resp.Name},
		{"Network", resp.ASN},
		{"Country", resp.Country},
		{"Registry", resp.Registry},
		{"Range", resp.Range},
		{"IP", resp.IP},
	})

	if matches.Network != "" {
		// t.AppendRow([]interface{}{"Network", matches.Network})
		t.AppendRow([]interface{}{"Services", strings.Join(matches.Services, ",")})
		t.AppendRow([]interface{}{"Region", matches.Region})
	}

	t.AppendRow([]interface{}{"AllocatedAt", resp.AllocatedAt})
	t.AppendFooter(table.Row{"", ""})
	t.SetStyle(table.StyleColoredBright)
	t.Render()

}
