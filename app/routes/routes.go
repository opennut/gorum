// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tPublicApp struct {}
var PublicApp tPublicApp


func (_ tPublicApp) AddUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("PublicApp.AddUser", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tJobs struct {}
var Jobs tJobs


func (_ tJobs) Status(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Jobs.Status", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tApp struct {}
var App tApp



type tAdminApp struct {}
var AdminApp tAdminApp



type tHome struct {}
var Home tHome


func (_ tHome) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.Index", args).URL
}

func (_ tHome) Detail(
		ID int,
		name string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "ID", ID)
	revel.Unbind(args, "name", name)
	return revel.MainRouter.Reverse("Home.Detail", args).URL
}


type tAccounts struct {}
var Accounts tAccounts


func (_ tAccounts) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Accounts.Login", args).URL
}

func (_ tAccounts) LoginProccess(
		email string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Accounts.LoginProccess", args).URL
}

func (_ tAccounts) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Accounts.Register", args).URL
}

func (_ tAccounts) RegisterPost(
		email string,
		username string,
		name string,
		password string,
		password2 string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "password2", password2)
	return revel.MainRouter.Reverse("Accounts.RegisterPost", args).URL
}

func (_ tAccounts) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Accounts.Logout", args).URL
}


type tDashboard struct {}
var Dashboard tDashboard


func (_ tDashboard) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Dashboard.Index", args).URL
}


