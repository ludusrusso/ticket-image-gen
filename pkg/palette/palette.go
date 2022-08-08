package palette

type ColorPalette struct {
	V100 string
	V200 string
	V300 string
	V400 string
	V500 string
	V600 string
	V700 string
	V800 string
	V900 string
}

type Palettes map[string]ColorPalette

var ColorPalettes = Palettes{
	"gray": ColorPalette{
		V100: "#f7fafc",
		V200: "#edf2f7",
		V300: "#e2e8f0",
		V400: "#cbd5e0",
		V500: "#a0aec0",
		V600: "#718096",
		V700: "#4a5568",
		V800: "#2d3748",
		V900: "#1a202c",
	},
	"red": ColorPalette{
		V100: "#fff5f5",
		V200: "#fed7d7",
		V300: "#feb2b2",
		V400: "#fc8181",
		V500: "#f56565",
		V600: "#e53e3e",
		V700: "#c53030",
		V800: "#9b2c2c",
		V900: "#742a2a",
	},
	"orange": ColorPalette{
		V100: "#fffaf0",
		V200: "#feebc8",
		V300: "#fbd38d",
		V400: "#f6ad55",
		V500: "#ed8936",
		V600: "#dd6b20",
		V700: "#c05621",
		V800: "#9c4221",
		V900: "#7b341e",
	},
	"yellow": ColorPalette{
		V100: "#fffff0",
		V200: "#fefcbf",
		V300: "#faf089",
		V400: "#f6e05e",
		V500: "#ecc94b",
		V600: "#d69e2e",
		V700: "#b7791f",
		V800: "#975a16",
		V900: "#744210",
	},
	"green": ColorPalette{
		V100: "#f0fff4",
		V200: "#c6f6d5",
		V300: "#9ae6b4",
		V400: "#68d391",
		V500: "#48bb78",
		V600: "#38a169",
		V700: "#2f855a",
		V800: "#276749",
		V900: "#22543d",
	},
	"teal": ColorPalette{
		V100: "#e6fffa",
		V200: "#b2f5ea",
		V300: "#81e6d9",
		V400: "#4fd1c5",
		V500: "#38b2ac",
		V600: "#319795",
		V700: "#2c7a7b",
		V800: "#285e61",
		V900: "#234e52",
	},
	"blue": ColorPalette{
		V100: "#ebf8ff",
		V200: "#bee3f8",
		V300: "#90cdf4",
		V400: "#63b3ed",
		V500: "#4299e1",
		V600: "#3182ce",
		V700: "#2b6cb0",
		V800: "#2c5282",
		V900: "#2a4365",
	},
	"indigo": ColorPalette{
		V100: "#ebf4ff",
		V200: "#c3dafe",
		V300: "#a3bffa",
		V400: "#7f9cf5",
		V500: "#667eea",
		V600: "#5a67d8",
		V700: "#4c51bf",
		V800: "#434190",
		V900: "#3c366b",
	},
	"purple": ColorPalette{
		V100: "#faf5ff",
		V200: "#e9d8fd",
		V300: "#d6bcfa",
		V400: "#b794f4",
		V500: "#9f7aea",
		V600: "#805ad5",
		V700: "#6b46c1",
		V800: "#553c9a",
		V900: "#44337a",
	},
	"pink": ColorPalette{
		V100: "#fff5f7",
		V200: "#fed7e2",
		V300: "#fbb6ce",
		V400: "#f687b3",
		V500: "#ed64a6",
		V600: "#d53f8c",
		V700: "#b83280",
		V800: "#97266d",
		V900: "#702459",
	},
}
