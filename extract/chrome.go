package extract

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/storage"
	"github.com/chromedp/chromedp"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func chromeCookies(b browser) ([]Cookie, error) {
	paths, err := b.GetPaths()
	if err != nil {
		return nil, fmt.Errorf("failed to find db path: %w", err)
	}
	return getChromeCookies(paths.UserDataPath, paths.ExecPath)
}

// getChromeCookies returns all cookies found in the browser's cookie database.
// The reason why chromedp is used is that chromium changed the way that its cookies are encrypted,
// making it impossible to go straight to the sqlite database. See https://security.googleblog.com/2024/07/improving-security-of-chrome-cookies-on.html.
// execPath is necessary because certain Chromium-based browsers such as Microsoft Edge won't work otherwise.
// If the browser in question is open, this function will return an empty slice.
func getChromeCookies(userDataDir, execPath string) ([]Cookie, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserDataDir(userDataDir),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
	)
	if execPath != "" {
		opts = append(opts, chromedp.ExecPath(execPath))
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Get all cookies from the browser
	var cookies []Cookie
	err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		storageCookies, err := storage.GetCookies().Do(ctx)
		if err != nil {
			return err
		}
		for _, c := range storageCookies {
			cookies = append(cookies, Cookie{
				Name:     c.Name,
				Value:    c.Value,
				Domain:   c.Domain,
				Path:     c.Path,
				Expires:  int64(c.Expires),
				Secure:   c.Secure,
				HttpOnly: c.HTTPOnly,
			})
		}
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to get cookies: %w", err)
	}
	return cookies, nil
}
