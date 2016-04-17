package easyscripts

type linePart interface {
	String(ctx *scriptContext) string
}
