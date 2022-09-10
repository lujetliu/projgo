package v1

import (
	"blog/global"
	"blog/pkg/app"
	"blog/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

/*

	Swagger 基于标准的 OpenAPI 规范进行设计的, 只要照着这套规范去编写你的
	注解或通过扫描代码去生成注解, 就能生成统一标准的接口文档和一系列
	Swagger 工具

	从功能使用上来讲, OpenAPI 规范能够帮助我们描述一个 API 的基本信息, 比如：
	- 有关该 API 的描述
	- 可用路径（/资源）
	- 在每个路径上的可用操作（获取/提交…）
	- 每个操作的输入/输出格式

	Swagger 相关的工具集会根据 OpenAPI 规范去生成各式各类的与接口相关联的内容,
	常见的流程是编写注解 =》调用生成库-》生成标准描述文件 =》生成/导入到对应的
	Swagger 工具;

	安装Swagger
	$ go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
	$ go get -u github.com/swaggo/gin-swagger@v1.2.0
	$ go get -u github.com/swaggo/files
	$ go get -u github.com/alecthomas/template


	针对项目里的 API 接口进行注解的编写, 以下是注解规范:
	@Summary	摘要
	@Produce	API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等
	@Param	参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释
	@Success	响应成功，从左到右分别为：状态码、参数类型、数据类型、注释
	@Failure	响应失败，从左到右分别为：状态码、参数类型、数据类型、注释
	@Router	路由，从左到右分别为：路由地址，HTTP 方法

	在项目根目录下:
	$ swag init

	2022/09/01 22:20:29 Generate swagger docs....
	2022/09/01 22:20:29 Generate general API Info, search dir:./
	2022/09/01 22:20:29 Generating errcode.Error
	2022/09/01 22:20:29 Generating model.Tag
	2022/09/01 22:20:29 create docs.go at  docs/docs.go
	2022/09/01 22:20:29 create swagger.json at  docs/swagger.json
	2022/09/01 22:20:29 create swagger.yaml at  docs/swagger.yaml

	在执行命令完毕后会在 docs 文件夹生成 docs.go、swagger.json、swagger.yaml
	三个文件;
*/

func (t Tag) Get(c *gin.Context) {}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state, default=1" binding:"oneof=0 1"`
	}{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}
