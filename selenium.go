package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	// FireFox драйвер без указания версии
	caps := selenium.Capabilities{"browserName": "firefox"}

	wd, _ := selenium.NewRemote(caps, "")
	defer wd.Quit()

	// Открываем страницу
	_ = wd.Get("https://google.ru/")

	// Находим поле и вводим текст
	elem, _ := wd.FindElement(selenium.ByCSSSelector, ".gLFyf.gsfi")
	_ = elem.Clear()
	_ = elem.SendKeys("golang")

	time.Sleep(time.Millisecond * 10)

	// Находим кнопку и кликаем по ней
	btn, _ := wd.FindElement(selenium.ByCSSSelector, ".aajZCb .gNO89b")
	_ = btn.Click()

	// Делаем скрин результата
	screen, _ := wd.Screenshot()

	// Записываем результат в файл
	file, err := os.Create("test.png")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	_, _ = file.Write(screen)
}
