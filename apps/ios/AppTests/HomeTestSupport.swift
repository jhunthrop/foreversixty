@testable import ForeverSixty
import Foundation
import Models
import Networking
import UIKit

/// The contract fixtures and the repo's icons, read in place for the F2 tests.
enum HomeFixtures {
    static let repoRoot = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // AppTests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root

    static let directory = repoRoot.appending(path: "docs/contracts/fixtures/app")

    /// The home fixture's character: Velaren, an Orc Beast Mastery Hunter.
    static let velaren = HomeCharacter(
        key: "us/normal/velaren", region: "us", ruleset: "normal", name: "Velaren", race: "orc",
        className: "hunter", specName: "Beast Mastery", level: 60, guild: "Olympus XXVII", faction: "horde"
    )

    /// 2026-10-29 23:15 UTC: two hours after the fixture's sync.
    static let now = Date(timeIntervalSince1970: 1_793_315_700)

    static func data(_ name: String) throws -> Data {
        try Data(contentsOf: directory.appending(path: "\(name).json"))
    }

    /// An enveloped API fixture's `data`.
    static func home(_ name: String = "character.home", state: String? = nil) throws -> CharacterHome {
        var root = try JSONSerialization.jsonObject(with: data(name)) as? [String: Any] ?? [:]
        if let state, var payload = root["data"] as? [String: Any] {
            payload["state"] = state
            payload["upgrades"] = []
            payload["already_best"] = []
            payload["total_gain"] = NSNull()
            payload["upgrade_count"] = 0
            root["data"] = payload
        }
        let payload = try JSONSerialization.data(withJSONObject: root["data"] ?? [:])
        return try WireFormat.decoder().decode(CharacterHome.self, from: payload)
    }

    static func me() throws -> Me {
        let root = try JSONSerialization.jsonObject(with: data("me")) as? [String: Any] ?? [:]
        return try WireFormat.decoder().decode(Me.self, from: JSONSerialization.data(withJSONObject: root["data"] ?? [:]))
    }

    static func staticFile<T: Decodable>(_ name: String, as _: T.Type) throws -> T {
        try WireFormat.decoder().decode(T.self, from: data(name))
    }

    static func tooltipSources() throws -> ItemTooltipSources {
        let items = try staticFile("cdn.items.hunter", as: ItemsFile.self).items
        return try ItemTooltipSources(
            items: Dictionary(uniqueKeysWithValues: items.map { ($0.id, $0) }),
            sets: staticFile("cdn.sets", as: [ItemSet].self),
            bis: staticFile("cdn.bis.hunter-beast-mastery", as: BisFile.self)
        )
    }

    static func presenter(sims: [FinishedSim] = []) -> HomePresenter {
        HomePresenter(character: velaren, finishedSims: sims, time: HomeTime(now: now, calendar: utc))
    }

    static var utc: Calendar {
        var calendar = Calendar(identifier: .gregorian)
        calendar.timeZone = TimeZone(identifier: "UTC") ?? .gmt
        return calendar
    }
}

/// Item icons from the repo's own build folder (`data/builds/{build}/icons`), loaded at once so
/// a snapshot shows the real art.
struct RepoItemIconSource: ItemIconSource {
    func cached(icon: String, build: String) -> UIImage? {
        let url = HomeFixtures.repoRoot.appending(path: "data/builds/\(build)/icons/\(icon).webp")
        return (try? Data(contentsOf: url)).flatMap(UIImage.init(data:))
    }

    func load(icon: String, build: String) async -> UIImage? { cached(icon: icon, build: build) }
}
