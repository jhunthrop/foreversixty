@testable import ForeverSixty
import DesignSystem
import Models
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// Every F14 surface and state (board F14.1 to F14.4), at the default size and at AX5 for the
/// medium widgets (spec §12). The widget views are the extension's own, drawn here in their
/// system sizes; lock-screen widgets on a grey ground in white, as the system tints them.
@MainActor
@Suite("Widget snapshots")
struct WidgetSnapshotTests {
    private let zone = WidgetTime.chicago
    private let now = WidgetTime.now
    private let icons = WidgetIcons { icon in
        let url = HomeFixtures.repoRoot.appending(path: "data/builds/\(SimsTestData.build)/icons/\(icon).webp")
        return (try? Data(contentsOf: url)).flatMap(UIImage.init(data:))
    }

    private func snap(_ view: some View, _ name: String, size: CGSize, large: Bool = false) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = large ? .accessibilityExtraExtraExtraLarge : .large
        }
        let framed = view
            .frame(width: size.width, height: size.height, alignment: .top)
            .environment(\.dynamicTypeSize, large ? .accessibility5 : .large)
            .preferredColorScheme(.dark)
            .padding(16)
            .background(Color(red: 0.16, green: 0.12, blue: 0.09))
        assertSnapshot(
            of: framed,
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .fixed(width: size.width + 32, height: size.height + 32), traits: traits),
            testName: name
        )
    }

    private func raid(_ snapshot: WidgetSnapshot?, at moment: Date? = nil) -> WidgetFigures {
        WidgetFigures(snapshot, now: moment ?? now, zone: zone)
    }

    private func lastRaid(_ snapshot: WidgetSnapshot?) -> String? {
        snapshot?.widget.lastRaid.map { Copy.widgetLastRaid(zone: $0.zoneName, date: WidgetClock(zone: zone).dayMonth($0.startsAt)) }
    }

    private func small(_ snapshot: WidgetSnapshot?, _ name: String, at moment: Date? = nil) {
        let view = WidgetFrame(size: WidgetMetrics.small) { NextRaidSmallView(figures: raid(snapshot, at: moment), lastRaid: lastRaid(snapshot)) }
        snap(view, "ios-raid-small-\(name)", size: WidgetMetrics.small)
    }

    private func medium(_ snapshot: WidgetSnapshot?, _ name: String, large: Bool = false) {
        let view = WidgetFrame(size: large ? WidgetMetrics.medium : WidgetMetrics.medium) {
            NextRaidMediumView(figures: raid(snapshot), lastRaid: lastRaid(snapshot), raid: snapshot?.widget.nextRaid, characterKey: nil)
        }
        snap(view, "ios-raid-medium-\(name)\(large ? "-ax5" : "")", size: large ? WidgetMetrics.medium : WidgetMetrics.medium, large: large)
    }

    private func mine(_ state: String?, spec: String? = "Elemental") -> [String: JSONValue] {
        ["my": .object(["state": state.map(JSONValue.string) ?? .null, "spec": spec.map(JSONValue.string) ?? .null, "role": .string("dps")])]
    }

    private func noRaid() throws -> WidgetSnapshot {
        let base = try WidgetFixtures.widget("widget.get")
        return WidgetSnapshot(widget: Widget(
            generatedAt: base.generatedAt, character: base.character, guild: base.guild, nextRaid: nil,
            lastRaid: base.lastRaid, upgrades: base.upgrades, live: base.live
        ), fetchedAt: now)
    }

    @Test("next raid, small: every state")
    func smallStates() throws {
        small(try WidgetFixtures.snapshot(), "in")
        small(try WidgetFixtures.snapshot(patch: mine("late")), "late")
        small(try WidgetFixtures.snapshot(patch: mine("bench")), "bench")
        small(try WidgetFixtures.snapshot(patch: mine("out")), "out")
        small(try WidgetFixtures.snapshot(patch: mine(nil)), "not-signed-up")
        small(try WidgetFixtures.snapshot(patch: mine(nil, spec: nil)), "choose-spec")
        small(try WidgetFixtures.snapshot(patch: ["status": .string("locked"), "can_sign_up": .bool(false)]), "locked")
        small(try WidgetFixtures.snapshot(patch: ["status": .string("cancelled"), "can_sign_up": .bool(false)]), "cancelled")
        small(try WidgetFixtures.snapshot(fetchedAt: now.addingTimeInterval(-14 * 3600)), "stale")
        small(try WidgetFixtures.snapshot("widget.get.no-guild"), "no-guild")
        small(try noRaid(), "no-raid")
        small(nil, "signed-out")
    }

    @Test("next raid, medium: each answer checked, the failure line, and AX5")
    func mediumStates() throws {
        medium(try WidgetFixtures.snapshot(), "in")
        medium(try WidgetFixtures.snapshot(patch: mine("late")), "late")
        medium(try WidgetFixtures.snapshot(patch: mine("out")), "out")
        medium(try WidgetFixtures.snapshot(patch: mine(nil)), "not-signed-up")
        medium(try WidgetFixtures.snapshot(patch: mine(nil, spec: nil)), "choose-spec")
        medium(try WidgetFixtures.snapshot(patch: ["status": .string("locked"), "can_sign_up": .bool(false)]), "locked")
        var failed = try WidgetFixtures.snapshot()
        failed.notSavedUntil = now.addingTimeInterval(5)
        failed.notSavedState = "out"
        medium(failed, "not-saved")
        medium(try WidgetFixtures.snapshot(), "in", large: true)
        medium(failed, "not-saved", large: true)
    }

    @Test("next upgrade: small, medium, AX5, no gear")
    func upgrades() throws {
        let snapshot = try WidgetFixtures.snapshot()
        let figures = UpgradeFigures(snapshot)
        let className = snapshot.widget.character?.className
        snap(WidgetFrame(size: WidgetMetrics.small) { UpgradeSmallView(figures: figures, className: className, icons: icons) },
             "ios-upgrade-small", size: WidgetMetrics.small)
        snap(WidgetFrame(size: WidgetMetrics.medium) { UpgradeMediumView(figures: figures, className: className, icons: icons) },
             "ios-upgrade-medium", size: WidgetMetrics.medium)
        let tall = WidgetMetrics.medium
        snap(WidgetFrame(size: tall) { UpgradeMediumView(figures: figures, className: className, icons: icons) },
             "ios-upgrade-medium-ax5", size: tall, large: true)
        let noGear = UpgradeFigures(try WidgetFixtures.snapshot("widget.get.no-guild"))
        snap(WidgetFrame(size: WidgetMetrics.small) { UpgradeSmallView(figures: noGear, className: "hunter", icons: icons) },
             "ios-upgrade-small-no-gear", size: WidgetMetrics.small)
    }

    @Test("lock screen: inline, rectangular, both circulars, the upgrade rectangle")
    func lockScreen() throws {
        let figures = raid(try WidgetFixtures.snapshot())
        func lock(_ view: some View, _ name: String, _ size: CGSize) {
            snap(WidgetFrame(size: size, card: false) { view.foregroundStyle(.white).frame(width: size.width, height: size.height) },
                 "ios-lock-\(name)", size: size)
        }
        lock(RaidInlineView(figures: figures), "inline", WidgetMetrics.inline)
        lock(RaidRectangularView(figures: figures), "rectangular", WidgetMetrics.rectangular)
        lock(RaidSignedUpCircularView(figures: figures), "circular-signed-up", WidgetMetrics.circular)
        lock(RaidCountdownCircularView(figures: figures), "circular-countdown", WidgetMetrics.circular)
        lock(UpgradeRectangularView(figures: UpgradeFigures(try WidgetFixtures.snapshot())), "upgrade-rectangular", WidgetMetrics.rectangular)
    }

    private func liveRaid(_ minutesOut: TimeInterval, status: String = "open") throws -> (RaidActivityAttributes, RaidActivityAttributes.ContentState, Date) {
        let raid = try #require(try WidgetFixtures.snapshot(patch: ["signed_up": .number(38), "status": .string(status)]).widget.nextRaid)
        return (RaidActivityAttributes(raid, characterName: "Lightningbolltz"), .init(raid, myState: "in"), raid.startsAt.addingTimeInterval(-minutesOut))
    }

    @Test("the raid Live Activity at 38 and 5 minutes, cancelled, and the Island")
    func raidLive() throws {
        let lockSize = CGSize(width: 370, height: 170)
        for (minutes, name) in [(38.0 * 60 + 12, "38-min"), (5.0 * 60, "5-min")] {
            let (attributes, state, moment) = try liveRaid(minutes)
            snap(RaidLiveLockView(attributes: attributes, state: state, fixedNow: moment, zone: zone).frame(width: lockSize.width)
                .background(FSColor.raised, in: .rect(cornerRadius: 22)), "ios-live-raid-\(name)", size: lockSize)
        }
        let (cancelledAttributes, cancelled, moment) = try liveRaid(20 * 60, status: "cancelled")
        snap(RaidLiveLockView(attributes: cancelledAttributes, state: cancelled, fixedNow: moment, zone: zone).frame(width: lockSize.width)
            .background(FSColor.raised, in: .rect(cornerRadius: 22)), "ios-live-raid-cancelled", size: lockSize)

        let (attributes, state, at38) = try liveRaid(38 * 60 + 12)
        let island = CGSize(width: 370, height: 170)
        snap(RaidIslandExpandedView(attributes: attributes, state: state, fixedNow: at38, zone: zone).padding(20).frame(width: island.width)
            .background(.black, in: .rect(cornerRadius: 44)), "ios-island-expanded", size: island)
        let compact = CGSize(width: 250, height: 37)
        snap(HStack { SealMark(); Spacer(); LiveCountdown(target: state.targetDate, fixedNow: at38, font: .system(.footnote, design: .monospaced, weight: .semibold)) }
            .padding(.horizontal, 12).frame(width: compact.width, height: compact.height)
            .background(.black, in: .capsule), "ios-island-compact", size: compact)
    }

    @Test("the sim Live Activity running, finished and failed")
    func simLive() {
        let attributes = SimActivityAttributes(simID: "k3v9q2m8x1ab", title: "Elemental · Top Gear")
        let size = CGSize(width: 370, height: 120)
        let states: [(String, SimActivityAttributes.ContentState)] = [
            ("running", .init(state: "running", done: 40, total: 64, etaSec: 120, dps: nil)),
            ("finished", .init(state: "finished", done: 64, total: 64, etaSec: nil, dps: 412)),
            ("failed", .init(state: "failed", done: 0, total: nil, etaSec: nil, dps: nil)),
        ]
        for (name, state) in states {
            snap(SimLiveLockView(attributes: attributes, state: state).frame(width: size.width)
                .background(FSColor.raised, in: .rect(cornerRadius: 22)), "ios-live-sim-\(name)", size: size)
        }
    }

    @Test("You › Live surfaces")
    func settings() {
        let suite = "widgets-snap-\(UUID().uuidString)"
        let defaults = UserDefaults(suiteName: suite) ?? .standard
        let view = NavigationStack {
            ScrollView { LiveSurfaceSettingsView(settings: LiveSurfaceSettings(suiteName: suite), showsSim: true) }
                .background(FSColor.bg)
                .navigationTitle(Copy.liveSettings)
                .navigationBarTitleDisplayMode(.inline)
        }
        snap(view, "ios-settings-live-surfaces", size: CGSize(width: 402, height: 360))
        snap(view, "ios-settings-live-surfaces-ax5", size: CGSize(width: 402, height: 600), large: true)
        defaults.removePersistentDomain(forName: suite)
    }
}
