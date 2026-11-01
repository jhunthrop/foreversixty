@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

@MainActor
@Suite("Planner editor")
struct PlannerEditorTests {
    let file: TalentsFile
    let index: PlannerIndex

    init() throws {
        file = try PlannerTestData.shaman()
        index = PlannerIndex(file)
    }

    private func id(_ name: String) -> Int { PlannerTestData.talentID(file, name) }

    private func model(
        _ launch: PlannerLaunch = PlannerLaunch(classSlug: "shaman", raceSlug: "orc", level: 60),
        source: FakePlannerSource? = nil, drafts: MemoryPlannerDraftStore = MemoryPlannerDraftStore(),
        haptics: RecordingPlannerHaptics = RecordingPlannerHaptics()
    ) async throws -> PlannerEditorModel {
        let model = PlannerEditorModel(
            launch: launch, dependencies: .test(source: try source ?? FakePlannerSource.shaman(), drafts: drafts, haptics: haptics)
        )
        await model.load()
        return model
    }

    @Test("a tap adds with a selection tick, a long press removes with a medium impact")
    func addRemoveHaptics() async throws {
        let haptics = RecordingPlannerHaptics()
        let model = try await model(haptics: haptics)
        model.add(id("Convection"))
        model.remove(id("Convection"))
        #expect(model.state?.order == [])
        #expect(haptics.played == [.segmentTick, .longPress])
    }

    @Test("a refusal changes nothing, plays REJECT and says why in the strip")
    func refusal() async throws {
        let haptics = RecordingPlannerHaptics()
        let model = try await model(haptics: haptics)
        let before = model.state
        model.add(id("Elemental Mastery"))
        #expect(model.state?.order == before?.order)
        #expect(model.state?.canUndo == false)
        #expect(haptics.played == [.reject])
        guard case let .refusal(line)? = model.state?.line else { Issue.record("no refusal"); return }
        #expect(line.hasPrefix("Tier "))
        #expect(model.state?.focus?.name == "Elemental Mastery")
    }

    @Test("the level's budget comes first: level 10 holds one point")
    func levelBudget() async throws {
        let model = try await model(PlannerLaunch(classSlug: "shaman", raceSlug: "orc", level: 10))
        model.add(id("Convection"))
        model.add(id("Convection"))
        #expect(model.state?.order.count == 1)
        #expect(model.state?.line == .refusal(Copy.planRefuseNoneLeft))
        #expect(model.state?.left == 0)
    }

    @Test("the level never drops below the points spent plus 9")
    func levelFloor() async throws {
        let model = try await model()
        for _ in 0..<5 { model.add(id("Convection")) }
        model.setLevel(10)
        #expect(model.state?.level == 14)
        model.setLevel(70)
        #expect(model.state?.level == 60)
        #expect(PlannerRules.minimumLevel(spent: 0) == 10)
    }

    @Test("undo steps back through the last 100 changes, one point a step")
    func undoLimit() throws {
        var state = PlannerState(index: index, raceSlug: "orc")
        let first = id("Convection")
        for _ in 0..<5 { _ = state.add(first) }
        for _ in 0..<60 {
            _ = state.remove(first)
            _ = state.add(first)
        }
        #expect(state.history.count == PlannerState.undoLimit)
        for _ in 0..<150 { state.undo() }
        #expect(!state.canUndo)
        // 125 changes, 100 kept: undoing them all lands before change 26, a removal from 5 points.
        #expect(state.order.count == 5)
    }

    @Test("undo restores exactly the order before each change")
    func undoExact() throws {
        var state = PlannerState(index: index)
        _ = state.add(id("Convection"))
        _ = state.add(id("Concussion"))
        state.undo()
        #expect(state.order == [id("Convection")])
        state.reset()
        #expect(state.order == [])
        state.undo()
        #expect(state.order == [id("Convection")])
    }

    @Test("Reset clears every point in one undoable step and saves the draft")
    func reset() async throws {
        let drafts = MemoryPlannerDraftStore()
        let model = try await model(drafts: drafts)
        model.add(id("Convection"))
        model.add(id("Concussion"))
        model.reset()
        #expect(model.state?.spent == 0)
        #expect(drafts.drafts["shaman"]?.order == [])
        model.undo()
        #expect(model.state?.spent == 2)
    }

    @Test("the draft is saved as you go and restored when Plan a build reopens")
    func draftRestore() async throws {
        let drafts = MemoryPlannerDraftStore()
        let source = try FakePlannerSource.shaman()
        let first = try await model(source: source, drafts: drafts)
        first.add(id("Convection"))
        first.setRace("troll")
        first.setLevel(30)
        let reopened = try await model(PlannerLaunch(classSlug: "shaman", raceSlug: "orc", level: 60), source: source, drafts: drafts)
        #expect(reopened.state?.order == [id("Convection")])
        #expect(reopened.state?.raceSlug == "troll")
        #expect(reopened.state?.level == 30)
    }

    @Test("a draft survives a force quit: UserDefaults round-trips it")
    func defaultsDraft() throws {
        let suite = "planner-tests-\(UUID().uuidString)"
        let defaults = try #require(UserDefaults(suiteName: suite))
        defer { defaults.removePersistentDomain(forName: suite) }
        let draft = PlannerDraft(treeVersion: "1", raceSlug: "orc", level: 42, order: [1, 2, 2])
        UserDefaultsPlannerDraftStore(defaults: defaults).save(draft, classSlug: "shaman")
        #expect(UserDefaultsPlannerDraftStore(defaults: UserDefaults(suiteName: suite) ?? defaults).load(classSlug: "shaman") == draft)
        #expect(UserDefaultsPlannerDraftStore(defaults: defaults).load(classSlug: "mage") == nil)
    }

    @Test("an FS1 code loads one legal order and the band note")
    func fs1Load() async throws {
        let code = "FS1:\(PlannerTestData.build):shaman:troll:5510331300001350/03/0533520000000000:"
        let model = try await model(PlannerLaunch(classSlug: "shaman", source: .code(code), origin: .guide))
        #expect(model.state?.spent == 51)
        #expect(model.state?.order == PlannerTestData.elementalOrder(index))
        #expect(model.state?.raceSlug == "troll")
        #expect(model.state?.line == .note(Copy.planLoadedNote(band: TalentGuide.label(band: 60))))
        #expect(model.state?.left == 0)
    }

    @Test("a copy of a shared build loads its points; an illegal order is rebuilt from its ranks")
    func copyLoad() throws {
        let legal = PlannerTestData.elementalOrder(index)
        let copy = PlannerState.opening(
            PlannerLaunch(classSlug: "shaman", source: .copy(order: legal)), index: index, draft: nil, legalRaces: ["orc"]
        )
        #expect(copy.order == legal)
        let shuffled = Array(legal.reversed())
        let rebuilt = PlannerState.opening(
            PlannerLaunch(classSlug: "shaman", source: .copy(order: shuffled)), index: index, draft: nil, legalRaces: ["orc"]
        )
        #expect(PlannerRules.validate(index, order: rebuilt.order).isEmpty)
        #expect(PlannerRules.ranks(rebuilt.order) == PlannerRules.ranks(legal))
    }

    @Test("the race sheet lists the class's legal, non-placeholder races")
    func races() async throws {
        let model = try await model()
        #expect(model.races.map(\.slug) == ["orc", "dwarf", "tauren", "troll", "windshaper-skyborne"])
        #expect(model.raceName == "Orc")
    }

    @Test("the share body: class, race, the active trees, the order, no gear, the trimmed title")
    func shareBody() async throws {
        let source = try FakePlannerSource.shaman()
        let haptics = RecordingPlannerHaptics()
        let model = try await model(source: source, haptics: haptics)
        model.add(id("Convection"))
        await model.share(title: "  Elemental raid  ")
        #expect(source.saved == [BuildInput(
            classId: 7, raceId: 2, treeVersion: file.build, pointOrder: [id("Convection")], gear: [:], title: "Elemental raid"
        )])
        #expect(model.share == .shared(try #require(URL(string: "https://foreversixty.gg/b/k7q2m9xa"))))
        #expect(haptics.played.last == .confirm)
        #expect(model.shareInput(title: "")?.title == nil)
        #expect(model.shareInput(title: String(repeating: "x", count: 80))?.title?.count == 60)
    }

    @Test("share errors: 429, offline and a 400 field sentence under the editor")
    func shareErrors() async throws {
        let source = try FakePlannerSource.shaman()
        let model = try await model(source: source)
        source.saveResult = .failure(.server(status: 429, body: ErrorBody(code: "rate_limited", message: "slow"), requestID: nil))
        await model.share(title: "")
        #expect(model.share == .failed(Copy.planRateLimited))
        source.saveResult = .failure(.transport("offline"))
        await model.share(title: "")
        #expect(model.share == .failed(Copy.planOffline))
        let fields = ["point_order[3]": "Later", "point_order[1]": "Tier 1 of Elemental needs 5 points in Elemental first"]
        source.saveResult = .failure(.server(status: 400, body: ErrorBody(code: "invalid", message: "bad", fields: fields), requestID: nil))
        await model.share(title: "")
        #expect(model.share == .idle)
        #expect(model.state?.line == .refusal("Tier 1 of Elemental needs 5 points in Elemental first"))
    }

    @Test("the keyboard walks the talents and adds and removes on the strip's talent")
    func keyboard() async throws {
        let model = try await model()
        #expect(model.state?.focus?.id == index.allTalents.first?.id)
        model.moveFocus(by: 1)
        #expect(model.state?.focus?.id == index.allTalents[1].id)
        model.moveFocus(by: -2)
        #expect(model.state?.focus?.id == index.allTalents.last?.id)
        model.moveFocus(by: 1)
        model.addFocused()
        #expect(model.state?.spent == 1)
        model.removeFocused()
        #expect(model.state?.spent == 0)
    }

    @Test("a missing talent file fails the load")
    func failed() async throws {
        let source = FakePlannerSource()
        let model = try await model(PlannerLaunch(classSlug: "mage"), source: source)
        #expect(model.phase == .failed)
    }

    @Test("cell states: gold where a point can go, open with none left, locked behind a tier")
    func cellStates() throws {
        let order = PlannerTestData.elementalOrder(index)
        let ranks = PlannerRules.ranks(order)
        let convection = try #require(index.byID[id("Convection")])
        let mastery = try #require(index.byID[id("Elemental Mastery")])
        #expect(PlannerRules.cellState(convection, index: index, order: [], ranks: [:], left: 51) == .available)
        #expect(PlannerRules.cellState(convection, index: index, order: [], ranks: [:], left: 0) == .open)
        #expect(PlannerRules.cellState(mastery, index: index, order: [], ranks: [:], left: 51) == .locked)
        #expect(PlannerRules.cellState(convection, index: index, order: order, ranks: ranks, left: 0) == .maxed)
    }
}

@MainActor
@Suite("Shared builds and Your builds")
struct SharedBuildTests {
    @Test("the viewer draws the build's own trees, its head and meta")
    func viewer() async throws {
        let source = try FakePlannerSource.shaman()
        let index = PlannerIndex(try PlannerTestData.shaman())
        source.builds["k7q2m9xa"] = PlannerTestData.build(order: PlannerTestData.elementalOrder(index))
        let model = SharedBuildModel(id: "k7q2m9xa", source: source, cache: MemoryRaidCache())
        await model.load()
        guard case let .loaded(content) = model.phase else { Issue.record("not loaded"); return }
        #expect(content.summary.head == "Orc Shaman 30 / 3 / 18")
        #expect(content.summary.viewerMeta(timeZone: .gmt) == "Elemental raid · Saved 4 Nov · 12 views")
        #expect(content.file?.build == PlannerTestData.build)
        #expect(content.summary.raceSlug == "orc")
    }

    @Test("an unknown id, offline with no copy, offline with a saved copy")
    func states() async throws {
        let source = try FakePlannerSource.shaman()
        let cache = MemoryRaidCache()
        let unknown = SharedBuildModel(id: "nope", source: source, cache: cache)
        await unknown.load()
        #expect(unknown.phase == .notFound)
        source.buildError = .transport("offline")
        let offline = SharedBuildModel(id: "k7q2m9xa", source: source, cache: cache)
        await offline.load()
        #expect(offline.phase == .offline)
        cache.write(PlannerTestData.build(order: []), key: "build-k7q2m9xa", at: Date())
        await offline.load()
        guard case .loaded = offline.phase else { Issue.record("the saved copy did not draw"); return }
    }

    @Test("swipe removes the row and calls DELETE; Undo re-posts the same build")
    func removeUndo() async throws {
        let source = try FakePlannerSource.shaman()
        let first = PlannerTestData.build(id: "a", order: [], title: "First")
        let second = PlannerTestData.build(id: "b", order: [], title: nil, views: 1)
        source.mine = [first, second]
        let model = YourBuildsModel(source: source)
        await model.load()
        guard case let .loaded(rows) = model.phase else { Issue.record("not loaded"); return }
        #expect(rows.map(\.rowTitle) == ["First", "Orc Shaman 0 / 0 / 0"])
        #expect(rows[1].rowMeta(timeZone: .gmt) == "Orc Shaman · 0 / 0 / 0 · 4 Nov")
        await model.remove(rows[0])
        #expect(source.removed == ["a"])
        #expect(model.removed?.summary.id == "a")
        guard case let .loaded(left) = model.phase else { Issue.record("not loaded"); return }
        #expect(left.map(\.id) == ["b"])
        await model.undo()
        #expect(source.saved == [BuildInput(classId: 7, raceId: 2, treeVersion: PlannerTestData.build, pointOrder: [], gear: [:], title: "First")])
        guard case let .loaded(back) = model.phase else { Issue.record("not loaded"); return }
        #expect(back.map(\.id) == ["a", "b"])
        #expect(model.removed == nil)
    }

    @Test("a failed DELETE puts the row back and says so")
    func removeFails() async throws {
        let source = try FakePlannerSource.shaman()
        source.mine = [PlannerTestData.build(id: "a", order: [])]
        source.removeError = .transport("offline")
        let model = YourBuildsModel(source: source)
        await model.load()
        guard case let .loaded(rows) = model.phase else { Issue.record("not loaded"); return }
        await model.remove(rows[0])
        guard case let .loaded(after) = model.phase else { Issue.record("not loaded"); return }
        #expect(after.map(\.id) == ["a"])
        #expect(model.error == Copy.planRemoveFailed)
    }

    @Test("links route to the editor, Your builds and the viewer only while planner is on")
    func router() throws {
        let flag = PlannerSwitch()
        let router = PlannerRouter(dependencies: .test(source: FakePlannerSource())) { flag.on }
        #expect(router.open(.planner(className: "warrior", code: "FS1:x:warrior:orc:0/0/0:", race: "orc")))
        #expect(router.editor?.classSlug == "warrior")
        #expect(router.editor?.source == .code("FS1:x:warrior:orc:0/0/0:"))
        #expect(router.open(.sharedBuild(id: "k7q2m9xa")))
        #expect(router.page == .shared("k7q2m9xa"))
        #expect(router.editor == nil)
        #expect(router.open(.builds))
        #expect(router.page == .builds)
        flag.on = false
        router.page = nil
        #expect(!router.open(.sharedBuild(id: "x")))
        #expect(router.page == nil)
    }
}

@MainActor
private final class PlannerSwitch {
    var on = true
}
