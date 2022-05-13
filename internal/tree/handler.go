package tree

import (
	"binary-tree-max-path-sum/internal/api"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

type TreeHandler struct {
}

func NewTreeHandler(r *gin.RouterGroup) {
	h := &TreeHandler{}
	r.POST("/max-path-sum", h.calculateMaxPathSum)
}

func (t *TreeHandler) calculateMaxPathSum(c *gin.Context) {
	requestBody := &api.Request{}

	if err := c.Bind(&requestBody); err != nil {
		c.Header("code", strconv.Itoa(http.StatusBadRequest)) //TODO
		c.JSON(http.StatusBadRequest, err)
	}

	if err := requestBody.Validate(strfmt.NewFormats()); err != nil {
		c.Header("code", strconv.Itoa(http.StatusForbidden)) //TODO
		c.JSON(http.StatusBadRequest, err)
	}
	log.Println(*requestBody.Tree)
	root := ResponseToBinaryTree(*requestBody.Tree)
	log.Println(root)
	result := MaxSum(root)

	c.Header("code", strconv.Itoa(http.StatusOK))
	c.JSON(http.StatusOK, result)

}
