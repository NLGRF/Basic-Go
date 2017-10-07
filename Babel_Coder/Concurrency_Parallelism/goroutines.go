func searchFromFolder(keyword string, folder string) {
  // ทำการค้นหา
}

func search(keyword string) {
  folders := [3]string{"Document", "Image", "Library"}

  for _, folder := range folders {
    // ตรงนี้
    go searchFromFolder(keyword, folder)
  }
}

func main() {
  search("dog")
}
