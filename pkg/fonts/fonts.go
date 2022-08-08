package fonts

import (
	_ "embed"

	"github.com/golang/freetype/truetype"
	"github.com/sirupsen/logrus"
)

var (
	//go:embed Roboto-Black.ttf
	robotoBlackTTF []byte

	//go:embed Roboto-Regular.ttf
	robotoRegularlTTF []byte

	//go:embed Roboto-Italic.ttf
	robotoItalicRegularTTF []byte

	//go:embed Roboto-Medium.ttf
	robotoMediumTTF []byte

	//go:embed Roboto-MediumItalic.ttf
	robotoMediumItalicTTF []byte
)

var (
	RobotoBlack         *truetype.Font
	RobotoRegularl      *truetype.Font
	RobotoItalicRegular *truetype.Font
	RobotMedium         *truetype.Font
	RobotMediumItalic   *truetype.Font
)

func init() {
	RobotoBlack = mustLoadFont(robotoBlackTTF)
	RobotoRegularl = mustLoadFont(robotoRegularlTTF)
	RobotoItalicRegular = mustLoadFont(robotoItalicRegularTTF)
	RobotMedium = mustLoadFont(robotoMediumTTF)
	RobotMediumItalic = mustLoadFont(robotoMediumItalicTTF)
}

func mustLoadFont(ttf []byte) *truetype.Font {
	f, err := truetype.Parse(ttf)
	if err != nil {
		logrus.Fatalf("failed to parse font: %s", err)
	}
	return f
}
