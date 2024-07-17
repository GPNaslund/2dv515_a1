package main;

import "gn222gq/rec-sys/internal"

func main() {
  app := internal.NewApp().Create()
  app.Listen(":8080") 
}
