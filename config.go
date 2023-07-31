package captcha

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "Captcha"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

type Config struct {
	Long     uint    // 验证码长度
	Width    uint    // 验证码宽度
	Height   uint    // 验证码高度
	MaxScrew float64 // MaxSkew max absolute skew factor of a single digit.
	DotCount int     // Number of background circles.
}
