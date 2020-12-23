package pkg1

import (
	"fmt"
	"gocms/pkg/config"
	"gocms/pkg/logger"
)

func init() {

}

func Echo() {
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	ok, err := config.Casbin.Enforce(sub, obj, act)

	if err != nil {
		logger.PanicError(err, "validate permission", true)
	}

	if ok == true {
		fmt.Println("permission success")
	} else {
		fmt.Println("permission error")
	}
}
