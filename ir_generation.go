package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TextDecoration struct {
	Bold      int
	Italic    int
	Underline int
	StrikeOut int
}

func mapTextDecoration(input string) TextDecoration {
	output := TextDecoration{
		Bold:      0,
		Italic:    0,
		Underline: 0,
		StrikeOut: 0,
	}
	inputProperties := strings.Split(input, " ")

	m := map[string]int{
		"bold":         0,
		"italic":       0,
		"underline":    0,
		"line-through": 0,
	}

	if len(inputProperties) > 0 && len(inputProperties) < 5 {
		for _, property := range inputProperties {
			_, ok := m[property]
			if ok {
				m[property] = 1
			}
		}
		output = TextDecoration{
			Bold:      m["bold"],
			Italic:    m["italic"],
			Underline: m["underline"],
			StrikeOut: m["line-through"],
		}
	}
	return output
}

type IR1Transform struct {
	ScaleX float64 `json:"scaleX"`
	ScaleY float64 `json:"scaleY"`
	Rotate float64 `json:"rotate"`
}

type IR1StyleProps struct {
	FontFamily            string        `json:"fontFamily"`
	FontSize              int           `json:"fontSize"`
	Color                 string        `json:"color"`
	KaraokeSecondaryColor string        `json:"karaokeSecondaryColor"`
	TextStrokeColor       string        `json:"textStrokeColor"`
	BoxBackgroundColor    string        `json:"boxBackgroundColor"`
	TextDecoration        string        `json:"textDecoration"`
	Transform             IR1Transform  `json:"transform"`
	LetterSpacing         int           `json:"letterSpacing"`
	BorderStyle           string        `json:"borderStyle"`
	TextStrokeWidth       int           `json:"textStrokeWidth"`
	TextShadowDepth       int           `json:"textShadowDepth"`
	Alignment             string        `json:"alignment"`
	MarginLeft            int           `json:"marginLeft"`
	MarginRight           int           `json:"marginRight"`
	MarginVertical        int           `json:"marginVertical"`
	Encoding              string        `json:"encoding"`
}

type IR1Input struct {
	GlobalStyle []map[string]IR1StyleProps `json:"globalStyle"`
	Subtitles   []json.RawMessage          `json:"subtitles"`
}

func parseNamedColor(s string) (string, bool) {
	named := map[string]string{
		"yellow":  "&H0000FFFF",
		"red":     "&H000000FF",
		"lime":    "&H0000FF00",
		"green":   "&H00008000",
		"cyan":    "&H00FFFF00",
		"magenta": "&H00FF00FF",
	}
	hex, ok := named[strings.ToLower(s)]
	return hex, ok
}

func parseCSSColor(s string) (string, error) {
	s = strings.TrimPrefix(s, "rgba(")
	s = strings.TrimSuffix(s, ")")
	parts := strings.Split(s, ",")
	if len(parts) != 4 {
		return "", fmt.Errorf("invalid rgba: %s", s)
	}
	r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return "", err
	}
	g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return "", err
	}
	b, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		return "", err
	}
	a, err := strconv.ParseFloat(strings.TrimSpace(parts[3]), 64)
	if err != nil {
		return "", err
	}
	alpha := int((1.0 - a) * 255)
	return fmt.Sprintf("&H%02X%02X%02X%02X", alpha, b, g, r), nil
}

func parseColor(s string) string {
	if hex, ok := parseNamedColor(s); ok {
		return hex
	}
	hex, err := parseCSSColor(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: failed to parse color %q: %v\n", s, err)
		return "&H00FFFFFF"
	}
	return hex
}

func mapAlignment(s string) int {
	m := map[string]int{
		"bottom-left":   1,
		"bottom-center": 2,
		"bottom-right":  3,
		"middle-left":   4,
		"middle-center": 5,
		"middle-right":  6,
		"top-left":      7,
		"top-center":    8,
		"top-right":     9,
	}
	if v, ok := m[s]; ok {
		return v
	}
	fmt.Fprintf(os.Stderr, "warning: unknown alignment %q, using 2 (bottom-center)\n", s)
	return 2
}

func mapBorderStyle(s string) int {
	m := map[string]int{
		"outline-shadow": 1,
		"opaque-box":     3,
	}
	if v, ok := m[s]; ok {
		return v
	}
	fmt.Fprintf(os.Stderr, "warning: unknown borderStyle %q, using 1 (outline+shadow)\n", s)
	return 1
}

func mapEncoding(s string) int {
	m := map[string]int{
		"ansi":  0,
		"utf-8": 1,
	}
	if v, ok := m[strings.ToLower(s)]; ok {
		return v
	}
	if v, err := strconv.Atoi(s); err == nil {
		return v
	}
	fmt.Fprintf(os.Stderr, "warning: unknown encoding %q, using 0 (ANSI)\n", s)
	return 0
}

func convertGlobalStyle(name string, props IR1StyleProps) GlobalStyle {
	td := mapTextDecoration(props.TextDecoration)
	return GlobalStyle{
		Name:            name,
		Fontname:        props.FontFamily,
		Fontsize:        props.FontSize,
		PrimaryColour:   parseColor(props.Color),
		SecondaryColour: parseColor(props.KaraokeSecondaryColor),
		OutlineColour:   parseColor(props.TextStrokeColor),
		BackColour:      parseColor(props.BoxBackgroundColor),
		Bold:            td.Bold,
		Italic:          td.Italic,
		Underline:       td.Underline,
		StrikeOut:       td.StrikeOut,
		ScaleX:          int(props.Transform.ScaleX * 100),
		ScaleY:          int(props.Transform.ScaleY * 100),
		Spacing:         props.LetterSpacing,
		Angle:           int(props.Transform.Rotate),
		BorderStyle:     mapBorderStyle(props.BorderStyle),
		Outline:         props.TextStrokeWidth,
		Shadow:          props.TextShadowDepth,
		Alignment:       mapAlignment(props.Alignment),
		MarginL:         props.MarginLeft,
		MarginR:         props.MarginRight,
		MarginV:         props.MarginVertical,
		Encoding:        mapEncoding(props.Encoding),
	}
}

func JsonGeneration(inputPath string, outputPath string) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("error reading input:", err)
		return
	}

	var input IR1Input
	if err := json.Unmarshal(data, &input); err != nil {
		fmt.Println("error unmarshalling input:", err)
		return
	}

	var styles []GlobalStyle
	for _, entry := range input.GlobalStyle {
		for name, props := range entry {
			styles = append(styles, convertGlobalStyle(name, props))
		}
	}

	output := IR{
		GlobalStyle: styles,
		Events:      []Event{},
	}

	outData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Println("error marshalling output:", err)
		return
	}

	if err := os.WriteFile(outputPath, outData, 0644); err != nil {
		fmt.Println("error writing output:", err)
		return
	}

	fmt.Println("successfully generated", outputPath)
}
