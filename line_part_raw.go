package easyscripts

type linePartRaw string

func (l linePartRaw) String(ctx *scriptContext) string {
	return string(l)
}
