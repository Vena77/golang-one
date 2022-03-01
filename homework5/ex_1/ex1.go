package main

import (
  "fmt"
  "sync"
)

const count = 1000

func main() {
  var (
     counter int

     // Создаем экземпляр
     wg = sync.WaitGroup{}
  )

  // Инициализируем семафор исходным состоянием
  wg.Add(count)
  for i := 0; i < count; i += 1 {
     go func() {
        counter += 1

        // Выполняем декремент семафора
        wg.Done()
     }()
  }
  // Ждем обнуления семафора
  wg.Wait()

  // Выводим показание общего счетчика
  fmt.Println(counter)
}
