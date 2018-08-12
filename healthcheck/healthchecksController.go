package healthcheck

import (
	"github.com/kataras/iris"
	"github.com/pedromsmoreira/go-simple-rest-api/database"
)

type HealthCheckController struct {
	RedisDb database.Repository
}

func NewHealthCheckController(redis database.Repository) *HealthCheckController {
	return &HealthCheckController{
		RedisDb: redis,
	}
}

func (hc *HealthCheckController) Shallow(ctx iris.Context) {
	hcs := make([]Shallow, 0)
	pong, err := hc.RedisDb.Ping()

	if err != nil {
		hcs = append(hcs, NewShallow("Redis", "Server is down.", false))
		ctx.JSON(ShallowHealthChecks{HealthChecks: hcs})
	}

	hcs = append(hcs, NewShallow("Redis", pong, true))

	ctx.JSON(ShallowHealthChecks{HealthChecks: hcs})
}

func (hc *HealthCheckController) Deep(ctx iris.Context) {}
