package admin_controller

import (
	"github.com/a20070322/go_fast_admin/utils/response"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
)

type AdminAutoCurd struct {
}

type GetDefaultConfigRep struct {
	WorkPath string `json:"work_path"`
	ModelName string `json:"model_name"`
}

func (c AdminAutoCurd) GetDefaultConfig(ctx *gin.Context) {
	var rep GetDefaultConfigRep
	var err error
	rep.WorkPath, err = os.Getwd()
	if err != nil {
		response.Fail(ctx, http.StatusLoopDetected, err.Error(), nil)
	}
	f, err := ioutil.ReadFile(path.Join(rep.WorkPath,"go.mod"))
	re := regexp.MustCompile(`^module (.*)\n`)
	params := re.FindStringSubmatch(string(f))
	rep.ModelName = params[1]
	response.Success(ctx, "ok", rep)
}
