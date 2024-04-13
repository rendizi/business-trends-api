package db

import "errors"

type User struct {
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Password       string `json:"password"`
	IIN            string `json:"iin"`
	City           string `json:"city"`
	Direction      string `json:"direction"`
	Email          string `json:"email"`
	Description    string `json:"description"`
	RevenueHistory struct {
		MonthAgo      string `json:"month-ago"`
		TwoMonthAgo   string `json:"two-month-ago"`
		ThreeMonthAgo string `json:"three-month-ago"`
		SixMonthAgo   string `json:"six-month-ago"`
		YearAgo       string `json:"year-ago"`
	} `json:"revenue-history"`
	Revenue string `json:"revenue"`
}

func InsertUser(user User) error {
	query := `INSERT INTO users(
                  username,surname,iin,city,direction, password, email, description, revenue
                  ) values($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	revenue := user.RevenueHistory.MonthAgo + ";" + user.RevenueHistory.TwoMonthAgo + ";" +
		user.RevenueHistory.ThreeMonthAgo + ";" + user.RevenueHistory.SixMonthAgo + ";" +
		user.RevenueHistory.YearAgo
	_, err := db.Exec(query,
		user.Name, user.Surname, user.IIN, user.Direction, user.Password, user.Email,
		user.Description, revenue)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(email string) (User, error) {
	query := `SELECT direction, description, city, revenue FROM users WHERE email = $1`
	var direction, description, city, revenue string
	err := db.QueryRow(query, email).Scan(&direction, &description, &city, &revenue)
	if err != nil {
		return User{}, err
	}
	return User{
		Direction:   direction,
		Description: description,
		City:        city,
		Revenue:     revenue,
	}, nil
}

func ValidatePassword(email, password string) error {
	query := `SELECT password FROM users WHERE email = $1`
	var storedPassword string
	err := db.QueryRow(query, email).Scan(&storedPassword)
	if err != nil {
		return err
	}
	if storedPassword != password {
		return errors.New("incorrect password")
	}
	return nil
}
