package fiberotel

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
)

func FromCtx(ctx *fiber.Ctx) context.Context {
	otelCtx := ctx.Locals(LocalsCtxKey).(context.Context)
	return otelCtx
}

func SpanFromCtx(ctx *fiber.Ctx) trace.Span {
	return trace.SpanFromContext(FromCtx(ctx))
}
