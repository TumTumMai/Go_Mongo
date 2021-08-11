package model

type User struct {
	Name string `bson:"name"`
	City string `bson:"city,omitempty"` ///omitemptyจะไม่สร้างถ้าไม่ส่งค่าไป
	Age  int    `bson:"age,omitempty"`
}
