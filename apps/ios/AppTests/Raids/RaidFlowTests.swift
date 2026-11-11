@testable import ForeverSixty
import DesignSystem
import Foundation
import Models
import Networking
import SwiftUI
import Testing

@MainActor
@Suite("Raid flow parity")
struct RaidFlowParityTests {
    static let device = CGSize(width: 393, height: 852)
    typealias T = RaidFlowTestData

    private func root(_ content: RaidRootContent, _ screen: String, testName: String = #function) async throws {
        let view = RaidRootView(content: content).environment(\.itemIconSource, RepoItemIconSource()).environment(\.raidIconBuild, HomeFlowTestData.build)
        try await ParityCheck.assert(view, screen: screen, size: Self.device, testName: testName)
    }

    @Test("ios-raid") func member() async throws { try await root(T.member(), "raid") }
    @Test("ios-raid-signup") func signup() async throws { try await root(T.member(late: true), "raid-signup") }
    @Test("ios-raid-officer") func officer() async throws { try await root(T.member(officer: true), "raid-officer") }
    @Test("ios-raid-live") func live() async throws { try await root(T.live, "raid-live") }
    @Test("ios-raid-no-raid") func noRaid() async throws { try await root(T.noRaid, "raid-no-raid") }
    @Test("ios-raid-out") func signedOut() async throws { try await root(.signedOut, "raid-out") }

    @Test("ios-raid-unanswered")
    func unanswered() async throws {
        let view = Self.pushed(T.notAnswered.title) {
            NotAnsweredView(content: T.notAnswered, build: HomeFlowTestData.build)
        }
        try await ParityCheck.assert(view, screen: "raid-unanswered", size: Self.device)
    }

    @Test("ios-raid-applicant")
    func applicant() async throws {
        let content = T.applicant(accepted: false)
        try await ParityCheck.assert(Self.pushed(content.name) { ApplicantView(content: content) }, screen: "raid-applicant", size: Self.device)
    }

    @Test("ios-raid-applicant-accepted")
    func applicantAccepted() async throws {
        let content = T.applicant(accepted: true)
        try await ParityCheck.assert(Self.pushed(content.name) { ApplicantView(content: content) }, screen: "raid-applicant-accepted", size: Self.device)
    }

    private func night(_ answered: Bool, _ screen: String, testName: String = #function) async throws {
        let content = T.night(answered: answered)
        let view = Self.pushed(content.title) {
            RaidNightPageView(content: content, build: HomeFlowTestData.build)
                .toolbar {
                    ToolbarItem(placement: .topBarTrailing) {
                        Button {} label: { Label(Copy.raidMenuShare, systemImage: "square.and.arrow.up") }
                    }
                }
        }
        try await ParityCheck.assert(view, screen: screen, size: Self.device, testName: testName)
    }

    /// A raid flow page as the app pushes it over the Raid tab: in its stack, on the N3 bar with the back
    /// control (Android's RaidPagesParityTest wraps FlowBackBar the same way), with the repo's icons.
    private static func pushed(_ title: String, @ViewBuilder _ content: @escaping () -> some View) -> some View {
        PushedScreen(back: AppTab.guild.title) { content().pushedBandBar(title) }
            .environment(\.itemIconSource, RepoItemIconSource())
    }

    @Test("ios-raid-night-live") func nightLive() async throws { try await night(true, "raid-night-live") }
    @Test("ios-raid-night-signup") func nightSignup() async throws { try await night(false, "raid-night-signup") }
}

/// The Raid tab's words and rules, from the API's shapes.
@MainActor
@Suite("Raid flow")
struct RaidFlowLogicTests {
    private let presenter = RaidFlowPresenter(now: RaidTime.now, clock: RaidTime.clock)

    @Test("kinds name the spec and class once; lists read `A, B and C`")
    func words() {
        #expect(RaidFlowPresenter.kind(spec: "Holy", className: "priest") == "Holy Priest")
        #expect(RaidFlowPresenter.kind(spec: "Holy Paladin", className: "paladin") == "Holy Paladin")
        #expect(RaidFlowPresenter.kind(spec: nil, className: "rogue") == "Rogue")
        #expect(RaidFlowPresenter.list(["Mirethil", "Selka", "Aurelith"]) == "Mirethil, Selka and Aurelith")
        #expect(RaidFlowPresenter.list(["Selka"]) == "Selka")
    }

    @Test("no guild: the recruiting heading takes its article")
    func noGuild() {
        let row = RecruitingGuild(guildId: 2, name: "Olympus XXVII", faction: "alliance", schedule: "Wed and Sat, 8:00 pm CST", needs: "damage", crestUrl: nil)
        let fury = RaidFlowPresenter.recruiting([row], kind: "Fury Warrior")
        #expect(fury.heading == "Guilds recruiting a Fury Warrior")
        #expect(fury.rows.first?.line == "Alliance · Wed and Sat, 8:00 pm CST · needs damage")
        #expect(RaidFlowPresenter.recruiting([], kind: "Arms Warrior").heading == "Guilds recruiting an Arms Warrior")
    }

    private func person(_ name: String, late: Int? = nil) -> RaidNightPerson {
        RaidNightPerson(name: name, className: "priest", state: late == nil ? "in" : "late", lateMinutes: late, note: "")
    }

    @Test("role rows: the late minutes on the root; the gap and `and N more` on the raid page")
    func roles() {
        let healers = [person("Mirethil"), person("Selka"), person("Brightwen", late: 15)]
        #expect(presenter.roleLine("Healers", healers, want: 8, lateShown: true, short: false)
            == HomeLine(title: "Healers 3 of 8", sub: "Mirethil, Selka, Brightwen (Late, 15 min)"))
        #expect(presenter.roleLine("Healers", healers, want: 8, lateShown: false, short: true).sub == "Mirethil, Selka, Brightwen · 5 short")
        let damage = (1 ... 11).map { person("D\($0)") }
        #expect(presenter.roleLine("Damage", damage, want: 30, lateShown: false, short: false).sub == "D1, D2, D3 and 8 more")
    }

    @Test("the Ready check: every item, a miss's place to get it, the export's age")
    func ready() throws {
        let check = try RaidFixture.load(ReadyCheck.self, "character.ready")
        let ready = try #require(presenter.ready(check))
        #expect(ready.rows.count == check.items.count)
        #expect(ready.rows.filter { $0.mark == .ok }.allSatisfy { $0.fix == nil })
        #expect(ready.stamp?.hasPrefix("From your last log-out, ") == true)
        #expect(presenter.ready(ReadyCheck(exportedAt: nil, items: [])) == nil)
    }

    @Test("the applicant: kind and level, the numbers, raid ready in words, the history")
    func applicant() throws {
        let page = try RaidFixture.load(GuildOfficerApplications.self, "guild.applications")
        let row = try #require(page.rows.first)
        let now = row.createdAt.addingTimeInterval(120)
        let content = RaidFlowPresenter(now: now, clock: RaidTime.clock).applicant(row, guild: "Olympus XXVII", accepted: false)
        #expect(content.facts == ["Holy Paladin · 60", "Alliance Dwarf", "applied 2 min ago"])
        #expect(content.figures.map(\.value) == ["74", "612", "63"])
        #expect(content.figures[1].unit == "HPS")
        #expect(content.ready.map(\.title) == ["Attuned to Onyxia's Lair", "Enchants", "Best in slot gap"])
        #expect(content.ready[1].sub == "2 of 4 slots · cloak and boots missing")
        #expect(content.history.first?.sub.hasPrefix("14 raids · last raid ") == true)
        #expect(content.accepted == nil)
        let accepted = RaidFlowPresenter(now: now).applicant(row, guild: "Olympus XXVII", accepted: true)
        #expect(accepted.accepted == "Accepted into Olympus XXVII. Invite Test Applicant in game:")
        #expect(accepted.invite == "/ginvite Test Applicant")
    }

    @Test("the pending applicants, newest first, carry their ids")
    func pending() throws {
        let page = try RaidFixture.load(GuildOfficerApplications.self, "guild.applications")
        let lines = ApplicantLine.pending(page)
        #expect(lines.first?.id == 41)
        #expect(lines.map(\.appliedAt) == lines.map(\.appliedAt).sorted(by: >))
    }

    // MARK: Not answered

    /// The officer's raid, healers short by 2: Mira (a healer) and two damage unanswered, Varkas a
    /// healer on the bench.
    private func shortRaid() throws -> RaidView {
        var json = try #require(JSONSerialization.jsonObject(with: RaidFixture.data("guild.raid.officer")) as? [String: Any])
        var data = try #require(json["data"] as? [String: Any])
        data["comp_target"] = ["tanks": 3, "healers": 8, "melee": 6, "ranged": 8, "source": "guild"]
        data["not_signed_raiders"] = [
            ["character_key": "us/pvp/zulmara", "name": "Zulmara", "class": "hunter", "account_key": "a1", "discord_linked": true, "role": "dps", "last_week": "in", "discord_dm_url": "https://discord.com/users/1"],
            ["character_key": "us/pvp/mira", "name": "Mira", "class": "priest", "account_key": "a2", "discord_linked": true, "role": "healer", "last_week": "late"],
            ["character_key": "us/pvp/brokk", "name": "Brokk", "class": "warrior", "account_key": "a3", "discord_linked": false, "role": "dps", "last_week": NSNull()],
        ]
        var signups = try #require(data["signups"] as? [[String: Any]])
        for index in signups.indices where signups[index]["name"] as? String == "Varkas" {
            signups[index]["role"] = "healer"
            signups[index]["class"] = "priest"
            signups[index]["spec_name"] = "Holy"
        }
        data["signups"] = signups
        json["data"] = data
        let bytes = try JSONSerialization.data(withJSONObject: json)
        return try #require(try WireFormat.decoder().decode(Envelope<RaidView>.self, from: bytes).data)
    }

    @Test("Not answered: the short role first, its mention, the bench that closes the gap")
    func notAnswered() throws {
        let raid = try shortRaid()
        let content = NotAnsweredPresenter(raid: raid, guild: "Olympus XXVII", marks: [:], clock: RaidTime.clock).content(filter: .notAnswered)
        #expect(content.rows.map(\.name) == ["Mira", "Zulmara", "Brokk"])
        #expect(content.rows.map(\.sub).last == "Arms Warrior · new this week")
        #expect(content.rows.first?.sub.hasSuffix("was Late last week") == true)
        #expect(content.facts.last == "healers short by 2")
        #expect(content.chips.first?.title == "Not answered 3")
        #expect(content.mention?.role == "Mention the 1 healer in Discord")
        #expect(content.mention?.all == "Mention all 3")
        #expect(content.mention?.preview.hasPrefix("@Mira: Onyxia's Lair is ") == true)
        #expect(content.listMeta == "3 of \(raid.signups.count + raid.notSigned.count), healers first")
        #expect(content.bench.map(\.name) == ["Varkas"])
        #expect(content.benchTitle == "Healers on the bench")
        #expect(content.benchNote == "If Mira and Varkas are In, healers are full.")
    }

    @Test("a mark stays on its row with Undo and moves the counts; the mention then names everyone left")
    func marked() throws {
        let raid = try shortRaid()
        let presenter = NotAnsweredPresenter(raid: raid, guild: "Olympus XXVII", marks: ["us/pvp/mira": .signedIn], clock: RaidTime.clock)
        let content = presenter.content(filter: .notAnswered)
        #expect(content.rows.first?.marked == .signedIn)
        #expect(content.chips.first?.title == "Not answered 2")
        #expect(content.chips[1].title == "In \(raid.signups.filter { $0.state == "in" }.count + 1)")
        #expect(presenter.shortRole?.missing == 1)
        #expect(content.mention?.role == nil)
        #expect(content.mention?.preview.hasPrefix("@Zulmara @Brokk: ") == true)
    }

    @Test("other chips list the sign-ups in that state, without row actions")
    func filters() throws {
        let raid = try shortRaid()
        let bench = NotAnsweredPresenter(raid: raid, guild: "Olympus XXVII", marks: [:], clock: RaidTime.clock).content(filter: .bench)
        #expect(bench.rows.map(\.name) == ["Varkas"])
        #expect(!bench.actionable)
        #expect(bench.mention == nil)
        #expect(bench.bench.isEmpty)
    }

    @Test("last week's answer in words")
    func lastWeek() {
        #expect(NotAnsweredPresenter.lastWeek("in") == "was In last week")
        #expect(NotAnsweredPresenter.lastWeek("missed") == "missed last week")
        #expect(NotAnsweredPresenter.lastWeek(nil) == "new this week")
    }

    @Test("Not answered's mark waits the hold, Undo cancels it, then the answer is sent")
    func markHold() async throws {
        let service = FakeRaidService()
        let model = NotAnsweredModel(guild: 2, raid: 41, service: service, ping: FakeRaidPing(), hold: .milliseconds(50))
        model.mark("us/pvp/mira", .signedIn)
        model.undo("us/pvp/mira")
        #expect(model.marks.isEmpty)
        model.mark("us/pvp/zulmara", .bench)
        #expect(model.marks["us/pvp/zulmara"] == .bench)
        try await Task.sleep(for: .milliseconds(400))
        #expect(service.signups.map(\.key) == ["us/pvp/zulmara"])
        #expect(service.signups.first?.request.state == "bench")
    }

    @Test("Mention says how many it reached; with no Discord bot the raid's page opens")
    func mention() async {
        let sent = NotAnsweredModel(guild: 2, raid: 41, service: FakeRaidService(), ping: FakeRaidPing(), hold: .seconds(4))
        #expect(await sent.mention(role: "healer") == false)
        #expect(sent.mentionStatus == "Mentioned 2 healers in Discord.")
        let noBot = APIError.server(status: 409, body: ErrorBody(code: NotAnsweredModel.noBot, message: ""), requestID: nil)
        let bare = NotAnsweredModel(guild: 2, raid: 41, service: FakeRaidService(), ping: FakeRaidPing(failure: noBot), hold: .seconds(4))
        #expect(await bare.mention(role: nil) == true)
        #expect(bare.mentionStatus == Copy.raidNaNoBot)
    }

    @Test("Discord DM shows only for a member who lets officers DM them")
    func directMessage() throws {
        let content = NotAnsweredPresenter(raid: try shortRaid(), guild: "Olympus XXVII", marks: [:], clock: RaidTime.clock)
            .content(filter: .notAnswered)
        #expect(content.rows.filter { $0.dm != nil }.map(\.name) == ["Zulmara"])
    }

    @Test("the root: under 60 with no guild, the recruiting guilds and one quiet line (never a block)")
    func composer() throws {
        let me = try RaidFixture.load(Me.self, "me")
        let character = try #require(me.mainCharacter)
        let composer = RaidRootComposer(presenter: presenter)
        let level = try #require(character.level)
        guard case let .noGuild(recruiting) = composer.content(.noGuild(character: character), extras: RaidRootExtras()) else {
            Issue.record("expected the recruiting guilds under 60")
            return
        }
        #expect(recruiting.raidsOpen == (level < 60 ? "Raids open at 60 · you're \(level)" : nil))
        guard case .noGuild = composer.content(.noGuild(character: nil), extras: RaidRootExtras()) else {
            Issue.record("expected the recruiting guilds")
            return
        }
        #expect(composer.content(.signedOut, extras: RaidRootExtras()) == .signedOut)
    }

    @Test("the root: a guild member under 60 sees the guild's raid, not the leveler")
    func guildMemberUnder60() async throws {
        let tab = RaidsTabModel(service: FakeRaidService(), cache: NoRaidCache(), now: { RaidTime.now })
        await tab.load()
        guard case let .loaded(content) = tab.state else { Issue.record("not loaded"); return }
        #expect((content.character.level ?? RaidFlowPresenter.maxLevel) < RaidFlowPresenter.maxLevel)
        var extras = RaidRootExtras()
        extras.ready = try RaidFixture.load(ReadyCheck.self, "character.ready")
        guard case let .guild(guild) = RaidRootComposer(presenter: presenter).content(tab.state, extras: extras) else {
            Issue.record("expected the guild's raid")
            return
        }
        // Before you pull is a 60's check (spec 2026-11-10-guild-tab.md 2.1).
        #expect(guild.ready == nil)
    }
}

/// Mentions as many as asked, or fails with an error.
struct FakeRaidPing: RaidPingService {
    var mentioned = 2
    var failure: APIError?

    func ping(guild: Int64, raid: Int64, role: String?) async throws(APIError) -> RaidPingResult {
        if let failure { throw failure }
        return RaidPingResult(
            posted: mentioned > 0, mentioned: mentioned, noDiscord: 0, noDiscordNames: [], nextAllowedAt: nil, recentlyMentioned: nil
        )
    }
}
