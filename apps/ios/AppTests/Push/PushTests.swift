@testable import ForeverSixty
import Foundation
import Models
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

    @Test("a new applicant opens its review; Decline reads the same guild and application")
    func applicant() {
        let data = ["guild_id": "2", "application_id": "17"]
        #expect(PushLanding.destination(kind: "guild.applicant", link: nil, data: data) == .link(.guildApplicant(guild: 2, application: 17)))
        #expect(PushLanding.applicant(data: data)?.guild == 2)
        #expect(PushLanding.applicant(data: data)?.application == 17)
        #expect(PushLanding.applicant(data: ["guild_id": "2"]) == nil)
        #expect(PushLanding.destination(kind: "guild.applicant", link: nil, data: ["guild_id": "2"]) == nil)
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

    /// The contract fixture's APNs payload as userInfo.
    private func apns(_ name: String) throws -> [AnyHashable: Any] {
        let data = try Data(contentsOf: RaidFixture.directory.appending(path: "\(name).json"))
        return try #require(try JSONSerialization.jsonObject(with: data) as? [AnyHashable: Any])
    }

    @Test("the officer pushes' categories carry the server's actions: Mention on RAID_SHORT, Accept before Decline")
    @MainActor
    func officerCategories() throws {
        let categories = RaidNotificationCenter.categories
        for name in ["push.raid-short.apns", "push.guild-applicant.apns"] {
            let info = try apns(name)
            let aps = try #require(info["aps"] as? [String: Any])
            let category = try #require(categories.first { $0.identifier == aps["category"] as? String }, "\(name)'s category is registered")
            let served = (info[RaidMentionRequest.actionsKey] as? [[String: Any]] ?? []).compactMap { $0["id"] as? String }
            #expect(category.actions.map(\.identifier) == served, "\(name)")
        }
        let short = try #require(categories.first { $0.identifier == RaidPushCategory.short })
        #expect(short.actions.allSatisfy { !$0.options.contains(.foreground) }, "Mention answers in place")
    }

    @Test("Mention reads the raid and the short role from the push; Accept and Decline the application")
    func officerAnswers() throws {
        let short = try apns("push.raid-short.apns")
        #expect(OfficerPushAnswer(action: RaidPushCategory.actionMention, userInfo: short)
            == .mention(try #require(RaidMentionRequest(userInfo: short))))
        let request = try #require(RaidMentionRequest(userInfo: short))
        #expect(request.guild == 2)
        #expect(request.raid == 7)
        #expect(request.role == "healer")
        var cooled = short
        cooled[RaidMentionRequest.actionsKey] = nil
        #expect(OfficerPushAnswer(action: RaidPushCategory.actionMention, userInfo: cooled) == .unreadable, "no action inside the cooldown")
        let applicant = try apns("push.guild-applicant.apns")
        #expect(OfficerPushAnswer(action: RaidPushCategory.actionAccept, userInfo: applicant) == .accept(guild: 2, application: 31))
        #expect(OfficerPushAnswer(action: RaidPushCategory.actionDecline, userInfo: applicant) == .decline(guild: 2, application: 31))
        #expect(OfficerPushAnswer(action: RaidPushCategory.actionIn, userInfo: applicant) == nil)
    }

    /// The fixture ping's answer: 3 mentioned, nobody left out.
    private func pingResult() throws -> RaidPingResult {
        try RaidFixture.load(RaidPingResult.self, "guild.raid.ping")
    }

    /// A centre whose officer actions log their calls and whose answers are recorded, not posted.
    @MainActor
    private func officerCenter(
        _ calls: OfficerCalls, accepts: Bool = true, declines: Bool = true, ping: RaidPingResult?
    ) -> (RaidNotificationCenter, PostedAnswers) {
        let center = RaidNotificationCenter()
        let posted = PostedAnswers()
        center.use(answerer: nil, officer: OfficerPushActions(
            accept: { guild, application in await calls.record("accept \(guild)/\(application)"); return accepts },
            decline: { guild, application in await calls.record("decline \(guild)/\(application)"); return declines },
            mention: { guild, raid, role in await calls.record("mention \(guild)/\(raid) \(role)"); return ping }
        ), post: { answer, reply in posted.add(answer, reply) })
        return (center, posted)
    }

    private func reply(_ action: String, _ name: String, title: String) throws -> PushReply {
        PushReply(action: action, identifier: name, thread: "t-\(name)", title: title, userInfo: try apns(name))
    }

    @Test("round 5, row 3: Mention, Accept and Decline run in place and their result replaces the push")
    @MainActor
    func officerResults() async throws {
        let calls = OfficerCalls()
        let (center, posted) = officerCenter(calls, ping: try pingResult())
        var opened: [IntentDestination] = []
        center.onOpen = { opened.append($0) }
        let applicant = "New applicant: Mirelle, Holy Priest 60"
        await center.respond(try reply(RaidPushCategory.actionMention, "push.raid-short.apns", title: "Onyxia's Lair"))
        await center.respond(try reply(RaidPushCategory.actionAccept, "push.guild-applicant.apns", title: applicant))
        await center.respond(try reply(RaidPushCategory.actionDecline, "push.guild-applicant.apns", title: applicant))
        #expect(await calls.log == ["mention 2/7 healer", "accept 2/31", "decline 2/31"])
        #expect(opened.isEmpty, "answered in place, nothing opens")
        #expect(posted.answers == [
            .init(title: "Mentioned 3 healers in Discord.", body: "Onyxia's Lair"),
            .init(title: Copy.pushApplicantAccepted, body: applicant),
            .init(title: Copy.pushApplicantDeclined, body: applicant),
        ])
        #expect(posted.replaced == ["push.raid-short.apns", "push.guild-applicant.apns", "push.guild-applicant.apns"])
        #expect(posted.threads.first == "t-push.raid-short.apns", "the result keeps the push's thread")
        let kept = try #require(posted.tapThrough.first)
        #expect(kept["kind"] == .text("raid.short"), "a tap on the result still lands on Not answered")
        #expect(kept["actions"] == nil, "the result carries no actions")
    }

    @Test("round 5, row 3: a failed or unreadable answer replaces the push with the failure line")
    @MainActor
    func officerFailures() async throws {
        let calls = OfficerCalls()
        let (center, posted) = officerCenter(calls, accepts: false, declines: false, ping: nil)
        await center.respond(try reply(RaidPushCategory.actionMention, "push.raid-short.apns", title: "Onyxia's Lair"))

        let named = try apns("push.guild-applicant.apns")
        let title = "New applicant: Mirelle, Holy Priest 60"
        for action in [RaidPushCategory.actionAccept, RaidPushCategory.actionDecline] {
            await center.respond(PushReply(action: action, identifier: "a", thread: "", title: title, userInfo: named))
        }
        var cooled = try apns("push.raid-short.apns")
        cooled[RaidMentionRequest.actionsKey] = nil
        await center.respond(PushReply(action: RaidPushCategory.actionMention, identifier: "c", thread: "", title: "Onyxia's Lair", userInfo: cooled))
        var unnamed = try apns("push.guild-applicant.apns")
        unnamed["data"] = ["guild_id": "2", "application_id": "31"]
        await center.respond(PushReply(action: RaidPushCategory.actionAccept, identifier: "u", thread: "", title: title, userInfo: unnamed))
        #expect(posted.answers == [
            .init(title: "Could not mention the healers", body: "Onyxia's Lair. Open the raid to try again."),
            .init(title: "Could not accept Mirelle", body: "Mirelle, Holy Priest 60. Open Forever Sixty to try again."),
            .init(title: "Could not decline Mirelle", body: "Mirelle, Holy Priest 60. Open Forever Sixty to try again."),
            // No role to name: the plain title, the same retry line.
            .init(title: Copy.commonFailed, body: "Onyxia's Lair. Open the raid to try again."),
            // A push without the applicant's data reads its own title.
            .init(title: Copy.commonFailed, body: "\(title). Open Forever Sixty to try again."),
        ])
        #expect(await calls.log == ["mention 2/7 healer", "accept 2/31", "decline 2/31", "accept 2/31"], "an unreadable Mention calls nothing")
        #expect(posted.tapThrough.first?["kind"] == .text("raid.short"), "a tap on the failure opens what the push did")
    }

    @Test("round 4: UIKit's completion for a response runs on the main thread")
    func completionOnMain() async {
        let onMain = await withCheckedContinuation { (continuation: CheckedContinuation<Bool, Never>) in
            let done = MainThreadCompletion { continuation.resume(returning: Thread.isMainThread) }
            Task.detached { await done.run() }
        }
        #expect(onMain)
    }
}

/// The answers a centre posted in place of their pushes, in order.
@MainActor
private final class PostedAnswers {
    private(set) var answers: [RaidPushOutcome.Confirmation] = []
    private(set) var replaced: [String] = []
    private(set) var threads: [String] = []
    private(set) var tapThrough: [[String: PushTapValue]] = []

    func add(_ answer: RaidPushOutcome.Confirmation, _ reply: PushReply) {
        answers.append(answer)
        replaced.append(reply.identifier)
        threads.append(reply.thread)
        tapThrough.append(reply.tapThrough)
    }
}

/// What the officer actions were asked to do, in order.
private actor OfficerCalls {
    private(set) var log: [String] = []

    func record(_ call: String) { log.append(call) }
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
