@testable import ForeverSixty
import DesignSystem
import Models
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// The widget and Live Activity states the parity galleries do not draw (WidgetParityTests holds
/// the mocked ones): each raid answer, the spec ask, locked and cancelled, stale, signed out, the
/// press that did not save, the lock-screen tints, the Live Activity's cancelled and sim states.
@MainActor
@Suite("Widget snapshots")
struct WidgetSnapshotTests {
    private let zone = WidgetTime.chicago
    private let now = WidgetTime.now
    private let icons = WidgetMockData.icons

    private func snap(_ view: some View, _ name: String, size: CGSize, large: Bool = false, ground: Color? = nil, keyWindow: Bool = false) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = large ? .accessibilityExtraExtraExtraLarge : .large
        }
        let framed = view
            .frame(width: size.width, height: size.height, alignment: .top)
            .environment(\.dynamicTypeSize, large ? .accessibility5 : .large)
            .preferredColorScheme(.dark)
            .padding(16)
            .background(ground ?? Color(red: 0.16, green: 0.12, blue: 0.09))
        let layout = SwiftUISnapshotLayout.fixed(width: size.width + 32, height: size.height + 32)
        assertSnapshot(
            of: framed,
            as: keyWindow
                ? .keyWindowImage(layout: layout, traits: traits)
                : .image(precision: SnapshotTolerance.precision, perceptualPrecision: SnapshotTolerance.perceptualPrecision, layout: layout, traits: traits),
            testName: name
        )
    }

    private func glance(_ snapshot: WidgetSnapshot?, at moment: Date? = nil) -> RaidGlance {
        RaidGlance(snapshot, now: moment ?? now, zone: zone)
    }

    private func small(_ snapshot: WidgetSnapshot?, _ name: String) {
        snap(WidgetFrame(size: WidgetMetrics.small) { NextRaidSmallView(glance: glance(snapshot)) }, "ios-raid-small-\(name)", size: WidgetMetrics.small)
    }

    private func medium(_ snapshot: WidgetSnapshot?, _ name: String, large: Bool = false) {
        let view = WidgetFrame(size: WidgetMetrics.medium) {
            RaidAnswerMediumView(glance: glance(snapshot), raid: snapshot?.widget.nextRaid, characterKey: nil)
        }
        snap(view, "ios-raid-medium-\(name)\(large ? "-ax5" : "")", size: WidgetMetrics.medium, large: large)
    }

    private func mine(_ state: String?, spec: String? = "Elemental") -> [String: JSONValue] {
        ["my": .object(["state": state.map(JSONValue.string) ?? .null, "spec": spec.map(JSONValue.string) ?? .null, "role": .string("dps")])]
    }

    @Test("next raid, small: the states the gallery does not draw")
    func smallStates() throws {
        small(try WidgetFixtures.snapshot(patch: mine(nil)), "not-answered")
        small(try WidgetFixtures.snapshot(patch: ["status": .string("cancelled"), "can_sign_up": .bool(false)]), "cancelled")
        small(try WidgetFixtures.snapshot(fetchedAt: now.addingTimeInterval(-14 * 3600)), "stale")
        small(nil, "signed-out")
    }

    @Test("next raid, medium: each answer selected, the spec ask, locked, the failed press, and AX5")
    func mediumStates() throws {
        medium(try WidgetFixtures.snapshot(patch: mine("late")), "late")
        medium(try WidgetFixtures.snapshot(patch: mine("bench")), "bench")
        medium(try WidgetFixtures.snapshot(patch: mine(nil)), "not-answered")
        medium(try WidgetFixtures.snapshot(patch: mine(nil, spec: nil)), "choose-spec")
        medium(try WidgetFixtures.snapshot(patch: ["status": .string("locked"), "can_sign_up": .bool(false)]), "locked")
        var failed = try WidgetFixtures.snapshot()
        failed.notSavedUntil = now.addingTimeInterval(5)
        failed.notSavedState = "out"
        medium(failed, "not-saved")
        medium(try WidgetFixtures.snapshot(), "in", large: true)
    }

    @Test("next upgrade, small: no gear yet")
    func upgrades() throws {
        let noGear = UpgradeGlance(try WidgetFixtures.snapshot("widget.get.no-guild"))
        snap(WidgetFrame(size: WidgetMetrics.small) { UpgradeSmallView(glance: noGear, icons: icons) }, "ios-upgrade-small-no-gear", size: WidgetMetrics.small)
    }

    @Test("lock screen: inline, rectangular, both circulars, the upgrade rectangle, three tints each")
    func lockScreen() throws {
        let snapshot = try WidgetFixtures.snapshot()
        let figures = WidgetFigures(snapshot, now: now, zone: zone)
        // A snapshot cannot run the system's tint pass: each lock widget is drawn white, vibrant and
        // accented as an approximation of what `widgetAccentable` marks.
        func lock(_ view: some View, _ name: String, _ size: CGSize) {
            snap(WidgetFrame(size: size, card: false) { view.foregroundStyle(.white).frame(width: size.width, height: size.height) },
                 "ios-lock-\(name)", size: size)
            snap(WidgetFrame(size: size, card: false) {
                view.environment(\.widgetRenderingMode, .vibrant).foregroundStyle(.white.opacity(0.8)).frame(width: size.width, height: size.height)
            }, "ios-lock-\(name)-vibrant", size: size, ground: Color(red: 0.32, green: 0.33, blue: 0.38))
            snap(WidgetFrame(size: size, card: false) {
                view.environment(\.widgetRenderingMode, .accented).foregroundStyle(Color(red: 0.55, green: 0.78, blue: 1.0)).frame(width: size.width, height: size.height)
            }, "ios-lock-\(name)-accented", size: size, ground: Color(red: 0.08, green: 0.12, blue: 0.22))
        }
        lock(RaidInlineView(glance: glance(snapshot)), "inline", WidgetMetrics.inline)
        lock(RaidRectangularView(glance: glance(snapshot)), "rectangular", WidgetMetrics.rectangular)
        lock(RaidSignedUpCircularView(figures: figures), "circular-signed-up", WidgetMetrics.circular)
        lock(RaidCountdownCircularView(figures: figures), "circular-countdown", WidgetMetrics.circular)
        lock(UpgradeRectangularView(glance: UpgradeGlance(snapshot)), "upgrade-rectangular", WidgetMetrics.rectangular)
    }

    @Test("the raid Live Activity: cancelled, unanswered without the counts, and the expanded Island at AX5")
    func raidLive() throws {
        let size = CGSize(width: 361, height: 120)
        let attributes = WidgetMockData.attributes
        var cancelled = WidgetMockData.live
        cancelled.status = "cancelled"
        // A raid whose counts the server has not sent: `21 of 40 signed up`.
        var plain = RaidActivityAttributes.ContentState(WidgetMockData.raid, myState: nil)
        plain.inCount = nil
        plain.late = nil
        plain.tanks = nil
        plain.healers = nil
        for (state, name) in [(cancelled, "cancelled"), (plain, "unanswered")] {
            snap(RaidLiveLockView(attributes: attributes, state: state, icons: icons, fixedNow: WidgetMockData.liveNow, zone: zone)
                .frame(width: size.width).background(FSColor.raised, in: .rect(cornerRadius: LiveMetrics.radius)),
                 "ios-live-raid-\(name)", size: size)
        }
        let island = CGSize(width: 361, height: 200)
        snap(RaidIslandExpandedView(attributes: attributes, state: WidgetMockData.live, icons: icons, fixedNow: WidgetMockData.liveNow)
            .padding(20).frame(width: island.width).background(.black, in: .rect(cornerRadius: 36)),
             "ios-island-expanded", size: island, large: true)
    }

    @Test("the sim Live Activity finished and failed")
    func simLive() {
        let attributes = SimActivityAttributes(simID: "k3v9q2m8x1ab", title: "Elemental Shaman, 60 raid")
        let size = CGSize(width: 361, height: 100)
        let states: [(String, SimActivityAttributes.ContentState)] = [
            ("finished", .init(state: "finished", done: 64, total: 64, etaSec: nil, dps: 412)),
            ("failed", .init(state: "failed", done: 0, total: nil, etaSec: nil, dps: nil)),
        ]
        for (name, state) in states {
            snap(SimLiveLockView(attributes: attributes, state: state).frame(width: size.width)
                .background(FSColor.raised, in: .rect(cornerRadius: LiveMetrics.radius)), "ios-live-sim-\(name)", size: size)
        }
    }

    @Test("Live surfaces settings")
    func settings() {
        let suite = "widgets-snap-\(UUID().uuidString)"
        let defaults = UserDefaults(suiteName: suite) ?? .standard
        let view = NavigationStack {
            ScrollView { LiveSurfaceSettingsView(settings: LiveSurfaceSettings(suiteName: suite), showsSim: true) }
                .background(FSColor.bg)
                .navigationTitle(Copy.liveSettings)
                .navigationBarTitleDisplayMode(.inline)
        }
        // Drawn in a key window so the switches render with their knobs and on state.
        snap(view, "ios-settings-live-surfaces", size: CGSize(width: 402, height: 360), keyWindow: true)
        snap(view, "ios-settings-live-surfaces-ax5", size: CGSize(width: 402, height: 600), large: true, keyWindow: true)
        defaults.removePersistentDomain(forName: suite)
    }

    @Test("the home screen on a light and a dark wallpaper")
    func wallpapers() {
        let snapshot = WidgetMockData.snapshot("guild")
        let raid = RaidGlance(snapshot, now: WidgetMockData.now, zone: zone)
        let upgrade = UpgradeGlance(snapshot)
        let screen = CGSize(width: 402, height: 420)
        func home(_ wallpaper: [Color]) -> some View {
            VStack(spacing: 18) {
                HStack(spacing: 22) {
                    WidgetFrame(size: WidgetMetrics.small) { NextRaidSmallView(glance: raid) }
                    WidgetFrame(size: WidgetMetrics.small) { UpgradeSmallView(glance: upgrade, icons: icons) }
                }
                WidgetFrame(size: WidgetMetrics.medium) { GlanceMediumView(raid: raid, upgrade: upgrade, icons: icons) }
            }
            .frame(width: screen.width, height: screen.height)
            .background(LinearGradient(colors: wallpaper, startPoint: .top, endPoint: .bottom))
        }
        snap(home([Color(red: 0.93, green: 0.87, blue: 0.78), Color(red: 0.72, green: 0.82, blue: 0.92)]), "ios-home-wallpaper-light", size: screen)
        snap(home([Color(red: 0.24, green: 0.15, blue: 0.08), Color(red: 0.05, green: 0.08, blue: 0.16)]), "ios-home-wallpaper-dark", size: screen)
    }
}
