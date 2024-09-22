package main

import (
	"bytes"
	"fmt"
	"github.com/phin1x/go-ipp"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	// create a new ipp client
	client := ipp.NewIPPClient("192.168.1.50/ipp", 631, "", "", true)
	printThings(client)
	// tryJob()
	tryOtherJob()
}

func printThings(client *ipp.IPPClient) {
	// print file
	res, err := client.PrintFile("/Users/colearendt/Documents/blank.pdf", "DELLA51348", map[string]interface{}{})
	if err != nil {
		fmt.Errorf("Error printing: %v\n", err)
	}
	fmt.Printf("Response: %s\n", strconv.Itoa(res))

}

func tryJob() {
	// define a ipp request
	req := ipp.NewRequest(ipp.OperationGetJobs, 1)
	req.OperationAttributes[ipp.AttributeWhichJobs] = "completed"
	req.OperationAttributes[ipp.AttributeMyJobs] = ""
	req.OperationAttributes[ipp.AttributeFirstJobID] = 0
	req.OperationAttributes[ipp.AttributeRequestingUserName] = "test"

	// encode request to bytes
	payload, err := req.Encode()
	if err != nil {
		panic(err)
	}

	// send ipp request to remote server via http
	httpReq, err := http.NewRequest("POST", "http://192.168.1.50:631/ipp", bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	// set ipp headers
	httpReq.Header.Set("Content-Length", strconv.Itoa(len(payload)))
	httpReq.Header.Set("Content-Type", ipp.ContentTypeIPP)

	httpClient := http.DefaultClient
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		panic(err)
	}
	defer httpResp.Body.Close()

	// response must be 200 for a successful operation
	// other possible http codes are:
	// - 500 -> server error
	// - 426 -> sever requests a encrypted connection
	// - 401 -> forbidden -> need authorization header or user is not permitted
	if httpResp.StatusCode != 200 {
		fmt.Printf("Response: %v\n", httpResp)
		body, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			fmt.Printf("Error reading body: %v\n", err)
		}
		fmt.Printf("Body: %v\n", string(body))
		panic("non 200 response from server")
	}

	// decode ipp response
	resp, err := ipp.NewResponseDecoder(httpResp.Body).Decode(nil)
	if err != nil {
		panic(err)
	}

	// check if the response status is "ok"
	if resp.StatusCode == ipp.StatusOk {
		panic(resp.StatusCode)
	}

	// do something with the returned data
	for _, job := range resp.JobAttributes {
		fmt.Printf("Job details: %v\n", job)
		// ...
	}
}

func tryOtherJob() {
	// define a ipp request
	// req := ipp.NewRequest(ipp.OperationGetOutputDeviceAttributes, 1)
	req := ipp.NewRequest(ipp.OperationGetPrinterAttributes, 1)
	// req := ipp.NewRequest(ipp.OperationCreatePrinter, 1)
	// req.OperationAttributes[ipp.AttributeWhichJobs] = ""
	// req.OperationAttributes[ipp.AttributeMyJobs] = ""
	// req.OperationAttributes[ipp.AttributeFirstJobID] = 1
	req.OperationAttributes[ipp.AttributePrinterURI] = "http://192.168.1.50:631/ipp"
	// req.OperationAttributes[ipp.AttributeRequestingUserName] = "test"

	// encode request to bytes
	payload, err := req.Encode()
	if err != nil {
		panic(err)
	}

	// send ipp request to remote server via http
	httpReq, err := http.NewRequest("POST", "http://192.168.1.50:631/ipp", bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	// set ipp headers
	httpReq.Header.Set("Content-Length", strconv.Itoa(len(payload)))
	httpReq.Header.Set("Content-Type", ipp.ContentTypeIPP)

	httpClient := http.DefaultClient
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		panic(err)
	}
	defer httpResp.Body.Close()

	// response must be 200 for a successful operation
	// other possible http codes are:
	// - 500 -> server error
	// - 426 -> sever requests a encrypted connection
	// - 401 -> forbidden -> need authorization header or user is not permitted
	if httpResp.StatusCode != 200 {
		fmt.Printf("Response: %v\n", httpResp)
		body, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			fmt.Printf("Error reading body: %v\n", err)
		}
		fmt.Printf("Body: %v\n", string(body))
		panic("non 200 response from server")
	}

	// decode ipp response
	resp, err := ipp.NewResponseDecoder(httpResp.Body).Decode(nil)
	if err != nil {
		panic(err)
	}

	// check if the response status is "ok"
	if resp.StatusCode == ipp.StatusOk {
		panic(resp.StatusCode)
	}

	// do something with the returned data
	for _, job := range resp.JobAttributes {
		fmt.Printf("Job details: %v\n", job)
		// ...
	}
}
