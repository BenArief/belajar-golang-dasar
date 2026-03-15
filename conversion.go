package main

func main(){
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	println(nilai32)
	println(nilai64)
	println(nilai16)

	var name = "Budi Iskandar"
	var e = name[0]
	var eString = string(e)
	println(name)
	println(e)
	println(eString)
}