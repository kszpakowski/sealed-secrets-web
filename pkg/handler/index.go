package handler

import (
	"net/http"

	"github.com/bakito/sealed-secrets-web/pkg/config"
	"github.com/bakito/sealed-secrets-web/pkg/marshal"
	"github.com/bakito/sealed-secrets-web/pkg/seal"
	"github.com/bakito/sealed-secrets-web/pkg/version"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	sealer             seal.Sealer
	indexHTML          string
	disableLoadSecrets bool
	marshaller         marshal.Marshaller
	filter             *config.FieldFilter
}

func New(indexHTML string, sealer seal.Sealer, cfg *config.Config) *Handler {
	return &Handler{
		marshaller: cfg.Marshaller,
		sealer:     sealer,
		indexHTML:  indexHTML,
		filter:     cfg.FieldFilter,
	}
}

func (h *Handler) Index(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")
	c.String(http.StatusOK, h.indexHTML)
}

func (h *Handler) RedirectToIndex(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "/")
	ctx.Abort()
}

func (h *Handler) Version(c *gin.Context) {
	c.JSONP(http.StatusOK, map[string]string{"version": version.Version, "build": version.Build})
}
