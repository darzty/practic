package main

import (
	"fmt"
	"math"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Status string
	Pages  int
}

type Point struct {
	X float64
	Y float64
}

type UserProfile struct {
	Username     string
	Age          int
	FriendsCount int
	IsVerified   bool
	Rating       float64
}

type Lesson struct {
	Subject string
	Room    int
	Teacher string
}

const (
	Online       = "Online"
	Offline      = "Offline"
	Away         = "Away"
	DoNotDisturb = "Do Not Disturb"
)

const (
	Monday    = "Monday"
	Wednesday = "Wednesday"
)

func main() {
	// Задача 1
	rent := 95000.0
	newRent := rent * 1.10
	fmt.Printf("Задача 1: Новая стоимость аренды: %.2f\n", newRent)

	// Задача 2
	laptopsCost := 6 * 55480.0
	monitorsCost := 3 * 21830.0
	miceCost := 11 * 890.0
	keyboardsCost := 5 * 1560.0
	totalBudget := laptopsCost + monitorsCost + miceCost + keyboardsCost
	fmt.Printf("Задача 2: Общая сумма для бюджета: %.2f руб.\n", totalBudget)

	// Задача 3
	totalStorage := 5000.0
	fileSize := 256.0
	filesCount := int(totalStorage / fileSize)
	remainingSpace := totalStorage - float64(filesCount)*fileSize
	fmt.Printf("Задача 3: Можно разместить файлов: %d, останется места: %.2f Гб\n", filesCount, remainingSpace)

	// Задача 4
	var fahrenheit float64
	fmt.Print("Задача 4: Введите температуру по Фаренгейту: ")
	fmt.Scan(&fahrenheit)
	celsius := 5.0 / 9.0 * (fahrenheit - 32.0)
	fmt.Printf("Температура по Цельсию: %.2f\n", celsius)

	// Задача 5
	var radius float64
	fmt.Print("Задача 5: Введите радиус клумбы: ")
	fmt.Scan(&radius)
	length := 2 * math.Pi * radius
	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("Длина окружности: %.2f, Площадь круга: %.2f\n", length, area)

	// Задача 6
	var initialSum, rate float64
	var years int
	fmt.Print("Задача 6: Введите начальную сумму вклада, годовую ставку (%%) и количество лет через пробел: ")
	fmt.Scan(&initialSum, &rate, &years)
	totalDeposit := initialSum * math.Pow(1+rate/100.0, float64(years))
	fmt.Printf("Итоговая сумма вклада: %.2f\n", totalDeposit)

	// Задача 7
	book1 := Book{ID: 1, Title: "Go Programming", Author: "Donovan", Year: 2015, Status: "доступна", Pages: 380}
	book2 := Book{ID: 2, Title: "Clean Code", Author: "Martin", Year: 2008, Status: "доступна", Pages: 464}
	book3 := Book{ID: 3, Title: "Code Complete", Author: "McConnell", Year: 2004, Status: "доступна", Pages: 960}
	book1.Status = "выдана"
	fmt.Printf("Задача 7: Книга '%s' изменила статус на: %s\n", book1.Title, book1.Status)
	_ = book2
	_ = book3

	// Задача 8
	p1 := Point{X: 0, Y: 0}
	p2 := Point{X: 3, Y: 4}
	distance := math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
	fmt.Printf("Задача 8: Расстояние между точками: %.2f\n", distance)

	// Задача 9
	var purchaseAmount float64
	fmt.Print("Задача 9: Введите сумму покупки: ")
	fmt.Scan(&purchaseAmount)
	discountedAmount := purchaseAmount * 0.80
	fmt.Printf("Сумма покупки со скидкой 20%%: %.2f\n", discountedAmount)

	// Задача 10
	var a, b int
	fmt.Print("Задача 10: Введите два целых числа (a и b) для деления через пробел: ")
	fmt.Scan(&a, &b)
	exactDiv := float64(a) / float64(b)
	roundedNearest := math.Round(exactDiv)
	roundedFloor := math.Floor(exactDiv)
	fmt.Printf("Точный результат: %.2f, Округление до ближайшего: %.2f, Округление вниз: %.2f\n", exactDiv, roundedNearest, roundedFloor)

	// Задача 11
	userMe := UserProfile{Username: "student_go", Age: 18, FriendsCount: 15, IsVerified: true, Rating: 4.8}
	userFriend1 := UserProfile{Username: "gamer_sanya", Age: 18, FriendsCount: 22, IsVerified: false, Rating: 4.5}
	userFriend2 := UserProfile{Username: "timur_dota", Age: 19, FriendsCount: 30, IsVerified: true, Rating: 4.9}
	fmt.Printf("Задача 11: Профили: %+v, %+v, %+v\n", userMe, userFriend1, userFriend2)

	// Задача 12
	statuses := map[string]string{
		"Alex":  Online,
		"Sanya": Away,
		"Timur": DoNotDisturb,
		"Roman": Offline,
	}
	statuses["Alex"] = Offline
	statuses["Sanya"] = Online
	fmt.Println("Задача 12: Текущие статусы пользователей в чате:")
	for name, status := range statuses {
		if status == Online {
			fmt.Printf("- %s сейчас онлайн\n", name)
		}
	}

	// Задача 13
	schedule := map[string][]Lesson{
		Monday: {
			{Subject: "Programming", Room: 404, Teacher: "Ivanov I.I."},
			{Subject: "Databases", Room: 302, Teacher: "Petrov P.P."},
		},
		Wednesday: {
			{Subject: "Assembler", Room: 201, Teacher: "Smirnov S.S."},
			{Subject: "Math", Room: 105, Teacher: "Sidorov V.V."},
		},
	}
	selectedDay := Monday
	fmt.Printf("Задача 13: Расписание на %s:\n", selectedDay)
	for _, lesson := range schedule[selectedDay] {
		fmt.Printf("- Предмет: %s, Аудитория: %d, Преподаватель: %s\n", lesson.Subject, lesson.Room, lesson.Teacher)
	}
}
