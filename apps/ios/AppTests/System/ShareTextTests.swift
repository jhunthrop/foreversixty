@testable import ForeverSixty
import Foundation
import Models
import Testing

@Suite("Share messages and links")
struct ShareTextTests {
    @Test("a character")
    func character() {
        let item = ShareText.character(key: SystemStory.key, name: "Lightningbolltz", spec: "Elemental", className: "shaman", level: 60)
        #expect(item == ShareItem(
            url: URL(string: "https://foreversixty.gg/character/us/pvp/lightningbolltz")!,
            text: "Lightningbolltz, Elemental Shaman, level 60"
        ))
        #expect(ShareText.character(key: SystemStory.key, name: "Lightningbolltz", spec: nil, className: nil, level: nil).text == "Lightningbolltz")
    }

    @Test("a guide")
    func guide() throws {
        let item = ShareText.guide(try GuidesTestData.elemental())
        #expect(item.url.absoluteString == "https://foreversixty.gg/guides/shaman/elemental")
        #expect(item.text == "Elemental Shaman guide")
    }

    @Test("the tier lists")
    func tiers() {
        #expect(ShareText.tiers(.dps) == ShareItem(url: URL(string: "https://foreversixty.gg/tiers")!, text: "DPS tier list"))
        #expect(ShareText.tiers(.tank).url.absoluteString == "https://foreversixty.gg/tiers/tank")
        #expect(ShareText.tiers(.healer).url.absoluteString == "https://foreversixty.gg/tiers/healer")
        #expect(ShareText.tiers(.healer).text == "Healer tier list")
    }

    @Test("a public or unlisted report shares; guild and private do not")
    func report() {
        let date = RaidTime.date("2026-12-15T01:58:00Z")
        let item = ShareText.report(id: "ewlrx3enxbeb", zone: "Onyxia's Lair", visibility: "public", createdAt: date, zoneOfPhone: SystemStory.zone)
        #expect(item == ShareItem(url: URL(string: "https://foreversixty.gg/reports/ewlrx3enxbeb")!, text: "Onyxia's Lair · 14 Dec"))
        #expect(ShareText.report(id: "x", zone: "Z", visibility: "unlisted", createdAt: date) != nil)
        #expect(ShareText.report(id: "x", zone: "Z", visibility: "guild", createdAt: date) == nil)
        #expect(ShareText.report(id: "x", zone: "Z", visibility: "private", createdAt: date) == nil)
    }
}

@MainActor
@Suite("Haptics: the facade records every moment")
struct HapticsTests {
    @Test("a saved sign-up plays selection then success; a refusal selection then error")
    func signup() async throws {
        let haptics = RecordingHaptics()
        haptics.selection()
        haptics.success()
        haptics.selection()
        haptics.error()
        haptics.play(.impactMedium)
        #expect(haptics.moments == [.selection, .success, .selection, .error, .impactMedium])
    }
}
