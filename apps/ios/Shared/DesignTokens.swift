// PLACEHOLDER. Lane T replaces this file with one generated from
// web/src/styles/tokens.css (checked in CI so web, iOS and Android never drift).
// Do not hand-edit tokens here; until the generator lands these few values are
// copied from tokens.css so the scaffold renders in the site's palette.
// Sources/DesignSystem/Theme.swift is the only reader; adapt it when the generated names land.

public enum DesignTokens {
    /// sRGB hex, 0xRRGGBB.
    public enum Color {
        public static let bg: UInt32 = 0x07090D
        public static let raised: UInt32 = 0x0D111A
        public static let line: UInt32 = 0x262E40
        public static let text: UInt32 = 0xE9E4D8
        public static let strong: UInt32 = 0xF2EEE4
        public static let muted: UInt32 = 0x9A9484
        public static let gold: UInt32 = 0xE5B955
    }

    /// Points.
    public enum Spacing {
        public static let xs: Double = 4
        public static let sm: Double = 8
        public static let md: Double = 16
        public static let lg: Double = 24
    }

    /// Points.
    public enum Radius {
        public static let card: Double = 8
    }
}
