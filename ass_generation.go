package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type GlobalStyle struct {
	Name            string `json:"Name"`
	Fontname        string `json:"Fontname"`
	Fontsize        int    `json:"Fontsize"`
	PrimaryColour   string `json:"PrimaryColour"`
	SecondaryColour string `json:"SecondaryColour"`
	OutlineColour   string `json:"OutlineColour"`
	BackColour      string `json:"BackColour"`
	Bold            int    `json:"Bold"`
	Italic          int    `json:"Italic"`
	Underline       int    `json:"Underline"`
	StrikeOut       int    `json:"StrikeOut"`
	ScaleX          int    `json:"ScaleX"`
	ScaleY          int    `json:"ScaleY"`
	Spacing         int    `json:"Spacing"`
	Angle           int    `json:"Angle"`
	BorderStyle     int    `json:"BorderStyle"`
	Outline         int    `json:"Outline"`
	Shadow          int    `json:"Shadow"`
	Alignment       int    `json:"Alignment"`
	MarginL         int    `json:"MarginL"`
	MarginR         int    `json:"MarginR"`
	MarginV         int    `json:"MarginV"`
	Encoding        int    `json:"Encoding"`
}

type Event struct {
	Layer   int    `json:"Layer"`
	Start   string `json:"Start"`
	End     string `json:"End"`
	Style   string `json:"Style"`
	Name    string `json:"Name"`
	MarginL int    `json:"MarginL"`
	MarginR int    `json:"MarginR"`
	MarginV int    `json:"MarginV"`
	Effect  string `json:"Effect"`
	Text    string `json:"Text"`
}

type IR struct {
	GlobalStyle []GlobalStyle `json:"globalStyle"`
	Events      []Event       `json:"events"`
}

const (
	styleFormat = "Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding"
	eventFormat = "Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text"
)

var scriptInfo = `[Script Info]
ScriptType: v4.00+
PlayResX: 1080
PlayResY: 1920
WrapStyle: 0
ScaledBorderAndShadow: yes
`

func parseIR(path string) (IR, error) {
	byteValue, err := os.ReadFile(path)
	if err != nil {
		return IR{}, fmt.Errorf("error reading file: %w", err)
	}

	var ir IR
	if err := json.Unmarshal(byteValue, &ir); err != nil {
		return IR{}, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return ir, nil
}

func buildStyles(styles []GlobalStyle) string {
	styleLines := make([]string, len(styles))
	for i, style := range styles {
		styleLines[i] = fmt.Sprintf("Style: %s,%s,%d,%s,%s,%s,%s,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d",
			style.Name,
			style.Fontname,
			style.Fontsize,
			style.PrimaryColour,
			style.SecondaryColour,
			style.OutlineColour,
			style.BackColour,
			style.Bold,
			style.Italic,
			style.Underline,
			style.StrikeOut,
			style.ScaleX,
			style.ScaleY,
			style.Spacing,
			style.Angle,
			style.BorderStyle,
			style.Outline,
			style.Shadow,
			style.Alignment,
			style.MarginL,
			style.MarginR,
			style.MarginV,
			style.Encoding,
		)
	}
	return strings.Join(styleLines, "\n")
}

func buildEvents(events []Event) string {
	dialogueLines := make([]string, len(events))
	for i, event := range events {
		dialogueLines[i] = fmt.Sprintf("Dialogue: %d,%s,%s,%s,%s,%d,%d,%d,,%s",
			event.Layer,
			event.Start,
			event.End,
			event.Style,
			event.Name,
			event.MarginL,
			event.MarginR,
			event.MarginV,
			event.Text,
		)
	}
	return strings.Join(dialogueLines, "\n")
}

func generateAss(ir IR) string {
	var sb strings.Builder
	sb.WriteString(scriptInfo)
	sb.WriteString("\n[V4+ Styles]\n")
	sb.WriteString("Format: " + styleFormat + "\n")
	sb.WriteString(buildStyles(ir.GlobalStyle) + "\n")
	sb.WriteString("\n[Events]\n")
	sb.WriteString("Format: " + eventFormat + "\n")
	sb.WriteString(buildEvents(ir.Events))
	return sb.String()
}

func AssGeneration(inputPath string, outputPath string) {
	ir, err := parseIR(inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	output := generateAss(ir)

	if err := os.WriteFile(outputPath, []byte(output), 0644); err != nil {
		fmt.Println("error when writing output.ass", err)
		return
	}

	fmt.Println("successfully generated", outputPath)
}
