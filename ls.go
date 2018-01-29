package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

func main () {
  //ioutil.ReadDir(パス名)は([]os.FileInfo, error)の二つの返り値を持つ
  //特定ディレクトリのファイル一覧のsliceを取得
  files, err := ioutil.ReadDir(".")
    //エラー処理
    if err != nil{
      fmt.Printf("エラーがおきました\n")
      os.Exit(1)
      return
    }

  //os.FileInfoのName()でエントリー名を出力
  //ファイル一覧を出力
  for i := 0; i < len(files); i++ {
    dir := files[i]
    fmt.Printf("%s\n", dir.Name())
  }
  os.Exit(0)
}
