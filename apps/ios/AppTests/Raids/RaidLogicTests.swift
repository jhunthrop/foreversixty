@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// The F4 test list, one per behaviour; Android's RaidLogicTest carries the same names.
@Suite("Raid times")
struct RaidClockTests {
    let clock = RaidTime.clock

    @Test("clocks read 7:45 pm in the viewer's zone, and the date block THU 10 DEC")
    func clocks() throws {
        let invite = RaidTime.date("2026-12-20T00:45:00Z")
        #expect(clock.clock(invite) == "6:45 pm")
        let block = clock.dateBlock(invite)
        #expect(block.weekday == "SAT" && block.day == "19" && block.month == "DEC")
        #expect(clock.longDate(invite) == "Saturday 19 December")
    }

    @Test("the countdown reads like the site's: in, starting now, started, ended")
    func countdown() throws {
        let start = RaidTime.date("2026-12-20T00:45:00Z")
        let end = RaidTime.date("2026-12-20T04:00:00Z")
        #expect(clock.countdown(to: start, end: end, now: RaidTime.now) == "in 2 days 2 h")
        #expect(clock.countdown(to: start, end: end, now: start.addingTimeInterval(-30)) == "starting now")
        #expect(clock.countdown(to: start, end: end, now: start.addingTimeInterval(1200)) == "started 20 min ago")
        #expect(clock.countdown(to: start, end: end, now: end.addingTimeInterval(3600)) == "ended 1 h ago")
        #expect(RaidClock.span(5 * 3600 + 20 * 60) == "5 h 20 min")
    }

    @Test("VoiceOver hears whole words, never h")
    func spanWords() throws {
        #expect(RaidClock.spanWords(2 * 86400 + 3 * 3600) == "2 days 3 hours")
        #expect(RaidClock.spanWords(3600 + 60) == "1 hour 1 minute")
    }

    @Test("a planned row's when: tonight, tomorrow, weekday, then the date")
    func when() throws {
        let now = RaidTime.date("2026-12-17T18:00:00Z")
        #expect(clock.when(RaidTime.date("2026-12-18T02:00:00Z"), now: now) == "Tonight 8:00 pm")
        #expect(clock.when(RaidTime.date("2026-12-19T02:00:00Z"), now: now) == "Tomorrow 8:00 pm")
        #expect(clock.when(RaidTime.date("2026-12-20T01:00:00Z"), now: now) == "Saturday 7:00 pm")
        #expect(clock.when(RaidTime.date("2027-01-08T02:00:00Z"), now: now) == "Thu 7 Jan, 8:00 pm")
    }

    @Test("the guild-time line shows only when the offsets differ")
    func guildTime() throws {
        let start = RaidTime.date("2026-12-20T01:00:00Z")
        #expect(!clock.offsetDiffers(at: start, guildZone: RaidTime.chicago))
        let berlin = RaidClock(zone: try #require(TimeZone(identifier: "Europe/Berlin")))
        #expect(berlin.offsetDiffers(at: start, guildZone: RaidTime.chicago))
        #expect(berlin.guildClock(start, guildZone: RaidTime.chicago).hasPrefix("7:00 pm"))
    }
}

@Suite("Raid facts")
struct RaidFactsTests {
    @Test("the count is In plus Late, with the bar only when the size is known")
    func count() throws {
        let facts = RaidFacts(raid: try RaidFixture.raid(), now: RaidTime.now)
        #expect(facts.countLine == "24 of 40 signed up")
        #expect(facts.ratio == "24/40")
        #expect(facts.breakdown == "3 tanks · 6 healers · 14 DPS · 1 late")
        let barrow = try #require(RaidFixture.load(GuildRaids.self, "guild.raids").planned?[1])
        let noSize = RaidFacts(raid: barrow, now: RaidTime.now)
        #expect(noSize.countLine == "11 signed up")
        #expect(noSize.fill == nil)
    }

    @Test("one tank reads 1 tank; no late clause when nobody is late")
    func singularParts() throws {
        let ony = try #require(RaidFixture.load(GuildRaids.self, "guild.raids").planned?[2])
        #expect(RaidFacts(raid: ony, now: RaidTime.now).breakdown == "1 tank · 2 healers · 5 DPS")
    }

    @Test("a raid past its end reads Ended and is no longer signable")
    func ended() throws {
        let after = RaidTime.date("2026-12-21T00:00:00Z")
        let facts = RaidFacts(raid: try RaidFixture.raid(), now: after)
        #expect(facts.chip == .ended)
        #expect(!facts.isSignable)
    }

    @Test("the countdown runs to invites, else to the pull")
    func countdownTarget() throws {
        #expect(RaidFacts(raid: try RaidFixture.raid(), now: RaidTime.now).countdownTarget == RaidTime.date("2026-12-20T00:45:00Z"))
        #expect(RaidFacts(raid: try RaidFixture.raid("guild.raid.no-invite"), now: RaidTime.now).countdownTarget == RaidTime.date("2026-12-20T01:00:00Z"))
    }

    @Test("officer tools come from viewer.can only")
    func officerTools() throws {
        #expect(!RaidFacts(raid: try RaidFixture.raid(), now: RaidTime.now).showsOfficerTools)
        let officer = RaidFacts(raid: try RaidFixture.raid("guild.raid.officer"), now: RaidTime.now)
        #expect(officer.canLock && officer.canMove && officer.showsOfficerTools)
    }

    @Test("a planned row's chip: Cancelled, Locked, Ended, else the viewer's answer, else Not signed up")
    func rowChip() throws {
        let planned = try #require(RaidFixture.load(GuildRaids.self, "guild.raids").planned)
        #expect(RaidFacts(raid: planned[0], now: RaidTime.now).rowChip(currentKey: nil) == .signed("In · DPS"))
        #expect(RaidFacts(raid: planned[1], now: RaidTime.now).rowChip(currentKey: nil) == .notSigned)
        #expect(RaidFacts(raid: planned[2], now: RaidTime.now).rowChip(currentKey: nil) == .status(.locked))
        #expect(RaidFacts(raid: try RaidFixture.raid("guild.raid.cancelled"), now: RaidTime.now).rowChip(currentKey: nil) == .status(.cancelled))
    }

    @Test("the next raid is the earliest signable one; Also planned lists the rest")
    func nextRaid() throws {
        let planned = try #require(RaidFixture.load(GuildRaids.self, "guild.raids").planned)
        let split = PlannedRaids(planned.reversed(), now: RaidTime.now)
        #expect(split.next?.id == 31)
        #expect(split.also.map(\.id) == [32, 33])
    }
}

@Suite("Roster")
struct RosterTests {
    @Test("groups in order: Tanks, Healers, DPS, Late, Bench, Out; class order, then name")
    func groups() throws {
        let roster = RosterGroups(raid: try RaidFixture.raid())
        #expect(roster.groups.map(\.kind) == [.tanks, .healers, .dps, .late, .bench, .out])
        #expect(roster.groups[0].rows.map(\.name) == ["Ironhide", "Kraggor", "Bramblehoof"])
        #expect(roster.groups[2].rows.first?.name == "Draknar")
        #expect(roster.notSigned.map(\.name) == ["Brokk", "Zulmara", "Sennah"])
    }

    @Test("an empty role says so; an empty Late, Bench or Out is left out")
    func emptyGroups() throws {
        let barrow = try #require(RaidFixture.load(GuildRaids.self, "guild.raids").planned?[1])
        let roster = RosterGroups(raid: barrow)
        #expect(roster.groups.map(\.kind) == [.tanks, .healers, .dps])
        #expect(roster.groups[0].kind.emptyLine == "No tanks yet.")
    }

    @Test("an officer's move shows at once in the new group")
    func overrides() throws {
        let roster = RosterGroups(raid: try RaidFixture.raid(), overrides: ["us/pvp/kraggor": .bench])
        #expect(roster.groups.first { $0.kind == .bench }?.rows.map(\.name).contains("Kraggor") == true)
    }

    @Test("the classes strip counts In and Late per class, and reads every class to VoiceOver")
    func classes() throws {
        let strip = ClassStrip(raid: try RaidFixture.raid())
        #expect(strip.entries.map(\.count) == [5, 0, 2, 2, 4, 4, 2, 2, 3])
        #expect(strip.accessibilityLabel.hasPrefix("Classes coming: Warrior 5, Hunter 2"))
        #expect(strip.accessibilityLabel.hasSuffix("no Paladins"))
    }

    @Test("the move menu offers the spec's items per place")
    func moveMenu() {
        #expect(PlaceMenu.items(for: .signedIn).map(\.1) == [.bench, .out])
        #expect(PlaceMenu.items(for: nil).map(\.1) == [.signedIn, .bench])
        #expect(PlaceMenu.items(for: .bench)[1].2)
    }

    @Test("Copy the list: header, count, then each group's names")
    func copyList() throws {
        let text = RosterCopy.text(raid: try RaidFixture.raid(), clock: RaidTime.clock, now: RaidTime.now)
        let lines = text.components(separatedBy: "\n")
        #expect(lines[0].hasPrefix("Onyxia's Lair · Sat 19 Dec, 7:00 pm CST · <t:"))
        #expect(lines[1] == "24 of 40: 3 tanks, 6 healers, 14 DPS, 1 late")
        #expect(lines[2] == "Tanks: Ironhide, Kraggor, Bramblehoof")
        #expect(lines.last == "Out: Wyrmtalon, Snaggle, Felgrin")
    }
}

@MainActor
@Suite("Sign-up")
struct SignupControllerTests {
    let key = "us/pvp/obnoxious-yell"

    @Test("a press draws at once, saves, says Saved. and plays selection then success")
    func pressSaves() async throws {
        let service = FakeRaidService()
        let haptics = RecordingHaptics()
        let controller = try makeController(service: service, haptics: haptics)
        await controller.press(key, .late)
        #expect(service.signups.first?.request.state == "late")
        #expect(controller.block(key).status == .saved("Saved."))
        #expect(haptics.calls == ["selection", "success"])
    }

    @Test("the PUT carries the stored note, since the PUT writes it")
    func keepsNote() async throws {
        let service = FakeRaidService()
        let controller = try makeController(service: service)
        controller.editNote(key, "Late from work")
        await controller.press(key, .bench)
        #expect(service.signups.first?.request.note == "Late from work")
    }

    @Test("a refused press restores the previous answer with the API's sentence (F4.11)")
    func refusedReverts() async throws {
        let service = FakeRaidService()
        service.signupAnswer = { _ throws(APIError) in throw (try? RaidFixture.error("guild.signup.error-raid-locked", status: 409)) ?? .unexpectedStatus(0) }
        let haptics = RecordingHaptics()
        let controller = try makeController(service: service, haptics: haptics)
        await controller.press(key, .late)
        #expect(controller.block(key).shown == .signedIn)
        #expect(controller.block(key).status == .error("This raid was locked while you were signing up."))
        #expect(haptics.calls == ["selection", "error"])
    }

    @Test("every refusal maps to its sentence")
    func refusals() throws {
        #expect(SignupRefusal.sentence(for: try RaidFixture.error("guild.signup.error-raid-cancelled", status: 409)) == "This raid was cancelled while you were signing up.")
        #expect(SignupRefusal.sentence(for: try RaidFixture.error("guild.signup.error-rate-limited", status: 429)) == "Too many sign-up changes. Wait a minute, then try again.")
        #expect(SignupRefusal.sentence(for: try RaidFixture.error("guild.signup.error-plan-required", status: 403)) == "Not saved. This guild's officer tools have ended.")
        #expect(SignupRefusal.sentence(for: try RaidFixture.error("guild.signup.error-forbidden", status: 403)) == "You can only sign up your own characters.")
        #expect(SignupRefusal.sentence(for: .transport("offline")) == "You are offline. Sign-ups need a connection.")
        #expect(SignupRefusal.sentence(for: .unexpectedStatus(500)) == "That did not save. Try again.")
    }

    @Test("a lost connection disables the segments with the offline line")
    func offline() async throws {
        let service = FakeRaidService()
        service.signupAnswer = { _ throws(APIError) in throw .transport("offline") }
        let controller = try makeController(service: service)
        await controller.press(key, .out)
        #expect(controller.offline)
        #expect(!controller.segmentsEnabled(key))
    }

    @Test("In with no saved spec asks which spec; the pick sends it and says it is kept")
    func specAsk() async throws {
        let service = FakeRaidService()
        let ask = "us/pvp/lightningbolltz"
        let controller = try makeController(try RaidFixture.raid("guild.raid.ask-spec"), service: service)
        await controller.press(ask, .signedIn)
        #expect(controller.specAsk == .ask(characterKey: ask, state: .signedIn))
        #expect(controller.block(ask).pending == .signedIn)
        #expect(service.signups.isEmpty)
        let elemental = try #require(controller.character(ask)?.specOptions?.first { $0.spec == "shaman-elemental" })
        await controller.pickSpec(elemental)
        #expect(service.signups.first?.request.spec == "shaman-elemental")
        #expect(controller.block(ask).status == .saved("Saved. Elemental is kept for Lightningbolltz."))
    }

    @Test("closing the spec sheet drops the press; Out never asks")
    func specSheetClose() async throws {
        let service = FakeRaidService()
        let ask = "us/pvp/lightningbolltz"
        let controller = try makeController(try RaidFixture.raid("guild.raid.ask-spec"), service: service)
        await controller.press(ask, .bench)
        controller.closeSpecSheet()
        #expect(controller.block(ask).pending == nil && controller.specAsk == nil && service.signups.isEmpty)
        await controller.press(ask, .out)
        #expect(controller.specAsk == nil)
        #expect(service.signups.first?.request.state == "out")
    }

    @Test("Change spec resends the current answer with the new spec")
    func changeSpec() async throws {
        let service = FakeRaidService()
        let controller = try makeController(service: service)
        controller.changeSpec(key)
        #expect(controller.specAsk == .change(characterKey: key))
        let arms = try #require(controller.character(key)?.specOptions?.first { $0.spec == "warrior-arms" })
        await controller.pickSpec(arms)
        #expect(service.signups.first?.request.state == "in")
        #expect(service.signups.first?.request.spec == "warrior-arms")
    }

    @Test("moved names the other character on its own block (R3)")
    func moved() async throws {
        let service = FakeRaidService()
        let answer = SignupAnswer(
            raid: try RaidFixture.raid("guild.raid.two-characters"),
            moved: try RaidFixture.load(SignupAnswer.self, "guild.signup.moved").moved
        )
        service.signupAnswer = { _ in answer }
        let controller = try makeController(try RaidFixture.raid("guild.raid.two-characters"), service: service)
        await controller.press(key, .signedIn)
        #expect(controller.block("us/pvp/yellfury").status == .info("Yellfury moved to Bench."))
    }

    @Test("a locked raid disables a member's segments and keeps an officer's")
    func locked() throws {
        let member = try makeController(try RaidFixture.raid("guild.raid.locked"))
        #expect(!member.segmentsEnabled(key))
        #expect(member.restingStatus(key) == .info("Sign-ups are locked. An officer can still change yours."))
        let cancelled = try makeController(try RaidFixture.raid("guild.raid.cancelled"))
        #expect(!cancelled.segmentsEnabled(key))
        let officer = try makeController(try RaidFixture.raid("guild.raid.officer-locked"))
        #expect(officer.segmentsEnabled(key))
    }

    @Test("an unverified character has no controls")
    func waitingCharacter() throws {
        let controller = try makeController(try RaidFixture.raid("guild.raid.two-characters"))
        #expect(!controller.segmentsEnabled("us/pvp/vorgash"))
    }

    @Test("the reminder ask shows after a first In, at most twice, 30 days apart")
    func reminderAsk() async throws {
        let store = MemoryAskStore()
        let controller = try makeController(undetermined: true, askStore: store)
        await controller.press(key, .signedIn)
        #expect(controller.showsReminderAsk && store.shown == 1)
        #expect(!ReminderAskPolicy(shown: 1, lastShown: RaidTime.now).mayAsk(now: RaidTime.now.addingTimeInterval(86400)))
        #expect(ReminderAskPolicy(shown: 1, lastShown: RaidTime.now).mayAsk(now: RaidTime.now.addingTimeInterval(31 * 86400)))
        #expect(!ReminderAskPolicy(shown: 2, lastShown: nil).mayAsk(now: RaidTime.now))
        let denied = try makeController(undetermined: false)
        await denied.press(key, .signedIn)
        #expect(!denied.showsReminderAsk)
    }

    @Test("an officer's move sends the row's state and role; a failure restores the row")
    func officerMove() async throws {
        let service = FakeRaidService()
        let officer = try RaidFixture.raid("guild.raid.officer")
        service.signupAnswer = { _ in SignupAnswer(raid: officer, moved: nil) }
        let controller = try makeController(officer, service: service)
        let kraggor = RosterGroups(raid: controller.raid).groups[0].rows.first { $0.name == "Kraggor" }!
        await controller.move(kraggor, to: .bench)
        #expect(service.signups.first?.key == "us/pvp/kraggor")
        #expect(service.signups.first?.request.state == "bench" && service.signups.first?.request.role == "tank")
        service.signupAnswer = { _ throws(APIError) in throw .unexpectedStatus(500) }
        await controller.move(kraggor, to: .out)
        #expect(controller.rosterOverrides.isEmpty)
        #expect(controller.rowErrors["us/pvp/kraggor"] == "That did not save. Try again.")
    }

    @Test("a member cannot move anyone")
    func memberCannotMove() async throws {
        let service = FakeRaidService()
        let controller = try makeController(service: service)
        await controller.move(RosterGroups(raid: controller.raid).groups[0].rows[0], to: .bench)
        #expect(service.signups.isEmpty)
    }

    @Test("Lock sign-ups locks and toasts who was told")
    func lock() async throws {
        let service = FakeRaidService()
        let controller = try makeController(try RaidFixture.raid("guild.raid.officer"), service: service)
        await controller.toggleLock()
        #expect(service.locks == ["lock"])
        #expect(controller.toast == "Sign-ups locked. 23 In and 1 on the Bench were told.")
    }
}

@MainActor
@Suite("Raids tab")
struct RaidsTabModelTests {
    func model(_ service: FakeRaidService, cache: any RaidCache = NoRaidCache()) -> RaidsTabModel {
        RaidsTabModel(service: service, cache: cache, now: { RaidTime.now })
    }

    @Test("signed out on a 401")
    func signedOut() async throws {
        let service = FakeRaidService()
        service.meAnswer = { () throws(APIError) in throw .unauthorized(requestID: nil) }
        let tab = model(service)
        await tab.load()
        #expect(tab.state == .signedOut)
    }

    @Test("a character in no guild shows No guild yet with its name")
    func noGuild() async throws {
        let tab = model(FakeRaidService())
        tab.currentKey = "us/pvp/yellfury"
        await tab.load()
        #expect(tab.state == .noGuild(name: "Yellfury"))
    }

    @Test("an unverified member waits for approval")
    func waiting() async throws {
        let service = FakeRaidService()
        let me = try RaidFixture.load(Me.self, "me.waiting")
        service.meAnswer = { me }
        let tab = model(service)
        await tab.load()
        guard case .waiting = tab.state else { Issue.record("expected waiting, got \(tab.state)"); return }
    }

    @Test("loaded: the planned raids, three raid nights at most, the officer's waiting count")
    func loaded() async throws {
        let tab = model(FakeRaidService())
        await tab.load()
        guard case let .loaded(content) = tab.state else { Issue.record("not loaded"); return }
        #expect(content.planned.count == 3)
        #expect(content.waiting == 2)
        #expect(content.head.line == "Horde · US PvP · 24 on the roster")
    }

    @Test("offline with a cached answer shows it under the strip; with none, the error")
    func offline() async throws {
        let cache = MemoryRaidCache()
        let service = FakeRaidService()
        let warm = model(service, cache: cache)
        await warm.load()
        service.raidsAnswer = { () throws(APIError) in throw .transport("offline") }
        await warm.load()
        guard case let .loaded(content) = warm.state else { Issue.record("not loaded"); return }
        #expect(content.offlineSince == RaidTime.now)
        let cold = model(service)
        await cold.load()
        #expect(cold.state == .failed)
    }
}

@Suite("Raid notifications")
struct RaidPushTests {
    func payload(_ extra: [String: String] = [:]) throws -> RaidPushPayload {
        let raw = try #require(JSONSerialization.jsonObject(with: RaidFixture.data("push.raid-posted.apns")) as? [String: Any])
        var userInfo = raw
        var data = try #require(raw["data"] as? [String: String])
        data["raid_id"] = "31"
        data.merge(extra) { $1 }
        userInfo["data"] = data
        return try #require(RaidPushPayload(userInfo: userInfo))
    }

    @Test("the payload names the raid, the character and the answers")
    func parse() throws {
        let push = try payload()
        #expect(push.guild == 2 && push.raid == 31)
        #expect(push.characterKey == "us/pvp/obnoxious-yell")
        #expect(!push.askSpec)
        #expect(push.answers["in"]?.title == "You are In for Onyxia's Lair")
        #expect(RaidPushCategory.state(forAction: "signup_late") == .late)
    }

    @Test("In from the lock screen sends the PUT for the chosen character, keeping its note")
    func answersIn() async throws {
        let service = FakeRaidService()
        let outcome = await RaidPushAnswerer(service: service).answer(try payload(), state: .signedIn)
        #expect(service.signups.first?.key == "us/pvp/obnoxious-yell")
        #expect(service.signups.first?.request.state == "in")
        guard case let .answered(answer) = outcome else { Issue.record("not answered"); return }
        #expect(answer.title == "You are In for Onyxia's Lair")
    }

    @Test("a moved character is named after the answer (R3)")
    func moved() async throws {
        let service = FakeRaidService()
        let moved = try RaidFixture.load(SignupAnswer.self, "guild.signup.moved")
        service.signupAnswer = { _ in moved }
        let outcome = await RaidPushAnswerer(service: service).answer(try payload(), state: .signedIn)
        guard case let .answered(answer) = outcome else { Issue.record("not answered"); return }
        #expect(answer.body.hasSuffix(" Yellfury moved to Bench."))
    }

    @Test("a locked raid posts the locked answer; a failure posts the failed one")
    func refused() async throws {
        let service = FakeRaidService()
        service.signupAnswer = { _ throws(APIError) in throw (try? RaidFixture.error("guild.signup.error-raid-locked", status: 409)) ?? .unexpectedStatus(0) }
        let locked = await RaidPushAnswerer(service: service).answer(try payload(), state: .late)
        #expect(locked == .answered(.init(title: "Sign-ups for Onyxia's Lair are locked", body: "An officer can still change yours.")))
        service.raidAnswer = { () throws(APIError) in throw .transport("offline") }
        let failed = await RaidPushAnswerer(service: service).answer(try payload(), state: .late)
        guard case let .answered(answer) = failed else { Issue.record("not answered"); return }
        #expect(answer.title == "That did not save")
    }

    @Test("a character with no spec opens the raid page with the spec sheet")
    func asksSpec() async throws {
        let service = FakeRaidService()
        let ask = try RaidFixture.raid("guild.raid.ask-spec")
        service.raidAnswer = { ask }
        let outcome = await RaidPushAnswerer(service: service).answer(try payload(["character_key": "us/pvp/lightningbolltz"]), state: .signedIn)
        #expect(outcome == .open(.raid(guild: 2, raid: 31, spec: "ask", state: nil)))
        #expect(service.signups.isEmpty)
    }

    @Test("with no pick and two unsigned characters, the tap opens the page")
    func noPick() throws {
        let push = try payload(["character_key": ""])
        #expect(RaidPushAnswerer.character(for: push, in: try RaidFixture.raid("guild.raid.two-characters"))?.characterKey == "us/pvp/obnoxious-yell")
        #expect(RaidPushAnswerer.character(for: push, in: try RaidFixture.raid("guild.raid.ask-spec"))?.characterKey == "us/pvp/lightningbolltz")
    }
}

@MainActor
@Suite("Raid links")
struct RaidRouterTests {
    @Test("a raid link pushes the raid page; night pushes the page then the night; spec=ask presses In")
    func routes() throws {
        let router = RaidsRouter()
        router.open(.raid(guild: 2, raid: 41, spec: "ask", state: nil))
        #expect(router.path == [.raid(guild: 2, raid: 41, roster: false, open: RaidOpenAction(state: .signedIn))])
        router.open(.raidNight(guild: 2, raid: 41))
        #expect(router.path == [.raid(guild: 2, raid: 41, roster: false, open: nil), .night(raid: 41)])
        router.open(.guild(guild: 2))
        #expect(router.path.isEmpty)
        #expect(!RaidsRouter.handles(.plan))
    }
}

@Suite("Notes")
struct NotesTextTests {
    @Test("notes_html keeps bold and https links, drops other tags")
    func notes() throws {
        let html = "<p>Bring <strong>fire resistance</strong>. See <a href=\"https://foreversixty.gg/guides/onyxia\">the guide</a>.</p>\n<p>Second &amp; last</p>"
        let text = NotesText.attributed(html, base: .body, bold: .headline, link: .yellow)
        #expect(String(text.characters) == "Bring fire resistance. See the guide.\nSecond & last")
        #expect(text.runs.contains { $0.link == URL(string: "https://foreversixty.gg/guides/onyxia") })
    }
}
