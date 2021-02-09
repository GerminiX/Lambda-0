package base_common

func StartUp()  {
	initConfig()
	initKeys()
	createDbSession()
	addIndexesToDataBase()
}