@testable import ForeverSixty
import DesignSystem
import Models
import Networking
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// F15's surfaces on iPhone 17 (402 x 874) at the default size and at AX5, from the spec's one
/// story (Lightningbolltz, Onyxia's Lair, Thursday 8 pm, 25 of 40). The PNGs double as the review
/// captures in design/reviews/2026-11-01-mobile-integration/build/ios. Siri's sheet, Spotlight,
/// the alert and the share sheet are the system's: these draw our own words in stand-in frames.
@MainActor
@Suite("System integration snapshots")
struct SystemSnapshotTests {
    private let now = RaidTime.date("2026-12-09T03:41:00Z")
    private let answers = SystemStory.answers

    private func snap(_ view: some View, _ name: String, large: Bool = false) {
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
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: config), traits: traits),
            testName: name
        )
    }

    private func both(_ view: some View, _name: String) {
        snap(view, "\(_name)-ios")
        snap(view, "\(_name)-large-ios", large: true)
    }

    // MARK: F15.1, F15.2 Siri and Shortcuts

    private func siri(_ said: String, _ card: IntentAnswerCard) -> some View {
        VStack(alignment: .leading, spacing: 16) {
            Text(said).font(ScaledType.lead).foregroundStyle(FSColor.strong)
            card
            Spacer(minLength: 0)
        }
        .padding(16)
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
        .background(FSColor.bg)
    }

    @Test("F15.1 the voice confirmation, the answer and the refusal")
    func signup() throws {
        let start = RaidTime.date(SystemStory.thursday)
        let confirm = answers.confirm(name: "Lightningbolltz", state: "in", zone: "Onyxia's Lair", startsAt: start)
        both(siri("Sign me up for tonight's raid", IntentAnswerCard(text: confirm, buttons: (Copy.intentConfirmAction, Copy.commonCancel))), _name: "f15-siri-confirm")
        let saved = answers.signedUp(state: "in", raid: try SystemStory.raid(), characterName: "Lightningbolltz", role: "dps")
        both(siri("Sign me up for tonight's raid", IntentAnswerCard(text: saved)), _name: "f15-siri-signedup")
        both(siri("Sign me up for tonight's raid", IntentAnswerCard(text: Copy.raidLocked)), _name: "f15-siri-refused")
    }

    @Test("F15.2 next upgrade, next raid, who is coming")
    func reads() throws {
        let home = try HomeFixtures.home("character.home.yours")
        let upgrade = try #require(UpgradeAnswer(home: home, characterName: "Lightningbolltz"))
        let iconURL = HomeFixtures.repoRoot.appending(path: "data/builds/\(SimsTestData.build)/icons/\(upgrade.icon ?? "").webp")
        let image = (try? Data(contentsOf: iconURL)).flatMap(UIImage.init(data:))
        both(siri("What's my next upgrade", IntentAnswerCard(text: upgrade.sentence, snippet: upgrade, snippetImage: image)), _name: "f15-siri-upgrade")
        let next = try #require(SystemStory.widget().nextRaid)
        both(siri("When is the next raid", IntentAnswerCard(text: answers.nextRaid(next))), _name: "f15-siri-nextraid")
        let raid = try SystemStory.raid(counts: ["in": 23, "late": 1, "tanks": 3, "healers": 6, "dps": 14])
        both(siri("Who is coming to raid", IntentAnswerCard(text: answers.whoIsComing(raid))), _name: "f15-siri-whoiscoming")
    }

    // MARK: F15.3 Spotlight

    @Test("F15.3 the Spotlight entries")
    func spotlight() throws {
        let me = try RaidFixture.load(Me.self, "me")
        let raids = try [SystemStory.raid(id: 7), SystemStory.raid(id: 8, startsAt: "2026-12-18T02:00:00Z")]
        let spec = try #require(SpecCatalog.bundled().spec("shaman-elemental"))
        let catalog = SpotlightCatalog.build(
            own: me.characters, watched: [], planned: raids, guildID: 2,
            guides: [.init(spec: spec, savedAt: RaidTime.date("2026-09-24T15:00:00Z"))], answers: answers, now: SystemStory.now
        )
        both(SpotlightPreview(catalog: catalog), _name: "f15-spotlight-entries")
    }

    // MARK: F15.4 Offline data

    private func offlineData(_ report: OfflineReport, online: Bool = true, update: OfflineUpdateState = .idle) -> some View {
        NavigationStack {
            OfflineDataView(
                report: report, update: update, online: online, checkedAt: now, words: OfflineWords(calendar: chicago), now: now,
                keepClassesReady: .constant(true), wifiOnly: .constant(true), updateNow: {}, remove: {}
            )
        }
    }

    private var chicago: Calendar {
        var calendar = Calendar(identifier: .gregorian)
        calendar.timeZone = SystemStory.zone
        return calendar
    }

    private var savedReport: OfflineReport {
        let build = "1.60.1.70707"
        let saved = now.addingTimeInterval(-600)
        return OfflineReport.build(files: [
            SavedStaticFile(relativePath: "\(build)/guides/shaman-elemental.json", bytes: 1_300_000, savedAt: saved),
            SavedStaticFile(relativePath: "\(build)/guides/warrior-fury.json", bytes: 1_100_000, savedAt: saved),
            SavedStaticFile(relativePath: "\(build)/bis/shaman-elemental.json", bytes: 3_100_000, savedAt: saved),
            SavedStaticFile(relativePath: "\(build)/talents/shaman.json", bytes: 600_000, savedAt: saved),
            SavedStaticFile(relativePath: "\(build)/search.json", bytes: 400_000, savedAt: saved),
        ], iconBytes: 41_000_000, iconsUpdated: saved)
    }

    @Test("F15.4 Offline data: with data, empty, offline, and the remove confirmation")
    func offlineDataScreen() {
        both(offlineData(savedReport), _name: "f15-offlinedata-data")
        both(offlineData(OfflineReport(rows: [])), _name: "f15-offlinedata-empty")
        snap(offlineData(savedReport, online: false), "f15-offlinedata-offline-ios")
        snap(offlineData(savedReport, update: .upToDate), "f15-offlinedata-uptodate-ios")
        both(AlertPreview(message: Copy.offlineDataRemoveConfirm, action: Copy.offlineDataRemoveAction), _name: "f15-offlinedata-remove")
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
        both(try homeStrip(offline: "Thu 9:41 pm"), _name: "f15-strip-offline")
        both(try homeStrip(stale: "9:41 pm"), _name: "f15-strip-stale")
        let guide = try GuidesTestData.elemental()
        let reader = NavigationStack {
            GuideReaderView(
                phase: .loaded(GuideContent(guide: guide, build: GuidesTestData.build), offlineSince: nil),
                newData: .ready("1.60.1.70800"), now: GuidesTestData.now
            )
            .background(FSColor.bg)
        }
        both(reader, _name: "f15-strip-newdata")
    }

    // MARK: Share

    @Test("the share sheet text: a character, a guide, a report")
    func shares() throws {
        let character = ShareText.character(key: SystemStory.key, name: "Lightningbolltz", spec: "Elemental", className: "shaman", level: 60)
        snap(SharePreview(item: character), "f15-share-character-ios")
        snap(SharePreview(item: ShareText.guide(try GuidesTestData.elemental())), "f15-share-guide-ios")
        let report = try #require(ShareText.report(
            id: "ewlrx3enxbeb", zone: "Onyxia's Lair", visibility: "public", createdAt: RaidTime.date("2026-12-15T01:58:00Z"),
            zoneOfPhone: SystemStory.zone
        ))
        snap(SharePreview(item: report), "f15-share-report-ios")
        snap(SharePreview(item: ShareText.tiers(.healer)), "f15-share-tiers-ios")
    }

    // MARK: Follow-ups

    @Test("Search with the Watching group and the Rankings link")
    func searchWatching() throws {
        let context = SearchTestData.context
        let watching = [
            CharacterSearchRow(key: "us/pvp/morvane", region: "us", ruleset: "pvp", name: "Morvane", className: "warlock", level: 60),
            CharacterSearchRow(key: "us/pvp/kraggor", region: "us", ruleset: "pvp", name: "Kraggor", className: "warrior", level: 60),
        ]
        let corpus = try SearchTestData.corpus()
        let forCharacter = SearchForCharacter(context, specs: try BisTestData.catalog())
        let view = SearchChromeHost {
            SearchContentView(
                hasQuery: false, query: "", scope: .all, index: .ready(corpus), results: .empty,
                characters: .idle, recent: [], forCharacter: forCharacter,
                viewer: context.viewer, build: SearchTestData.build, watching: watching, nativeRankings: true
            )
        }
        both(view, _name: "f15-search-watching")
    }

    @Test("the Guides root bar with the search button")
    func guidesBar() throws {
        let statics = StaticDataStore(client: StaticDataClient(transport: FixtureTransport(directory: HomeFixtures.directory)), directory: nil)
        let navigator = SearchNavigator(dependencies: SearchDependencies(
            index: SearchIndexReader(store: statics), catalog: ItemCatalog(store: statics), specs: try BisTestData.catalog(),
            characters: FakeCharacterSearch(.success([])), recent: MemoryRecentSearches(),
            build: { SearchTestData.build }, context: { SearchTestData.context }, sleep: { _ in }
        ))
        let index = try GuidesTestData.index()
        let root = NavigationStack {
            GuidesRootView(
                segment: .guides, tiersOn: true,
                homePhase: .loaded(GuidesHome(build: GuidesTestData.build, index: index, saved: GuidesTestData.saved()), offlineSince: nil),
                character: nil, now: GuidesTestData.now
            )
            .navigationTitle(Copy.guidesTitle).navigationBarTitleDisplayMode(.inline).toolbar(removing: .title)
            .searchButton()
        }
        .environment(\.searchNavigator, navigator)
        snap(root, "f15-guides-searchbutton-ios")
    }
}

/// Spotlight's list as the system draws it under the app's name: our own rows' words.
private struct SpotlightPreview: View {
    let catalog: SpotlightCatalog

    var body: some View {
        VStack(alignment: .leading, spacing: 10) {
            Text(Copy.appName).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            ForEach(catalog.characters.prefix(2), id: \.id) { row($0, icon: "person.crop.circle") }
            ForEach(catalog.guides, id: \.id) { row($0, icon: "book") }
            ForEach(catalog.raids, id: \.id) { row($0, icon: "calendar") }
            Spacer(minLength: 0)
        }
        .padding(16)
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
        .background(FSColor.bg)
    }

    private func row(_ row: SpotlightRow, icon: String) -> some View {
        HStack(spacing: 12) {
            Image(systemName: icon).foregroundStyle(FSColor.gold).frame(width: 36, height: 36)
                .background(FSColor.raised, in: .rect(cornerRadius: 8))
            VStack(alignment: .leading, spacing: 2) {
                Text(row.title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                Text(row.subtitle).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
            .fixedSize(horizontal: false, vertical: true)
        }
    }
}

/// The share sheet's header as the system draws it: the message over the link.
private struct SharePreview: View {
    let item: ShareItem

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            Text(item.text).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
            Text(item.url.absoluteString).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            Spacer(minLength: 0)
        }
        .padding(16)
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
        .background(FSColor.bg)
    }
}

/// The remove confirmation's words, in a system alert's frame.
private struct AlertPreview: View {
    let message: String
    let action: String

    var body: some View {
        VStack(spacing: 14) {
            Text(message).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong).multilineTextAlignment(.center)
                .fixedSize(horizontal: false, vertical: true)
            HStack(spacing: 10) {
                Text(Copy.commonCancel).frame(maxWidth: .infinity, minHeight: 44).background(FSColor.line, in: .capsule)
                Text(action).frame(maxWidth: .infinity, minHeight: 44).background(FSColor.line, in: .capsule).foregroundStyle(FSColor.wipe)
            }
            .font(ScaledType.bodyStrong)
        }
        .padding(20)
        .background(FSColor.raised, in: .rect(cornerRadius: 28))
        .padding(32)
        .frame(maxWidth: .infinity, maxHeight: .infinity)
        .background(Color.black.opacity(0.6))
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
