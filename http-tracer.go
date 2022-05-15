package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tcnksm/go-httpstat"
)

func main() {

	var urls, queris, cluster, application string
	var numberOfRequests, sleepTime int
	flag.StringVar(&urls, "u", "", "Comma separated list of URLs to test")
	flag.IntVar(&numberOfRequests, "n", 1, "Number of requests made")
	flag.StringVar(&queris, "q", "", "Comma separated list of queries to run")
	flag.IntVar(&sleepTime, "s", 0, "Cool off time in ms between consecutive call to same server")
	flag.StringVar(&cluster, "c", "not-set", "Details from where job is run")
	flag.StringVar(&application, "a", "not-set", "Application which is tested")
	flag.Parse()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8081", nil)
	}()

	urlsArray := strings.Split(urls, ",")
	queriesArray := strings.Split(queris, ",")
	responseStats := make(map[string]map[int]int64)
	httpStats := make(map[string]map[string]map[string]int64)

	for _, url := range urlsArray {
		responseStats[url] = make(map[int]int64)
		httpStats[url] = make(map[string]map[string]int64)
		httpStats[url]["max"] = make(map[string]int64)
		httpStats[url]["min"] = make(map[string]int64)
		httpStats[url]["cumulative"] = make(map[string]int64)
	}

	for i := 0; i < numberOfRequests; i++ {
		for _, url := range urlsArray {
			query := queriesArray[rand.Intn(len(queriesArray))]
			req, err := http.NewRequest("GET", url+query, nil)
			if err != nil {
				log.Fatal(err)
			}

			var result httpstat.Result
			ctx := httpstat.WithHTTPStat(req.Context(), &result)
			req = req.WithContext(ctx)

			client := http.DefaultClient
			res, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}

			if val, ok := responseStats[url][res.StatusCode]; ok {
				responseStats[url][res.StatusCode] = val + 1
			} else {
				responseStats[url][res.StatusCode] = 1
			}

			if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
				log.Fatal(err)
			}
			res.Body.Close()
			result.End(time.Now())

			// DNS LookUp
			DNSTime := result.DNSLookup.Milliseconds()
			if val, ok := httpStats[url]["max"][DNSLookup]; ok {
				if val < DNSTime {
					httpStats[url]["max"][DNSLookup] = DNSTime
				}
				if val > DNSTime {
					httpStats[url]["min"][DNSLookup] = DNSTime
				}
				httpStats[url]["cumulative"][DNSLookup] = httpStats[url]["cumulative"][DNSLookup] + DNSTime
			} else {
				httpStats[url]["max"][DNSLookup] = DNSTime
				httpStats[url]["min"][DNSLookup] = DNSTime
				httpStats[url]["cumulative"][DNSLookup] = DNSTime
			}

			//TCP connection
			TCPTime := result.TCPConnection.Milliseconds()
			if val, ok := httpStats[url]["max"][TCPConnection]; ok {
				if val < TCPTime {
					httpStats[url]["max"][TCPConnection] = TCPTime
				}
				if val > TCPTime {
					httpStats[url]["min"][TCPConnection] = TCPTime
				}
				httpStats[url]["cumulative"][TCPConnection] = httpStats[url]["cumulative"][TCPConnection] + TCPTime
			} else {
				httpStats[url]["max"][TCPConnection] = TCPTime
				httpStats[url]["min"][TCPConnection] = TCPTime
				httpStats[url]["cumulative"][TCPConnection] = TCPTime
			}

			//TLS HandShake
			TLSTime := result.TLSHandshake.Milliseconds()
			if val, ok := httpStats[url]["max"][TLSHandshake]; ok {
				if val < TLSTime {
					httpStats[url]["max"][TLSHandshake] = TLSTime
				}
				if val > TLSTime {
					httpStats[url]["min"][TLSHandshake] = TLSTime
				}
				httpStats[url]["cumulative"][TLSHandshake] = httpStats[url]["cumulative"][TLSHandshake] + TLSTime
			} else {
				httpStats[url]["max"][TLSHandshake] = TLSTime
				httpStats[url]["min"][TLSHandshake] = TLSTime
				httpStats[url]["cumulative"][TLSHandshake] = TLSTime
			}

			//Server Processing
			serverProcessingTime := result.ServerProcessing.Milliseconds()
			if val, ok := httpStats[url]["max"][serverProcessing]; ok {
				if val < serverProcessingTime {
					httpStats[url]["max"][serverProcessing] = serverProcessingTime
				}
				if val > serverProcessingTime {
					httpStats[url]["min"][serverProcessing] = serverProcessingTime
				}
				httpStats[url]["cumulative"][serverProcessing] = httpStats[url]["cumulative"][serverProcessing] + serverProcessingTime
			} else {
				httpStats[url]["max"][serverProcessing] = serverProcessingTime
				httpStats[url]["min"][serverProcessing] = serverProcessingTime
				httpStats[url]["cumulative"][serverProcessing] = serverProcessingTime
			}

			//Content Transfer
			contentTransferTime := result.ContentTransfer(time.Now()).Milliseconds()
			if val, ok := httpStats[url]["max"][contentTransfer]; ok {
				if val < contentTransferTime {
					httpStats[url]["max"][contentTransfer] = contentTransferTime
				}
				if val > contentTransferTime {
					httpStats[url]["min"][contentTransfer] = contentTransferTime
				}
				httpStats[url]["cumulative"][contentTransfer] = httpStats[url]["cumulative"][contentTransfer] + contentTransferTime
			} else {
				httpStats[url]["max"][contentTransfer] = contentTransferTime
				httpStats[url]["min"][contentTransfer] = contentTransferTime
				httpStats[url]["cumulative"][contentTransfer] = contentTransferTime
			}

			//Name Lookup
			nameLookupTime := result.NameLookup.Milliseconds()
			if val, ok := httpStats[url]["max"][nameLookup]; ok {
				if val < nameLookupTime {
					httpStats[url]["max"][nameLookup] = nameLookupTime
				}
				if val > nameLookupTime {
					httpStats[url]["min"][nameLookup] = nameLookupTime
				}
				httpStats[url]["cumulative"][nameLookup] = httpStats[url]["cumulative"][nameLookup] + nameLookupTime
			} else {
				httpStats[url]["max"][nameLookup] = nameLookupTime
				httpStats[url]["min"][nameLookup] = nameLookupTime
				httpStats[url]["cumulative"][nameLookup] = nameLookupTime
			}

			//Connect
			connectTime := result.Connect.Milliseconds()
			if val, ok := httpStats[url]["max"][connect]; ok {
				if val < connectTime {
					httpStats[url]["max"][connect] = connectTime
				}
				if val > connectTime {
					httpStats[url]["min"][connect] = connectTime
				}
				httpStats[url]["cumulative"][connect] = httpStats[url]["cumulative"][connect] + connectTime
			} else {
				httpStats[url]["max"][connect] = connectTime
				httpStats[url]["min"][connect] = connectTime
				httpStats[url]["cumulative"][connect] = connectTime
			}

			//Pre Transfer
			preTransferTime := result.Pretransfer.Milliseconds()
			if val, ok := httpStats[url]["max"][preTransfer]; ok {
				if val < preTransferTime {
					httpStats[url]["max"][preTransfer] = preTransferTime
				}
				if val > preTransferTime {
					httpStats[url]["min"][preTransfer] = preTransferTime
				}
				httpStats[url]["cumulative"][preTransfer] = httpStats[url]["cumulative"][preTransfer] + preTransferTime
			} else {
				httpStats[url]["max"][preTransfer] = preTransferTime
				httpStats[url]["min"][preTransfer] = preTransferTime
				httpStats[url]["cumulative"][preTransfer] = preTransferTime
			}

			//Start Transfer
			startTransferTime := result.StartTransfer.Milliseconds()
			if val, ok := httpStats[url]["max"][startTransfer]; ok {
				if val < startTransferTime {
					httpStats[url]["max"][startTransfer] = startTransferTime
				}
				if val > startTransferTime {
					httpStats[url]["min"][startTransfer] = startTransferTime
				}
				httpStats[url]["cumulative"][startTransfer] = httpStats[url]["cumulative"][startTransfer] + startTransferTime
			} else {
				httpStats[url]["max"][startTransfer] = startTransferTime
				httpStats[url]["min"][startTransfer] = startTransferTime
				httpStats[url]["cumulative"][startTransfer] = startTransferTime
			}

			//Total
			totalTime := result.Total(time.Now()).Milliseconds()
			if val, ok := httpStats[url]["max"][total]; ok {
				if val < totalTime {
					httpStats[url]["max"][total] = totalTime
				}
				if val > totalTime {
					httpStats[url]["min"][total] = totalTime
				}
				httpStats[url]["cumulative"][total] = httpStats[url]["cumulative"][total] + totalTime
			} else {
				httpStats[url]["max"][total] = totalTime
				httpStats[url]["min"][total] = totalTime
				httpStats[url]["cumulative"][total] = totalTime
			}

			setPromMetrics(cluster, url, application, httpStats, responseStats)

		}
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
	}
	for _, url := range urlsArray {
		fmt.Println("\n-----")
		fmt.Println(url)
		fmt.Println("-----")
		fmt.Println("HTTP Response Code Stats:")
		fmt.Println(responseStats[url])
		fmt.Println("\nHTTP Latency Stats:")
		fmt.Println(httpStats[url])
		fmt.Println("")
	}
	time.Sleep(2 * time.Minute)
}
