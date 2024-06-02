package persist

import (
	"fmt"
	"project/railroad/model"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			s := item.(*model.Character)
			s.CharacterPrint()
			fmt.Println()
			//log.Printf("Item Saver: got item #%d: %v", itemCount, s)
			itemCount++
		}
	}()
	return out
}
