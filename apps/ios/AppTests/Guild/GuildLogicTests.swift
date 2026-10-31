@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// The F16 test list; Android's guild extras tests carry the same names.
@Suite("Roster")
struct RosterListTests {
    private func list() throws -> RosterList { RosterList(home: try GuildFixture.home()) }

    private func entry(_ name: String) throws -> RosterEntry {
        try #require(try list().entries.first { $0.name == name })
    }

    @Test("the counts are the API's: 38 on the roster, 4 no account, 2 waiting")
    func counts() throws {
        let counts = try list().counts
        #expect(counts.onRoster == 38 && counts.inGameNoAccount == 4 && counts.waiting == 2)
    }

    @Test("the role chips read members[].role, never a spec word")
    func roleChips() throws {
        let list = try list()
        #expect(list.visible(role: .all, query: "", sort: .rank).count == 38)
        #expect(list.visible(role: .tanks, query: "", sort: .rank).count == 4)
        #expect(list.visible(role: .healers, query: "", sort: .rank).count == 10)
        #expect(list.visible(role: .dps, query: "", sort: .rank).count == 24)
        #expect(list.visible(role: .tanks, query: "", sort: .rank).allSatisfy { $0.role == "tank" })
    }

    @Test("search filters by name, ignoring case and accents")
    func search() throws {
        let found = try list().visible(role: .all, query: "KRAG", sort: .rank)
        #expect(found.map(\.name) == ["Kraggor"])
        #expect(try list().visible(role: .healers, query: "kraggor", sort: .rank).isEmpty)
        #expect(try list().visible(role: .all, query: "  ", sort: .rank).count == 38)
    }

    @Test("rank sorts leader, officer, member, then by name")
    func sortRank() throws {
        let names = try list().visible(role: .all, query: "", sort: .rank).map(\.rank)
        #expect(names.first == "leader")
        #expect(Array(names.prefix(3)) == ["leader", "officer", "officer"])
        #expect(names.last == "member")
        let members = try list().visible(role: .all, query: "", sort: .rank).filter { $0.rank == "member" }.map(\.name)
        #expect(members == members.sorted { $0.localizedCaseInsensitiveCompare($1) == .orderedAscending })
    }

    @Test("class sorts in the game's order, name sorts A to Z, level sorts high to low")
    func sortOthers() throws {
        let byClass = try list().visible(role: .all, query: "", sort: .className)
        #expect(byClass.first?.className == "warrior")
        #expect(byClass.last?.className == "druid")
        let byName = try list().visible(role: .all, query: "", sort: .name).map(\.name)
        #expect(byName == byName.sorted { $0.localizedCaseInsensitiveCompare($1) == .orderedAscending })
        let byLevel = try list().visible(role: .all, query: "", sort: .level)
        #expect(byLevel.last?.name == "Brokkar")
        #expect(byLevel.first?.level == 60)
    }

    @Test("attendance sorts most nights present first; no attendance sorts last")
    func sortAttendance() throws {
        let sorted = try list().visible(role: .all, query: "", sort: .attendance)
        #expect(sorted.first?.standing == .attendance(present: 8, nights: 8))
        let tail = sorted.suffix(6)
        #expect(tail.allSatisfy { $0.standing == .waiting || $0.standing == .noAccount })
    }

    @Test("attendance and the supporter mark join by character key")
    func join() throws {
        let kraggor = try entry("Kraggor")
        #expect(kraggor.standing == .attendance(present: 8, nights: 8))
        #expect(kraggor.supporter)
        let bolt = try entry("Lightningbolltz")
        #expect(bolt.standing == .attendance(present: 6, nights: 8))
        #expect(bolt.attendanceCaption == "6 of 8")
        #expect(!(try entry("Mendara")).supporter)
    }

    @Test("a waiting character shows Waiting; a name with no account has no attendance and no page")
    func waitingAndNoAccount() throws {
        #expect(try entry("Voidsong").standing == .waiting)
        let gruzz = try entry("Gruzz")
        #expect(gruzz.standing == .noAccount)
        #expect(gruzz.attendanceCaption == nil)
        #expect(gruzz.pageURL == nil)
        #expect(try entry("Kraggor").pageURL?.absoluteString == "https://foreversixty.gg/character/us/pvp/kraggor")
    }

    @Test("a guild with no logged nights says so; the bar never has more than 8 segments")
    func nights() throws {
        #expect(RosterStanding(present: 0, nights: 0) == .noNights)
        #expect(RosterStanding(present: 3, nights: 5) == .attendance(present: 3, nights: 5))
        let bar = AttendanceBar(present: 12, nights: 12)
        #expect(bar.segments == 8 && bar.filled == 8)
        #expect(AttendanceBar(present: 2, nights: 5).segments == 5)
    }

    @Test("the meta reads the full spec name, level and rank word")
    func meta() throws {
        #expect(try entry("Kraggor").meta == "Protection Warrior · 60 · Guild Master")
        #expect(try entry("Mendara").meta == "Holy Priest · 60 · Officer")
        #expect(try entry("Lightningbolltz").meta == "Elemental Shaman · 60 · Member")
        #expect(RosterEntry.rankWord("leader") == "Guild Master")
        #expect(RosterEntry.rankWord("officer") == "Officer")
        #expect(RosterEntry.rankWord("member") == "Member")
    }

    @Test("a null spec name falls back to the class name")
    func metaFallback() throws {
        let member = GuildMember(name: "Ash", characterKey: "us/pvp/ash", state: "approved", rank: "member", className: "mage", level: 58)
        let entry = RosterEntry(member: member, row: nil)
        #expect(entry.meta == "Mage · 58 · Member")
    }

    @Test("VoiceOver reads the row in words; the bar is not read")
    func accessibility() throws {
        #expect(try entry("Kraggor").accessibilityLabel == "Kraggor, Protection Warrior, level 60, Guild Master, attended 8 of 8 recent raid nights")
        #expect(try entry("Voidsong").accessibilityLabel == "Voidsong, Shadow Priest, level 57, Member, waiting for approval")
        #expect(try entry("Gruzz").accessibilityLabel == "Gruzz, Fury Warrior, level 54, Member, No account")
    }
}

@Suite("Raid history")
struct HistoryNightTests {
    let clock = RaidTime.clock
    let yours: Set<String> = [GuildFixture.viewerKey]

    private func night(_ index: Int) throws -> HistoryNight {
        let rows = try GuildFixture.historyPage1().rows
        return HistoryNight(row: rows[index], yourKeys: yours)
    }

    @Test("bosses are the distinct fight names; wipes read no wipes or n wipes")
    func facts() throws {
        #expect(try night(0).factsLine == "1 boss · 1 kill · no wipes")
        #expect(try night(1).factsLine == "5 bosses · 5 kills · 1 wipe")
        #expect(try night(3).factsLine == "1 boss · 1 kill · 1 wipe")
    }

    @Test("a night of 10 or fewer reads players, a raid reads raiders; lengths in h and m")
    func size() throws {
        #expect(try night(0).sizeLine == "40 raiders · 52 m")
        #expect(try night(1).sizeLine == "10 players · 1 h 48 m")
        #expect(try night(2).sizeLine == "5 players · 2 h 11 m")
        #expect(try night(3).sizeLine == "39 raiders · 1 h 11 m")
    }

    @Test("the date block is the night's start, created_at minus its length, in the device zone")
    func start() throws {
        let night = try night(3)
        let block = clock.dateBlock(night.start)
        #expect(block.weekday == "SUN" && block.day == "27" && block.month == "DEC")
    }

    @Test("You were there when one of your characters is in present[]")
    func there() throws {
        #expect(try night(0).youWereThere)
        #expect(!(try night(2)).youWereThere)
        #expect(!HistoryNight(row: try GuildFixture.historyPage1().rows[0], yourKeys: []).youWereThere)
    }

    @Test("a tap opens your raid night when you were there and it has a raid, else the report")
    func target() throws {
        #expect(try night(0).target == .night(raid: 40))
        #expect(try night(1).target == .report(URL(string: "https://foreversixty.gg/reports/g2k4der64mlv")!))
        #expect(try night(5).target == .report(URL(string: "https://foreversixty.gg/reports/jz2vv7qhmvgk")!))
    }

    @Test("VoiceOver reads the night in words")
    func accessibility() throws {
        #expect(try night(3).accessibilityLabel(clock: clock)
            == "Onyxia's Lair, Sunday 27 December, 1 boss, 1 kill, 1 wipe, 39 raiders, 1 hour 11 minutes. You were there.")
        #expect(try night(2).accessibilityLabel(clock: clock)
            == "Blackrock Depths, Tuesday 29 December, 6 bosses, 6 kills, no wipes, 5 players, 2 hours 11 minutes")
    }
}

@MainActor
@Suite("Raid history paging")
struct HistoryStoreTests {
    private func store(_ service: FakeGuildExtrasService, cache: any RaidCache = MemoryRaidCache()) -> HistoryStore {
        HistoryStore(guild: 2, service: service, cache: cache, now: { RaidTime.now })
    }

    @Test("the first page loads 20; Load more appends the next and ends at no cursor")
    func paging() async throws {
        let service = FakeGuildExtrasService()
        let store = store(service)
        await store.load()
        #expect(store.content?.rows.count == 20)
        #expect(store.content?.hasMore == true)
        await store.loadMore()
        #expect(store.content?.rows.count == 23)
        #expect(store.content?.hasMore == false)
        #expect(service.cursors == [nil, "MjAyNi0xMi0xMFQyMDo1NTowMFp8bjNhcDVrdzd0aml5"])
        await store.loadMore()
        #expect(service.cursors.count == 2)
    }

    @Test("a failed next page keeps the rows and says so under them")
    func moreFails() async throws {
        let service = FakeGuildExtrasService()
        let store = store(service)
        await store.load()
        service.historyAnswer = { _ throws(APIError) in throw .transport("down") }
        await store.loadMore()
        #expect(store.moreFailed)
        #expect(store.content?.rows.count == 20)
    }

    @Test("an error with nothing cached is the error state; offline shows the cache under the strip")
    func failures() async throws {
        let failing = FakeGuildExtrasService()
        failing.historyAnswer = { _ throws(APIError) in throw .transport("down") }
        let empty = store(failing)
        await empty.load()
        #expect(empty.phase == .failed)

        let cache = MemoryRaidCache()
        let first = store(FakeGuildExtrasService(), cache: cache)
        await first.load()
        let offline = store(failing, cache: cache)
        await offline.load()
        #expect(offline.content?.offlineSince != nil)
        #expect(offline.content?.rows.count == 20)
    }
}

@MainActor
@Suite("Roster store")
struct RosterStoreTests {
    @Test("the roster loads from the guild home, and an error with nothing cached is the error state")
    func load() async throws {
        let store = RosterStore(guild: 2, service: FakeGuildExtrasService(), cache: MemoryRaidCache(), now: { RaidTime.now })
        await store.load()
        #expect(store.list?.entries.count == 38)
        #expect(store.onRoster == 38)
        #expect(store.viewerKey == GuildFixture.viewerKey)

        let failing = FakeGuildExtrasService()
        failing.homeAnswer = { () throws(APIError) in throw .transport("down") }
        let broken = RosterStore(guild: 2, service: failing, cache: MemoryRaidCache(), now: { RaidTime.now })
        await broken.load()
        #expect(broken.phase == .failed)
    }
}
