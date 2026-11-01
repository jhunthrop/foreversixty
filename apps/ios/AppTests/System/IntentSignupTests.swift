@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

@Suite("Sign up by voice and Shortcuts")
struct IntentSignupTests {
    private func signup(_ calls: FakeIntentCalls, enabled: Bool = true) -> IntentSignup {
        IntentSignup(calls: calls, answers: SystemStory.answers, now: { SystemStory.now }, enabled: { enabled })
    }

    private func calls(_ planned: [RaidView]? = nil) throws -> FakeIntentCalls {
        let calls = FakeIntentCalls()
        calls.planned = try planned ?? [SystemStory.raid()]
        return calls
    }

    @Test("by voice: asks once, then answers with the server's count")
    func voice() async throws {
        let calls = try calls()
        let asker = ScriptedAsker()
        let reply = try await signup(calls).run(state: "in", character: nil, askFirst: true, asker: asker)
        #expect(asker.questions == ["Sign Lightningbolltz up as In for Onyxia's Lair, Thursday 8 pm?"])
        #expect(reply == IntentReply(sentence: "You are In for Onyxia's Lair, Thursday at 8 pm, as DPS on Lightningbolltz. 25 of 40 signed up.", saved: true))
        #expect(calls.signups.map(\.request.state) == ["in"])
    }

    @Test("from Shortcuts, the Action button or the Control: never asks")
    func shortcut() async throws {
        let calls = try calls()
        let asker = ScriptedAsker()
        let reply = try await signup(calls).run(state: "late", character: nil, askFirst: false, asker: asker)
        #expect(asker.questions.isEmpty)
        #expect(reply.saved)
        #expect(calls.signups.first?.request.spec == nil)
    }

    @Test("Cancel on the confirmation saves nothing and says nothing")
    func cancel() async throws {
        let calls = try calls()
        let asker = ScriptedAsker()
        asker.confirms = false
        let reply = try await signup(calls).run(state: "in", character: nil, askFirst: true, asker: asker)
        #expect(reply == .cancelled)
        #expect(calls.signups.isEmpty)
    }

    @Test("two raids within 24 hours ask Which raid?")
    func which() async throws {
        let first = try SystemStory.raid(id: 9, startsAt: "2026-12-09T02:00:00Z")
        let second = try SystemStory.raid(id: 10, startsAt: "2026-12-09T05:00:00Z")
        let calls = try calls([first, second])
        let asker = ScriptedAsker()
        asker.choice = 1
        _ = try await signup(calls).run(state: "in", character: nil, askFirst: false, asker: asker)
        #expect(asker.questions == ["Which raid?"])
        #expect(asker.options == [["Onyxia's Lair · Tue 8 pm", "Onyxia's Lair · Tue 11 pm"]])
        #expect(calls.signups.map(\.raid) == [10])
    }

    @Test("no raid spec asks the spec question with the specs as options")
    func spec() async throws {
        let calls = try calls([SystemStory.raid(raidSpec: nil)])
        let asker = ScriptedAsker()
        asker.choice = 2
        _ = try await signup(calls).run(state: "in", character: nil, askFirst: false, asker: asker)
        #expect(asker.questions == ["Which spec is Lightningbolltz on OLYMPUS XXVII?"])
        #expect(asker.options == [["Elemental", "Enhancement", "Restoration"]])
        #expect(calls.signups.first?.request.spec == "shaman-restoration")
    }

    @Test("Out never asks the spec")
    func outNoSpec() async throws {
        let calls = try calls([SystemStory.raid(raidSpec: nil)])
        let asker = ScriptedAsker()
        _ = try await signup(calls).run(state: "out", character: nil, askFirst: false, asker: asker)
        #expect(asker.questions.isEmpty)
    }

    @Test("a locked raid answers the app's refusal")
    func locked() async throws {
        let calls = try calls()
        let locked = try RaidFixture.error("guild.signup.error-raid-locked", status: 409)
        calls.signupAnswer = { _ throws(APIError) in throw locked }
        let reply = try await signup(calls).run(state: "in", character: nil, askFirst: false, asker: ScriptedAsker())
        #expect(reply == .refused("Sign-ups are locked. An officer can still change yours."))
    }

    @Test("a failed save answers That did not save")
    func failed() async throws {
        let calls = try calls()
        calls.signupAnswer = { _ throws(APIError) in throw .unexpectedStatus(500) }
        let reply = try await signup(calls).run(state: "in", character: nil, askFirst: false, asker: ScriptedAsker())
        #expect(reply == .refused("That did not save. Try again."))
    }

    @Test("switch off, signed out, no raid")
    func refusals() async throws {
        let off = try await signup(try calls(), enabled: false).run(state: "in", character: nil, askFirst: false, asker: ScriptedAsker())
        #expect(off == .refused(Copy.intentOff))
        let signedOut = try calls()
        signedOut.widgetAnswer = { () throws(APIError) in throw .unauthorized(requestID: nil) }
        #expect(try await signup(signedOut).run(state: "in", character: nil, askFirst: false, asker: ScriptedAsker()) == .refused(Copy.intentSignedOut))
        let empty = try calls([])
        #expect(try await signup(empty).run(state: "in", character: nil, askFirst: false, asker: ScriptedAsker()) == .refused("No raid planned."))
    }
}

@Suite("Read-only intents")
struct IntentQueriesTests {
    private func queries(_ calls: FakeIntentCalls = FakeIntentCalls(), enabled: Bool = true) -> IntentQueries {
        IntentQueries(calls: calls, answers: SystemStory.answers, enabled: { enabled })
    }

    @Test("next raid and who is coming")
    func raidAnswers() async throws {
        let calls = FakeIntentCalls()
        calls.raidAnswer = { () throws(APIError) in
            do { return try SystemStory.raid(counts: ["in": 23, "late": 1, "tanks": 3, "healers": 6, "dps": 14]) } catch { throw .missingFixture("\(error)") }
        }
        #expect(await queries(calls).nextRaid(character: nil) == "Onyxia's Lair, Thursday at 8 pm. You are In. 24 of 40 signed up.")
        #expect(await queries(calls).whoIsComing(character: nil) == "24 of 40 signed up: 3 tanks, 6 healers, 14 DPS, 1 late.")
    }

    @Test("next upgrade: the snippet, or a sentence")
    func upgrade() async throws {
        let reply = await queries().nextUpgrade(character: nil)
        guard case let .upgrade(answer, build) = reply else { Issue.record("no upgrade"); return }
        #expect(answer.itemName == "Ironfeather Breastplate")
        #expect(!build.isEmpty)
        #expect(await queries(enabled: false).nextUpgrade(character: nil) == .sentence(Copy.intentOff))
        let offline = FakeIntentCalls()
        offline.widgetAnswer = { () throws(APIError) in throw .transport("offline") }
        #expect(await queries(offline).nextRaid(character: nil) == Copy.intentOffline)
    }
}
