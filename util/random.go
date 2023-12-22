package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
 

	  rand.Seed(time.Now().Unix())
	
}

func RandomInt(min , max int64)int64{

	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {


var sb strings.Builder

for i :=0 ; i<n ;i++{
 
	v :=  byte(RandomInt(65 ,90))
	sb.WriteByte(v)
}


return  sb.String()
}

