@testable import ForeverSixty
import Foundation
import Routing
import Testing

@Suite("Push landings (flows spec F4)")
struct PushLandingTests {
    private let raidData = ["guild_id": "2", "raid_id": "41"]

    @Test("raid events open the raid page; raid.short opens Not answered; raid.starting opens raid night")
    func raids() {
        #expect(PushLanding.destination(kind: "raid.short", link: nil, data: raidData) == .raid(.notAnswered(guild: 2, raid: 41)))
        for kind in ["raid.posted", "raid.signup_nudge", "raid.changed", "roster.pending", "raid.late_change"] {
            #expect(PushLanding.destination(kind: kind, link: nil, data: raidData) == .link(.raid(guild: 2, raid: 41)))
        }
        #expect(PushLanding.destination(kind: "raid.starting", link: nil, data: raidData) == .link(.raidNight(guild: 2, raid: 41)))
    }

    @Test("the night's summary opens its report")
    func summary() {
        let data = raidData.merging(["report_id": "r9"]) { a, _ in a }
        #expect(PushLanding.destination(kind: "raid.summary", link: nil, data: data) == .link(.report(id: "r9")))
    }

    @Test("sims and reports land on the site link's screen")
    func webLinks() throws {
        let sim = try #require(URL(string: "https://foreversixty.gg/sim/abc"))
        let report = try #require(URL(string: "https://foreversixty.gg/reports/r1"))
        #expect(PushLanding.destination(kind: "sim.finished", link: sim, data: [:]) == .link(.simPage(id: "abc")))
        #expect(PushLanding.destination(kind: "report.ready", link: report, data: [:]) == .link(.reportPage(id: "r1")))
        #expect(PushLanding.destination(kind: "coach.ready", link: report, data: [:]) == .link(.reportPage(id: "r1")))
    }

    @Test("loot and gear events open the gear list; a linked character opens itself")
    func gear() {
        for kind in ["loot.awarded", "bis.changed", "band.entered", "game.milestone"] {
            #expect(PushLanding.destination(kind: kind, link: nil, data: ["item_id": "19364"]) == .tab(.gear))
        }
        #expect(PushLanding.destination(kind: "character.linked", link: nil, data: ["character_key": "us/forever/thandrel"])
            == .character(key: "us/forever/thandrel"))
    }

    @Test("billing and seat events open the plan")
    func account() {
        #expect(PushLanding.destination(kind: "billing.failed", link: nil, data: [:]) == .link(.plan))
        #expect(PushLanding.destination(kind: "premium.seat", link: nil, data: [:]) == .link(.plan))
    }

    @Test("an unknown kind falls back to its link, else nothing")
    func fallback() throws {
        let character = try #require(URL(string: "https://foreversixty.gg/character/us/forever/kaelra"))
        #expect(PushLanding.destination(kind: "something.new", link: character, data: [:])
            == .link(.characterPage(region: "us", ruleset: "forever", name: "kaelra")))
        #expect(PushLanding.destination(kind: "something.new", link: nil, data: [:]) == nil)
        let elsewhere = try #require(URL(string: "https://example.com/sim/abc"))
        #expect(PushLanding.destination(kind: nil, link: elsewhere, data: [:]) == nil)
    }

    @Test("the APNs userInfo shape: top-level kind and link, data strings")
    func userInfo() {
        let info: [AnyHashable: Any] = ["aps": [:] as [String: Any], "kind": "raid.posted", "data": raidData]
        #expect(PushLanding.destination(userInfo: info) == .link(.raid(guild: 2, raid: 41)))
    }

    @Test("Bench is a background answer like In, Late and Out")
    @MainActor
    func bench() {
        #expect(RaidPushCategory.signup(forAction: RaidPushCategory.actionBench) == PushSignup(state: "bench"))
        let signup = RaidNotificationCenter.categories.first { $0.identifier == RaidPushCategory.signup }
        #expect(signup?.actions.map(\.identifier) == ["signup_in", "signup_late", "signup_bench", "signup_out"])
    }
}

@Suite("Push token registration")
@MainActor
struct PushRegistrarTests {
    final class Calls: PushTokenRegistering, @unchecked Sendable {
        var tokens: [String] = []
        var fails = false
        func register(token: String) async throws {
            if fails { throw URLError(.notConnectedToInternet) }
            tokens.append(token)
        }
    }

    private func registrar(authorized: Bool, calls: Calls, apns: @escaping @MainActor () -> Void = {}) -> PushRegistrar {
        let registrar = PushRegistrar(isAuthorized: { authorized }, registerWithAPNs: apns)
        registrar.calls = calls
        return registrar
    }

    private let token = Data([0xAB, 0x01, 0xFF])

    @Test("signed in and allowed: asks APNs, then sends the hex token once")
    func registers() async {
        let calls = Calls()
        var asked = 0
        let registrar = registrar(authorized: true, calls: calls) { asked += 1 }
        await registrar.sync(signedIn: true)
        await registrar.received(deviceToken: token)
        await registrar.sync(signedIn: true)
        #expect(asked == 2)
        #expect(calls.tokens == ["ab01ff"])
    }

    @Test("never registers without permission or while signed out")
    func waits() async {
        let calls = Calls()
        var asked = 0
        let denied = registrar(authorized: false, calls: calls) { asked += 1 }
        await denied.sync(signedIn: true)
        let signedOut = registrar(authorized: true, calls: calls) { asked += 1 }
        await signedOut.sync(signedIn: false)
        await signedOut.received(deviceToken: token)
        #expect(asked == 0)
        #expect(calls.tokens.isEmpty)
    }

    @Test("a new sign-in sends the token again; a failure retries on the next sync")
    func resends() async {
        let calls = Calls()
        let registrar = registrar(authorized: true, calls: calls)
        calls.fails = true
        await registrar.sync(signedIn: true)
        await registrar.received(deviceToken: token)
        calls.fails = false
        await registrar.sync(signedIn: true)
        await registrar.sync(signedIn: false)
        await registrar.sync(signedIn: true)
        #expect(calls.tokens == ["ab01ff", "ab01ff"])
    }
}
