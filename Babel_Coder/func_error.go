package main

import (
	"errors"
	"fmt"
)

func main() {
	// เพราะว่าฟังก์ชันคืนค่าสองค่า เราจึงประกาศตัวแปรมารองรับได้พร้อมกันสองตัว
	result, err := divide(5, 3)

    // ตรวจสอบก่อนว่ามี error ไหม ถ้ามีก็จบโปรแกรมไปแบบไม่ค่อยสวยด้วย Exit(1)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(result)
}

// คืนค่า float และ error ออกไปพร้อมกันจากฟังก์ชัน
func divide(dividend float32, divisor float32) (float32, error) {
	if divisor == 0.0 {
		err := errors.New("Division by zero!")
		return 0.0, err
	}

	return dividend / divisor, nil
}
