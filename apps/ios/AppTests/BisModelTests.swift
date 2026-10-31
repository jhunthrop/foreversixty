@testable import ForeverSixty
import Foundation
import Models
import Networking
import Synchronization
import Testing

/// A data source answering from the contract fixtures, with switches for offline and errors.
private final class FakeBisSource: BisDataSource, Sendable {
    struct Behaviour {
        var bis: Result<CachedFile<BisFile>, APIError>
        var character: AccountCharacter?
        var savedBis: CachedFile<BisFile>?
    }

    private let state: Mutex<Behaviour>
    let items: ItemsFile
    let classes: [ClassInfo]

    init(_ behaviour: Behaviour) throws {
        state = Mutex(behaviour)
        items = try BisTestData.items()
        classes = try BisTestData.classes()
    }

    func set(_ change: (inout Behaviour) -> Void) { state.withLock { change(&$0) } }

    func activeBuild() async -> String { BisTestData.build }

    func bis(build _: String, spec _: String) async throws(APIError) -> CachedFile<BisFile> {
        try state.withLock { $0.bis }.get()
    }

    func items(build _: String, classSlug _: String) async throws(APIError) -> CachedFile<ItemsFile> {
        CachedFile(value: items, savedAt: BisTestData.now, isFromCache: false)
    }

    func classes(build _: String) async throws(APIError) -> CachedFile<[ClassInfo]> {
        CachedFile(value: classes, savedAt: BisTestData.now, isFromCache: false)
    }

    func cachedBis(build _: String, spec _: String) -> CachedFile<BisFile>? { state.withLock { $0.savedBis } }
    func cachedItems(build _: String, classSlug _: String) -> CachedFile<ItemsFile>? { nil }
    func cachedClasses(build _: String) -> CachedFile<[ClassInfo]>? { nil }
    func currentCharacter() async -> AccountCharacter? { state.withLock { $0.character } }
    func keyDates() async -> [KeyDate] { (try? BisTestData.fixture("cdn.dates")) ?? [] }
}

private final class MemoryPreferences: BisPreferences, @unchecked Sendable {
    var lastSelection: BisSelection?
    var hasOpenedPicker = false
}

@MainActor
@Suite("BiS model")
struct BisModelTests {
    private let file: BisFile
    private let main: AccountCharacter
    private let catalog: SpecCatalog

    init() throws {
        file = try BisTestData.bis()
        main = try #require(try BisTestData.me().mainCharacter)
        catalog = try BisTestData.catalog()
    }

    private func fresh() -> Result<CachedFile<BisFile>, APIError> {
        .success(CachedFile(value: file, savedAt: BisTestData.now, isFromCache: false))
    }

    @Test("signed in: the character's list, marked Yours, with the worn pick tagged")
    func signedIn() async throws {
        let source = try FakeBisSource(.init(bis: fresh(), character: main))
        let model = BisModel(source: source, catalog: catalog, preferences: MemoryPreferences())
        await model.start()
        #expect(model.context.isYours)
        #expect(model.context.selection.range == BisRange(band: 50, preset: "bare"))
        #expect(model.isPickerPresented == false)
        #expect(model.context.raidsOpening?.label == "Dec 9")
        guard case let .loaded(list, offlineSince) = model.phase else { Issue.record("not loaded"); return }
        #expect(offlineSince == nil)
        #expect(list.rows.count == 17)
    }

    @Test("signed out, first visit: Warrior preselected and the picker open, once")
    func signedOutFirstVisit() async throws {
        let source = try FakeBisSource(.init(bis: fresh(), character: nil))
        let preferences = MemoryPreferences()
        let model = BisModel(source: source, catalog: catalog, preferences: preferences)
        await model.start()
        #expect(model.isPickerPresented)
        #expect(model.context.selection.classSlug == "warrior")
        #expect(model.context.isYours == false)
        let again = BisModel(source: source, catalog: catalog, preferences: preferences)
        await again.start()
        #expect(again.isPickerPresented == false)
    }

    @Test("a choice applies at once and Back returns to the character's list")
    func choiceAndBack() async throws {
        let source = try FakeBisSource(.init(bis: fresh(), character: main))
        let model = BisModel(source: source, catalog: catalog, preferences: MemoryPreferences())
        await model.start()
        await model.select(model.context.selection.withRange(BisRange(band: 60, preset: "raid")))
        #expect(model.context.isYours == false)
        let back = try #require(model.context.backTarget)
        await model.select(back.selection)
        #expect(model.context.isYours)
    }

    @Test("offline with a saved copy: the list under the strip, stamped with the save time")
    func offlineSaved() async throws {
        let saved = Date(timeIntervalSince1970: 1_793_300_000)
        let source = try FakeBisSource(.init(
            bis: .success(CachedFile(value: file, savedAt: saved, isFromCache: true)),
            character: main
        ))
        let model = BisModel(source: source, catalog: catalog, preferences: MemoryPreferences())
        await model.start()
        guard case let .loaded(_, offlineSince) = model.phase else { Issue.record("not loaded"); return }
        #expect(offlineSince == saved)
    }

    @Test("offline and never fetched: the not-on-your-phone sentence")
    func offlineNeverFetched() async throws {
        let source = try FakeBisSource(.init(bis: .failure(.transport("offline")), character: main))
        let model = BisModel(source: source, catalog: catalog, preferences: MemoryPreferences())
        await model.start()
        #expect(model.phase == .notCached)
    }

    @Test("a CDN error is the error state, and Try again recovers")
    func errorAndRetry() async throws {
        let source = try FakeBisSource(.init(bis: .failure(.unexpectedStatus(500)), character: main))
        let model = BisModel(source: source, catalog: catalog, preferences: MemoryPreferences())
        await model.start()
        #expect(model.phase == .failed)
        source.set { $0.bis = self.fresh() }
        await model.refresh()
        guard case .loaded = model.phase else { Issue.record("not loaded after retry"); return }
    }

    @Test("a range the file does not publish says so")
    func noList() async throws {
        let source = try FakeBisSource(.init(bis: fresh(), character: main))
        let model = BisModel(source: source, catalog: catalog, preferences: MemoryPreferences())
        await model.start()
        await model.select(model.context.selection.withFaction("alliance"))
        #expect(model.phase == .noList(offlineSince: nil))
    }
}

@Suite("Offline stamp")
struct OfflineStampTests {
    private var calendar: Calendar {
        var calendar = Calendar(identifier: .gregorian)
        calendar.timeZone = TimeZone.gmt
        return calendar
    }

    @Test("today, this week, else the date", arguments: [
        (0.0, "9:41 pm"),
        (86_400.0 * 1, "Thu 9:41 pm"),
        (86_400.0 * 6, "Sat 9:41 pm"),
        (86_400.0 * 100, "22 Jul, 9:41 pm"),
    ])
    func stamp(ago: Double, expected: String) {
        let saved = BisTestData.now.addingTimeInterval(-ago)
        #expect(OfflineStamp.format(saved, now: BisTestData.now, calendar: calendar) == expected)
    }
}
