package user

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/SKilliu/novels-backend/internal/db/models"
	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Autogenerated struct {
	IP            string  `json:"ip"`
	Hostname      string  `json:"hostname"`
	Type          string  `json:"type"`
	ContinentCode string  `json:"continent_code"`
	ContinentName string  `json:"continent_name"`
	CountryCode   string  `json:"country_code"`
	CountryName   string  `json:"country_name"`
	RegionCode    string  `json:"region_code"`
	RegionName    string  `json:"region_name"`
	City          string  `json:"city"`
	Zip           string  `json:"zip"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Location      struct {
		GeonameID int    `json:"geoname_id"`
		Capital   string `json:"capital"`
		Languages []struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Native string `json:"native"`
		} `json:"languages"`
		CountryFlag             string `json:"country_flag"`
		CountryFlagEmoji        string `json:"country_flag_emoji"`
		CountryFlagEmojiUnicode string `json:"country_flag_emoji_unicode"`
		CallingCode             string `json:"calling_code"`
		IsEu                    bool   `json:"is_eu"`
	} `json:"location"`
	TimeZone struct {
		ID               string `json:"id"`
		CurrentTime      string `json:"current_time"`
		GmtOffset        int    `json:"gmt_offset"`
		Code             string `json:"code"`
		IsDaylightSaving bool   `json:"is_daylight_saving"`
	} `json:"time_zone"`
	Currency struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Plural       string `json:"plural"`
		Symbol       string `json:"symbol"`
		SymbolNative string `json:"symbol_native"`
	} `json:"currency"`
	Connection struct {
		Asn int    `json:"asn"`
		Isp string `json:"isp"`
	} `json:"connection"`
	Security struct {
		IsProxy     bool        `json:"is_proxy"`
		ProxyType   interface{} `json:"proxy_type"`
		IsCrawler   bool        `json:"is_crawler"`
		CrawlerName interface{} `json:"crawler_name"`
		CrawlerType interface{} `json:"crawler_type"`
		IsTor       bool        `json:"is_tor"`
		ThreatLevel string      `json:"threat_level"`
		ThreatTypes interface{} `json:"threat_types"`
	} `json:"security"`
}

// @Summary Guest sign in
// @Tags Authentication
// @Consume application/json
// @Param JSON body dto.GuestSignInRequest true "Body for guest sign in"
// @Description Sign in like a guest (without progress saving)
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/guest-registration [post]
func (h *Handler) GuestSignIn(c echo.Context) error {
	var req dto.GuestSignInRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse signup request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}

	if req.DeviceID == "" {
		h.log.WithError(err).Error("empty device ID")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	user, err := h.usersDB.GetByDeviceID(req.DeviceID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			uid := uuid.New().String()

			randomName, err := utils.GenerateName("guest")
			if err != nil {
				h.log.WithError(err).Error("failed to generate random name")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			err = h.usersDB.Insert(models.User{
				ID:       uid,
				Username: randomName,
				// Email:        fmt.Sprintf("%s has no email", randomName),
				DeviceID:     req.DeviceID,
				DateOfBirth:  time.Now().Unix(),
				IsRegistered: false,
			})
			if err != nil {
				h.log.WithError(err).Error("failed to create new user")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			token, err := utils.GenerateJWT(uid, "user", h.authKey)
			if err != nil {
				h.log.WithError(err).Error("failed to generate jwt token")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			resp := dto.AuthResponse{
				ID:       uid,
				Username: randomName,
				// Email:    fmt.Sprintf("%s has no email", randomName),
				Token: token,
			}

			return c.JSON(http.StatusOK, resp)
		default:
			h.log.WithError(err).Error("failed to get user form db by device ID")
			return c.JSON(http.StatusBadRequest, errs.InternalServerErr)
		}
	}

	token, err := utils.GenerateJWT(user.ID, "user", h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to generate jwt token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	resp := dto.AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		// Email:    user.Email,
		Token:       token,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Membership:  user.Membership,
		AvatarData:  user.AvatarData,
		Rate:        user.Rate,
	}

	return c.JSON(http.StatusOK, resp)
}
