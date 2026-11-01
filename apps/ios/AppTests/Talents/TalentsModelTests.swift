@testable import ForeverSixty
import Foundation
import Models
import Networking
import Synchronization
import Testing

/// A source answering from the fixtures, with switches for each file's outcome.
private final class FakeTalentsSource: TalentsSource, Sendable {
    struct Behaviour {
        var talents: Result<CachedFile<TalentsFile>, APIError>
        var olderTalents: Result<CachedFile<TalentsFile>, APIError>
        var bis: Result<CachedFile<BisFile>, APIError>
    }

    private let state: Mutex<Behaviour>

    init(_ behaviour: Behaviour) { state = Mutex(behaviour) }

    func activeBuild() async -> String { BisTestData.build }

    func talents(build: String, classSlug _: String) async throws(APIError) -> CachedFile<TalentsFile> {
        try state.withLock { build == BisTestData.build ? $0.talents : $0.olderTalents }.get()
    }

    func bis(build _: String, spec _: String) async throws(APIError) -> CachedFile<BisFile> {
        try state.withLock { $0.bis }.get()
    }
}

/// How the Talents screen loads (spec F6 §4, §6). The same list runs on Android (TalentsModelTest.kt).
@MainActor
@Suite("Talents model")
struct TalentsModelTests {
    private let warrior: TalentsFile
    private let fury: BisFile
    private let catalog: SpecCatalog
    private let yell: HomeCharacter

    init() throws {
        warrior = try TalentTestData.warrior()
        fury = try TalentTestData.fury()
        catalog = try BisTestData.catalog()
        let me = try BisTestData.me()
        yell = HomeCharacter(try #require(me.characters.first { $0.name == "Obnoxious Yell" }))
    }

    private func fresh<T>(_ value: T) -> Result<CachedFile<T>, APIError> {
        .success(CachedFile(value: value, savedAt: BisTestData.now, isFromCache: false))
    }

    private func model(
        _ character: HomeCharacter? = nil,
        talents: Result<CachedFile<TalentsFile>, APIError>? = nil,
        older: Result<CachedFile<TalentsFile>, APIError> = .failure(.unexpectedStatus(404)),
        bis: Result<CachedFile<BisFile>, APIError>? = nil
    ) -> TalentsModel {
        let source = FakeTalentsSource(.init(talents: talents ?? fresh(warrior), olderTalents: older, bis: bis ?? fresh(fury)))
        return TalentsModel(source: source, character: TalentCharacter(character ?? yell), catalog: catalog)
    }

    private func presenter(_ model: TalentsModel) throws -> (TalentsPresenter, Date?) {
        guard case let .loaded(presenter, since) = model.phase else {
            Issue.record("not loaded: \(model.phase)")
            throw CancellationError()
        }
        return (presenter, since)
    }

    private func with(_ character: HomeCharacter, level: Int? = nil, spec: String?? = nil, dataBuild: String? = nil,
                      talents: BuildTalents?? = nil) -> HomeCharacter {
        HomeCharacter(
            key: character.key, region: character.region, ruleset: character.ruleset, name: character.name,
            race: character.race, className: character.className, specName: spec ?? character.specName,
            level: level ?? character.level, guild: character.guild, faction: character.faction, gear: character.gear,
            talents: talents ?? character.talents, dataBuild: dataBuild ?? character.dataBuild
        )
    }

    @Test("The fixture character compares with the 50 to 59 guide build")
    func loadsTheBand() async throws {
        let model = model()
        await model.load()
        let (presenter, since) = try presenter(model)
        #expect(presenter.status.sentence == "Your build has every point of the Fury 50 to 59 guide build.")
        #expect(presenter.notes == ["+8 beyond the guide"])
        #expect(presenter.tabs(.yours).map(\.points) == [17, 32, 0])
        #expect(since == nil)
    }

    @Test("At 60 the raid entry is the guide")
    func levelSixtyReadsRaid() async throws {
        let model = model(with(yell, level: 60))
        await model.load()
        let (presenter, _) = try presenter(model)
        #expect(presenter.guide == .build(talents: TalentTestData.band60, band: 60))
    }

    @Test("Saved copies set the offline stamp")
    func offline() async throws {
        let saved = CachedFile(value: warrior, savedAt: BisTestData.now, isFromCache: true)
        let model = model(talents: .success(saved))
        await model.load()
        #expect(try presenter(model).1 == BisTestData.now)
    }

    @Test("A talent file that cannot load is the error state")
    func failed() async {
        let model = model(talents: .failure(.transport("offline")))
        await model.load()
        #expect(model.phase == .failed)
    }

    @Test("A missing BiS file is no build; a failed one is the guide-failed line")
    func guideOutcomes() async throws {
        let missing = model(bis: .failure(.unexpectedStatus(404)))
        await missing.load()
        #expect(try presenter(missing).0.guide == .none)
        let broken = model(bis: .failure(.transport("offline")))
        await broken.load()
        #expect(try presenter(broken).0.guide == .failed)
    }

    @Test("Under 20 and with no spec there is no guide to read")
    func noGuide() async throws {
        let young = model(with(yell, level: 15))
        await young.load()
        #expect(try presenter(young).0.guide == .none)
        let unknown = model(with(yell, spec: .some(nil)))
        await unknown.load()
        #expect(try presenter(unknown).0.guide == .noSpec)
    }

    @Test("An export of zeros is no export")
    func zeroExport() async throws {
        let model = model(with(yell, talents: .some(BuildTalents(trees: ["0", "0", "0"], points: [0, 0, 0]))))
        await model.load()
        #expect(try presenter(model).0.export == .none)
    }

    @Test("An export from another build draws only when its trees are laid out the same")
    func olderBuild() async throws {
        let same = model(with(yell, dataBuild: "1.60.1.70532"), older: fresh(warrior))
        await same.load()
        #expect(try presenter(same).0.yourRanks != nil)
        let changed = model(with(yell, dataBuild: "1.60.1.70532"), older: fresh(try TalentTestData.shaman()))
        await changed.load()
        #expect(try presenter(changed).0.export == .stale)
        let gone = model(with(yell, dataBuild: "1.60.1.70532"))
        await gone.load()
        #expect(try presenter(gone).0.export == .stale)
    }
}
