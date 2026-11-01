import Foundation
import Models
import Networking
import Observation

/// The saved item icons, as Offline data reads and removes them.
struct IconCacheAccess: Sendable {
    let bytes: @Sendable () -> Int
    let lastSaved: @Sendable () -> Date?
    let removeAll: @Sendable () -> Void

    static var live: IconCacheAccess {
        let icons = RemoteItemIconSource.shared
        return IconCacheAccess(bytes: { icons.diskBytes }, lastSaved: { icons.lastSaved }, removeAll: { icons.removeAll() })
    }
}

/// What `Update now` reports under its button.
enum OfflineUpdateState: Equatable, Sendable {
    case idle
    case updating
    case upToDate
    case failed
}

/// You › Offline data (spec §5.7 F15.4): the saved files by category, the two switches, Update
/// now and Remove downloaded data, all over the one static data store.
@MainActor
@Observable
final class OfflineDataModel {
    private(set) var report = OfflineReport(rows: [])
    private(set) var update: OfflineUpdateState = .idle
    private(set) var checkedAt: Date?
    var keepClassesReady: Bool { didSet { settings.keepClassesReady = keepClassesReady } }
    var wifiOnly: Bool { didSet { settings.wifiOnly = wifiOnly } }

    private let store: StaticDataStore
    private let icons: IconCacheAccess
    private let latestBuild: @Sendable () async -> String?
    private let settings: OfflineSettings
    private let connectivity: any Connectivity
    private let now: @Sendable () -> Date

    init(
        store: StaticDataStore, icons: IconCacheAccess, latestBuild: @escaping @Sendable () async -> String?,
        settings: OfflineSettings, connectivity: any Connectivity, now: @escaping @Sendable () -> Date
    ) {
        self.store = store
        self.icons = icons
        self.latestBuild = latestBuild
        self.settings = settings
        self.connectivity = connectivity
        self.now = now
        keepClassesReady = settings.keepClassesReady
        wifiOnly = settings.wifiOnly
        checkedAt = settings.checkedAt
    }

    var isOnline: Bool { connectivity.isConnected }

    func load() async {
        report = OfflineReport.build(files: await store.savedFiles(), iconBytes: icons.bytes(), iconsUpdated: icons.lastSaved())
    }

    /// Refreshes the active build and the saved files: a newer build fetches the same files
    /// under it and drops the old one; nothing newer says everything is up to date.
    func updateNow() async {
        guard update != .updating else { return }
        update = .updating
        guard let build = await latestBuild() else {
            update = .failed
            return
        }
        let saved = await store.savedFiles()
        let stale = saved.filter { StaticFileLayout.build(ofRelative: $0.relativePath) != build }
        for file in stale where OfflineCategory.of(relativePath: file.relativePath) != nil {
            let path = "/data/\(build)/" + file.relativePath.split(separator: "/").dropFirst().joined(separator: "/")
            _ = try? await store.file(StaticResource<JSONValue>(path: path))
        }
        if !stale.isEmpty { await store.prune(keeping: build) }
        checkedAt = now()
        settings.checkedAt = checkedAt
        update = stale.isEmpty ? .upToDate : .idle
        await load()
    }

    static func live(_ environment: AppEnvironment) -> OfflineDataModel {
        let store = environment.staticData
        return OfflineDataModel(
            store: store, icons: .live,
            latestBuild: { await store.latestBuild() },
            settings: OfflineSettings(), connectivity: PathConnectivity.shared, now: { Date() }
        )
    }

    /// Remove downloaded data: the store's files and the icons. They download again on open.
    func removeAll() async {
        await store.removeAll()
        icons.removeAll()
        update = .idle
        await load()
    }
}
