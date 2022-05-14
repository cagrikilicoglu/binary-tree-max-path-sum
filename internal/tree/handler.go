package tree

import (
	"binary-tree-max-path-sum/internal/api"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

func NewTreeHandler(r *gin.RouterGroup) {
	r.POST("/max-path-sum", calculateMaxPathSum)
}

// calculateMaxPathSum calculates the maximum path sum of a given binary tree
func calculateMaxPathSum(c *gin.Context) {
	requestBody := &api.Request{}

	if err := c.Bind(&requestBody); err != nil {

		respondWithJSON(c, http.StatusBadRequest, ("Request body cannot be binded"))
		return
	}

	if err := requestBody.Validate(strfmt.NewFormats()); err != nil {

		respondWithJSON(c, http.StatusForbidden, errors.New("Request body is not valid"))
		return
	}

	root := CreateBinaryTree(*requestBody.Tree)
	result := MaxSum(root)
	respondWithJSON(c, http.StatusOK, result)

}

// respondWithJSON creates a structured response for http request
func respondWithJSON(c *gin.Context, code int, payload interface{}) {
	c.Header("code", strconv.Itoa(code))
	c.JSON(code, payload)
}
