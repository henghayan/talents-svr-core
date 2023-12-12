package context

var ModuleCtx *ModuleContext

type ModuleContext struct {
	DB   interface{}
	Data interface{}
}

func (ctx *ModuleContext) Init(db interface{}, data interface{}) {
	ctx.DB = db
	ctx.Data = data
}
