@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// F5.3 (spec §5.5): the result's words and gear, and the store's seen and offline rules.
@MainActor
@Suite("Sim result")
struct SimResultTests {
    private let specs = SpecCatalog.bundled()

    private func result() throws -> SimResult { try RaidFixture.load(SimResult.self, "sims.result") }

    private func listRow(_ id: String) throws -> SimRow? {
        try RaidFixture.load(SimList.self, "sims.list").rows.first { $0.simId == id }
    }

    @Test func plainRun_drawsTheMeanAndTheErrorLine() throws {
        let model = SimResultModel(id: "km3xibdrsiyr", result: try result(), row: try listRow("km3xibdrsiyr"), specs: specs)
        #expect(model.title == "Tuesday")
        #expect(model.kind == "Sim")
        #expect(model.meta == "Fury Warrior · level 60")
        #expect(model.answer == .figure("101.4"))
        #expect(model.errorLine == "± 0.1 · 3,000 iterations · engine 1f542def3")
        #expect(model.breakdownURL.absoluteString == "https://foreversixty.gg/sim/km3xibdrsiyr")
    }

    @Test func untitled_isNamedBySpec() throws {
        let model = SimResultModel(id: "x", result: try SimsTestData.result(untitled: true), row: nil, specs: specs)
        #expect(model.title == "Fury")
    }

    @Test func otherKinds_showTheRowHeadline() throws {
        let row = SimRow(simId: "g", kind: "gear", headline: "+12 DPS from Robe of the Archmage")
        let model = SimResultModel(id: "g", result: try result(), row: row, specs: specs)
        #expect(model.kind == "Top Gear")
        #expect(model.answer == .sentence("+12 DPS from Robe of the Archmage"))
    }

    @Test func gear_isInFrameOrder() throws {
        let model = SimResultModel(id: "x", result: try result(), row: nil, specs: specs)
        #expect(model.gear.map(\.slot) == ["head", "main_hand"])
        #expect(model.gear.allSatisfy { !$0.changed })
    }

    @Test func load_marksSeenOnce() async throws {
        let service = FakeSimsService()
        var seen: [String] = []
        let store = SimResultStore(
            id: "km3xibdrsiyr", row: nil, service: service, cache: MemoryRaidCache(), specs: specs,
            now: { SimsLogicTests.now }, markSeen: { seen.append($0) }
        )
        await store.load()
        await store.load()
        #expect(seen == ["km3xibdrsiyr"])
        guard case let .loaded(model, offline) = store.phase else { Issue.record("not loaded"); return }
        #expect(model.title == "Tuesday")
        #expect(offline == nil)
    }

    @Test func offline_drawsTheCachedResult_elseFails() async throws {
        let cache = MemoryRaidCache()
        let online = FakeSimsService()
        await SimResultStore(id: "a", row: nil, service: online, cache: cache, specs: specs, now: { SimsLogicTests.now }, markSeen: { _ in }).load()
        let offline = FakeSimsService()
        offline.resultAnswer = { _ throws(APIError) in throw .transport("offline") }
        offline.listAnswer = { _, _ throws(APIError) in throw .transport("offline") }
        let cached = SimResultStore(id: "a", row: nil, service: offline, cache: cache, specs: specs, now: { SimsLogicTests.now }, markSeen: { _ in })
        await cached.load()
        guard case let .loaded(_, since) = cached.phase else { Issue.record("not loaded"); return }
        #expect(since == SimsLogicTests.now)
        let fresh = SimResultStore(id: "b", row: nil, service: offline, cache: cache, specs: specs, now: { SimsLogicTests.now }, markSeen: { _ in })
        await fresh.load()
        #expect(fresh.phase == .failed)
    }
}
