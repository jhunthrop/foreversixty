@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

@Suite("Intent answers: every spec sentence verbatim")
struct IntentAnswersTests {
    private let answers = SystemStory.answers

    @Test("signed up")
    func signedUp() throws {
        let raid = try SystemStory.raid()
        #expect(answers.signedUp(state: "in", raid: raid, characterName: "Lightningbolltz", role: "dps")
            == "You are In for Onyxia's Lair, Thursday at 8 pm, as DPS on Lightningbolltz. 25 of 40 signed up.")
    }

    @Test("the voice confirmation")
    func confirm() {
        let start = RaidTime.date(SystemStory.thursday)
        #expect(answers.confirm(name: "Lightningbolltz", state: "in", zone: "Onyxia's Lair", startsAt: start)
            == "Sign Lightningbolltz up as In for Onyxia's Lair, Thursday 8 pm?")
        #expect(answers.confirm(name: "Lightningbolltz", state: "late", zone: "Onyxia's Lair", startsAt: start)
            == "Sign Lightningbolltz up as Late for Onyxia's Lair, Thursday 8 pm?")
    }

    @Test("next raid, answered and unanswered")
    func nextRaid() throws {
        let raid = try #require(SystemStory.widget().nextRaid)
        #expect(answers.nextRaid(raid) == "Onyxia's Lair, Thursday at 8 pm. You are In. 24 of 40 signed up.")
        let open = try #require(WidgetFixtures.widget("widget.get", patch: [
            "my": .object(["state": .null, "spec": .string("Elemental"), "role": .string("dps")]),
        ]).nextRaid)
        #expect(answers.nextRaid(open) == "Onyxia's Lair, Thursday at 8 pm. You have not signed up. 24 of 40 signed up.")
    }

    @Test("a cancelled next raid")
    func cancelled() throws {
        let raid = try #require(WidgetFixtures.widget("widget.get", patch: ["status": .string("cancelled")]).nextRaid)
        #expect(answers.nextRaid(raid) == "Onyxia's Lair on Thursday is cancelled.")
    }

    @Test("who is coming, with and without anyone late")
    func whoIsComing() throws {
        let raid = try SystemStory.raid(counts: ["in": 23, "late": 1, "tanks": 3, "healers": 6, "dps": 14])
        #expect(answers.whoIsComing(raid) == "24 of 40 signed up: 3 tanks, 6 healers, 14 DPS, 1 late.")
        let onTime = try SystemStory.raid(counts: ["in": 20, "late": 0, "tanks": 1, "healers": 5, "dps": 14])
        #expect(answers.whoIsComing(onTime) == "20 of 40 signed up: 1 tank, 5 healers, 14 DPS.")
    }

    @Test("the Spotlight raid title")
    func raidTitle() {
        #expect(answers.raidTitle(zone: "Onyxia's Lair", startsAt: RaidTime.date(SystemStory.thursday)) == "Onyxia's Lair · Thu 8 pm")
        #expect(answers.raidTitle(zone: "Onyxia's Lair", startsAt: RaidTime.date("2026-12-11T02:30:00Z")) == "Onyxia's Lair · Thu 8:30 pm")
    }

    @Test("refusals reuse the app's sentences")
    func refusals() throws {
        #expect(IntentAnswers.refusal(try RaidFixture.error("guild.signup.error-raid-locked", status: 409))
            == "Sign-ups are locked. An officer can still change yours.")
        #expect(IntentAnswers.refusal(.unexpectedStatus(500)) == "That did not save. Try again.")
        #expect(IntentAnswers.refusal(.transport("offline")) == Copy.intentOffline)
    }

    @Test("the next upgrade snippet's lines")
    func upgrade() throws {
        let home = try HomeFixtures.home("character.home.yours")
        let answer = try #require(UpgradeAnswer(home: home, characterName: "Lightningbolltz"))
        #expect(answer.itemName == "Ironfeather Breastplate")
        #expect(answer.meta == "Chest · item level 58 · Leatherworking")
        #expect(answer.gain == "+3.8 DPS for Lightningbolltz. Replaces Vest of Elements.")
    }
}

@Suite("Which raid")
struct RaidPickTests {
    private let now = SystemStory.now

    @Test("the next raid within 24 hours, else the next planned one")
    func pick() throws {
        let later = try SystemStory.raid(id: 8, startsAt: "2026-12-15T02:00:00Z")
        let thursday = try SystemStory.raid(id: 7)
        #expect(RaidPick.pick([later, thursday], now: now) == .one(thursday))
        let tonight = try SystemStory.raid(id: 9, startsAt: "2026-12-09T02:00:00Z")
        #expect(RaidPick.pick([thursday, tonight], now: now) == .one(tonight))
    }

    @Test("two within 24 hours ask")
    func ask() throws {
        let first = try SystemStory.raid(id: 9, startsAt: "2026-12-09T02:00:00Z")
        let second = try SystemStory.raid(id: 10, startsAt: "2026-12-09T05:00:00Z")
        #expect(RaidPick.pick([second, first], now: now) == .ask([first, second]))
    }

    @Test("cancelled and ended raids are skipped; nothing left is none")
    func none() throws {
        let cancelled = try SystemStory.raid(id: 7, status: "cancelled")
        let ended = try SystemStory.raid(id: 6, startsAt: "2026-12-01T02:00:00Z")
        #expect(RaidPick.pick([cancelled, ended], now: now) == .none)
    }
}
