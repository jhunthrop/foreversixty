import Foundation
import Networking

/// The last good answer per endpoint, on disk (spec 5.0 Caching): Application Support, excluded
/// from backup. A cached answer draws at once and sits under the offline strip when a refresh
/// cannot connect.
protocol RaidCache: Sendable {
    func read<Value: Decodable>(_ type: Value.Type, key: String) -> Cached<Value>?
    func write(_ value: some Encodable, key: String, at date: Date)
    /// Sign-out: nothing of the last account's guild stays on the device.
    func removeAll()
}

struct Cached<Value> {
    let value: Value
    let savedAt: Date
}

struct DiskRaidCache: RaidCache {
    let directory: URL

    static func standard() -> DiskRaidCache {
        let base = FileManager.default.urls(for: .applicationSupportDirectory, in: .userDomainMask).first
            ?? FileManager.default.temporaryDirectory
        return DiskRaidCache(directory: base.appending(path: "raids", directoryHint: .isDirectory))
    }

    private func file(_ key: String) -> URL {
        directory.appending(path: key.replacingOccurrences(of: "/", with: "_") + ".json")
    }

    func read<Value: Decodable>(_ type: Value.Type, key: String) -> Cached<Value>? {
        guard let data = try? Data(contentsOf: file(key)),
              let saved = try? WireFormat.decoder().decode(SavedAt.self, from: data),
              let value = try? WireFormat.decoder().decode(Wrapped<Value>.self, from: data)
        else { return nil }
        return Cached(value: value.value, savedAt: saved.savedAt)
    }

    func write(_ value: some Encodable, key: String, at date: Date) {
        do {
            try FileManager.default.createDirectory(at: directory, withIntermediateDirectories: true)
            var url = file(key)
            let data = try WireFormat.encoder().encode(Wrapper(savedAt: date, value: value))
            try data.write(to: url, options: .atomic)
            var values = URLResourceValues()
            values.isExcludedFromBackup = true
            try url.setResourceValues(values)
        } catch {
            // A cache that cannot write only costs the next cold start its instant draw.
        }
    }

    func removeAll() {
        try? FileManager.default.removeItem(at: directory)
    }

    private struct SavedAt: Decodable { let savedAt: Date }
    private struct Wrapped<Value: Decodable>: Decodable { let value: Value }
    private struct Wrapper<Value: Encodable>: Encodable {
        let savedAt: Date
        let value: Value
    }
}

/// No cache: tests and previews.
struct NoRaidCache: RaidCache {
    func read<Value: Decodable>(_ type: Value.Type, key: String) -> Cached<Value>? { nil }
    func write(_ value: some Encodable, key: String, at date: Date) {}
    func removeAll() {}
}
