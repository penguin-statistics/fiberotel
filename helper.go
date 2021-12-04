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

func StartTracerFromCtx(ctx *fiber.Ctx, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return Tracer.Start(FromCtx(ctx), spanName, opts...)
}

func SpanFromCtx(ctx *fiber.Ctx) trace.Span {
	return trace.SpanFromContext(FromCtx(ctx))
}
