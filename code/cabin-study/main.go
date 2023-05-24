import "github.com/casbin/casbin/v2"

func main() {

	e, err := casbin.NewEnforcer("deploy/model.conf", "deploy/policy.csv")

}