package main

func realModifySlice(slice []int) {
	slice[0] = 9
}
func realModifyMap(myMap map[string]string) {
	myMap["name"] = "changeName"
}
func realModifyChannel(channel chan int) {
	channel <- 9
}
