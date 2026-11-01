@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

@Suite("Offline data: sizes, days and rows")
struct OfflineDataTests {
    private var calendar: Calendar {
        var calendar = Calendar(identifier: .gregorian)
        calendar.timeZone = SystemStory.zone
        return calendar
    }

    private var words: OfflineWords { OfflineWords(calendar: calendar) }
    /// Tuesday 8 December, 9:41 pm in Chicago.
    private let now = RaidTime.date("2026-12-09T03:41:00Z")
    private let build = "1.60.1.70707"

    @Test("sizes: one decimal under 10 MB, whole numbers above, KB below a megabyte")
    func sizes() {
        #expect(OfflineWords.size(2_400_000) == "2.4 MB")
        #expect(OfflineWords.size(410_000) == "410 KB")
        #expect(OfflineWords.size(41_000_000) == "41 MB")
        #expect(OfflineWords.size(9_960_000) == "10.0 MB")
        #expect(OfflineWords.size(0) == "0 KB")
        #expect(OfflineWords.size(120) == "1 KB")
    }

    @Test("days: today, yesterday, else the date")
    func days() {
        #expect(words.day(now.addingTimeInterval(-3600), now: now) == "today")
        #expect(words.day(now.addingTimeInterval(-86400), now: now) == "yesterday")
        #expect(words.day(RaidTime.date("2026-09-24T15:00:00Z"), now: now) == "24 Sep")
    }

    @Test("the footer")
    func footer() {
        #expect(words.footer(checkedAt: now, now: now)
            == "Checked today at 9:41 pm. Your characters, sign-ups and sims are not stored here; they load when you are online.")
        #expect(words.footer(checkedAt: nil, now: now).hasPrefix("Not checked yet."))
    }

    @Test("rows by kind with the saved classes; no build ids or file counts")
    func rows() {
        let saved = now.addingTimeInterval(-600)
        let files = [
            SavedStaticFile(relativePath: "\(build)/guides/shaman-elemental.json", bytes: 1_200_000, savedAt: saved),
            SavedStaticFile(relativePath: "\(build)/guides/warrior-fury.json", bytes: 1_100_000, savedAt: saved),
            SavedStaticFile(relativePath: "\(build)/guides/index.json", bytes: 100_000, savedAt: saved),
            SavedStaticFile(relativePath: "\(build)/bis/shaman-elemental.json", bytes: 3_100_000, savedAt: saved),
            SavedStaticFile(relativePath: "\(build)/talents/shaman.json", bytes: 600_000, savedAt: saved),
            SavedStaticFile(relativePath: "\(build)/search.json", bytes: 400_000, savedAt: saved),
            SavedStaticFile(relativePath: "active-build.json", bytes: 100, savedAt: saved),
        ]
        let report = OfflineReport.build(files: files, iconBytes: 41_000_000, iconsUpdated: saved)
        #expect(report.rows.map(\.category) == [.guides, .bis, .talents, .search, .icons])
        let lines = report.rows.map { words.meta($0, now: now) }
        #expect(lines == [
            "Shaman, Warrior · 2.4 MB · Updated today",
            "Shaman · 3.1 MB · Updated today",
            "Shaman · 600 KB · Updated today",
            "400 KB · Updated today",
            "41 MB",
        ])
        #expect(OfflineWords.size(report.totalBytes) == "48 MB")
        #expect(lines.allSatisfy { !$0.contains(build) })
    }

    @Test("nothing saved is the empty state")
    func empty() {
        #expect(OfflineReport.build(files: [], iconBytes: 0, iconsUpdated: nil).isEmpty)
    }

    @Test("the switches default on and persist")
    func settings() throws {
        let defaults = try #require(UserDefaults(suiteName: "offline-\(UUID().uuidString)"))
        let settings = OfflineSettings(defaults: defaults)
        #expect(settings.keepClassesReady)
        #expect(settings.wifiOnly)
        settings.wifiOnly = false
        #expect(!OfflineSettings(defaults: defaults).wifiOnly)
        #expect(settings.mayDownload(on: FixedConnectivity(isConnected: true, isExpensive: true)))
        settings.wifiOnly = true
        #expect(!settings.mayDownload(on: FixedConnectivity(isConnected: true, isExpensive: true)))
        #expect(settings.mayDownload(on: FixedConnectivity(isConnected: true, isExpensive: false)))
    }
}

@MainActor
@Suite("Offline data: the one store")
struct OfflineDataModelTests {
    private let directory = FileManager.default.temporaryDirectory.appending(path: "offline-\(UUID().uuidString)")

    private func model(_ store: StaticDataStore, latest: String?, removed: @escaping @Sendable () -> Void = {}) throws -> OfflineDataModel {
        let defaults = try #require(UserDefaults(suiteName: "offline-\(UUID().uuidString)"))
        return OfflineDataModel(
            store: store, icons: IconCacheAccess(bytes: { 0 }, lastSaved: { nil }, removeAll: removed),
            latestBuild: { latest }, settings: OfflineSettings(defaults: defaults),
            connectivity: FixedConnectivity(), now: { SystemStory.now }
        )
    }

    @Test("the store's files and the screen read one directory; Remove empties it")
    func oneStore() async throws {
        let store = StaticDataStore(client: StaticDataClient(transport: CountingTransport()), directory: directory)
        defer { try? FileManager.default.removeItem(at: directory) }
        _ = try await store.file(.getSearchIndex(build: SearchTestData.build))
        _ = try await store.files.load(.getGuide(build: SearchTestData.build, spec: "shaman-elemental"))
        let model = try model(store, latest: SearchTestData.build)
        await model.load()
        #expect(model.report.rows.map(\.category) == [.guides, .search])
        await model.updateNow()
        #expect(model.update == .upToDate)
        #expect(model.checkedAt == SystemStory.now)
        await model.removeAll()
        #expect(model.report.isEmpty)
        #expect(await store.sizeOnDisk() == 0)
    }

    @Test("Update now with no answer fails and keeps the files")
    func offlineUpdate() async throws {
        let store = StaticDataStore(client: StaticDataClient(transport: CountingTransport()), directory: directory)
        defer { try? FileManager.default.removeItem(at: directory) }
        _ = try await store.file(.getSearchIndex(build: SearchTestData.build))
        let model = try model(store, latest: nil)
        await model.updateNow()
        #expect(model.update == .failed)
        #expect(await store.sizeOnDisk() > 0)
    }
}
