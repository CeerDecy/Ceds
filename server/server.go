package main

import "Ceds/log"

var banner = `
   ______          ____           ___     
  / ____/__  ___  / __ \___  ____/ (_)____
 / /   / _ \/ _ \/ /_/ / _ \/ __  / / ___/
/ /___/  __/  __/ _. _/  __/ /_/ / (__  ) 
\____/\___/\___/_/ |_|\___/\____/_/____/   
`

func main() {
	println(banner)

	log.Log.Info("Starting server", "")
	//tcp.ListenAndServe()
}
