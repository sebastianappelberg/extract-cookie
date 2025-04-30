package extract

type Cookie struct {
	Name     string
	Value    string
	Domain   string
	Path     string
	Expires  int64
	Secure   bool
	HttpOnly bool
}

type Result struct {
	Cookies []Cookie `json:"cookies"`
}

func Cookies(browser string) (Result, error) {
	b, err := getBrowser(browser)
	if err != nil {
		return Result{}, err
	}
	cookies, err := b.GetCookies()
	if err != nil {
		return Result{}, err
	}
	return Result{Cookies: cookies}, err
}
