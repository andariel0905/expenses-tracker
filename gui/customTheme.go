package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyTheme struct{}

func (m MyTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 233, G: 242, B: 208, A: 0}

	case theme.ColorNameButton:
		return color.White

	case theme.ColorNameForeground:
		return color.RGBA{R: 99, G: 145, B: 42, A: 200}

	case theme.ColorNameSeparator:
		return color.RGBA{R: 99, G: 145, B: 42, A: 200}

	case theme.ColorNameHover:
		return color.RGBA{R: 229, G: 220, B: 225, A: 100}
		//		return color.RGBA{R: 99, G: 145, B: 42, A: 0}

	case theme.ColorNameInputBackground:
		// Cambiar el color de fondo de los controles de entrada (como Select)
		return color.White

	default:
		return theme.DefaultTheme().Color(name, theme.VariantLight)
	}
}

func (m MyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(theme.IconNameHome)
}

func (m MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
