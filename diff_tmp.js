const a = {
  "Default": {
    // Font family used to render the text
    // ASS equivalent: Fontname
    // Possible values: any installed font name (e.g., "Arial", "Verdana", "Times New Roman")
    "fontFamily": "Arial",

    // Font size in pixels
    // ASS equivalent: Fontsize
    // Possible values: any positive number (commonly 10–100+ depending on resolution)
    "fontSize": 20,

    // Primary text color (main visible color)
    // ASS equivalent: PrimaryColour (&HAABBGGRR format)
    // Possible values: CSS color formats (rgba, hex, etc.)
    "color": "rgba(255, 255, 255, 1)",

    // Secondary color used mainly for karaoke effects (progressive highlighting)
    // ASS equivalent: SecondaryColour
    // Possible values: any color; only used in karaoke/timing effects
    "karaokeSecondaryColor": "rgba(255, 0, 0, 1)",

    // Color of the text outline (stroke)
    // ASS equivalent: OutlineColour
    // Possible values: any color
    "textStrokeColor": "rgba(0, 0, 0, 1)",

    // Background color for box-style subtitles or shadow base
    // ASS equivalent: BackColour
    // Possible values: any color; visible mainly when BorderStyle = 3
    "boxBackgroundColor": "rgba(0, 0, 0, 1)",

    // General text decoration shorthand
    // ASS equivalent: a combination of Bold, Italic, Underline, StrikeOut
    // Possible values: "none", "bold italic underline line-through" or remove the ones you want
    "textDecoration": "none", // "bold italic underline line-through"

    // Transformations applied to the text
    // ASS equivalents: ScaleX, ScaleY, Angle
    "transform": {
      // Horizontal scaling factor
      // ASS equivalent: ScaleX (percentage, 100 = normal)
      // Possible values: number (1 = 100%)
      "scaleX": 1,

      // Vertical scaling factor
      // ASS equivalent: ScaleY (percentage, 100 = normal)
      // Possible values: number (1 = 100%)
      "scaleY": 1,

      // Rotation angle in degrees
      // ASS equivalent: Angle
      // Possible values: any number (positive = clockwise)
      "rotate": 0
    },

    // Space between characters
    // ASS equivalent: Spacing
    // Possible values: number (can be negative or positive)
    "letterSpacing": 0,

    // Defines how borders/background are rendered
    // ASS equivalent: BorderStyle
    // Possible values:
    //   "outline-shadow" → BorderStyle = 1 (outline + shadow)
    //   "opaque-box" → BorderStyle = 3 (solid background box)
    "borderStyle": "outline-shadow",

    // Thickness of the text outline (stroke)
    // ASS equivalent: Outline
    // Possible values: number (pixels)
    "textStrokeWidth": 2,

    // Shadow depth (distance offset)
    // ASS equivalent: Shadow
    // Possible values: number (pixels)
    "textShadowDepth": 0,

    // Combined alignment (vertical + horizontal)
    // ASS equivalent: Alignment (1–9 numpad layout)
    // Possible values:
    //   "bottom-left", "bottom-center", "bottom-right"
    //   "middle-left", "middle-center", "middle-right"
    //   "top-left", "top-center", "top-right"
    "alignment": "middle-center",

    // Left margin from screen edge
    // ASS equivalent: MarginL
    // Possible values: number (pixels)
    "marginLeft": 10,

    // Right margin from screen edge
    // ASS equivalent: MarginR
    // Possible values: number (pixels)
    "marginRight": 10,

    // Vertical margin (top or bottom depending on alignment)
    // ASS equivalent: MarginV
    // Behavior:
    //   - bottom-aligned → distance from bottom
    //   - top-aligned → distance from top
    // Possible values: number (pixels)
    "marginVertical": 540,

    // Character encoding used for text rendering
    // ASS equivalent: Encoding
    // Possible values:
    //   "ansi", "utf-8", or codepage identifiers (e.g., 0, 1, 128, etc.)
    // Note: not a visual property
    "encoding": "ansi"
  }
}