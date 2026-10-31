@testable import ForeverSixty
import Testing

/// The generated accessors read the bundled String Catalog (apps/shared/copy/en.json).
@Suite("Copy catalog")
struct CopyTests {
    @Test("a plain key reads its text, not its key")
    func plain() {
        #expect(Copy.tabHome == "Home")
        #expect(Copy.devicesThis == "This iPhone")
    }

    @Test("placeholders fill in order")
    func placeholders() {
        #expect(Copy.raidCount(n: 24, size: 40) == "24 of 40 signed up")
        #expect(Copy.homeUpgradesReplaces(item: "Dark Iron Rifle", n: 55) == "Replaces Dark Iron Rifle, ilvl 55")
    }

    @Test("plurals pick their form from the count")
    func plurals() {
        #expect(Copy.talentsUnspent(n: 1) == "1 point unspent.")
        #expect(Copy.talentsUnspent(n: 3) == "3 points unspent.")
        #expect(Copy.homeAnswerUpgrades(n: 1) == "1 upgrade")
        #expect(Copy.homeAnswerUpgrades(n: 2) == "2 upgrades")
    }
}
