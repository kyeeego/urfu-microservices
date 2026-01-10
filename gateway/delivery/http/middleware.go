package http

import "github.com/gin-gonic/gin"

func (h *Handler) JwtAuthorize(c *gin.Context) {
	status, _, err := h.http.Get(h.cfg.UsersUrl+"/authorize", map[string]string{
		"Authorization": c.GetHeader("Authorization"),
	})

	if err != nil || status != 200 {
		c.AbortWithStatusJSON(403, map[string]string{
			"error": "invalid JWT",
		})
	}
}
