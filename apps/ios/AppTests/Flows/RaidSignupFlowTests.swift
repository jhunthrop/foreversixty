@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// Changing the answer on the Raid tab root (flows spec 4.2, F1), as Android's RaidSignupFlowTest:
/// RaidsTabModel, RaidAnswerModel and RaidRootComposer over the echoing fixture service. The step
/// captures are in FlowSnapshotTests (`raid-signup-NN-step`).
@MainActor
@Suite("Raid tab sign-up flow")
struct RaidSignupFlowTests {
    private static let late30 = 30
    private static let saved: [String] = ["selection", "success"]

    private func loaded() async throws -> RaidTabFlow {
        let flow = try RaidTabFlow()
        await flow.load()
        try #require(flow.head != nil, "the Raid tab draws no next raid")
        return flow
    }

    @Test("In, Late (its minutes open), 30 min (sent, Late · 30 min, the chips stay with 30 chosen), Bench and Out")
    func answers() async throws {
        let flow = try await loaded()
        var head = try #require(flow.head)
        #expect(head.selected == .signedIn)
        #expect(!head.lateOpen)
        #expect(head.status == nil)
        #expect(head.answersEnabled)

        await flow.answer(.late)
        head = try #require(flow.head)
        #expect(flow.service.sent.map(\.state) == ["late"])
        #expect(flow.service.sent.last?.lateMinutes == nil)
        #expect(head.selected == .late)
        #expect(head.lateOpen, "Late opens its minutes")
        #expect(head.lateMinutes == nil)
        #expect(head.status == .saved(Copy.commonSaved))
        #expect(flow.haptics.calls == Self.saved)

        await flow.late(Self.late30)
        head = try #require(flow.head)
        #expect(flow.service.sent.last?.state == "late")
        #expect(flow.service.sent.last?.lateMinutes == Self.late30)
        #expect(head.selected == .late)
        #expect(head.lateOpen, "the minutes stay open, as Android")
        #expect(head.lateMinutes == Self.late30, "30 is the chosen chip")
        #expect(head.selected == .late, "AS5: the Segmented is the answer; no chip repeats it")

        await flow.answer(.bench)
        head = try #require(flow.head)
        #expect(head.selected == .bench)
        #expect(!head.lateOpen)
        #expect(head.lateMinutes == nil)
        #expect(flow.service.sent.last?.state == "bench")
        #expect(flow.service.sent.last?.lateMinutes == nil)

        await flow.answer(.out)
        head = try #require(flow.head)
        #expect(head.selected == .out)
        #expect(flow.service.sent.map(\.state) == ["late", "late", "bench", "out"])
        #expect(flow.service.requests.allSatisfy { $0.key == FlowPlayer.key })
        #expect(flow.haptics.calls == Array(repeating: Self.saved, count: 4).flatMap(\.self))
    }

    @Test("a write in flight draws its answer at once; a failure returns to the answer before it with the error and Try again")
    func failure() async throws {
        let flow = try await loaded()
        let gate = FlowGate()
        flow.service.gate = gate
        flow.service.failWith = .unexpectedStatus(500)
        let write = flow.startAnswer(.late)
        #expect(await flowSettle { flow.service.inFlight == 1 })
        #expect(flow.head?.selected == .late, "the press draws before the write returns")
        #expect(flow.head?.lateOpen == true)
        #expect(flow.head?.busy == true)
        // A second press while the write is out waits.
        await flow.answer(.out)
        #expect(flow.service.requests.count == 1)

        gate.open()
        await write.value
        let head = try #require(flow.head)
        #expect(head.selected == .signedIn, "a failed write returns to In")
        #expect(!head.lateOpen)
        #expect(!head.busy)
        #expect(head.status == .error(Copy.commonFailed))
        #expect(flowDraws(flow.root, Copy.commonFailed))
        #expect(flow.answers.canRetry)
        #expect(flow.haptics.calls == ["selection", "error"])

        // Try again resends the failed answer.
        flow.service.gate = nil
        flow.service.failWith = nil
        await flow.retry()
        #expect(flow.head?.selected == .late)
        #expect(flow.head?.status == .saved(Copy.commonSaved))
        #expect(flow.service.sent.map(\.state) == ["late", "late"])
        #expect(!flow.answers.canRetry)
    }

    @Test("offline the answers disable with the offline sentence, under the strip, and a tap sends nothing")
    func offline() async throws {
        let flow = try await loaded()
        flow.service.offline = true
        await flow.load()
        let guild = try #require(flow.guild)
        #expect(guild.offlineStamp != nil, "the offline strip says why")
        let head = try #require(guild.next)
        #expect(head.selected == .signedIn)
        #expect(!head.answersEnabled)
        #expect(head.status == .info(Copy.raidOffline))

        let before = flow.service.requests.count
        await flow.answer(.bench)
        #expect(flow.head?.selected == .signedIn)
        #expect(flow.service.requests.count == before)
        #expect(flow.haptics.calls == ["selection", "error"])
    }

    @Test("a lost connection on the write disables the answers until the tab reloads online")
    func writeOffline() async throws {
        let flow = try await loaded()
        flow.service.failWith = .transport("offline")
        await flow.answer(.late)
        #expect(flow.head?.answersEnabled == false)
        #expect(flow.head?.status == .error(Copy.raidOffline))
        flow.service.failWith = nil
        await flow.load()
        #expect(flow.head?.answersEnabled == true)
    }
}
