package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"sort"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
	Uuid     string
	Imported bool
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

/*
func stampa(lista []*User) {
	for _, x := range lista {
		println(x.ID)
		println(x.Name)
		println(x.Username)
		println(x.Email)
		println(x.Address.Street)
		println(x.Address.Suite)
		println(x.Address.City)
		println(x.Address.Zipcode)
		println(x.Address.Geo.Lat)
		println(x.Address.Geo.Lng)
		println(x.Phone)
		println(x.Website)
		println(x.Company.Name)
		println(x.Company.CatchPhrase)
		println(x.Company.Bs)
		println(x.uuid)
		println(x.imported)
		println()
	}
}
*/
func setUuid(lista []*User) {
	for _, i := range lista {
		i.Uuid = uuid.NewString()
	}
}

func setImported(lista []*User) {
	for _, i := range lista {
		i.Imported = true
	}
}

func getLista() []*User {
	listaSeria := make([]*User, 0)
	listaJson, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		panic(err)
	}
	defer listaJson.Body.Close()

	if err := json.NewDecoder(listaJson.Body).Decode(&listaSeria); err != nil {
		panic(err)
	}

	coseUtili, err := ioutil.ReadAll(listaJson.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(coseUtili, &listaSeria)

	setImported(listaSeria)
	setUuid(listaSeria)
	newMail(listaSeria)

	sort.Slice(listaSeria, func(i, j int) bool { return listaSeria[i].Name < listaSeria[j].Name })
	//sorting dello slice listaSeria attraverso nome

	return listaSeria
}

func newMail(x []*User) {
	for _, i := range x {
		email := strings.SplitAfter(i.Email, "@")
		email[0] = strings.ToLower(email[0])
		email[0] = strings.ReplaceAll(email[0], "_", ".")
		email[1] = "madeinapp.net"
		i.Email = strings.Join(email, "")
	}
}

func stampaFIERAAA(lista []*User) *os.File {
	file, err := os.Create("file.csv")
	check(err)
	defer file.Close()
	write := csv.NewWriter(file)
	write.Comma = ';'
	defer write.Flush()
	for _, l := range lista {
		listaString := make([]string, 0)
		listaString = append(listaString, strconv.FormatInt(l.ID, 10))
		listaString = append(listaString, l.Name)
		listaString = append(listaString, l.Username)
		listaString = append(listaString, l.Address.Street)
		listaString = append(listaString, l.Address.Suite)
		listaString = append(listaString, l.Address.City)
		listaString = append(listaString, l.Address.Zipcode)
		listaString = append(listaString, l.Address.Geo.Lat)
		listaString = append(listaString, l.Address.Geo.Lng)
		listaString = append(listaString, l.Phone)
		listaString = append(listaString, l.Website)
		listaString = append(listaString, l.Company.Name)
		listaString = append(listaString, l.Company.CatchPhrase)
		listaString = append(listaString, l.Company.Bs)
		listaString = append(listaString, l.Uuid)
		listaString = append(listaString, strconv.FormatBool(l.Imported))
		write.Write(listaString)
	}
	return file
}

func main() {
	e := echo.New()

	e.GET("/users", func(c echo.Context) error {
		users := getLista()
		return c.JSON(200, users)
	})
	e.GET("/users/file", func(c echo.Context) error {
		file := stampaFIERAAA(getLista())
		return c.File(file.Name())
	})

	e.Use(middleware.CORS())

	e.Logger.Fatal(e.Start(":8080"))

	// lista := getLista()
	// newMail(lista)
	// stampaFIERAAA(lista)
}
