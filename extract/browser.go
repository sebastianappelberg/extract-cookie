package extract

import "fmt"

const (
	Chrome browser = iota + 1
	Chromium
	Edge
	Firefox
)

type browser int

func (b browser) String() string {
	switch b {
	case Chrome:
		return "Chrome"
	case Chromium:
		return "Chromium"
	case Edge:
		return "Edge"
	case Firefox:
		return "Firefox"
	default:
		return "Unknown"
	}
}

func (b browser) GetCookies() ([]Cookie, error) {
	switch b {
	case Chromium, Chrome, Edge:
		return chromeCookies(b)
	case Firefox:
		return firefoxCookies()
	default:
		return nil, fmt.Errorf("unsupported browser: %s", b)
	}
}

func getBrowser(browser string) (browser, error) {
	switch browser {
	case "chrome":
		return Chrome, nil
	case "chromium":
		return Chromium, nil
	case "edge":
		return Edge, nil
	case "firefox":
		return Firefox, nil
	default:
		return 0, fmt.Errorf("unsupported browser: %s", browser)
	}
}
