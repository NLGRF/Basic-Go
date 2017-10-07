// Go ไม่มี inheritance & ไม่มี class

// Composition over Inheritance 

class Devise {
  ip: string;
  location: string; // ที่จัดเก็บ
}

class Device {
  ip: string;
  location: string;
}

class Printer extends Device {
  print() {
    // เป็นปริ้นเตอร์ ก็ต้องปริ้นงานได้ซิ
  }
}

class Device {
  ip: string;
  location: string;
}

class Harddisk extends Device {
  store() {
    // จัดเก็บข้อมูล
  }
}
