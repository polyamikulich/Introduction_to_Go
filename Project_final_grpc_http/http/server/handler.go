package server

import (
	dto2 "Go_Project/accounts/dto"
	"Go_Project/accounts/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto2.CreateAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	response := dto2.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()

	if _, ok := h.accounts[name]; !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	h.guard.RUnlock()

	delete(h.accounts, name)

	return c.NoContent(http.StatusOK)
}

func (h *Handler) PatchAccount(c echo.Context) error {
	var request dto2.PatchAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}
	if request.Amount == 0 {
		return c.String(http.StatusBadRequest, "empty amount")
	}

	h.guard.Lock()

	account, ok := h.accounts[request.Name]

	h.guard.Unlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	account.Amount += request.Amount

	return c.NoContent(http.StatusOK)
}

func (h *Handler) ChangeAccount(c echo.Context) error {
	var request dto2.ChangeAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}
	if len(request.NewName) == 0 {
		return c.String(http.StatusBadRequest, "empty new name")
	}

	h.guard.Lock()

	account, ok := h.accounts[request.Name]

	h.guard.Unlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	if _, ok := h.accounts[request.NewName]; ok {
		return c.String(http.StatusConflict, "account with new name already exists")
	}

	account.Name = request.NewName
	delete(h.accounts, request.Name)
	h.accounts[request.NewName] = account

	return c.NoContent(http.StatusOK)
}
