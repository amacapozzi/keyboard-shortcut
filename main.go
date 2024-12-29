package main

import (
	"fmt"
	"log"

	"github.com/micmonay/keybd_event"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func main() {
	mainthread.Init(fn)

}
func fn() {

	for {

		hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyK)
		err := hk.Register()
		if err != nil {
			log.Fatalf("hotkey: failed to register hotkey: %v", err)
			return
		}

		log.Printf("hotkey: %v is registered\n", hk)
		<-hk.Keydown()
		log.Printf("hotkey: %v is down\n", hk)
		<-hk.Keyup()
		log.Printf("hotkey: %v is up\n", hk)
		msg, err := pressInsertKey()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(msg)
		hk.Unregister()
		log.Printf("hotkey: %v is unregistered\n", hk)

	}
}

func pressInsertKey() (string, error) {
	kbf, err := keybd_event.NewKeyBonding()

	if err != nil {
		return "", err
	}

	kbf.SetKeys(keybd_event.VK_INSERT)
	return "Insert pressed", nil
}
