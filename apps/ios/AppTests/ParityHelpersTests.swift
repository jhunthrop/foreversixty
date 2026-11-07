import DesignSystem
import SwiftUI
import Testing

/// The parity helpers themselves, on a small screen with known regions, before any screen
/// carries them (design/parity.md "Native").
@MainActor
@Suite("Parity helpers")
struct ParityHelpersTests {
    private static let size = CGSize(width: 393, height: 400)

    private var screen: some View {
        VStack(alignment: .leading, spacing: 0) {
            Text("Play your class better.")
                .frame(width: 361, height: 40, alignment: .leading)
                .parityRegion(.headTitle)
            VStack(spacing: 0) {
                ForEach(["Warrior", "Hunter"], id: \.self) { name in
                    Text(name).frame(width: 393, height: 56).parityItem(.row)
                }
            }
            .parityRegion(.section("Pick your class"))
        }
        .padding(.leading, 16)
        .frame(width: Self.size.width, height: Self.size.height, alignment: .topLeading)
        .parityContent()
    }

    private func reference(_ regions: [String: ParityRegions.Reference.Region]) -> ParityRegions.Reference {
        ParityRegions.Reference(device: .init(width: 393, height: 852), regions: regions)
    }

    @Test func sectionRegionsUseTheWebSlug() {
        #expect(ParityRegionName.section("What to stack").rawValue == "section-what-to-stack")
        #expect(ParityRegionName.section("Top 3 for Fury Warrior, 60").rawValue == "section-top-3-for-fury-warrior-60")
    }

    @Test func aMatchingScreenHolds() {
        let capture = ParityRegions.capture(of: screen, size: Self.size)
        let want = reference([
            "head-title": .init(box: [16, 0, 361, 40], inventory: [:]),
            "section-pick-your-class": .init(box: [16, 40, 393, 112], inventory: ["row": 2]),
        ])
        #expect(ParityRegions.compare(capture, with: want) == [])
    }

    @Test func driftAndMissingRegionsFail() {
        let capture = ParityRegions.capture(of: screen, size: Self.size)
        let want = reference([
            "head-title": .init(box: [16, 30, 361, 40], inventory: [:]),
            "section-pick-your-class": .init(box: [16, 40, 393, 112], inventory: ["row": 3]),
            "countdowns": .init(box: [16, 200, 361, 80], inventory: [:]),
        ])
        #expect(ParityRegions.compare(capture, with: want) == [
            "countdowns: in the mock, missing from the build",
            "head-title: y 0, the mock has 30 (within 8)",
            "section-pick-your-class: inventory row 2, the mock has row 3",
        ])
    }

    @Test func textReadsInOrder() {
        #expect(ParityText.lines(of: screen, size: Self.size) == ["Play your class better.", "Warrior", "Hunter"])
    }
}
