package gg.foreversixty.app.guild

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.guild.logic.GuildRank
import gg.foreversixty.app.guild.logic.HistoryCopy
import gg.foreversixty.app.guild.logic.HistoryPager
import gg.foreversixty.app.guild.logic.HistoryTarget
import gg.foreversixty.app.guild.logic.Loadable
import gg.foreversixty.app.guild.logic.MoreState
import gg.foreversixty.app.guild.logic.NightFacts
import gg.foreversixty.app.guild.logic.RosterCopy
import gg.foreversixty.app.guild.logic.RosterEntries
import gg.foreversixty.app.guild.logic.RosterList
import gg.foreversixty.app.guild.logic.RosterQuery
import gg.foreversixty.app.guild.logic.RosterRole
import gg.foreversixty.app.guild.logic.RosterSort
import gg.foreversixty.app.guild.logic.RosterStanding
import gg.foreversixty.app.guild.logic.RosterStore
import gg.foreversixty.app.model.RaidNightRow
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.raids.MemoryRaidCache
import gg.foreversixty.app.raids.RaidTime
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.fail
import gg.foreversixty.app.raids.resources
import java.io.IOException
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

@RunWith(AndroidJUnit4::class)
class GuildLogicTest {
    private val entries = RosterEntries.of(GuildFixture.home())
    private fun names(query: RosterQuery) = RosterList.apply(entries, query).map { it.name }
    private fun entry(name: String) = entries.first { it.name == name }

    // Roster (F16.1)

    @Test
    fun `every member is a row, the waiting and the no-account ones included`() {
        assertEquals(GuildFixture.home().members.size, entries.size)
        assertEquals(RosterStanding.Waiting, entry("Voidsong").standing)
        assertEquals(RosterStanding.NoAccount, entry("Gruzz").standing)
        assertFalse(entry("Gruzz").opensPage)
    }

    @Test
    fun `attendance and the supporter mark join from the roster by character key`() {
        assertEquals(RosterStanding.Attended(8, 8), entry("Kraggor").standing)
        assertTrue(entry("Kraggor").supporter)
        assertEquals(RosterStanding.Attended(3, 8), entry("Fernhoof").standing)
        assertFalse(entry("Mendara").supporter)
    }

    @Test
    fun `a roster row with no logged nights says so`() {
        val home = GuildFixture.home()
        val row = home.roster.first { it.name == "Kraggor" }
        val noNights = home.copy(roster = listOf(row.copy(attendance = row.attendance.copy(present = 0, nights = 0))))
        assertEquals(RosterStanding.NoNights, RosterEntries.of(noNights).first { it.name == "Kraggor" }.standing)
    }

    @Test
    fun `the bar never draws more than 8 segments`() {
        val home = GuildFixture.home()
        val row = home.roster.first { it.name == "Kraggor" }
        val wide = home.copy(roster = listOf(row.copy(attendance = row.attendance.copy(present = 9, nights = 12))))
        assertEquals(RosterStanding.Attended(8, 8), RosterEntries.of(wide).first { it.name == "Kraggor" }.standing)
    }

    @Test
    fun `role chips read members role, never a spec word`() {
        assertEquals(listOf("Kraggor", "Thokgar", "Brokkar", "Ursoka"), names(RosterQuery(RosterRole.Tanks)))
        assertTrue(names(RosterQuery(RosterRole.Healers)).all { entry(it).role == "healer" })
        assertEquals(entries.count { it.role == "dps" }, names(RosterQuery(RosterRole.Dps)).size)
        assertEquals(entries.size, names(RosterQuery(RosterRole.All)).size)
    }

    @Test
    fun `search filters by name, any case`() {
        assertEquals(listOf("Zekra"), names(RosterQuery(search = "zek")))
        assertEquals(emptyList<String>(), names(RosterQuery(RosterRole.Tanks, search = "mendara")))
    }

    @Test
    fun `rank sort is leader, officers, members, each by name`() {
        val sorted = names(RosterQuery(sort = RosterSort.Rank))
        assertEquals(listOf("Kraggor", "Draknar", "Mendara", "Thokgar"), sorted.take(4))
        val members = sorted.drop(4)
        assertTrue(members.all { entry(it).rank == GuildRank.Member })
        assertEquals(members.sortedWith(String.CASE_INSENSITIVE_ORDER), members)
    }

    @Test
    fun `rows with the same sort key keep name order`() {
        val tanks = names(RosterQuery(RosterRole.Tanks, sort = RosterSort.Level))
        val sixty = tanks.filter { entry(it).level == 60 }
        assertEquals(sixty.sortedWith(String.CASE_INSENSITIVE_ORDER), sixty)
        val classes = RosterList.apply(entries, RosterQuery(sort = RosterSort.ClassName)).filter { it.className == "warrior" }.map { it.name }
        assertEquals(classes.sortedWith(String.CASE_INSENSITIVE_ORDER), classes)
    }

    @Test
    fun `class sort follows the game's class order`() {
        val sorted = RosterList.apply(entries, RosterQuery(sort = RosterSort.ClassName))
        assertEquals("warrior", sorted.first().className)
        assertEquals("druid", sorted.last().className)
    }

    @Test
    fun `name, attendance and level sorts`() {
        assertEquals("Arrowmane", names(RosterQuery(sort = RosterSort.Name)).first())
        val byAttendance = RosterList.apply(entries, RosterQuery(sort = RosterSort.Attendance))
        assertEquals(RosterStanding.Attended(8, 8), byAttendance.first().standing)
        assertEquals(RosterStanding.Attended(1, 8), byAttendance.first { it.name == "Quillara" }.standing)
        assertTrue(byAttendance.indexOfFirst { it.name == "Quillara" } < byAttendance.indexOfFirst { it.name == "Voidsong" })
        val byLevel = RosterList.apply(entries, RosterQuery(sort = RosterSort.Level))
        assertEquals("Brokkar", byLevel.last().name)
    }

    @Test
    fun `rank words and the meta line`() {
        val copy = RosterCopy(resources())
        // A line breaks only after a separator (non-breaking space before the dot), never in the rank.
        assertEquals("Protection Warrior\u00A0· 60\u00A0· Guild\u00A0Master", copy.meta(entry("Kraggor")))
        assertEquals("Holy Priest\u00A0· 60\u00A0· Officer", copy.meta(entry("Mendara")))
        assertEquals("Elemental Shaman\u00A0· 60\u00A0· Member", copy.meta(entry("Lightningbolltz")))
        assertEquals("Warrior\u00A0· Member", copy.meta(entry("Gruzz").copy(specName = null, level = null)))
        assertFalse(copy.meta(entry("Kraggor")).contains(" ·"))
        assertEquals("8 of 8 nights", copy.attendance(RosterStanding.Attended(8, 8)))
    }

    @Test
    fun `the row reads as one sentence`() {
        val copy = RosterCopy(resources())
        assertEquals(
            "Kraggor, Protection Warrior, level 60, Guild Master, attended 8 of 8 recent raid nights",
            copy.a11y(entry("Kraggor")),
        )
        assertEquals("Voidsong, Shadow Priest, level 57, Member, waiting for approval", copy.a11y(entry("Voidsong")))
        assertEquals("Gruzz, Fury Warrior, level 54, Member, No account", copy.a11y(entry("Gruzz")))
    }

    @Test
    fun `a row opens the character's web page`() {
        assertEquals("https://foreversixty.gg/character/us/pvp/kraggor", entry("Kraggor").webUrl)
        assertNull(entry("Gruzz").webUrl)
    }

    @Test
    fun `the roster draws the cache first, then the answer, and offline keeps the cache`() = runTest {
        val service = FakeGuildService()
        val cache = MemoryRaidCache()
        val store = RosterStore(2, service, cache) { RaidTime.now }
        store.load()
        assertTrue(store.state is Loadable.Loaded)
        assertEquals(38, store.onRoster)

        service.homeAnswer = { fail(ApiError.Network(IOException("offline"))) }
        val again = RosterStore(2, service, cache) { RaidTime.now.plusSeconds(60) }
        again.load()
        assertEquals(RaidTime.now, (again.state as Loadable.Loaded).offlineSince)
    }

    @Test
    fun `the roster fails with nothing cached`() = runTest {
        val service = FakeGuildService().apply { homeAnswer = { fail(ApiError.Network(IOException("offline"))) } }
        val store = RosterStore(2, service, MemoryRaidCache()) { RaidTime.now }
        store.load()
        assertEquals(Loadable.Failed, store.state)
        assertNull(store.onRoster)
    }

    // Raid history (F16.3)

    private val page1 = GuildFixture.history(1).rows
    private fun night(id: String): RaidNightRow = (page1 + GuildFixture.history(2).rows).first { it.id == id }

    private val ony = "ewlrx3enxbeb"
    private val brd = "k6opku5uoucu"
    private val ubrs = "5owictw7n3jw"

    @Test
    fun `a night starts at its log's creation minus its length`() {
        val facts = NightFacts(night(ony), GuildFixture.mine)
        assertEquals(Triple("MON", "14", "DEC"), clock().dateBlock(facts.startedAt))
        assertEquals(1, facts.hours)
        assertEquals(8, facts.minutes)
    }

    @Test
    fun `bosses are distinct fight names`() {
        assertEquals(1, NightFacts(night(ony), emptySet()).bosses)
        assertEquals(6, NightFacts(night(brd), emptySet()).bosses)
    }

    @Test
    fun `ten or fewer are players, more are raiders, and the size is never below a kill's`() {
        val copy = HistoryCopy(resources(), clock())
        assertEquals("5 players", copy.size(NightFacts(night(brd), emptySet())))
        assertEquals("10 players", copy.size(NightFacts(night(ubrs), emptySet())))
        assertEquals("38 raiders", copy.size(NightFacts(night(ony), emptySet())))
        val row = night(ony)
        val bigKill = row.copy(raiders = 30, fights = row.fights.map { it.copy(players = 39) })
        assertEquals("39 raiders", copy.size(NightFacts(bigKill, emptySet())))
    }

    @Test
    fun `the meta lines and length words`() {
        val copy = HistoryCopy(resources(), clock())
        assertEquals("1 boss · 1 kill · 2 wipes", copy.fights(NightFacts(night(ony), emptySet())))
        assertEquals("6 bosses · 6 kills · no wipes", copy.fights(NightFacts(night(brd), emptySet())))
        assertEquals("5 bosses · 5 kills · 1 wipe", copy.fights(NightFacts(night(ubrs), emptySet())))
        assertEquals("1 h 8 m", copy.length(NightFacts(night(ony), emptySet())))
        val short = night(ony).copy(durationMs = 52 * 60_000L)
        assertEquals("52 m", copy.length(NightFacts(short, emptySet())))
    }

    @Test
    fun `you were there when one of your characters is present`() {
        assertTrue(NightFacts(night(ony), GuildFixture.mine).wasThere)
        assertFalse(NightFacts(night(brd), GuildFixture.mine).wasThere)
        assertFalse(NightFacts(night(ony), emptySet()).wasThere)
    }

    @Test
    fun `a tap opens your raid night when you were there and it was planned, else the report`() {
        assertEquals(HistoryTarget.Night(29), NightFacts(night(ony), GuildFixture.mine).target)
        assertEquals(HistoryTarget.Report("https://foreversixty.gg/reports/$ony"), NightFacts(night(ony), emptySet()).target)
        assertEquals(HistoryTarget.Report("https://foreversixty.gg/reports/$brd"), NightFacts(night(brd), GuildFixture.mine).target)
        // There, but no planned raid: the report.
        assertEquals(HistoryTarget.Report("https://foreversixty.gg/reports/$ubrs"), NightFacts(night(ubrs), GuildFixture.mine).target)
    }

    @Test
    fun `the history row reads as one sentence`() {
        val copy = HistoryCopy(resources(), clock())
        assertEquals(
            "Onyxia's Lair, Monday 14 December, 1 boss, 1 kill, 2 wipes, 38 raiders, 1 hour 8 minutes. You were there.",
            copy.a11y(NightFacts(night(ony), GuildFixture.mine)),
        )
        assertEquals(
            "Blackrock Depths, Friday 18 December, 6 bosses, 6 kills, no wipes, 5 players, 2 hours 8 minutes",
            copy.a11y(NightFacts(night(brd), GuildFixture.mine)),
        )
    }

    @Test
    fun `history pages by cursor and ends when the cursor does`() = runTest {
        val service = FakeGuildService()
        val pager = HistoryPager(2, service, MemoryRaidCache()) { RaidTime.now }
        pager.load()
        assertEquals(20, (pager.state as Loadable.Loaded).value.size)
        assertEquals(MoreState.Idle, pager.more)
        pager.loadMore()
        assertEquals(23, (pager.state as Loadable.Loaded).value.size)
        assertEquals(MoreState.End, pager.more)
        assertEquals(listOf("history:-", "history:MjAyNi0xMi0xMFQyMzo0NTowMFp8bHQ0d3R4bG02ZGxq"), service.calls)
    }

    @Test
    fun `a failed next page keeps the rows and offers Try again`() = runTest {
        val service = FakeGuildService()
        val pager = HistoryPager(2, service, MemoryRaidCache()) { RaidTime.now }
        pager.load()
        service.historyAnswer = { fail(ApiError.Network(IOException("offline"))) }
        pager.loadMore()
        assertEquals(MoreState.Failed, pager.more)
        assertEquals(20, (pager.state as Loadable.Loaded).value.size)
        service.historyAnswer = { ok2() }
        pager.loadMore()
        assertEquals(MoreState.End, pager.more)
    }

    private fun ok2() = gg.foreversixty.app.raids.ok(GuildFixture.history(2))

    @Test
    fun `history fails with nothing cached`() = runTest {
        val service = FakeGuildService().apply { historyAnswer = { fail(ApiError.Network(IOException("offline"))) } }
        val pager = HistoryPager(2, service, MemoryRaidCache()) { RaidTime.now }
        pager.load()
        assertEquals(Loadable.Failed, pager.state)
    }
}
