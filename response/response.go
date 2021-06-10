package response

import (
	"github.com/labstack/echo/v4"
	"github.com/ulule/paging"
	"net/http"
)

type Meta struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

// APIError
type APIError struct {
	Code    int    `json:"code,omitempty"`
	Type    string `json:"type,omitempty"`
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type Paginator struct {
	Total  int64 `json:"total"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Link   Link  `json:"links"`
}

type Link struct {
	NextPageUrl string `json:"next_page_url"`
	PrevPageUrl string `json:"prev_page_url"`
}

type MetaPaginator struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Page    Paginator   `json:"page"`
}

type Single struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

type Paging struct {
	MetaPaginator MetaPaginator `json:"meta"`
	Data          interface{}   `json:"data,omitempty"`
}

func SingleData(c echo.Context, message string, data interface{}, error interface{}) error {
	return c.JSON(http.StatusOK, Single{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: message,
			Error:   error,
		},
		Data: data,
	})
}

func ListData(c echo.Context, message string, data interface{}, error interface{}) error {
	return c.JSON(http.StatusOK, Single{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: message,
			Error:   error,
		},
		Data: data,
	})
}

func DataWithoutMeta(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func NotFound(c echo.Context, message string, data interface{}, error interface{}) error {
	return c.JSON(http.StatusNotFound, Single{
		Meta: Meta{
			Code:    http.StatusNotFound,
			Message: message,
			Error:   error,
		},
		Data: data,
	})
}

func BadRequest(c echo.Context, message string, data interface{}, error interface{}) error {
	return c.JSON(http.StatusBadRequest, Single{
		Meta: Meta{
			Code:    http.StatusBadRequest,
			Message: message,
			Error:   error,
		},
		Data: data,
	})
}

func ValidationError(c echo.Context, message string, data interface{}, error interface{}) error {
	return c.JSON(http.StatusUnprocessableEntity, Single{
		Meta: Meta{
			Code:    http.StatusUnprocessableEntity,
			Message: message,
			Error:   error,
		},
		Data: data,
	})
}

func InternalServerError(c echo.Context, message string, data interface{}, error interface{}) error {
	return c.JSON(http.StatusInternalServerError, Single{
		Meta: Meta{
			Code:    http.StatusInternalServerError,
			Message: message,
			Error:   error,
		},
		Data: data,
	})
}

func Paginate(c echo.Context, message string, paginator *paging.OffsetPaginator, data interface{}, error interface{}) error {
	return c.JSON(http.StatusOK, Paging{
		MetaPaginator: MetaPaginator{
			Code:    http.StatusOK,
			Message: message,
			Error:   error,
			Page: Paginator{
				Total:  paginator.Count,
				Limit:  paginator.Limit,
				Offset: paginator.Offset,
				Link: Link{
					NextPageUrl: paginator.NextURI.String,
					PrevPageUrl: paginator.PreviousURI.String,
				},
			},
		},
		Data: data,
	})
}