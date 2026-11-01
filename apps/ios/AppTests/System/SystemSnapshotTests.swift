@testable import ForeverSixty
import DesignSystem
import Models
import Networking
import Routing
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// F15's surfaces on iPhone 17 (402 x 874) at the default size and at AX5, from the spec's one
/// story (Lightningbolltz, Onyxia's Lair, Thursday 8 pm, 25 of 40). The PNGs double as the review
/// captures in design/reviews/2026-11-01-mobile-integration/build/ios. Siri's sheet, Spotlight,
/// Shortcuts and Control Center are the system's: these draw our own words in stand-in frames
/// (the real share sheet, alert and Guides bar are captured from the simulator by
/// UITests/SystemCaptureTests).
@MainActor
@Suite("System integration snapshots")
struct SystemSnapshotTests {
    private let now = RaidTime.date("2026-12-09T03:41:00Z")
    private let answers = SystemStory.answers

    /// `keyWindow` draws through the window, so UIKit controls (the switches' knobs) render.
    private func snap(_ view: some View, _ name: String, large: Bool = false, keyWindow: Bool = false) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = large ? .accessibilityExtraExtraExtraLarge : .large
        }
        let config = ViewImageConfig(safeArea: SimsSnapshotTests.safeArea, size: SimsSnapshotTests.phone, traits: traits)
        let sized = view
            .environment(\.itemIconSource, RepoItemIconSource())
            .environment(\.dynamicTypeSize, large ? .accessibility5 : .large)
            .preferredColorScheme(.dark)
        assertSnapshot(
            of: sized,
            as: .image(drawHierarchyInKeyWindow: keyWindow, precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: config), traits: traits),
            testName: name
        )
    }

    private func both(_ view: some View, _ name: String, keyWindow: Bool = false) {
        snap(view, "\(name)-ios", keyWindow: keyWindow)
        snap(view, "\(name)-large-ios", large: true, keyWindow: keyWindow)
    }

    // MARK: F15.1, F15.2 Siri, Shortcuts and the Control

    private func siri(_ said: String, _ card: IntentAnswerCard) -> some View {
        VStack(alignment: .leading, spacing: 16) {
            Text(said).font(ScaledType.lead).foregroundStyle(FSColor.strong).fixedSize(horizontal: false, vertical: true)
            card
            Spacer(minLength: 0)
        }
        .padding(16)
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
        .background(FSColor.bg)
    }

    @Test("F15.1 the voice confirmation, the answer, the refusal and the two questions")
    func signup() throws {
        let start = RaidTime.date(SystemStory.thursday)
        let said = "Sign me up for tonight's raid"
        let confirm = answers.confirm(name: "Lightningbolltz", state: "in", zone: "Onyxia's Lair", startsAt: start)
        both(siri(said, IntentAnswerCard(text: confirm, buttons: (Copy.intentConfirmAction, Copy.commonCancel))), "f15-siri-confirm")
        let saved = answers.signedUp(state: "in", raid: try SystemStory.raid(), characterName: "Lightningbolltz", role: "dps")
        both(siri(said, IntentAnswerCard(text: saved)), "f15-siri-signedup")
        both(siri(said, IntentAnswerCard(text: Copy.raidLocked)), "f15-siri-refused")
        let raids = [answers.raidTitle(zone: "Onyxia's Lair", startsAt: RaidTime.date("2026-12-09T02:00:00Z")),
                     answers.raidTitle(zone: "Molten Core", startsAt: RaidTime.date("2026-12-09T05:00:00Z"))]
        both(siri(said, IntentAnswerCard(text: Copy.intentWhich, options: raids)), "f15-siri-which")
        let spec = Copy.intentSpec(name: "Lightningbolltz", guild: "OLYMPUS XXVII")
        both(siri(said, IntentAnswerCard(text: spec, options: ["Elemental", "Enhancement", "Restoration"])), "f15-siri-spec")
    }

    @Test("F15.2 next upgrade, next raid, who is coming")
    func reads() throws {
        let upgrade = try #require(UpgradeAnswer(home: try SystemStory.home(), characterName: "Lightningbolltz"))
        let iconURL = HomeFixtures.repoRoot.appending(path: "data/builds/\(SimsTestData.build)/icons/\(upgrade.icon ?? "").webp")
        let image = (try? Data(contentsOf: iconURL)).flatMap(UIImage.init(data:))
        both(siri("What's my next upgrade", IntentAnswerCard(text: upgrade.sentence, snippet: upgrade, snippetImage: image)), "f15-siri-upgrade")
        let next = try #require(SystemStory.widget().nextRaid)
        both(siri("When is the next raid", IntentAnswerCard(text: answers.nextRaid(next))), "f15-siri-nextraid")
        both(siri("Who is coming to raid", IntentAnswerCard(text: answers.whoIsComing(try SystemStory.raid()))), "f15-siri-whoiscoming")
    }

    @Test("a Shortcuts run and the Control: no question, the answer at once")
    func shortcutAndControl() throws {
        let saved = answers.signedUp(state: "in", raid: try SystemStory.raid(), characterName: "Lightningbolltz", role: "dps")
        both(siri(Copy.intentSignupTitle, IntentAnswerCard(text: saved)), "f15-shortcuts-run")
        both(ControlPreview(), "f15-control")
    }

    // MARK: F15.3 Spotlight

    @Test("F15.3 the Spotlight entries")
    func spotlight() throws {
        let watched = try RankingsTestData.watching().rows
        let raids = try [SystemStory.raid(id: 7), SystemStory.raid(id: 8, startsAt: "2026-12-18T02:00:00Z")]
        let spec = try #require(SpecCatalog.bundled().spec("shaman-elemental"))
        let catalog = SpotlightCatalog.build(
            own: try SystemStory.ownCharacters(), watched: watched, planned: raids, guildID: 2,
            guides: [.init(spec: spec, updated: "2026-11-01")], answers: answers, now: SystemStory.now
        )
        both(SpotlightPreview(catalog: catalog), "f15-spotlight-entries")
    }

    // MARK: F15.4 Offline data

    private func offlineData(_ report: OfflineReport, online: Bool = true, update: OfflineUpdateState = .idle) -> some View {
        NavigationStack {
            OfflineDataView(
                report: report, update: update, online: online, checkedAt: now, words: OfflineWords(calendar: chicago), now: now,
                keepClassesReady: .constant(true), wifiOnly: .constant(true), updateNow: {}, remove: {}
            )
            .navigationTitle(Copy.offlineDataTitle).navigationBarTitleDisplayMode(.inline).toolbar(removing: .title)
        }
    }

    private var chicago: Calendar {
        var calendar = Calendar(identifier: .gregorian)
        calendar.timeZone = SystemStory.zone
        return calendar
    }

    @Test("F15.4 Offline data: with data, empty, offline, up to date, and the remove confirmation")
    func offlineDataScreen() {
        let saved = SystemStory.offlineReport(savedAt: now.addingTimeInterval(-600))
        both(offlineData(saved), "f15-offlinedata-data", keyWindow: true)
        both(offlineData(OfflineReport(rows: [])), "f15-offlinedata-empty", keyWindow: true)
        both(offlineData(saved, online: false), "f15-offlinedata-offline", keyWindow: true)
        snap(offlineData(saved, update: .upToDate), "f15-offlinedata-uptodate-ios", keyWindow: true)
        both(AlertPreview(message: Copy.offlineDataRemoveConfirm, action: Copy.offlineDataRemoveAction), "f15-offlinedata-remove")
    }

    // MARK: F15.5 Strips

    private func homeStrip(offline: String? = nil, stale: String? = nil) throws -> some View {
        let home = try HomeFixtures.home()
        return HomeCaptureFrame {
            HomeAnswerView(
                phase: .loaded(home, offlineSince: nil), presenter: HomeFixtures.presenter(),
                build: "1.60.1.70658", offlineStamp: offline, staleStamp: stale
            )
        }
    }

    @Test("F15.5 each strip on a representative screen")
    func strips() throws {
        both(try homeStrip(offline: "Thu 9:41 pm"), "f15-strip-offline")
        both(try homeStrip(stale: "9:41 pm"), "f15-strip-stale")
        let guide = try GuidesTestData.elemental()
        let reader = NavigationStack(path: .constant([1])) {
            Color.clear.navigationTitle(Copy.guidesTitle).navigationDestination(for: Int.self) { _ in
                GuideReaderView(
                    phase: .loaded(GuideContent(guide: guide, build: GuidesTestData.build), offlineSince: nil),
                    newData: .ready("1.60.1.70800"), now: GuidesTestData.now
                )
                .background(FSColor.bg)
                .guideBar(guide)
            }
        }
        both(reader, "f15-strip-newdata")
    }

    // MARK: Search

    @Test("Search with the Watching group, the Rankings link and the in-app guide")
    func searchWatching() throws {
        let context = SearchTestData.context
        let watching = try RankingsTestData.watching().rows.map(SearchWords.row)
        let corpus = try SearchTestData.corpus()
        let forCharacter = SearchForCharacter(context, specs: try BisTestData.catalog())
        let view = SearchChromeHost {
            SearchContentView(
                hasQuery: false, query: "", scope: .all, index: .ready(corpus), results: .empty,
                characters: .idle, recent: [], forCharacter: forCharacter,
                viewer: context.viewer, build: SearchTestData.build, watching: watching, nativeRankings: true, nativeGuides: true
            )
        }
        both(view, "f15-search-watching")
    }

    // MARK: Round 2: the planner follow-ups

    @Test("F11.4 Open in planner under Builds seen")
    func openInPlanner() throws {
        let profile = try RankingsTestData.lightningbolltz()
        let ref = try #require(CharacterRef(key: profile.character.key, name: profile.character.name))
        let page = NavigationStack {
            ScrollView {
                CharacterView(
                    ref: ref, phase: .loaded(CharacterContent(profile: profile, offlineSince: nil)), items: [:],
                    account: RankingsTestData.account(), specs: RankingsTestData.specs, build: RankingsTestData.build,
                    now: RankingsTestData.now, clock: RankingsTestData.clock, actions: CharacterActions(openPlanner: { _ in })
                )
            }
            .background(FSColor.bg)
            .defaultScrollAnchor(.bottom)
        }
        both(page, "f15-char-open-planner")
    }

    private func sharedContent(_ plannerFile: TalentsFile) -> SharedBuildContent {
        let index = PlannerIndex(plannerFile)
        let order = PlannerTestData.elementalOrder(index)
        let summary = PlannerBuildSummary(
            build: PlannerTestData.build(order: order), classSlug: "shaman", className: "Shaman",
            raceName: "Orc", raceSlug: "orc", totals: PlannerRules.pointsPerTree(index, order: order)
        )
        return SharedBuildContent(summary: summary, file: plannerFile)
    }

    private func viewer(_ content: SharedBuildContent, compare: Bool) -> some View {
        NavigationStack {
            SharedBuildView(phase: .loaded(content), tree: .constant(0), timeZone: .gmt, compare: compare ? {} : nil)
                .toolbarBackground(FSColor.bg, for: .navigationBar)
        }
    }

    @Test("F13.5 Compare with mine: shown with an export of the class, hidden without")
    func compareButton() throws {
        let content = sharedContent(try PlannerTestData.shaman())
        both(viewer(content, compare: true), "f15-shared-build-compare-button")
        snap(viewer(content, compare: false), "f15-shared-build-compare-button-noexport-ios")
    }

    @Test("F13.5 the compare with the shared-build wording")
    func compareScreen() throws {
        let file = try PlannerTestData.shaman()
        let content = sharedContent(file)
        let character = try SystemStory.exporter(dropping: 3, file: file, order: content.summary.build.pointOrder)
        let presenter = SharedBuildCompare.presenter(content: content, file: file, character: character)
        let screen = NavigationStack { SharedBuildCompareScreen(presenter: presenter, characterName: character.name) }
        both(screen, "f15-shared-build-compare")
    }
}

/// Spotlight's list as the system draws it under the app's name: our own rows' words, the ringed
/// class crest on every character, and the app's shortcuts.
private struct SpotlightPreview: View {
    let catalog: SpotlightCatalog

    private static let icon: CGFloat = 40

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 12) {
                Text(Copy.appName).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                ForEach(catalog.characters.prefix(3), id: \.id) { entry in
                    row(entry) { ClassCrest(className: entry.classSlug, size: Self.icon) }
                }
                ForEach(catalog.guides, id: \.id) { entry in row(entry) { tile("book") } }
                ForEach(catalog.raids, id: \.id) { entry in row(entry) { tile("calendar") } }
                Text("Shortcuts").font(ScaledType.meta).foregroundStyle(FSColor.muted).padding(.top, 8)
                ForEach(["Sign me up for tonight's raid", "What's my next upgrade"], id: \.self) { phrase in
                    row(SpotlightRow(id: phrase, title: phrase, subtitle: Copy.appName)) { tile("checkmark.circle") }
                }
            }
            .padding(16)
            .frame(maxWidth: .infinity, alignment: .topLeading)
        }
        .background(FSColor.bg)
    }

    private func tile(_ symbol: String) -> some View {
        Image(systemName: symbol).foregroundStyle(FSColor.gold).frame(width: Self.icon, height: Self.icon)
            .background(FSColor.raised, in: .rect(cornerRadius: 8))
    }

    private func row(_ entry: SpotlightRow, @ViewBuilder icon: () -> some View) -> some View {
        HStack(alignment: .top, spacing: 12) {
            icon().frame(width: Self.icon, height: Self.icon)
            VStack(alignment: .leading, spacing: 2) {
                Text(entry.title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                Text(entry.subtitle).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
            .fixedSize(horizontal: false, vertical: true)
        }
    }
}

/// The Control Center control as the system draws it: the glyph and `In for the next raid`.
private struct ControlPreview: View {
    var body: some View {
        VStack(spacing: 10) {
            Image(systemName: "checkmark.circle").font(.system(size: 30)).foregroundStyle(FSColor.gold)
                .frame(width: 72, height: 72).background(Color.white.opacity(0.12), in: .circle)
            Text(Copy.intentControlTitle).font(ScaledType.meta).foregroundStyle(FSColor.strong)
                .multilineTextAlignment(.center).fixedSize(horizontal: false, vertical: true)
            Spacer(minLength: 0)
        }
        .padding(.top, 60)
        .frame(maxWidth: .infinity, maxHeight: .infinity)
        .background(Color.black)
    }
}

/// The remove confirmation's words in a system alert's frame: side by side while they fit,
/// stacked with equal heights at accessibility sizes.
private struct AlertPreview: View {
    let message: String
    let action: String
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        VStack(spacing: 14) {
            Text(message).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong).multilineTextAlignment(.center)
                .fixedSize(horizontal: false, vertical: true)
            let layout = typeSize.isAccessibilitySize ? AnyLayout(VStackLayout(spacing: 10)) : AnyLayout(HStackLayout(spacing: 10))
            layout {
                button(action, color: FSColor.wipe)
                button(Copy.commonCancel, color: FSColor.gold)
            }
            .font(ScaledType.bodyStrong)
        }
        .padding(20)
        .background(FSColor.raised, in: .rect(cornerRadius: 28))
        .padding(24)
        .frame(maxWidth: .infinity, maxHeight: .infinity)
        .background(Color.black.opacity(0.6))
    }

    private func button(_ title: String, color: Color) -> some View {
        Text(title).foregroundStyle(color).multilineTextAlignment(.center).fixedSize(horizontal: false, vertical: true)
            .padding(.vertical, 8)
            .frame(maxWidth: .infinity, minHeight: 48)
            .background(FSColor.line, in: .capsule)
    }
}

/// Search's chrome around the content, as F12 draws it.
private struct SearchChromeHost<Content: View>: View {
    @State private var query = ""
    @State private var scope = SearchScope.all
    @FocusState private var focused: Bool
    @ViewBuilder let content: () -> Content

    var body: some View {
        SearchChrome(query: $query, scope: $scope, focused: $focused, cancel: {}, content: content)
    }
}
