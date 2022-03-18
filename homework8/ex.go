package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	direct string
	mapx   = make(map[int]string)
	i      int
	mx     sync.RWMutex
	wg     = sync.WaitGroup{}
	remove = flag.Bool("remove", true, "Нужно ли удалить файл?")
	//help   = flag.String(help, "", "Программа находит дубликаты файлов.")

	// карта для хранения дубликатов файлов
	mapsame = make(map[string]string)
)

// Ввод папки.
func getWd() string {
	var dir string
	fmt.Println("Введите путь до папки.")
	_, err := fmt.Scan(&dir)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// ListByReadDir returns a map[int]string
func ListByReadDir(path string) {
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, val := range lst {
		if val.IsDir() {
			ListByReadDir(path + "/" + val.Name())
		} else {
			fmt.Println(val.Name(), path)
			mapx[i] = path + "/" + val.Name()
			i = i + 1
		}
	}
}

// compare1 look for the same files
func compare1() {
	wg.Add(len(mapx))
	for k, _ := range mapx {
		fi1, err := os.Stat(mapx[k])
		if err != nil {
			panic(err)
		}
		for kk, _ := range mapx {
			fi2, err := os.Stat(mapx[kk])
			if err != nil {
				panic(err)
			}
			if k != kk {
				if fi1.Name() == fi2.Name() &&
					fi1.Size() == fi2.Size() {
					fmt.Println(fi1.Name(), mapx[k], ";", fi2.Name(), mapx[kk])
					mapsame[fi1.Name()] = mapx[k]
				}
			}
		}
	}
}

//thesimilar look for the same files
func thesimilar() {
	wg.Add(len(mapx))
	for k, _ := range mapx {
		fi1, err := os.Stat(mapx[k])
		if err != nil {
			panic(err)
		}

		go func(k int, fi1 os.FileInfo) {
			defer wg.Done()
			similar(k, fi1)
		}(k, fi1)
	}
	wg.Wait()
}

//similar is part of the thesimilar
func similar(i int, fi os.FileInfo) {
	mx.RLock()
	defer mx.RUnlock()
	for kk, _ := range mapx {
		fi2, err := os.Stat(mapx[kk])
		if err != nil {
			panic(err)
		}
		if i != kk {
			if fi.Name() == fi2.Name() &&
				fi.Size() == fi2.Size() {
				fmt.Println(fi.Name(), mapx[i])
			}
		}
	}
}

func main() {
	var qwest string
	flag.Parse()
	//fmt.Println(*help)
	fmt.Println(*remove)
	//direct := "C:/Users/Acer 1/go/src/NewNew1"
	//fmt.Println(direct)
	direct = getWd()
	_, err := ioutil.ReadDir(direct)
	if err != nil {
		fmt.Println("Ваш путь неправильный, загружаем путь до локальной папки.")
		direct, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}
	i = 0
	ListByReadDir(direct)
	//fmt.Println(mapx)

	//thesimilar()
	compare1()
	if len(mapsame) == 0 {
		fmt.Println("Нет дубликатов файлов.")
	}
	if *remove && len(mapsame) > 0 {
		fmt.Println("Вы действительно хотите удалить дубликаты файлов? Да/Нет")
		fmt.Scan(&qwest)
		switch qwest {
		case "Да":
			//fmt.Println(mapsame)
			for _, val := range mapsame {
				os.Remove(val)
			}
		case "Нет":
			fmt.Scanln()
		default:
			fmt.Println("Неправильный ввод.")
		}
	}
}
