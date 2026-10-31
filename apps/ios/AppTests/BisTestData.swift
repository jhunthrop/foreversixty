@testable import ForeverSixty
import Foundation
import Models
import Networking

/// The repo's files, read in place (the simulator can read the host filesystem): the contract
/// fixtures, the spec catalogue and the generator's real output for the active build.
enum BisTestData {
    static let repo = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // AppTests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root

    static func fixture<T: Decodable>(_ name: String, as _: T.Type = T.self) throws -> T {
        let data = try Data(contentsOf: repo.appending(path: "docs/contracts/fixtures/app/\(name).json"))
        return try WireFormat.decoder().decode(T.self, from: data)
    }

    static func me() throws -> Me {
        struct Wrapped: Decodable { let data: Me }
        return try fixture("me", as: Wrapped.self).data
    }

    static func catalog() throws -> SpecCatalog {
        try SpecCatalog.decode(Data(contentsOf: repo.appending(path: "data/curated/specs.json")))
    }

    static let build = "1.60.1.70658"

    static func bis() throws -> BisFile { try fixture("cdn.bis.warrior-fury") }
    static func items() throws -> ItemsFile { try fixture("cdn.items.warrior") }
    static func classes() throws -> [ClassInfo] { try fixture("cdn.classes") }

    /// The fixture's level-60 Horde raid list for Fury.
    static let raid = BisSelection(
        classSlug: "warrior",
        spec: "warrior-fury",
        range: BisRange(band: 60, preset: "raid"),
        faction: "horde"
    )

    /// A fixed clock for the stamp and Updated lines: Fri 30 Oct 2026, 21:41 UTC.
    static let now = Date(timeIntervalSince1970: 1_793_396_460)
}
