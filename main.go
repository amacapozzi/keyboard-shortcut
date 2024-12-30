package main

import (
	"fmt"
	"log"
	"os"

	"github.com/micmonay/keybd_event"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
	"golang.org/x/sys/windows/registry"
)

func main() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	regPath := "Software\\Microsoft\\Windows\\CurrentVersion\\Run"

	_, registryError := setRegistryKey(registry.CURRENT_USER, regPath, ex)

	if registryError != nil {
		log.Panic(err)
	}

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
	err = kbf.Launching()
	if err != nil {
		return "", err
	}
	return "Insert pressed", nil
}

func setRegistryKey(registryKey registry.Key, registryPath string, value string) (string, error) {

	regKey, err := registry.OpenKey(registryKey, registryPath, registry.ALL_ACCESS)

	if err != nil {
		return "", err
	}
	defer regKey.Close()

	errorStringValue := regKey.SetStringValue("KeyMapper", value)

	if errorStringValue != nil {
		return "", errorStringValue
	}

	return "Registry key setted successfully", nil
}
