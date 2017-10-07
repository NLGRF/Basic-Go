package main

func main() {

  // อาร์เรย์ของข้อความจำนวนสามช่อง
 var names [3]string

names[0] = "Somchai"
names[1] = "Somsree"
names[2] = "Somset"

names := [3]string{"Somchai", "Somsree", "Somset"}

// การประกาศ slice เราไม่ต้องระบุจำนวนช่อง
// เพราะเราสามารถเพิ่ม element เข้าไปได้อย่างอิสระภายหลัง
var names []string

// เพิ่ม element เข้าไป
names = append(names, "Somchai")
names = append(names, "Somsree")
names = append(names, "Somset")

names := []string{"Somchai", "Somsree", "Somset"}

// for-loop => range

names := [3]string{"Somchai", "Somsree", "Somset"}

for index, name := range names {
	fmt.Println(index, name)
}

// ผลลัพธ์
// 0 Somchai
// 1 Somsree
// 2 Somset
}
