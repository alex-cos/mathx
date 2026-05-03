package mathx

type Integer interface {
	int | int8 | int16 | int32 | int64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Integer | Float
}

type DistributionInterval[F Float, I Integer] struct {
	Inf   F `json:"inf"`
	Sup   F `json:"sup"`
	Count I `json:"count"`
}
