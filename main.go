package main

import (
	"flag"
	"log"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	logger *log.Logger

	// IPMI credentials
	ipmiAddress  = flag.String("address", "192.168.2.200", "IPMI/iDRAC address")
	ipmiUser     = flag.String("user", "root", "IPMI user")
	ipmiPassword = flag.String("password", "calvin", "IPMI password")

	// Fan control parameters
	interval    = flag.Int("interval", 10, "Interval in seconds for updating fan speed")
	T0          = *flag.Float64("t0", 92.0, "Midpoint temperature in Celsius")
	k           = *flag.Float64("k", 0.5, "Steepness of the curve")
	L           = *flag.Float64("L", 100, "Maximum fan speed")
	minFanSpeed = *flag.Float64("minFanSpeed", 5, "Minimum fan speed")
)

func init() {
	logger = log.New(os.Stdout, "PID Controller: ", log.LstdFlags)

	flag.Parse()
}

func main() {
	interval := time.Duration(*interval) * time.Second

	for {
		updateFanSpeed()
		time.Sleep(interval)
	}
}

func updateFanSpeed() {
	logger.Println("Updating fan speed")
	temperature, err := getTemperature()
	if err != nil {
		logger.Printf("Error getting temperature: %v", err)
		return
	}

	logger.Printf("Temperature: %d\n", temperature)

	speed := getFanSpeed(temperature)
	logger.Printf("Setting fan speed to: %d", speed)

	if err := setFanSpeed(speed); err != nil {
		logger.Printf("Error setting fan speed: %v", err)
	}
}

func setFanSpeed(speed int) error {
	_, err := executeCommand("-I", "lanplus", "-H", *ipmiAddress, "-U", *ipmiUser, "-P", *ipmiPassword, "raw", "0x30", "0x30", "0x02", "0xff", strconv.FormatUint(uint64(speed), 16))
	return err
}

func getFanSpeed(temperature int) int {
	temp := float64(temperature)
	speed := L / (1 + math.Exp(-k*(temp-T0)))

	if speed < minFanSpeed {
		speed = minFanSpeed
	}

	return int(speed)
}

func getTemperature() (int, error) {
	output, err := executeCommand("-I", "lanplus", "-H", *ipmiAddress, "-U", *ipmiUser, "-P", *ipmiPassword, "sdr", "type", "temperature")
	if err != nil {
		return -1, err
	}

	return parseTemperature(output), nil
}

func executeCommand(args ...string) (string, error) {
	cmd := exec.Command("ipmitool", args...)
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

// parseTemperature parses the output of the `ipmitool sdr type temperature` command and returns the highest temperature
// found in the output.
func parseTemperature(output string) int {
	re := regexp.MustCompile(`\| ([0-9]+) degrees C`)
	matches := re.FindAllStringSubmatch(output, -1)

	highestTemp := -1
	for _, match := range matches {
		if len(match) == 2 {
			temp, err := strconv.Atoi(match[1])
			if err == nil && temp > highestTemp {
				highestTemp = temp
			}
		}
	}

	return highestTemp
}
