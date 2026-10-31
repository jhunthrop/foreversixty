import DesignSystem
import Testing

/// The press scale (spec §8): 0.98 while held, never under Reduce Motion or on a disabled button.
@Suite("Brand button press")
struct BrandButtonStyleTests {
    @Test("a held, enabled button scales to 0.98")
    func pressed() {
        #expect(BrandButtonStyle.scale(pressed: true, reduceMotion: false, enabled: true) == 0.98)
    }

    @Test("an idle button stays at 1")
    func idle() {
        #expect(BrandButtonStyle.scale(pressed: false, reduceMotion: false, enabled: true) == 1)
    }

    @Test("Reduce Motion skips the scale")
    func reduceMotion() {
        #expect(BrandButtonStyle.scale(pressed: true, reduceMotion: true, enabled: true) == 1)
    }

    @Test("a disabled button never scales")
    func disabled() {
        #expect(BrandButtonStyle.scale(pressed: true, reduceMotion: false, enabled: false) == 1)
    }
}
