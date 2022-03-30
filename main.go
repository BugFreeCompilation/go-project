package main

import (
	database "BugFreeCompilation/go-project/database"
	router "BugFreeCompilation/go-project/router"
	service "BugFreeCompilation/go-project/service"
)

var (
	mydb       database.Database = database.InitSQLite3()
	myservice  service.Service   = service.InitService(mydb)
	httprouter router.Router     = router.InitMuxRouter()
)

func main() {
	var port string = ":8000"

	httprouter.GET("/data", myservice.GetAll)
	httprouter.POST("/data", myservice.Add)
	httprouter.DELETE("/data", myservice.Delete)

	httprouter.SERVE(port)
}
