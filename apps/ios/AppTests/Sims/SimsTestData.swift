@testable import ForeverSixty
import Foundation
import Models
import Networking

/// F5's test data: the contract fixtures, plus a level-60 Fury set from the build's own BiS list
/// for the result boards. Every helper throws, so a missing file fails only its own test.
enum SimsTestData {
    static let build = "1.60.1.70658"
    static let now = SimsLogicTests.now

    /// The Fury warrior's level-60 raid list, wearing Cape of the Black Baron where Howler's Furs
    /// is best: the Top Gear board's one change.
    static let furyGear: [(String, Int)] = [
        ("head", 12640), ("neck", 22150), ("shoulder", 272105), ("back", 13340), ("chest", 252483),
        ("wrist", 19587), ("hands", 272095), ("waist", 20216), ("legs", 22385), ("feet", 21995),
        ("finger1", 234034), ("finger2", 19325), ("trinket1", 19289), ("trinket2", 249469),
        ("main_hand", 21520), ("off_hand", 21522), ("ranged", 22656),
    ]
    static let howlersFurs = 272414

    struct Failure: Error, CustomStringConvertible { let description: String }

    static func list() throws -> SimList { try RaidFixture.load(SimList.self, "sims.list") }

    static func content(offline: Date? = nil, server: [SimRow]? = nil, mine: [SimRow]? = nil) throws -> SimsContent {
        let rows = try list().rows
        let mineRows = mine ?? rows + [browserRow]
        return SimsContent(
            serverRows: server ?? rows, mine: mineRows, total: mineRows.count, page: 1, offlineSince: offline
        )
    }

    /// A browser save beside the fixture's server rows, as Your sims shows them.
    static let browserRow = SimRow(
        simId: "p2k7d0wq5nme", spec: "warrior-fury", kind: "run", dps: 851.9, headline: "851.9 DPS", state: "done",
        createdAt: RaidTime.date("2026-12-12T20:40:00Z"), title: "Bare 60", lane: "browser", seen: true
    )

    static func row(_ id: String) throws -> SimRow {
        guard let row = try list().rows.first(where: { $0.simId == id }) else { throw Failure(description: "no row \(id)") }
        return row
    }

    /// The fixture result with the Fury set, a 10,000-iteration run, and for `topGear` the
    /// winning swap into Howler's Furs.
    static func result(topGear: Bool = false, untitled: Bool = false) throws -> SimResult {
        var root = try JSONSerialization.jsonObject(with: RaidFixture.data("sims.result")) as? [String: Any] ?? [:]
        guard var data = root["data"] as? [String: Any],
              var request = data["request"] as? [String: Any],
              var character = request["character"] as? [String: Any]
        else { throw Failure(description: "sims.result has no request.character") }
        character["gear"] = furyGear.map { ["slot": $0.0, "item_id": $0.1] }
        request["character"] = character
        data["request"] = request
        data["dps"] = ["mean": 877.54, "stddev": 61.2, "error": 1.22, "min": 702.1, "max": 1041.8]
        data["iterations_run"] = 10000
        if untitled { data.removeValue(forKey: "title") }
        if topGear {
            data["combos"] = [[
                "substitutions": [["kind": "item", "slot": "back", "item_id": howlersFurs, "name": "Howler's Furs"]],
                "dps": ["mean": 881.6], "delta": ["mean": 4.07, "error": 1.1], "group": 0,
            ]]
        }
        root["data"] = data
        let payload = try JSONSerialization.data(withJSONObject: data)
        return try WireFormat.decoder().decode(SimResult.self, from: payload)
    }

    static let topGearRow = SimRow(
        simId: "g5tq8n2rk1wa", spec: "warrior-fury", kind: "gear", headline: "+4 DPS from Howler's Furs", state: "done",
        createdAt: RaidTime.date("2026-12-15T04:02:00Z"), title: "Cloak check", lane: "server",
        finishedAt: RaidTime.date("2026-12-15T04:05:40Z"), seen: true
    )

    /// The build's warrior items, as the gear grid's names and icons.
    static func looks() throws -> [Int: SimGearLook] {
        let url = HomeFixtures.repoRoot.appending(path: "data/builds/\(build)/items/warrior.json")
        let items = try WireFormat.decoder().decode(ItemsFile.self, from: Data(contentsOf: url)).items
        let wanted = Set(furyGear.map(\.1) + [howlersFurs])
        return Dictionary(uniqueKeysWithValues: items.filter { wanted.contains($0.id) }.map {
            ($0.id, SimGearLook(name: $0.name, icon: $0.icon, quality: $0.quality))
        })
    }
}
