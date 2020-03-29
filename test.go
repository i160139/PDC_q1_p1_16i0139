package quiz01PDC

import "fmt"

func InfoDesk() {
	showed := [3]bool{false, false, false}
	check := false
	menuitem := -1;
	for ;check == false; {

		fmt.Println("1 – Print Covid19 cases in Pakistan\n2 – Print Covid19 cases in SouthKorea\n3" +
			" – Print Covid19 cases in France\n4 – Print my personalized message to Coronavirus\n0 – Exit")

		fmt.Print("Please enter a number from menu=")
		fmt.Scan(&menuitem)

		if menuitem > -1 && menuitem < 5 {
			switch menuitem {
			case 1:
				fmt.Println("Total Cases: 1, ,526Total Deaths: 13, Total Recovered: 29")
				showed[menuitem-1] = true
				break
			case 2:
				fmt.Println("Total Cases: 9,583, Total Deaths: 152, Total Recovered: 5,033")
				showed[menuitem-1] = true
				break
			case 3:
				fmt.Println("Total Cases: 37,575, Total Deaths: 2,314, Total Recovered: 5,700")
				showed[menuitem-1] = true
				break
			case 4:
				fmt.Println("Corona please go")
				break
			case 0:
				if showed[0] && showed[1] && showed[2] {
					check = true
				}
				if check {
					return
				}else {
					fmt.Println("Please first see info of all countries")
				}
				break
			}
		}
	}
}