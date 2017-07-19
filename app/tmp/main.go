// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "github.com/opennut/gorum/app"
	controllers "github.com/opennut/gorum/app/controllers"
	tests "github.com/opennut/gorum/tests"
	controllers1 "github.com/revel/modules/jobs/app/controllers"
	_ "github.com/revel/modules/jobs/app/jobs"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers2 "github.com/revel/modules/testrunner/app/controllers"
	_ "github.com/wariosolis/modules/gorm/app"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.PublicApp)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "AddUser",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Jobs)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Status",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					28: []string{ 
						"entries",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					76: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					129: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			
		})
	
	revel.RegisterController((*controllers.AdminApp)(nil),
		[]*revel.MethodType{
			
		})
	
	revel.RegisterController((*controllers.Home)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					17: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Detail",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "ID", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					29: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Accounts)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					17: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "LoginProccess",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Register",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					51: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "RegisterPost",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password2", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Dashboard)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					14: []string{ 
					},
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"github.com/opennut/gorum/app/controllers.Accounts.LoginProccess": { 
			22: "email",
			23: "email",
			24: "password",
			30: "user.ID",
			32: "err",
		},
		"github.com/opennut/gorum/app/controllers.Accounts.RegisterPost": { 
			56: "email",
			57: "email",
			58: "username",
			59: "name",
			60: "password",
			61: "password2",
			62: "password",
		},
		"github.com/opennut/gorum/app/controllers.AdminApp.checkUser": { 
			70: "user",
			74: "user",
		},
		"github.com/opennut/gorum/app/controllers.App.checkUser": { 
			56: "user",
		},
		"github.com/opennut/gorum/app/controllers.Home.Detail": { 
			22: "name",
			23: "name",
			24: "name",
			25: "name",
			26: "name",
		},
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
