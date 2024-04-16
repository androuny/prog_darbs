package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"mypackage/pi"
	"net/http"
	"os"
	"strconv"
	"time"
    "syscall"
    "golang.org/x/sys/windows"

	"github.com/shirou/gopsutil/v3/cpu"
)

const DEBUG = false
const URL = "" // enter api endpoint for data submission

type Results struct {
    One_bil int `json:"one_bil"`
    Hundred_mil  int `json:"hundred_mil"`
    Ten_mil  int `json:"ten_mil"`
    One_mil  int `json:"one_mil"`
    CPU string `json:"CPU"`
}

func setup_console() {
    hwnd := windows.Handle(syscall.Stdout)
    var mode uint32
    err := windows.GetConsoleMode(hwnd, &mode)
    if err != nil {
        panic(err)
    }
    err = windows.SetConsoleMode(hwnd, mode|0x0004)
    if err != nil {
        panic(err)
    }
}

func writeJSONToFile(filename string, jsonBytes []byte) error {
    // Open the file for writing.
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    // Write the JSON data to the file.
    _, err = file.Write(jsonBytes)
    if err != nil {
        return err
    }

    return nil
}

func calc_time_for_throws_ms(throws int) int {
    start := time.Now()
    pi := pi.Calc_pi(throws)
    duration := time.Since(start)

    if DEBUG == true {
        fmt.Println("Calculated π:", pi)
        fmt.Println("Difference between math.Pi and calculated:", math.Pi - pi)
        fmt.Println("Throws: ", throws)
        fmt.Printf("Time elapsed: %v ms", duration.Milliseconds())
        fmt.Printf("-------------------------------------------------------------")
    }
    return int(duration.Milliseconds())
}



func main() {
    setup_console()

    cpuStat, _ := cpu.Info()
    var cpu_model = cpuStat[0].ModelName

    println(cpu_model)

    fmt.Println("Starting calculations... It may look stuck but its working")
    var one_bil_throws_ms = calc_time_for_throws_ms(1000000000)
    var hundred_mil_throws_ms = calc_time_for_throws_ms(100000000)
    var ten_mil_throws_ms = calc_time_for_throws_ms(10000000)
    var one_mil_throws_ms = calc_time_for_throws_ms(1000000)


    results := Results{One_bil: one_bil_throws_ms, Hundred_mil: hundred_mil_throws_ms, Ten_mil: ten_mil_throws_ms, One_mil: one_mil_throws_ms, CPU: cpu_model}

    if DEBUG == true {
        // Encode the struct to JSON format.
        jsonBytes, err := json.Marshal(results)
        if err != nil {
            fmt.Println("Error encoding JSON:", err)
            return
        }

        // Write the JSON data to a file.
        err = writeJSONToFile("results.json", jsonBytes)
        if err != nil {
            fmt.Println("Error writing JSON to file:", err)
            return
        }
    }



    // send the results to the server
    jsonData := map[string]string{
        "one_bil": strconv.Itoa(one_bil_throws_ms),
        "hundred_mil": strconv.Itoa(hundred_mil_throws_ms),
        "ten_mil": strconv.Itoa(ten_mil_throws_ms),
        "one_mil": strconv.Itoa(one_mil_throws_ms),
        "cpu": cpu_model,
    }
    jsonValue, _ := json.Marshal(jsonData)

    req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Println(err)
        return
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    fmt.Println("Server Response Status:", resp.Status)

}