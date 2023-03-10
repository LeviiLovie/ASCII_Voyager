package foo

import "fmt"

type MenuItem int

const (
	MenuItemNewGame MenuItem = iota
	MenuItemLoadGame
	MenuItemCredits
	MenuItemExit
)

var MenuTasks = []MenuItem{MenuItemNewGame, MenuItemLoadGame, MenuItemCredits, MenuItemExit}

func (m MenuItem) String() string {
	switch m {
	case MenuItemNewGame:
		return "New Game"
	case MenuItemLoadGame:
		return "Load Game"
	case MenuItemCredits:
		return "Credits"
	case MenuItemExit:
		return "Exit"
	default:
		return "Unknown"
	}
}

func (m MenuItem) Next() MenuItem {
	if m < MenuItemExit {
		return m + 1
	}
	return MenuItemNewGame
}

func (m MenuItem) Prev() MenuItem {
	if m > MenuItemNewGame {
		return m - 1
	}
	return MenuItemExit
}

var MenuCredits = []string{"LeviLovie", "Oto", "Learunaso"}

func MenuDrawTasks(chose MenuItem, x, y int) {
	for _, item := range MenuTasks {
		MoveCursor(x, y+int(item))
		if chose == item {
			fmt.Printf(TEXT_WHITE_BOLD + item.String() + TEXT_RESET)
		} else {
			fmt.Print(item.String())
		}
	}
}

func MenuDrawCredits(x, y int, prefix string) {
	for i, credit := range MenuCredits {
		MoveCursor(x, y+i)
		fmt.Printf(prefix, credit)
		fmt.Printf(prefix, credit)
	}
}

func MenuSetUp() {
	SetTerminalSize(TerminalWidth, TerminalHeight)
	ClearScreen()
	NotVisibleCursor()
}

func MenuDrawLogo() {
	WriteTextOnCenter("      _         ______      ______   _____   _____          ____   ____     ___     ____  ____        _          ______    ________   _______     ", TerminalWidth, 5)
	WriteTextOnCenter("     / \\      .' ____ \\   .' ___  | |_   _| |_   _|        |_  _| |_  _|  .'   `.  |_  _||_  _|      / \\       .' ___  |  |_   __  | |_   __ \\    ", TerminalWidth, 6)
	WriteTextOnCenter("    / _ \\     | (___ \\_| / .'   \\_|   | |     | |            \\ \\   / /   /  .-.  \\   \\ \\  / /       / _ \\     / .'   \\_|    | |_ \\_|   | |__) |   ", TerminalWidth, 7)
	WriteTextOnCenter("   / ___ \\     _.____`.  | |          | |     | |             \\ \\ / /    | |   | |    \\ \\/ /       / ___ \\    | |   ____    |  _| _    |  __ /    ", TerminalWidth, 8)
	WriteTextOnCenter(" _/ /   \\ \\_  | \\____) | \\ `.___.'\\  _| |_   _| |_             \\ ' /     \\  `-'  /    _|  |_     _/ /   \\ \\_  \\ `.___]  |  _| |__/ |  _| |  \\ \\_  ", TerminalWidth, 9)
	WriteTextOnCenter("|____| |____|  \\______.'  `.____ .' |_____| |_____|             \\_/       `.___.'    |______|   |____| |____|  `._____.'  |________| |____| |___| ", TerminalWidth, 10)
	WriteTextOnCenter("                                                                                                                                             ", TerminalWidth, 11)
}
