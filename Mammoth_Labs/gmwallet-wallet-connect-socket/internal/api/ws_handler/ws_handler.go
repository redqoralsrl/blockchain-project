package ws_handler

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	errorRes "gmwallet-connect-go/internal/api/error"
	"gmwallet-connect-go/internal/api/response"
	"gmwallet-connect-go/internal/config"
	"gmwallet-connect-go/internal/ws"
	"net/http"
	"time"
)

type JwtClaims struct {
	jwt.StandardClaims
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Handler struct {
	hub *ws.Hub
}

func NewHandler(h *ws.Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

func (h *Handler) ActiveRoom(c echo.Context) error {
	// create uuid
	newUUID := uuid.New().String()
	// 10 minutes expire room
	expired := time.Now().Add(5 * time.Minute)

	uniqueString := newUUID + expired.String()
	hasher := sha256.New()
	hasher.Write([]byte(uniqueString))
	hashBytes := hasher.Sum(nil)
	hash := hex.EncodeToString(hashBytes)

	// create room by uuid
	h.hub.Rooms[hash] = &ws.Room{
		Hash:    hash,
		Clients: make(map[string]*ws.Client),
		Expiry:  expired,
	}

	// jwt claims setting
	claims := JwtClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	conf := config.LoadConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(conf.JwtSecretKey))
	if err != nil {
		err = errorRes.NewAppError(http.StatusInternalServerError, response.InternalError, err.Error())
		return err
	}

	result := &ws.CreateRoomRes{
		Hash:   hash,
		Expiry: expired.String(),
		UUID:   newUUID,
		Jwt:    signedToken,
	}

	data := &response.Response[*ws.CreateRoomRes]{
		Data: &result,
	}

	res := response.NewApiResponse[response.Response[*ws.CreateRoomRes]](http.StatusOK, response.Success, nil, *data)

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) JoinRoom(c echo.Context) error {
	var input ws.JoinRoomReq
	if err := c.Bind(&input); err != nil {
		err = errorRes.NewAppError(http.StatusBadRequest, response.InvalidInput, err.Error())
		return err
	}

	// Jwt Checking
	if input.Jwt == "" {
		//errMsg := ws.Message{
		//	Topic: ws.JwtMissing,
		//	Code:  404,
		//}
		//msgBytes, _ := json.Marshal(errMsg)
		//_ = wsConn.WriteMessage(websocket.TextMessage, msgBytes)
		err := errorRes.NewAppError(http.StatusBadRequest, response.InvalidInput, "jwt required")
		return err
	}
	// JWT verify
	token, err := jwt.ParseWithClaims(input.Jwt, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		conf := config.LoadConfig()
		return []byte(conf.JwtSecretKey), nil
	})
	if err != nil || !token.Valid {
		err := errorRes.NewAppError(http.StatusBadRequest, response.InvalidInput, "token rejected")
		return err
	}

	wsConn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		err = errorRes.NewAppError(http.StatusInternalServerError, response.InternalError, err.Error())
		return err
	}
	defer wsConn.Close()

	cl := &ws.Client{
		UUID:    input.UUID,
		Conn:    wsConn,
		Hash:    input.Hash,
		Message: make(chan *ws.Message, 10),
	}
	defer cl.Close()

	m := &ws.Message{
		Hash: input.Hash,
		Code: 200,
		Data: "",
	}
	if input.UUID == "GM_WALLET" {
		m.Topic = ws.GmWalletConnected
	} else {
		m.Topic = ws.DAppConnected
	}

	h.hub.Register <- cl

	h.hub.Unicast <- m

	go cl.WriteMessage()
	for {
		err := cl.ReadMessage(h.hub)
		if err != nil {
			return err // Other errors
		}
	}
}
