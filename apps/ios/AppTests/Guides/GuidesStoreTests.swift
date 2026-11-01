@testable import ForeverSixty
import Foundation
import Models
import Testing

@MainActor
@Suite("Guides stores")
struct GuidesStoreTests {
    @Test func home_loadsTheIndexAndWhatIsSaved() async {
        let source = FixtureGuidesSource()
        source.cachedSpecs = ["shaman-elemental"]
        let store = GuidesHomeStore(source: source)
        await store.load()
        let home = store.phase.value
        #expect(home?.index.classes.count == 9)
        #expect(home?.saved.isSaved("shaman-elemental") == true)
        #expect(home?.saved.isSaved("mage-fire") == false)
    }

    @Test func home_failsInItsRegion() async {
        let source = FixtureGuidesSource()
        source.failing = true
        let store = GuidesHomeStore(source: source)
        await store.load()
        // The saved index draws first; a failed refresh keeps it.
        #expect(store.phase.value != nil)
    }

    @Test func guide_offlineAndNeverSaved_saysSo() async {
        let source = FixtureGuidesSource()
        source.offline = true
        let store = GuideStore(spec: "shaman-enhancement", source: source)
        await store.load()
        #expect(store.phase == .notCached)
    }

    @Test func guide_offlineAndSaved_drawsTheSavedCopy() async {
        let source = FixtureGuidesSource()
        source.offline = true
        source.cachedSpecs = ["shaman-elemental"]
        let store = GuideStore(spec: "shaman-elemental", source: source)
        await store.load()
        guard case let .loaded(content, since) = store.phase else {
            Issue.record("the saved guide did not draw")
            return
        }
        #expect(content.guide.specName == "Elemental")
        #expect(since == GuidesTestData.now)
    }

    @Test func guide_serverError_isTheError() async {
        let source = FixtureGuidesSource()
        source.failing = true
        let store = GuideStore(spec: "shaman-elemental", source: source)
        await store.load()
        #expect(store.phase == .failed)
    }

    @Test func newBuild_waitsForUpdate() async {
        let source = FixtureGuidesSource()
        let store = GuideStore(spec: "shaman-elemental", source: source)
        await store.load()
        let shown = store.phase
        source.latest = "1.60.1.70800"
        await store.checkForNewData()
        #expect(store.newData == .ready("1.60.1.70800"))
        #expect(store.phase == shown)
        await store.update()
        #expect(store.newData == .none)
        guard case let .loaded(content, _) = store.phase else {
            Issue.record("the update did not draw")
            return
        }
        #expect(content.build == "1.60.1.70800")
    }

    @Test func update_thatFails_saysSoAndKeepsTheGuide() async {
        let source = FixtureGuidesSource()
        let store = GuideStore(spec: "shaman-elemental", source: source)
        await store.load()
        source.latest = "1.60.1.70800"
        await store.checkForNewData()
        source.failing = true
        await store.update()
        #expect(store.newData == .failed("1.60.1.70800"))
        #expect(store.phase.value?.build == GuidesTestData.build)
    }

    @Test func tiers_failInTheirRegion() async {
        let source = FixtureGuidesSource()
        source.failing = true
        let store = TiersStore(source: source, choice: .fallback)
        await store.load()
        #expect(store.phase == .failed)
    }

    @Test func changes_loadDatesAndTheChangelog() async {
        let store = ChangesStore(source: FixtureGuidesSource(), now: { GuidesTestData.now })
        await store.load()
        #expect(store.dates?.map(\.label) == ["Name reservation", "First raids"])
        #expect(store.changelog.value?.entries.isEmpty == false)
    }
}
