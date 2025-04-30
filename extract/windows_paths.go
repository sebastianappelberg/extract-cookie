//go:build windows

package extract

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	localAppDataPath = os.Getenv("LOCALAPPDATA")
	appDataPath      = os.Getenv("APPDATA")
	programPath      = os.Getenv("PROGRAMFILES")
	programX86Path   = os.Getenv("PROGRAMFILES(X86)")

	firefoxProfilePath  = filepath.Join(appDataPath, "Mozilla", "Firefox", "Profiles", "*.default-release")
	firefoxCookieDbName = "cookies.sqlite"

	chromeUserDataPath = filepath.Join(localAppDataPath, "Google", "Chrome", "User Data")
	chromeCookiePath   = filepath.Join(chromeUserDataPath, "Default", "Network", "Cookies")
	chromeExecPath     = filepath.Join(programPath, "Google", "Chrome", "Application", "chrome.exe")

	chromiumUserDataPath = filepath.Join(localAppDataPath, "Chromium", "User Data")
	chromiumCookiePath   = filepath.Join(chromiumUserDataPath, "Default", "Network", "Cookies")
	chromiumExecPath     = filepath.Join(programPath, "Chromium", "Application", "chrome.exe")

	edgeUserDataPath = filepath.Join(localAppDataPath, "Microsoft", "Edge", "User Data")
	edgeCookiePath   = filepath.Join(edgeUserDataPath, "Default", "Network", "Cookies")
	edgeExecPath     = filepath.Join(programX86Path, "Microsoft", "Edge", "Application", "msedge.exe")
)

type Paths struct {
	ExecPath     string
	UserDataPath string
	CookieDbPath string
}

func (b browser) GetPaths() (Paths, error) {
	switch b {
	case Chrome:
		return Paths{
			UserDataPath: chromeUserDataPath,
			CookieDbPath: chromeCookiePath,
			ExecPath:     chromeExecPath,
		}, nil
	case Chromium:
		return Paths{
			UserDataPath: chromiumUserDataPath,
			CookieDbPath: chromiumCookiePath,
			ExecPath:     chromiumExecPath,
		}, nil
	case Edge:
		return Paths{
			ExecPath:     edgeExecPath,
			UserDataPath: edgeUserDataPath,
			CookieDbPath: edgeCookiePath,
		}, nil
	case Firefox:
		dbPath, err := firefoxCookieDBPath()
		if err != nil {
			return Paths{}, err
		}
		return Paths{
			CookieDbPath: dbPath,
		}, nil
	default:
		return Paths{}, fmt.Errorf("unsupported browser: %s", b)
	}
}

// firefoxCookieDBPath tries to find the cookies.sqlite belonging to the default profile.
func firefoxCookieDBPath() (string, error) {
	profilePath, err := findFirstMatch(firefoxProfilePath)
	if err != nil {
		return "", err
	}
	return filepath.Join(profilePath, firefoxCookieDbName), nil
}

func findFirstMatch(str string) (string, error) {
	p, err := filepath.Glob(str)
	if err != nil {
		return "", err
	}
	if len(p) > 0 {
		return p[0], nil
	}
	return "", fmt.Errorf("find %s failed", str)
}
