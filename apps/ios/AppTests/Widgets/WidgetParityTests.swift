@testable import ForeverSixty
import DesignSystem
import DesignTokens
import Models
import SwiftUI
import Testing
import UIKit

/// The mocks' data (design/mocks/ws_app.py `widgets`, `live_activity`): Thandrel, a Fury Warrior
/// in Olympus XXVII, Onyxia's Lair on Wed 9 Dec at 8:00 pm CST; Felsnap at 34 with no guild.
enum WidgetMockData {
    static let zone = TimeZone(identifier: "America/Chicago") ?? .gmt
    /// The mocks' clock (ws_home NOW).
    static let now = RaidTime.date("2026-11-06T18:00:00Z")
    static let onyxia = RaidTime.date("2026-12-10T02:00:00Z")

    static let icons = WidgetIcons { icon in
        let url = HomeFixtures.repoRoot.appending(path: "data/builds/\(BisTestData.build)/icons/\(icon).webp")
        return (try? Data(contentsOf: url)).flatMap(UIImage.init(data:))
    }

    private static func upgrades(_ name: String, quality: Int, icon: String, gain: Double) -> WidgetUpgrades {
        WidgetUpgrades(state: "ok", count: 3, totalGain: gain + 9, unit: "dps", rows: [
            WidgetUpgradeRow(slot: "head", slotLabel: "Head", pick: WidgetUpgradePick(id: 1, name: name, quality: quality, icon: icon, itemLevel: 62), gain: gain),
        ])
    }

    private static let thandrel = WidgetCharacter(key: "us/pvp/thandrel", name: "Thandrel", className: "warrior", specName: "Fury")
    private static let guild = WidgetGuild(id: 2, name: "Olympus XXVII")

    static let raid = WidgetNextRaid(
        raidId: 7, guildId: 2, guildName: "Olympus XXVII", zone: "raid:onyxias-lair", zoneName: "Onyxia's Lair",
        startsAt: onyxia, pullAt: onyxia, endsAt: onyxia.addingTimeInterval(3 * 3600), size: 40, signedUp: 21, status: "open",
        my: WidgetMine(state: "in", spec: "Fury", role: "dps"), canSignUp: true
    )

    static func snapshot(_ state: String) -> WidgetSnapshot {
        let widget: Models.Widget = switch state {
        case "guild":
            Models.Widget(generatedAt: now, firstRaidsAt: "2026-12-09", character: thandrel, guild: guild, nextRaid: raid,
                   upgrades: upgrades("Lionheart Helm", quality: 4, icon: "inv_helmet_36", gain: 14.5), live: WidgetLive(simRunsRunning: 0))
        case "noguild":
            Models.Widget(generatedAt: now, firstRaidsAt: "2026-12-09",
                   character: WidgetCharacter(key: "us/pvp/felsnap", name: "Felsnap", className: "hunter", specName: "Marksmanship"),
                   upgrades: upgrades("Felbough Arm", quality: 3, icon: "inv_spear_08", gain: 1.3), live: WidgetLive(simRunsRunning: 0))
        default:
            Models.Widget(generatedAt: now, firstRaidsAt: "2026-12-09", character: thandrel, guild: guild,
                   upgrades: upgrades("Lionheart Helm", quality: 4, icon: "inv_helmet_36", gain: 14.5), live: WidgetLive(simRunsRunning: 0))
        }
        return WidgetSnapshot(widget: widget, fetchedAt: now)
    }

    /// `open` is after the first raids opened.
    static func moment(_ state: String) -> Date {
        state == "open" ? RaidTime.date("2026-12-12T18:00:00Z") : now
    }

    static let attributes = RaidActivityAttributes(raid, characterName: "Thandrel")
    /// 42 minutes before the pull, 18 in and 3 late, 2 tanks and 5 heals.
    static var live: RaidActivityAttributes.ContentState {
        var state = RaidActivityAttributes.ContentState(raid, myState: "in")
        state.inCount = 18
        state.late = 3
        state.tanks = 2
        state.healers = 5
        return state
    }

    static var done: RaidActivityAttributes.ContentState {
        var state = live
        state.status = RaidActivityAttributes.ContentState.done
        state.kills = 1
        state.loot = RaidLiveLoot(name: "Deathbringer", quality: 4, icon: "inv_axe_09")
        state.reportReady = true
        return state
    }

    static let liveNow = onyxia.addingTimeInterval(-42 * 60)
}

/// The gallery pages the mocks draw the system surfaces on: the title, then each surface under its
/// label, every surface reporting its region. Test-only; the system draws the real ones.
struct WidgetGalleryPage<Content: View>: View {
    let title: String
    @ViewBuilder let content: () -> Content

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                // The mocks' app-only pages inset the title by the page's and the header's gutters.
                PageHeader(title: title).padding(.horizontal, FSSpace.gutterPhone)
                content()
            }
            .padding(.top, FSSpace.space2)
            .padding(.horizontal, FSSpace.gutterPhone)
            .padding(.bottom, FSSpace.space6)
            .parityContent()
        }
        .background(FSColor.bg)
    }
}

/// The mocks' `.lbl`: 12/16 bold muted, 22 above and 8 below.
struct GalleryLabel: View {
    let text: String

    init(_ text: String) {
        self.text = text
    }

    var body: some View {
        Text(text).type(WidgetType.label).foregroundStyle(FSColor.muted).padding(.top, 22).padding(.bottom, 8)
    }
}

/// The home screen widgets in the mocks' order.
struct WidgetGallery: View {
    let state: String

    var body: some View {
        let snapshot = WidgetMockData.snapshot(state)
        let raid = RaidGlance(snapshot, now: WidgetMockData.moment(state), zone: WidgetMockData.zone)
        let upgrade = UpgradeGlance(snapshot)
        let icons = WidgetMockData.icons
        WidgetGalleryPage(title: "Widgets") {
            GalleryLabel("Home screen widgets: small")
            HStack(spacing: 16) {
                WidgetFrame(size: WidgetMetrics.small) { NextRaidSmallView(glance: raid) }
                    .parityRegion(ParityRegionName("widget-raid-small"))
                WidgetFrame(size: WidgetMetrics.small) { UpgradeSmallView(glance: upgrade, icons: icons) }
                    .parityRegion(ParityRegionName("widget-upgrade-small"))
            }
            GalleryLabel("Medium")
            WidgetFrame(size: WidgetMetrics.medium) { GlanceMediumView(raid: raid, upgrade: upgrade, icons: icons) }
                .parityRegion(ParityRegionName("widget-medium"))
            if state == "guild" {
                GalleryLabel("Medium, answer the raid")
                WidgetFrame(size: WidgetMetrics.medium) {
                    RaidAnswerMediumView(glance: raid, raid: snapshot.widget.nextRaid, characterKey: nil)
                }
                .parityRegion(ParityRegionName("widget-answer"))
            }
            GalleryLabel("Lock screen")
            HStack(spacing: 10) {
                RaidInlineView(glance: raid).lineLimit(1)
                UpgradeRectangularView(glance: upgrade).fixedSize()
            }
            .font(.footnote.weight(.semibold))
            .foregroundStyle(FSColor.strong)
            .padding(16)
            .frame(maxWidth: .infinity)
            .background(FSColor.raised, in: .rect(cornerRadius: 24))
        }
    }
}

/// The Live Activity's surfaces in the mock's order.
struct LiveActivityGallery: View {
    var body: some View {
        let attributes = WidgetMockData.attributes
        let icons = WidgetMockData.icons
        let now = WidgetMockData.liveNow
        WidgetGalleryPage(title: "Live Activity") {
            GalleryLabel("Lock screen")
            RaidLiveLockView(attributes: attributes, state: WidgetMockData.live, icons: icons, fixedNow: now, zone: WidgetMockData.zone)
                .liveCard()
                .parityRegion(ParityRegionName("live-lock"))
            GalleryLabel("Raid done")
            RaidLiveLockView(attributes: attributes, state: WidgetMockData.done, icons: icons, fixedNow: now, zone: WidgetMockData.zone)
                .liveCard()
                .parityRegion(ParityRegionName("live-done"))
            GalleryLabel("Dynamic Island, compact")
            HStack {
                RaidIslandCompactLeading(attributes: attributes, state: WidgetMockData.live)
                Spacer()
                LiveCountdown(target: WidgetMockData.live.startDate, fixedNow: now, role: LiveType.island)
            }
            .padding(.horizontal, 14)
            .frame(width: 250, height: 37)
            .background(.black, in: .capsule)
            .parityRegion(ParityRegionName("island-compact"))
            .frame(maxWidth: .infinity)
            GalleryLabel("Dynamic Island, expanded")
            RaidIslandExpandedView(attributes: attributes, state: WidgetMockData.live, icons: icons, fixedNow: now)
                .padding(.horizontal, 20)
                .padding(.vertical, 16)
                .background(.black, in: .rect(cornerRadius: 36))
                .parityRegion(ParityRegionName("island-expanded"))
            GalleryLabel("Sim running")
            SimLiveLockView(
                attributes: SimActivityAttributes(simID: "k3v9q2m8x1ab", title: "Fury Warrior, 60 raid"),
                state: .init(state: "running", done: 64, total: 100, etaSec: 120, dps: nil)
            )
            .liveCard()
            .parityRegion(ParityRegionName("live-sim"))
        }
    }
}

private extension View {
    /// The lock screen's activity ground (the mock's `.la`).
    func liveCard() -> some View {
        frame(maxWidth: .infinity, alignment: .leading)
            .background(FSColor.raised.opacity(0.9), in: .rect(cornerRadius: LiveMetrics.radius))
    }
}

@MainActor
@Suite("Widget parity")
struct WidgetParityTests {
    static let device = CGSize(width: 393, height: 852)

    @Test("ios-widgets-guild")
    func guild() throws {
        try ParityCheck.assert(WidgetGallery(state: "guild"), screen: "widgets-guild", size: Self.device)
    }

    @Test("ios-widgets-noguild")
    func noGuild() throws {
        try ParityCheck.assert(WidgetGallery(state: "noguild"), screen: "widgets-noguild", size: Self.device)
    }

    @Test("ios-widgets-open")
    func open() throws {
        try ParityCheck.assert(WidgetGallery(state: "open"), screen: "widgets-open", size: Self.device)
    }

    @Test("ios-live-activity")
    func liveActivity() throws {
        try ParityCheck.assert(LiveActivityGallery(), screen: "live-activity", size: Self.device)
    }
}
