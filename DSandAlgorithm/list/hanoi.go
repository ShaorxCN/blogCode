package main

import (
	"log"
)

func hanoi(l []int,x,y,z string){
	length := len(l)
	if length == 1{
		log.Printf("move %d from %s to %s",l[0],x,z)
	}else{
		hanoi(l[:length-1],x,z,y)
		log.Printf("move %d from %s to %s",l[length-1],x,z)
		hanoi(l[:length-1],y,x,z)
	}


}

func main(){
	l := []int{1,2,3,4,5,6,7}

	hanoi(l,"x","y","z")
	log.Println("check")
	l2 := []int{1,2,3,4}
	hanoi(l2,"x","y","z")
}



