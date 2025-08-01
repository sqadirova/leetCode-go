package main

import (
	//"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Instance represents a Kubernetes worker instance.
type Instance struct {
	Capacity  int
	Remaining int
	Pods      []int
	ID        int
}

// ReadLinesFromStdin reads all lines from standard input until EOF or an error occurs.
// It returns a slice of strings, where each string is a line, and an an error if any.
//func ReadLinesFromStdin() ([]string, error) {
//	var lines []string
//	scanner := bufio.NewScanner(os.Stdin)
//
//	for scanner.Scan() {
//		line := scanner.Text()
//		lines = append(lines, line)
//	}
//
//	if scanner.Err() != nil {
//		return nil, scanner.Err()
//	}
//
//	return lines, nil
//}

func main() {
	// Read all lines from stdin using the provided helper function
	//lines, err := ReadLinesFromStdin()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//// Ensure we have at least 2 lines (capacity and pod requests)
	//if len(lines) < 2 {
	//	fmt.Fprintf(os.Stderr, "Error: Insufficient input. Expected at least 2 lines (capacity and pods).\n")
	//	os.Exit(1)
	//}

	// For demonstration purposes, I will use a hardcoded input
	lines := []string{
		"10",         // Instance type CPU capacity
		"6,1,10,7,4", // Pod CPU requests
	}

	// First line is instance type CPU capacity
	instanceCapacityStr := strings.TrimSpace(lines[0])
	instanceCapacity, err := strconv.Atoi(instanceCapacityStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing instance capacity '%s': %v\n", instanceCapacityStr, err)
		os.Exit(1)
	}

	// Second line is comma-separated pod CPU requests
	podRequestsStr := strings.TrimSpace(lines[1])
	podStrs := strings.Split(podRequestsStr, ",")

	var podRequests []int
	for _, s := range podStrs {
		s = strings.TrimSpace(s)
		if s == "" {
			continue // Skip empty strings
		}
		cpu, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing pod CPU request '%s': %v\n", s, err)
			os.Exit(1)
		}
		podRequests = append(podRequests, cpu)
	}

	// --- Container Scheduling Logic (same as before) ---

	// Sort pods in descending order (First Fit Decreasing heuristic)
	sort.Sort(sort.Reverse(sort.IntSlice(podRequests)))

	var instances []*Instance // List of active instances
	instanceCount := 0        // Counter for instance IDs

	// Schedule pods
	for _, podCPU := range podRequests {
		foundInstance := false
		for _, inst := range instances {
			if inst.Remaining >= podCPU {
				inst.Remaining -= podCPU
				inst.Pods = append(inst.Pods, podCPU)
				foundInstance = true
				break
			}
		}

		if !foundInstance {
			instanceCount++
			newInstance := &Instance{
				ID:        instanceCount,
				Capacity:  instanceCapacity,
				Remaining: instanceCapacity - podCPU,
				Pods:      []int{podCPU},
			}
			instances = append(instances, newInstance)
		}
	}

	// Ensure instances are processed in their creation order for output
	sort.Slice(instances, func(i, j int) bool {
		return instances[i].ID < instances[j].ID
	})

	// --- Generate Output ---
	for _, inst := range instances {
		// Join pod CPU requests for output
		podStrings := make([]string, len(inst.Pods))
		for j, p := range inst.Pods {
			podStrings[j] = strconv.Itoa(p)
		}
		podsOutput := strings.Join(podStrings, ",")

		// Determine status message
		status := ""
		if inst.Remaining == 0 {
			status = "full"
		} else {
			status = fmt.Sprintf("%d spare CPUs", inst.Remaining)
		}

		// Get instance number suffix (1st, 2nd, 3rd, etc.)
		instanceNumSuffix := ""
		switch inst.ID {
		case 1:
			instanceNumSuffix = "1st"
		case 2:
			instanceNumSuffix = "2nd"
		case 3:
			instanceNumSuffix = "3rd"
		default:
			instanceNumSuffix = fmt.Sprintf("%dth", inst.ID)
		}

		// Print the output line
		fmt.Printf("%s # %s instance, %s\n", podsOutput, instanceNumSuffix, status)
	}
}
