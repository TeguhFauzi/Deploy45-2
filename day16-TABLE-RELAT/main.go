package main

import (
	"context"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"main.go/connection"
	"main.go/middleware"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

// Initial file
type Template struct {
	templates *template.Template
}

// Rendering template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Struct Interface Data Dummy
type Blog struct {
	ID          int
	Title       string
	Description string
	Author      string
	PostDate    string
	Techno      []string
	Image       string
	Icon        string
	// Durat    string
	Start     string
	End       string
	Post_at   time.Time
	StartDate time.Time
	EndDate   time.Time
	Duration  string
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

var dummyD = []Blog{
	{
		Title:       "DW MOBILE APP-2021",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Deserunt alias consequatur.",
		Author:      "Michael Ipin",
		PostDate:    "12 Nov 2021 13:04",
		Image:       "proc",
		Icon: `<img src="/public/images/javascript.png"/>
					<img src="/public/images/react.png"/>`,
		Duration: "3 Week(s)",
	},
	{

		Title:       "DW MOBILE APP-2020",
		Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit.Fugiat placeat aut earum!",
		Author:      "Michael Ipin",
		PostDate:    "08 Feb 2020 15:04",
		Image:       "netflix",
		Icon: `<img src="/public/images/react.png"/>
					<img src="/public/images/node.png"/>
					<img src="/public/images/javascript.png"/>
					<img src="/public/images/angular.png"/>`,
		Duration: "2 Month(s)",
	},
}

func main() {
	connection.DatabaseConnection()

	e := echo.New()

	//FOlder Static
	e.Static("/public", "public")
	e.Static("/upload", "upload")

	//INITIAL SESSION USING ECHO
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session")))) //menyimpan sessi pada cookie(memori sementara)

	//Parsing menggunakan parseglob
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t

	// ROUTING
	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/user", user)
	e.GET("/contact", contact)
	e.GET("/blog-Detail/:id", blogDetail)
	e.GET("/delete-blog/:id", deleteblog)
	e.GET("/form-project", formProject)
	e.POST("/add-blog", middleware.UploadFile(AddBlog))
	e.GET("/delete-project/:id", deleteProject)
	e.GET("/blog-detail/:id", blogDet)
	e.GET("/edit-project/:id", editProject)
	e.POST("/update-project/:id", uProject)
	e.GET("/form-register", formregister)
	e.POST("/register", addregister)
	e.GET("/form-login", formlogin)
	e.POST("/login", login)
	e.GET("/logout", logout)
	fmt.Println("Server running on Port : 5000")
	e.Logger.Fatal(e.Start("Localhost:5000"))

}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func home(c echo.Context) error {
	//QUERY SELECT ==> Menjalankan perintah SQL
	data, _ := connection.Conn.Query(context.Background(), "SELECT tb_blog.id, title, description, image, post_at, postdate, techno, start_date, end_date, duration, tb_user.name as author_name FROM public.tb_blog LEFT JOIN tb_user ON tb_blog.author = tb_user.id ORDER BY id DESC;")

	var result []Blog //Ditampung dalam array
	for data.Next() { //method dari library pgx
		var each = Blog{}
		//SCAN DATA
		err := data.Scan(&each.ID, &each.Title, &each.Description, &each.Image, &each.Post_at, &each.PostDate, &each.Techno, &each.StartDate, &each.EndDate, &each.Duration, &each.Author) //membaca hasil query
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
		}
		fmt.Println(result)
		result = append(result, each)
	}

	flash := map[string]interface{}{
		"Dummy": dummyD,
		"Blogs": result,
	}
	return c.Render(http.StatusOK, "index.html", flash)
}


func user(c echo.Context) error {
	sess, _ := session.Get("session", c)

	// asumsikan userID disimpan dalam session dengan key "userID" dan tipe data int
	authorId := sess.Values["id"]

	// Do something with userLogin, userName, and userData variables
	if sess.Values["isLogin"] != true {
		userData.IsLogin = false //
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)

	}

	//QUERY SELECT
	data, err := connection.Conn.Query(context.Background(), "SELECT tb_blog.id, title, description, image, post_at, postdate, techno, start_date, end_date, duration, tb_user.name as author_name FROM public.tb_blog LEFT JOIN tb_user ON tb_blog.author = tb_user.id WHERE tb_blog.author = $1 ORDER BY id DESC", authorId)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
	}

	var result []Blog
	for data.Next() {
		var each = Blog{}
		//SCAN DATA
		err := data.Scan(&each.ID, &each.Title, &each.Description, &each.Image, &each.Post_at, &each.PostDate, &each.Techno, &each.StartDate, &each.EndDate, &each.Duration, &each.Author)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
		}

		result = append(result, each)
	}
	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["isLogin"],
		"FlashMessage": sess.Values["message"],
		"FlashName":    sess.Values["name"],
		"Dummy":        dummyD,
		"Blogs":        result,
	}
	delete(sess.Values, "status")
	delete(sess.Values, "message")
	sess.Save(c.Request(), c.Response())
	return c.Render(http.StatusOK, "user.html", flash)
}

func contact(c echo.Context) error {
	sess, _ := session.Get("session", c)
	flash := map[string]interface{}{
		"FlashStatus": sess.Values["isLogin"],
		"FlashName":   sess.Values["name"],
	}
	return c.Render(http.StatusOK, "contactme.html", flash)
}

/*
================================================================
BLOG-DETAIL-DUMMY-DATA
================================================================
*/
func blogDetail(c echo.Context) error {
	sess, _ := session.Get("session", c)
	id, _ := strconv.Atoi(c.Param("id")) // url params | dikonversikan dari string menjadi int/integer

	var Dummy = Blog{}

	for i, data := range dummyD {
		if id == i {
			Dummy = Blog{

				// ID:id,
				Title:       data.Title,
				Description: data.Description,
				Author:      data.Author,
				PostDate:    data.PostDate,
				Icon:        data.Icon,
				Image:       data.Image,
			}

		}
	}
	dataDetail := map[string]interface{}{
		"Blog":         Dummy,
		"FlashStatus":  sess.Values["isLogin"],
		"FlashMessage": sess.Values["message"],
		"FlashName":    sess.Values["name"],
	}
	return c.Render(http.StatusOK, "blog-d.html", dataDetail)
}

func formProject(c echo.Context) error {
	sess, _ := session.Get("session", c)
	flash := map[string]interface{}{
		"FlashName":   sess.Values["name"],
		"FlashStatus": sess.Values["isLogin"],
	}
	return c.Render(http.StatusOK, "myproject.html", flash)
}

/*
================================================================
ADD BLOG
================================================================
*/
func AddBlog(c echo.Context) error {
	sess, _ := session.Get("session", c)

	authorId := sess.Values["id"]
	title := c.FormValue("inTitle")
	description := c.FormValue("inDesc")
	image := c.Get("dataFile").(string)
	techno := c.Request().Form["techno"]
	//Parse Start
	startDate := c.FormValue("startdate")
	StartDate, _ := time.Parse("2006-01-02", startDate)
	//Parse End
	endDate := c.FormValue("enddate")
	EndDate, _ := time.Parse("2006-01-02", endDate)
	//Time Parse and Format
	loc, _ := time.LoadLocation("Asia/Jakarta")     //Deklarasi variabel lokasi untuk waktu
	Post_at := time.Now().In(loc)                   //Manipulasi waktu untuk dapat waktu sekarang sesuai dengan lokasi
	PostDate := Post_at.Format("02 Jan 2006 15:04") //Format waktu

	//Get Duration
	var Duration string
	durations := EndDate.Sub(StartDate)

	if durations.Hours()/24 < 7 {
		if durations.Hours()/24 == 1 {
			Duration = strconv.FormatFloat(durations.Hours()/24, 'f', 0, 64) + " Day"
		} else {
			Duration = strconv.FormatFloat(durations.Hours()/24, 'f', 0, 64) + " Day(s)"
		}
	} else if durations.Hours()/24/7 < 4 {
		if durations.Hours()/24/7 == 1 {
			Duration = strconv.FormatFloat(durations.Hours()/24/7, 'f', 0, 64) + " Week"
		} else {
			Duration = strconv.FormatFloat(durations.Hours()/24/7, 'f', 0, 64) + " Week(s)"
		}
	} else if durations.Hours()/24/30 < 12 {
		if durations.Hours()/24/30 == 1 {
			Duration = strconv.FormatFloat(durations.Hours()/24/30, 'f', 0, 64) + " Month"
		} else {
			Duration = strconv.FormatFloat(durations.Hours()/24/30, 'f', 0, 64) + " Month(s)"
		}
	} else {
		if durations.Hours()/24/30/12 == 1 {
			Duration = strconv.FormatFloat(durations.Hours()/24/30/12, 'f', 0, 64) + " Year"
		} else {
			Duration = strconv.FormatFloat(durations.Hours()/24/30/12, 'f', 0, 64) + " Year(s)"
		}
	}

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	// _ melakukan query initial _ karena returnnya tidak digunakan
	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog (title, description, image, post_at, postdate, techno, start_date, end_date, duration, author) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", title, description, image, Post_at, PostDate, techno, StartDate, EndDate, Duration, authorId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
	}

	// RETURN
	return c.Redirect(http.StatusMovedPermanently, "/user")
}

func deleteblog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	dummyD = append(dummyD[:id], dummyD[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/user")
}

/*
================================================================
DELETE-PROJECT
================================================================
*/
func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_blog WHERE id=$1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/user")
}

/*
================================================================
PROJECT-DETAILS
================================================================
*/
func blogDet(c echo.Context) error {
	sess, _ := session.Get("session", c)
	id, _ := strconv.Atoi(c.Param("id")) // url params | dikonversikan dari string menjadi int/integer

	var Blogdet = Blog{}
	err := connection.Conn.QueryRow(context.Background(), "SELECT tb_blog.id, title, description, image, post_at, postdate, techno, start_date, end_date, duration, tb_user.name as author FROM public.tb_blog LEFT JOIN tb_user ON tb_blog.author = tb_user.id WHERE tb_blog.id=$1;", id).Scan(&Blogdet.ID, &Blogdet.Title, &Blogdet.Description, &Blogdet.Image, &Blogdet.Post_at, &Blogdet.PostDate, &Blogdet.Techno, &Blogdet.StartDate, &Blogdet.EndDate, &Blogdet.Duration, &Blogdet.Author)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
	}

	dataDetail := map[string]interface{}{ //data yang akan digunakan/dikirimkan ke html menggunakan map interface
		"Blog":         Blogdet,
		"FlashStatus":  sess.Values["isLogin"],
		"FlashMessage": sess.Values["message"],
		"FlashName":    sess.Values["name"],
	}
	return c.Render(http.StatusOK, "blog-e.html", dataDetail)
}

/*
================================================================
Edit Project
================================================================
*/
func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var Edit = Blog{}
	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, description, techno, start_date, end_date FROM public.tb_blog WHERE id = $1", id).Scan(&Edit.ID, &Edit.Title, &Edit.Description, &Edit.Techno, &Edit.StartDate, &Edit.EndDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
	}

	var title string
	var desc string
	var techno []string
	// REVERSE
	title = Edit.Title
	Edit.Title = title

	// REVERSE
	desc = Edit.Description
	Edit.Description = desc

	// REVERSE
	techno = Edit.Techno
	Edit.Techno = techno

	// REVERSE
	startdate := Edit.StartDate.Format("2006-01-02")
	enddate := Edit.EndDate.Format("2006-01-02")
	Edit.Start = startdate
	Edit.End = enddate
	form := map[string]interface{}{
		"Blog": Edit,
	}

	return c.Render(http.StatusMovedPermanently, "edit.html", form)
}

/*
================================================================
ADD PROJECT
================================================================
*/
func uProject(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("id"))
	Title := c.FormValue("inTitle")
	Description := c.FormValue("inDesc")
	Techno := c.Request().Form["techno"]

	//===============TIME-UPDATE================
	//Parse Start
	startDate := c.FormValue("startdate")
	StartDate, _ := time.Parse("2006-01-02", startDate)
	//Parse End
	endDate := c.FormValue("enddate")
	EndDate, _ := time.Parse("2006-01-02", endDate)
	var Duration string
	durations := EndDate.Sub(StartDate)
	//Time Parse and Format
	loc, _ := time.LoadLocation("Asia/Jakarta")     //Deklarasi variabel lokasi untuk waktu
	Post_at := time.Now().In(loc)                   //Manipulasi waktu untuk dapat waktu sekarang sesuai dengan lokasi
	PostDate := Post_at.Format("02 Jan 2006 15:04") //Format waktu
	if durations.Hours()/24 < 7 {
		if durations.Hours()/24 == 1 {
			Duration = strconv.FormatFloat(durations.Hours()/24, 'f', 0, 64) + " Day"
		} else {
			Duration = strconv.FormatFloat(durations.Hours()/24, 'f', 0, 64) + " Day(s)"
		}
	} else if durations.Hours()/24/7 < 4 {
		if durations.Hours()/24/7 == 1 {
			Duration = strconv.FormatFloat(durations.Hours()/24/7, 'f', 0, 64) + " Week"
		} else {
			Duration = strconv.FormatFloat(durations.Hours()/24/7, 'f', 0, 64) + " Week(s)"
		}
	} else if durations.Hours()/24/30 < 12 {
		if durations.Hours()/24/30 == 1 {
			Duration = strconv.FormatFloat(durations.Hours()/24/30, 'f', 0, 64) + " Month"
		} else {
			Duration = strconv.FormatFloat(durations.Hours()/24/30, 'f', 0, 64) + " Month(s)"
		}
	} else {
		if durations.Hours()/24/30/12 == 1 {
			Duration = strconv.FormatFloat(durations.Hours()/24/30/12, 'f', 0, 64) + " Year"
		} else {
			Duration = strconv.FormatFloat(durations.Hours()/24/30/12, 'f', 0, 64) + " Year(s)"
		}
	}

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_blog SET title=$1, description=$2, techno=$3, start_date=$4, end_date=$5, duration=$6, postdate=$7 WHERE id=$8", Title, Description, Techno, StartDate, EndDate, Duration, PostDate, ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
	}
	return c.Redirect(http.StatusMovedPermanently, "/user")
}


func formregister(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}


//REGISTER FUNC
func addregister(c echo.Context) error {
	//MENGINILIASISASI VALUE APA SAJA YANG DIAMBIL DARI FORMVALUE MENGGUNAKAN PARSEFORM
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	//GET PASSWORD FROM FORM VALUE TAMPUNG KEDALAM VARIABEL KEMUDIAN DI DIGENERATE DAN DITAMPUNG DIDALAM []KARNA BERUPA STRING KEMUDIAN DIKONVERT KE BYTE SEBELUM DIHASING MENGGUNAKAN,10 RANDOM CHAR(PROSES SALT)
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (name,email,password) VALUES ($1, $2, $3)", name, email, passwordHash)
	if err != nil {
		redirectMessage(c, "Resgiter failed, please try again", false, "/form-register")
	}

	return redirectMessage(c, "Resgiter succes", true, "/form-login")
}

func formlogin(c echo.Context) error {
	sess, _ := session.Get("session", c)

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	return c.Render(http.StatusOK, "login.html", nil)
}

//LOGIN FUNCTION
func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return redirectMessage(c, "Email denied!", false, "/form-login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectMessage(c, "Password denied!", false, "/form-login")
	}
	sess, _ := session.Get("session", c) //Mengambil sesi yang sudah ada
	sess.Options.MaxAge = 10800          //3Hours session in seconds
	sess.Values["message"] = "Login Succes !"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response()) //Menyimpan sessi yang sudah diatur pada cookie

	return c.Redirect(http.StatusMovedPermanently, "/user")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1 //set sesi agar keluar tanpa dihapus
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func redirectMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, path)
}
