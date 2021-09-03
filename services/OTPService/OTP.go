package otpservice

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"auth/repository/redisP"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func GenerateRandNum(phone string) *big.Int {

	val, err := strconv.Atoi(phone)
	if err != nil {
		log.Fatal(err)
	}

	p, err := rand.Int(rand.Reader, big.NewInt(int64(val)))
	//fmt.Println("this is the value of mobile number:", p)
	s := prefix(p.String())
	v := s[:4]
	t := prefix(v)
	//v1 := int(p.Int64()) % 1e4
	if err != nil {
		log.Fatal(err)
	}
	otp, err := strconv.Atoi(t)
	if err != nil {
		log.Fatal(err)
	}
	//	v,err := rand.Int(rand.Reader,p
	//	fmt.Println("value of OTP", v)
	return big.NewInt(int64(otp))
}
func prefix(val string) string {
	if len(val) < 4 {
		fmt.Println("this is the otp with length less than 4", val)
		v, err := rand.Int(rand.Reader, big.NewInt(9))
		fmt.Println("prefix number", v)
		if err != nil {
			log.Fatal(err)
		}
		for i := (4 - len(val)); i > 0; i-- {
			val = v.String() + val
		}
		fmt.Println(val)

	}
	return val
}
func tester2() {
	var j, k int
	for i := 0; i < 50; i++ {
		p, err := rand.Int(rand.Reader, big.NewInt(time.Now().Unix()))
		//	fmt.Println(len(p.String()))
		if err != nil {
			log.Fatal(err)
		}
		v := GenerateRandNum(p.String())
		if len(v.String()) == 4 {
			j++
		}
		if len(v.String()) != 4 {

			k++
			fmt.Println("OTP which length not equal to 4 ", v)
		}

	}
	fmt.Println("times otp length is 4: ", j, "<------>", "times otp length was not equal to 4: ", k)
}

func GenerateUniqueId() uuid.UUID {
	//d := generateRandNum("7557588325")
	//	n := uuid.MustParse(d.String())
	//g, _ := uuid.NewRandomFromReader(rand.Reader)
	//fmt.Println(time.Now())
	e, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(e)
	//fmt.Println(w)
	return e
}

func StoreData(phone string) (*big.Int, uuid.UUID) {
	v := GenerateRandNum(phone)
	redisP.Store(phone, v.String())
	fmt.Println("this is running", v)
	val := GenerateUniqueId()
	redisP.Store(val.String(), phone)
	return v, val
	//Compare(v.String(), val.String())
}

func Compare(otp string, nonce string) bool {
	v := redisP.FindId(nonce)
	//a := database.Connect().Client.Get(context.TODO(), v).Val()

	if otp == v {
		fmt.Println("OTP found ", otp, v)
		return true
	} else {
		fmt.Println("Please enter a valid OTP")
		return false
	}

}
