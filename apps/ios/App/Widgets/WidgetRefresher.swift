import Auth
import BackgroundTasks
import Foundation
import Models
import Networking
import UIKit
import WidgetKit

/// Keeps the App Group snapshot current (spec §5.6 Refresh): on every foreground, after a sign-up
/// answered in the app, from the refresh task about every 30 minutes and on the silent
/// `widget.refresh` push; then every widget reloads and the Live Activities follow. `GET
/// /v1/me/widget` revalidates through URLSession's cache with its ETag, so an unchanged answer
/// is a 304. Signed out empties the container, so every surface says so.
@MainActor
final class WidgetRefresher {
    nonisolated static let taskID = "gg.foreversixty.app.widget-refresh"
    nonisolated static let interval: TimeInterval = 30 * 60

    private let calls: any WidgetCalls
    private let store: WidgetSnapshotStore
    private let icons: any ItemIconSource
    private let iconBuild: @MainActor () -> String?
    private let enabled: @MainActor () -> Bool
    private let now: @Sendable () -> Date
    private let reload: @Sendable () -> Void
    var afterRefresh: (@MainActor (WidgetSnapshot?) async -> Void)?

    init(
        calls: any WidgetCalls, store: WidgetSnapshotStore = .shared(), icons: any ItemIconSource = RemoteItemIconSource.shared,
        iconBuild: @escaping @MainActor () -> String?, enabled: @escaping @MainActor () -> Bool,
        now: @escaping @Sendable () -> Date = { Date() },
        reload: @escaping @Sendable () -> Void = { WidgetCenter.shared.reloadAllTimelines() }
    ) {
        self.calls = calls
        self.store = store
        self.icons = icons
        self.iconBuild = iconBuild
        self.enabled = enabled
        self.now = now
        self.reload = reload
    }

    /// One refresh for the account as the session holds it.
    func refresh(signedIn: Bool, me: Me?, current: String?) async {
        guard enabled() else { return }
        guard signedIn else {
            clear()
            return
        }
        if let me {
            store.saveCharacters(me.characters.map { WidgetCharacterChoice(key: $0.key, name: $0.name) })
        }
        let snapshot = await fetch(character: current, into: nil)
        for key in await configuredCharacters() where key != snapshot?.widget.character?.key {
            _ = await fetch(character: key, into: key)
        }
        reload()
        await afterRefresh?(snapshot ?? store.load())
    }

    /// Sign-out: the container empties and the widgets redraw signed out.
    func clear() {
        store.clear()
        reload()
    }

    @discardableResult
    private func fetch(character: String?, into file: String?) async -> WidgetSnapshot? {
        do throws(APIError) {
            let widget = try await calls.snapshot(character: character)
            let snapshot = WidgetSnapshot(widget: widget, fetchedAt: now())
            try? store.save(snapshot, character: file)
            await saveIcons(widget)
            return snapshot
        } catch {
            if case .unauthorized = error { store.clear() }
            return nil
        }
    }

    /// The extension cannot fetch while it draws, so the upgrade rows' icons go in the container.
    private func saveIcons(_ widget: Widget) async {
        guard let build = iconBuild(), let rows = widget.upgrades?.rows else { return }
        for row in rows {
            let url = store.iconURL(row.pick.icon)
            guard !FileManager.default.fileExists(atPath: url.path),
                  let image = await icons.load(icon: row.pick.icon, build: build), let data = image.pngData() else { continue }
            try? FileManager.default.createDirectory(at: url.deletingLastPathComponent(), withIntermediateDirectories: true)
            try? data.write(to: url, options: .atomic)
        }
    }

    /// The characters widgets on this phone are set to, besides the current one.
    private func configuredCharacters() async -> [String] {
        guard let infos = try? await WidgetCenter.shared.currentConfigurations() else { return [] }
        let keys = infos.compactMap { ($0.widgetConfigurationIntent(of: CharacterConfigurationIntent.self))?.characterKey }
        return Array(Set(keys)).sorted()
    }

    /// The next refresh task, about 30 minutes out; the system decides the exact moment.
    nonisolated static func schedule(after interval: TimeInterval = interval) {
        let request = BGAppRefreshTaskRequest(identifier: taskID)
        request.earliestBeginDate = Date(timeIntervalSinceNow: interval)
        try? BGTaskScheduler.shared.submit(request)
    }

    /// The silent push's payload (fixture surface.widget-refresh.apns): `{ aps, kind }`.
    nonisolated static func isRefreshPush(_ userInfo: [AnyHashable: Any]) -> Bool {
        userInfo["kind"] as? String == "widget.refresh"
    }
}

extension Notification.Name {
    /// A raid answer saved in the app: the widgets refresh at once.
    static let raidSignupSaved = Notification.Name("gg.foreversixty.raidSignupSaved")
}
