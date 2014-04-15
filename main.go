package main

func main() {
	// Begin by parsing the command line arguments
	parseFlags()

	// Print the startup message
	printStartupMessage()

	// Finally start the load balancer
	startProxy()
}
