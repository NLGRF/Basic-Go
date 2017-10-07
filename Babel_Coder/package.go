/*
-- src
   -- godubai
      -- main.go
      -- search
         -- main.go
*/

func searchFromFolder(keyword string, folder string, wg *sync.WaitGroup) {
  // ทำการค้นหา
  wg.Done()
}

func Run(keyword string) {
  folders := [3]string{"Document", "Image", "Library"}
  var wg sync.WaitGroup
  wg.Add(len(folders))
  for _, folder := range folders {
    go searchFromFolder(keyword, folder, &wg)
  }
  wg.Wait()
}

/*
import "godubai/search"

func main() {
  search.Run("dog")
}
*/
