@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// Answering from Home's raid card (flows spec 4.1, F1), as Android's HomeSignupFlowTest:
/// HomeFlowModel and HomeFlowPresenter over the echoing fixture service. The step captures are in
/// FlowSnapshotTests (`home-signup-NN-step`).
@MainActor
@Suite("Home card sign-up flow")
struct HomeSignupFlowTests {
    private func loaded() async throws -> HomeCardFlow {
        let flow = try HomeCardFlow()
        await flow.load()
        #expect(flow.model.phase == .loaded)
        guard case .raidCard? = flow.blocks.first else {
            Issue.record("Home draws no raid card: \(flow.blocks)")
            throw CancellationError()
        }
        return flow
    }

    @Test("In, Late (its minutes open), 30 min (sent with late_minutes 30, the card reads Late · 30 min)")
    func answers() async throws {
        let flow = try await loaded()
        var card = try #require(flow.card)
        #expect(card.selected == .signedIn)
        #expect(card.raidID == FlowPlayer.raid)
        #expect(!card.lateOpen)

        await flow.model.answer(.late)
        card = try #require(flow.card)
        #expect(flow.service.sent.map(\.state) == ["late"])
        #expect(flow.service.sent.last?.lateMinutes == nil)
        #expect(card.selected == .late)
        #expect(card.lateOpen, "Late opens its minutes")
        #expect(!card.answer.contains(SignupState.late.label), "AS5: the Segmented shows the answer; no sentence repeats it")
        #expect(flow.model.answerError == nil)

        await flow.model.late(30)
        card = try #require(flow.card)
        #expect(flow.service.sent.last?.state == "late")
        #expect(flow.service.sent.last?.lateMinutes == 30)
        #expect(card.selected == .late)
        #expect(card.lateMinutes == 30)
        #expect(card.lateOpen, "the minutes stay open with 30 chosen")
        #expect(card.status == .saved(Copy.commonSaved))
        #expect(flow.haptics.calls == ["selection", "success", "selection", "success"])
        #expect(flow.service.requests.allSatisfy { $0.key == FlowPlayer.key })
    }

    @Test("offline Home's answers disable with the offline sentence and send nothing")
    func offline() async throws {
        let flow = try await loaded()
        flow.model.offline = true
        let card = try #require(flow.card)
        #expect(!card.answersEnabled)
        #expect(card.status == .info(Copy.raidOffline))
        await flow.model.answer(.bench)
        #expect(flow.service.requests.isEmpty)
        #expect(flow.card?.selected == .signedIn)
    }

    @Test("a write in flight draws Late at once; the failure returns to In with the error and Try again")
    func failure() async throws {
        let flow = try await loaded()
        let gate = FlowGate()
        flow.service.gate = gate
        flow.service.failWith = .unexpectedStatus(500)
        let write = Task { await flow.model.answer(.late) }
        #expect(await flowSettle { flow.service.inFlight == 1 })
        #expect(flow.card?.selected == .late)
        #expect(flow.card?.lateOpen == true)

        gate.open()
        await write.value
        let card = try #require(flow.card)
        #expect(card.selected == .signedIn)
        #expect(!card.lateOpen)
        #expect(flow.model.answerError == Copy.commonFailed)
        #expect(card.status == .error(Copy.commonFailed))
        #expect(flowDraws(flow.blocks, Copy.commonFailed))
        #expect(flow.haptics.calls == ["selection", "error"])

        flow.service.gate = nil
        flow.service.failWith = nil
        await flow.model.retry()
        #expect(flow.card?.selected == .late)
        #expect(flow.card?.status == .saved(Copy.commonSaved))
    }
}
