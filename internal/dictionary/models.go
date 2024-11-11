package dictionary

type Country struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

type City struct {
	Id        int64  `db:"id"`
	CountryId int64  `db:"country_id"`
	Name      string `db:"name"`
}

type Hotel struct {
	Id         int64  `db:"id"`
	CityId     int64  `db:"city_id"`
	Name       string `db:"name"`
	Priority   int64  `db:"priority"`
	StarsCount int64  `db:"stars_count"`
}
