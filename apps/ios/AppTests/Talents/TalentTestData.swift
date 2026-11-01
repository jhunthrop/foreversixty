@testable import ForeverSixty
import Foundation
import Models
import Networking

/// The F6 tests' data: the contract fixtures (the real warrior talent file and the Fury BiS
/// bands of the active build), the active build's shaman file read in place, and named builds.
enum TalentTestData {
    static func warrior() throws -> TalentsFile { try BisTestData.fixture("cdn.talents.warrior") }
    static func fury() throws -> BisFile { try BisTestData.bis() }

    static func shaman() throws -> TalentsFile {
        let url = BisTestData.repo.appending(path: "data/builds/\(BisTestData.build)/talents/shaman.json")
        return try WireFormat.decoder().decode(TalentsFile.self, from: Data(contentsOf: url))
    }

    /// The Fury BiS file's 50 to 59 Horde guide build (41 points).
    static let band50 = "34200000000000000-05153105022011501-000000000000000000"
    /// The Fury BiS file's level-60 raid guide build (51 points).
    static let band60 = "34320003002000000-05153105022011501-200000000000000000"
    /// Obnoxious Yell's build in `me.json`: the level-60 path less Improved Bloodrage, 49 points.
    static let yoursTrees: [[Int]] = ["34320003002", "05153105022011501", "0"].map { TalentRules.ranks(ofTree: $0) }
    /// `band60` without Fury's 5-point second talent: 46 points, 5 of `band50`'s missing, 10 stray.
    static let strayAndShort = "34320003002000000-00153105022011501-200000000000000000"

    /// The spec's §4.4 vector: Elemental Shaman, level 60, Horde, band 60 raid.
    static let shamanGuide = "5510331300001350-030000000000000000-0533520000000000"
    static let shamanYours = "5512311100001350-030000000000000000-0433520000000000"
}
