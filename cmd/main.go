package main

import (
	"github.com/IKostarev/HFLabs_test/pkg"
)

func main() {
	url := "https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999"
	title, body := pkg.Parse(url)

	pkg.WriteToSheet(title, body)
}
