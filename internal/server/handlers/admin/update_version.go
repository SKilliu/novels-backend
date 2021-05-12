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

// @Summary Update client version
// @Security bearerAuth
// @Tags admin
// @Consume application/json
// @Param JSON body dto.UpdateVersionRequest true "body for a new version saving"
// @Description Update client version info for requested platform
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/admin/version [post]
func (h *Handler) UpdateVersion(c echo.Context) error {
	var req dto.UpdateVersionRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse update version request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}

	ymlFile, err := os.Open("./static/versions.yaml")
	if err != nil {
		h.log.WithError(err).Error("failed to open yaml file")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	defer ymlFile.Close()

	var versions = make(map[string]string)

	byteValue, err := ioutil.ReadAll(ymlFile)
	if err != nil {
		h.log.WithError(err).Error("failed to read bytes from yaml file")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	err = yaml.Unmarshal(byteValue, &versions)
	if err != nil {
		h.log.WithError(err).Error("failed to unmarshal yaml file")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if req.Platform == "android" {
		versions["android"] = req.Version
	}

	if req.Platform == "ios" {
		versions["ios"] = req.Version
	}

	bytes, err := yaml.Marshal(&versions)
	if err != nil {
		h.log.WithError(err).Error("failed to marshal versions to bytes")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	err = ioutil.WriteFile("./static/versions.yaml", bytes, 0644)
	if err != nil {
		h.log.WithError(err).Error("failed to write bytes into yaml file")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
