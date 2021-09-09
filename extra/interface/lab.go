package main

import "fmt"

type PersistenceSwitchConf struct {
	WhitelistUserIds []int64 `json:"whitelist_user_ids"`
	IsOpenAll        bool    `json:"is_open_all"`
}

func Parser() interface{} {

	//_ := PersistenceSwitchConf{IsOpenAll: true}
	return nil
}

func main() {
	p := Parser()
	c := p.(*PersistenceSwitchConf)
	fmt.Println(c)
}
