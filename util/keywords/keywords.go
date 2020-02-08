package keywords

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Keywords struct {
	SearchKeyword  string
	ProductKeyword string
}

func getKeywords(path string) []Keywords {
	csvFile, err := os.Open(path)

	defer csvFile.Close()

	if err != nil {
		log.Fatalln("Failed to open csv file", err)
	}

	r := csv.NewReader(csvFile)
	var extractedKeywords []Keywords

	for {
		keywordsFromCsv, err := r.ReadAll()
		if keywordsFromCsv == nil {
			break
		}
		if err != nil {
			log.Fatalln("Failed to read csv file", err)
		}
		for i, keyword := range keywordsFromCsv {
			if i == 0 {
				continue
			}
			keywords := Keywords{
				SearchKeyword:  keyword[1],
				ProductKeyword: keyword[3],
			}
			extractedKeywords = append(extractedKeywords, keywords)
		}
	}
	return extractedKeywords
}

func encodeKeywords(kws []Keywords) []string {
	var encodedKeywords []string

	for _, kw := range kws {
		sk := kw.SearchKeyword
		src, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(sk)),
			simplifiedchinese.GBK.NewEncoder()))

		encodedKeywords = append(encodedKeywords, url.QueryEscape(string(src)))
	}
	return encodedKeywords
}

func GetMainUrls() {
	var urls []string
	csvPath := "/Users/seungweonbaek/Projects/business/ingest-categories/data/alibaba13.csv"
	kws := getKeywords(csvPath)

	eks := encodeKeywords(kws)

	for _, ek := range eks {
		urls = append(urls, "https://s.1688.com/selloffer/offer_search.htm?quantityBegin=1&uniqfield=pic_tag_id&keywords="+ek+"&filt=y&netType=1%2C11&n=y&filt=y")
	}

	fmt.Println(urls)
}
