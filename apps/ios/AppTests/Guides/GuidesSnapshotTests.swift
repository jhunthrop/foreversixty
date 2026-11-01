@testable import ForeverSixty
import DesignSystem
import Models
import Networking
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// Every F9 screen and state on iPhone 17 (402 x 874), at the default size and at AX5, from the
/// W3-C1 fixtures. Recorded on the shared iOS 26.5 simulator like ReportsSnapshotTests; the PNGs
/// double as the review captures in design/reviews/2026-11-01-mobile-guides/build/ios.
/// The app draws dark only (ForeverSixtyApp forces `.dark`), so there is no light capture.
@MainActor
@Suite("Guides snapshots")
struct GuidesSnapshotTests {
    private let now = GuidesTestData.now
    private let build = GuidesTestData.build
    private var offlineSince: Date { now.addingTimeInterval(-2 * 3600) }

    private func snap(_ view: some View, _ name: String, height: CGFloat = 874, large: Bool = false) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = large ? .accessibilityExtraExtraExtraLarge : .large
        }
        let config = ViewImageConfig(safeArea: SimsSnapshotTests.safeArea, size: CGSize(width: SimsSnapshotTests.phone.width, height: height), traits: traits)
        let sized = view
            .environment(\.itemIconSource, RepoItemIconSource())
            .environment(\.dynamicTypeSize, large ? .accessibility5 : .large)
            .preferredColorScheme(.dark)
        assertSnapshot(
            of: sized,
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: config), traits: traits),
            testName: name
        )
    }

    /// The tab's root in its stack, as GuidesScreen draws it.
    private func root(_ view: GuidesRootView) -> some View {
        NavigationStack {
            view.navigationTitle(Copy.guidesTitle).navigationBarTitleDisplayMode(.inline).toolbar(removing: .title)
        }
    }

    /// A screen pushed over the Guides root, so the back control shows.
    private func pushed(_ guide: GuideFile? = nil, @ViewBuilder _ content: @escaping () -> some View) -> some View {
        NavigationStack(path: .constant([1])) {
            Color.clear.navigationTitle(Copy.guidesTitle).navigationDestination(for: Int.self) { _ in
                content().background(FSColor.bg).guideBar(guide)
            }
        }
    }

    private func home(saved: SavedGuides = GuidesTestData.saved(), offline: Date? = nil) throws -> StaticPhase<GuidesHome> {
        .loaded(GuidesHome(build: build, index: try GuidesTestData.index(), saved: saved), offlineSince: offline)
    }

    private func guidesRoot(_ phase: StaticPhase<GuidesHome>, character: AccountCharacter?) -> GuidesRootView {
        GuidesRootView(segment: .guides, tiersOn: true, homePhase: phase, character: character, now: now)
    }

    private func tiersRoot(_ choice: TierChoice, character: AccountCharacter?, phase: StaticPhase<TiersFile>? = nil) throws -> GuidesRootView {
        GuidesRootView(
            segment: .tiers, tiersOn: true, homePhase: try home(), character: character,
            tiersPhase: try phase ?? .loaded(GuidesTestData.tiers(), offlineSince: nil), choice: choice,
            yourSpec: TiersLogic.specKey(character, specs: SpecCatalog.bundled()), now: now
        )
    }

    private func guide(_ file: GuideFile, newData: NewDataState = .none) -> some View {
        pushed(file) {
            GuideReaderView(phase: .loaded(GuideContent(guide: file, build: build), offlineSince: nil), newData: newData, now: now)
        }
    }

    /// One section's region as the reader shows it scrolled there: the strip with that section
    /// current, then the section.
    private func region(_ file: GuideFile, _ section: GuideSection, @ViewBuilder _ content: () -> some View) -> some View {
        VStack(alignment: .leading, spacing: 0) {
            GuideSectionStrip(sections: GuideSection.of(file), current: section) { _ in }
            VStack(alignment: .leading, spacing: 0) {
                GuideSectionLabel(text: section.title)
                content()
            }
            .padding(.horizontal, GuidesLayout.gutter)
            Spacer(minLength: 0)
        }
        .frame(maxHeight: .infinity, alignment: .top)
        .background(FSColor.bg)
    }

    // MARK: F9.1 the tab

    @Test func guides_withCharacter() throws {
        snap(root(guidesRoot(try home(), character: GuidesTestData.lightningbolltz)), "f9-guides-ios")
    }

    @Test func guides_withCharacter_large() throws {
        snap(root(guidesRoot(try home(), character: GuidesTestData.lightningbolltz)), "f9-guides-large-ios", height: 2600, large: true)
    }

    @Test func guides_noCharacter() throws {
        snap(root(guidesRoot(try home(saved: .none), character: nil)), "f9-guides-nochar-ios")
    }

    @Test func guides_loading() {
        snap(root(guidesRoot(.loading, character: GuidesTestData.lightningbolltz)), "f9-loading-ios")
    }

    @Test func guides_error() {
        snap(root(guidesRoot(.failed, character: GuidesTestData.lightningbolltz)), "f9-guides-error-ios")
    }

    // MARK: F9.6 a class page

    @Test func classPage_offlineWithOneUnsaved() throws {
        let loaded = GuidesHome(build: build, index: try GuidesTestData.index(), saved: GuidesTestData.saved())
        snap(pushed {
            ScrollView { GuideClassView(classSlug: "shaman", home: loaded, offlineSince: offlineSince, now: now) { _ in } }
        }, "f9-class-offline-ios")
    }

    @Test func classPage_large() throws {
        let loaded = GuidesHome(build: build, index: try GuidesTestData.index(), saved: GuidesTestData.saved())
        snap(pushed {
            ScrollView { GuideClassView(classSlug: "shaman", home: loaded, offlineSince: nil, now: now) { _ in } }
        }, "f9-class-large-ios", height: 1400, large: true)
    }

    @Test func guide_notCached() {
        snap(pushed { GuideReaderView(phase: .notCached, now: now) }, "f9-guide-notcached-ios")
    }

    // MARK: F9.2 a guide

    @Test func guide_top() throws {
        snap(guide(try GuidesTestData.elemental()), "f9-guide-top-ios")
    }

    @Test func guide_whole() throws {
        snap(guide(try GuidesTestData.elemental()), "f9-guide-whole-ios", height: 5200)
    }

    @Test func guide_large() throws {
        snap(guide(try GuidesTestData.elemental()), "f9-guide-large-ios", height: 2400, large: true)
    }

    @Test func guide_tank() throws {
        snap(guide(try GuidesTestData.protection()), "f9-guide-protection-ios", height: 5600)
    }

    @Test func guide_talents() throws {
        let file = try GuidesTestData.elemental()
        snap(region(file, .talents) { GuideTalentsSection(talents: file.talents, build: build) }, "f9-guide-talents-ios", height: 1300)
    }

    @Test func guide_talents_large() throws {
        let file = try GuidesTestData.elemental()
        snap(region(file, .talents) { GuideTalentsSection(talents: file.talents, build: build) }, "f9-guide-talents-large-ios", height: 3600, large: true)
    }

    @Test func guide_rotation() throws {
        let file = try GuidesTestData.elemental()
        snap(region(file, .rotation) { GuideRotationSection(rotation: file.rotation, build: build) }, "f9-guide-rotation-ios", height: 1500)
    }

    @Test func guide_otherSections() throws {
        let file = try GuidesTestData.fire()
        snap(region(file, .races) {
            VStack(alignment: .leading, spacing: 0) {
                GuideStatsCard(guide: file)
                GuideSectionLabel(text: GuideSection.gear.title)
                GuideGearRow(guide: file) { _ in }
                GuideSectionLabel(text: GuideSection.races.title)
                GuideRacesCard(races: file.races)
                GuideSources(sources: file.sources).padding(.top, GuidesLayout.sectionGap)
            }
        }, "f9-guide-sections-ios", height: 1400)
    }

    @Test func guide_newData() throws {
        snap(guide(try GuidesTestData.elemental(), newData: .ready("1.60.1.70800")), "f9-guide-newdata-ios")
    }

    @Test func guide_updateFailed() throws {
        snap(guide(try GuidesTestData.elemental(), newData: .failed("1.60.1.70800")), "f9-guide-updatefailed-ios")
    }

    @Test func guide_loading() {
        snap(pushed { GuideReaderView(phase: .loading, now: now) }, "f9-guide-loading-ios")
    }

    @Test func guide_error() {
        snap(pushed { GuideReaderView(phase: .failed, now: now) }, "f9-guide-error-ios")
    }

    @Test func guide_offline() throws {
        let file = try GuidesTestData.elemental()
        snap(pushed(file) {
            GuideReaderView(phase: .loaded(GuideContent(guide: file, build: build), offlineSince: offlineSince), now: now)
        }, "f9-guide-offline-ios")
    }

    // MARK: F9.4 tier lists

    @Test func tiers_dps() throws {
        snap(root(try tiersRoot(TierChoice(role: .dps, faction: .horde), character: GuidesTestData.emberquill)), "f9-tiers-dps-ios", height: 2200)
    }

    @Test func tiers_first() throws {
        snap(root(try tiersRoot(TierChoice(role: .dps, faction: .horde), character: GuidesTestData.obnoxiousYell)), "f9-tiers-first-ios")
    }

    @Test func tiers_tank() throws {
        snap(root(try tiersRoot(TierChoice(role: .tank, faction: .horde), character: GuidesTestData.tank)), "f9-tiers-tank-ios")
    }

    @Test func tiers_healer() throws {
        snap(root(try tiersRoot(TierChoice(role: .healer, faction: .horde), character: GuidesTestData.healer)), "f9-tiers-healer-ios")
    }

    @Test func tiers_noCharacter() throws {
        snap(root(try tiersRoot(.fallback, character: nil)), "f9-tiers-nochar-ios")
    }

    @Test func tiers_large() throws {
        snap(root(try tiersRoot(TierChoice(role: .dps, faction: .horde), character: GuidesTestData.emberquill)), "f9-tiers-large-ios", height: 4200, large: true)
    }

    @Test func tiers_loading() throws {
        snap(root(try tiersRoot(.fallback, character: GuidesTestData.emberquill, phase: .loading)), "f9-tiers-loading-ios")
    }

    @Test func tiers_error() throws {
        snap(root(try tiersRoot(.fallback, character: GuidesTestData.emberquill, phase: .failed)), "f9-tiers-error-ios")
    }

    // MARK: F9.5 what changed

    private func changesRoot(_ changelog: StaticPhase<Changelog>) throws -> GuidesRootView {
        GuidesRootView(
            segment: .changes, tiersOn: true, homePhase: try home(), character: GuidesTestData.lightningbolltz,
            dates: KeyDatesLogic.rows(try GuidesTestData.dates(), now: now), changelog: changelog, now: now
        )
    }

    @Test func changes() throws {
        snap(root(try changesRoot(.loaded(try GuidesTestData.changelog(), offlineSince: nil))), "f9-changes-ios", height: 1400)
    }

    @Test func changes_large() throws {
        snap(root(try changesRoot(.loaded(try GuidesTestData.changelog(), offlineSince: nil))), "f9-changes-large-ios", height: 4000, large: true)
    }

    @Test func changes_error() throws {
        snap(root(try changesRoot(.failed)), "f9-changes-error-ios")
    }

    @Test func changes_entry() throws {
        let entry = try #require(try GuidesTestData.changelog().entries.first)
        snap(pushed { ScrollView { ChangelogEntryView(entry: entry) } }, "f9-changes-entry-ios")
    }

    @Test func guides_tiersOff_hasTwoSegments() throws {
        let view = GuidesRootView(segment: .guides, tiersOn: false, homePhase: try home(), character: GuidesTestData.lightningbolltz, now: now)
        #expect(GuidesSegment.available(tiersOn: false) == [.guides, .changes])
        snap(root(view), "f9-guides-tiersoff-ios")
    }
}
