package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func certExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return false
	}
	return !info.IsDir()
}
func readFile(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	bytes := []byte(string(content))
	return bytes
}
func decodeCert(bytes []byte) float64 {
	block, _ := pem.Decode(bytes)
	var cert *x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)
	expiryDate := cert.NotAfter
	expiryDateEpoch := float64(expiryDate.Unix())
	return expiryDateEpoch
}

func recordMetrics(expiryDateEpoch float64) {
	go func() {
		for {
			certExpiryDateMetric.Observe(expiryDateEpoch)
		}
	}()
}

var (
	certExpiryDateMetric = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "cert_expiry_date_metric",
		Help: "The expiry date of the certificate",
	})
)

func main() {
	filename := "/u01/app/oracle/product/12.1.0/dbhome_12102/network/admin/dbwallet/andrei.crt"
	promFile, err := os.Create("promFile.txt")
	if err != nil {
		fmt.Println("Could not create file promFile.txt")

		certExists(filename)
		bytes := readFile(filename)
		expiryDateEpoch := decodeCert(bytes)
		expiryDateString := strconv.FormatInt(int64(expiryDateEpoch), 10)
		certExpiryDate := "oracle_certificate_expiration_ca" + ": " + expiryDateString + "\n"
		// promFile.WriteString(certExpiryDate)
		recordMetrics(expiryDateEpoch)
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}
}
