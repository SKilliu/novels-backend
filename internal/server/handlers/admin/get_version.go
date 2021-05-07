package admin

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
)

// @Summary Get client version
// @Security bearerAuth
// @Tags admin
// @Consume application/json
// @Param platform query string true "requested platform"
// @Description Returns a client version for Andrion or iOS by platform name.
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.GetVersionResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/admin/version [get]
func (h *Handler) GetVersion(c echo.Context) error {
	var resp dto.GetVersionResponse
	platform := c.QueryParam("platform")

	if platform == "" {
		h.log.Error("empty platform param")
		return c.JSON(http.StatusBadRequest, errs.EmptyQueryParamErr)
	}

	ymlFile, err := os.Open("./static/versions.yaml")
	if err != nil {
		panic(err)
	}

	defer ymlFile.Close()

	var versions = make(map[string]string)

	byteValue, err := ioutil.ReadAll(ymlFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(byteValue, &versions)
	if err != nil {
		panic(err)
	}

	switch platform {
	case "andriod":
		resp.Version = versions["android"]
	case "ios":
		resp.Version = versions["ios"]
	default:
		h.log.Error("incorrect platform name in request")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.JSON(http.StatusOK, resp)
}
