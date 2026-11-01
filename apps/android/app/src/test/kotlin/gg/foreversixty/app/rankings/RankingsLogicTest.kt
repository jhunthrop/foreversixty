package gg.foreversixty.app.rankings

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.raids.MemoryRaidCache
import gg.foreversixty.app.raids.fail
import gg.foreversixty.app.raids.ok
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiMode
import gg.foreversixty.app.network.ClasspathFixtureSource
import gg.foreversixty.app.network.FixtureEngine
import gg.foreversixty.app.network.ForeverHttpClient
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.TokenProvider
import gg.foreversixty.app.rankings.logic.LiveRankingsService
import gg.foreversixty.app.rankings.logic.BoardPhase
import gg.foreversixty.app.rankings.logic.CharacterRoute
import gg.foreversixty.app.rankings.logic.Cohort
import gg.foreversixty.app.rankings.logic.Measure
import gg.foreversixty.app.rankings.logic.OwnRow
import gg.foreversixty.app.rankings.logic.Rankings
import gg.foreversixty.app.rankings.logic.RankingsCopy
import gg.foreversixty.app.rankings.logic.RankingsStore
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.reports.logic.ParseTier
import java.io.IOException
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** The one F11 rankings test list (spec §5.3), mirrored in iOS's RankingsLogicTests. */
@RunWith(AndroidJUnit4::class)
class RankingsLogicTest {
    private val now = RankingsTestData.now
    private fun copy() = RankingsCopy(resources(), RankingsTestData.zone)
    private val offline = ApiError.Network(IOException("offline"))
    private val phases = RankingsTestData.phases().phases
    private val encounters = RankingsTestData.encounters().encounters
    private val specs = RankingsTestData.specs()

    private fun filters(viewer: RankingsViewer = RankingsTestData.viewer()) =
        Rankings.defaults(viewer, specs, phases, now, "onyxia")

    private fun store(service: FakeRankingsService = FakeRankingsService(), viewer: RankingsViewer = RankingsTestData.viewer(), cache: MemoryRaidCache = MemoryRaidCache()) =
        RankingsStore(service, cache, viewer, specs) { now }

    // The rank's colour (web row-view.ts)

    @Test fun `first of a board is its 100th percentile and reads pink, never gold`() {
        assertEquals(100.0, Rankings.rankPercentile(1, 212), 0.0001)
        assertEquals(ParseTier.Pink, Rankings.rankTier(1, 212))
    }

    @Test fun `rank colours step down the ladder floors`() {
        assertEquals(ParseTier.Pink, Rankings.rankTier(3, 212))
        assertEquals(ParseTier.Pink, Rankings.rankTier(4, 212))
        assertEquals(ParseTier.Orange, Rankings.rankTier(5, 212))
        assertEquals(ParseTier.Orange, Rankings.rankTier(11, 212))
        assertEquals(ParseTier.Purple, Rankings.rankTier(14, 212))
        assertEquals(ParseTier.Blue, Rankings.rankTier(100, 212))
        assertEquals(ParseTier.Green, Rankings.rankTier(150, 212))
        assertEquals(ParseTier.Grey, Rankings.rankTier(204, 212))
    }

    @Test fun `an empty board does not divide by zero`() {
        assertEquals(ParseTier.Pink, Rankings.rankTier(1, 0))
    }

    @Test fun `ordinals`() {
        assertEquals(listOf("1st", "2nd", "3rd", "4th", "11th", "12th", "13th", "14th", "21st", "204th", "112th"),
            listOf(1, 2, 3, 4, 11, 12, 13, 14, 21, 204, 112).map(Rankings::ordinal))
    }

    // Crests

    @Test fun `a single-class board has no crest per row, a mixed board does`() {
        val rows = RankingsTestData.page().rows
        assertFalse(Rankings.showsCrests(rows))
        val mixed = rows.mapIndexed { i, r -> if (i == 1) r.copy(player = r.player.copy(className = "Mage")) else r }
        assertTrue(Rankings.showsCrests(mixed))
    }

    // Defaults and the query

    @Test fun `defaults follow the character, its cohort, ruleset, role's measure and the current phase`() {
        val f = filters()
        assertEquals(Cohort("shaman", "Elemental"), f.cohort)
        assertEquals(Measure.Damage, f.measure)
        assertEquals("normal", f.ruleset)
        assertEquals("", f.region)
        assertEquals("raids-1", f.phase)
        assertEquals("", f.faction)
        assertFalse(f.todayOnly)
    }

    @Test fun `a healer is judged on healing and a tank on damage taken`() {
        val resto = RankingsTestData.lightningbolltz.copy(specName = "Restoration")
        assertEquals(Measure.Healing, filters(RankingsTestData.viewer(resto)).measure)
        val prot = RankingsTestData.lightningbolltz.copy(classSlug = "warrior", specName = "Protection")
        assertEquals(Measure.Taken, filters(RankingsTestData.viewer(prot)).measure)
    }

    @Test fun `signed out the board starts on every class, every ruleset, damage`() {
        val f = filters(RankingsViewer.SignedOut)
        assertNull(f.cohort)
        assertEquals("", f.ruleset)
        assertEquals(Measure.Damage, f.measure)
    }

    @Test fun `the phase at a moment is the last phase begun`() {
        assertEquals("launch", Rankings.phaseAt(phases, java.time.Instant.parse("2026-11-05T00:00:00Z")))
        assertEquals("beta", Rankings.phaseAt(phases, java.time.Instant.parse("2026-11-04T22:59:59Z")))
        assertEquals("raids-1", Rankings.phaseAt(phases, now))
        assertNull(Rankings.phaseAt(emptyList(), now))
    }

    @Test fun `the phase chips are the phases begun, newest first`() {
        assertEquals(listOf("raids-1", "launch", "beta", "pre-beta"), Rankings.phaseChoices(phases, now))
    }

    @Test fun `the default encounter is the one most recently killed, else the first listed`() {
        val history = RankingsTestData.voltaris().history
        assertEquals("onyxia", Rankings.defaultEncounter(history, encounters))
        val wipeFirst = listOf(history[1].copy(encounterId = 227), history[2])
        assertEquals("emperor-dagran-thaurissan", Rankings.defaultEncounter(wipeFirst, encounters))
        assertEquals("onyxia", Rankings.defaultEncounter(emptyList(), encounters))
        assertNull(Rankings.defaultEncounter(emptyList(), emptyList()))
    }

    @Test fun `the query places the current character among one best kill per character`() {
        val q = Rankings.query(filters().copy(region = "us", faction = "horde", todayOnly = true), "us/normal/lightningbolltz", 2)
        assertEquals("onyxia", q.encounter)
        assertEquals("dps", q.metric)
        assertEquals("shaman", q.classSlug)
        assertEquals("Elemental", q.spec)
        assertEquals("raids-1", q.phase)
        assertEquals("us", q.region)
        assertEquals("normal", q.ruleset)
        assertEquals("horde", q.faction)
        assertEquals("today", q.since)
        assertEquals("character", q.distinct)
        assertEquals("us/normal/lightningbolltz", q.character)
        assertEquals(2, q.page)
    }

    @Test fun `an all choice is left out of the query`() {
        val q = Rankings.query(filters(RankingsViewer.SignedOut).copy(phase = ""), null, 1)
        assertNull(q.classSlug)
        assertNull(q.spec)
        assertNull(q.ruleset)
        assertNull(q.region)
        assertNull(q.faction)
        assertNull(q.phase)
        assertNull(q.since)
        assertNull(q.character)
    }

    // Your row: in place, pinned, jump

    @Test fun `your row on a loaded page is marked in place`() {
        val page = RankingsTestData.page()
        val own = Rankings.placement(page.rows, page.characterRank)
        assertTrue(own is OwnRow.InPlace)
        assertEquals(13, (own as OwnRow.InPlace).index)
        assertEquals(14, own.rank)
    }

    @Test fun `your row off the page is pinned with its true rank`() {
        val page = RankingsTestData.offPage()
        val own = Rankings.placement(page.rows, page.characterRank)
        assertTrue(own is OwnRow.Pinned)
        assertEquals(204, (own as OwnRow.Pinned).rank.rank)
        assertEquals("Sparkwick", own.rank.row.player.name)
    }

    @Test fun `no character rank, no own row`() {
        assertEquals(OwnRow.None, Rankings.placement(RankingsTestData.page().rows, null))
    }

    @Test fun `the You tag marks any of the account's characters`() {
        val rows = RankingsTestData.offPage().rows + RankingsTestData.offPage().characterRank!!.row
        val viewer = RankingsTestData.viewer()
        assertEquals(listOf("Lightningbolltz", "Sparkwick"), rows.filter { Rankings.isYours(it, viewer) }.map { it.player.name }.distinct().sorted())
        assertFalse(Rankings.isYours(rows.first(), RankingsViewer.SignedOut))
    }

    // Paging

    @Test fun `a next page is added de-duplicated by rank`() {
        val first = RankingsTestData.page().rows
        assertEquals(first, Rankings.merge(first, first))
        val next = first.map { it.copy(rank = it.rank + 100) }
        val merged = Rankings.merge(first, next)
        assertEquals(200, merged.size)
        assertEquals((1..200).toList(), merged.map { it.rank })
    }

    // Words

    @Test fun `the filter line names every filter, all regions and every ruleset in words`() {
        val names = mapOf("onyxia" to "Onyxia")
        assertEquals("Onyxia · Normal · All regions · Phase 1 · Damage", copy().filterLine(filters(), names))
        val wide = filters().copy(ruleset = "", region = "eu", phase = "launch", measure = Measure.Execution)
        assertEquals("Onyxia · Every ruleset · EU · Launch · Execution", copy().filterLine(wide, names))
        val narrowed = filters().copy(faction = "horde", todayOnly = true)
        assertEquals("Onyxia · Normal · All regions · Phase 1 · Damage · Horde · Today only", copy().filterLine(narrowed, names))
    }

    @Test fun `the answer card's unit says characters when deduped, else kills`() {
        val cohort = Cohort("shaman", "Elemental")
        assertEquals("of 212 Elemental Shaman characters", copy().answerUnit(212, cohort, deduped = true))
        assertEquals("of 212 Elemental Shaman kills", copy().answerUnit(212, cohort, deduped = false))
    }

    @Test fun `the meta names the character on the page and off it`() {
        val names = mapOf("onyxia" to "Onyxia")
        val on = copy().answerMeta(RankingsTestData.page().characterRank!!.row, filters(), names)
        assertEquals("Lightningbolltz · best kill on Onyxia, Normal, Phase 1.", on.text)
        assertEquals("Lightningbolltz", on.name)
        val off = copy().answerMeta(RankingsTestData.offPage().characterRank!!.row, filters(), names)
        assertEquals("Sparkwick · best kill on Onyxia, Normal, Phase 1.", off.text)
    }

    @Test fun `a row's age and each separator never start a line`() {
        val meta = copy().rowMeta(RankingsTestData.page().rows.first(), now)
        assertTrue(meta.contains("2\u00A0d\u00A0ago"))
        assertFalse(meta.contains(" · "))
        assertTrue(meta.contains("\u00A0· "))
        assertTrue(meta.contains("40\u00A0players"))
    }

    @Test fun `a row's meta, value and unit`() {
        val row = RankingsTestData.page().rows.first()
        assertEquals("Hordebreakers\u00A0· 40\u00A0players\u00A0· 3:13\u00A0· 2\u00A0d\u00A0ago", copy().rowMeta(row, now))
        assertEquals("462", Rankings.valueText(row, Measure.Damage))
        assertEquals("97%", Rankings.valueText(row.copy(value = 0.968), Measure.Execution))
        assertEquals("DPS", copy().unit(Measure.Damage))
        assertEquals("Executed", copy().unit(Measure.Execution))
        assertEquals("40\u00A0players\u00A0· 3:13\u00A0· 2\u00A0d\u00A0ago", copy().rowMeta(row.copy(guild = null), now))
    }

    @Test fun `a row reads as one sentence`() {
        val row = RankingsTestData.page().characterRank!!.row
        assertEquals(
            "Rank 14 of 212, Lightningbolltz, 398 damage per second, Hordebreakers, your character",
            copy().rowLabel(row, 212, Measure.Damage, yours = true),
        )
    }

    @Test fun `Home's line from one read`() {
        val names = mapOf("onyxia" to "Onyxia")
        assertEquals("14th of 212 Elemental Shaman characters on Onyxia", copy().homeLine(RankingsTestData.page(), filters(), names))
        assertEquals("Not ranked yet.", copy().homeLine(RankingsTestData.empty(), filters(), names))
    }

    @Test fun `a jump link reads the ordinal`() {
        assertEquals("Jump to your row, 14th", copy().jump(14))
    }

    // The store

    @Test fun `the store loads defaults then the board, with the character placed`() = runTest {
        val service = FakeRankingsService().apply { characterAnswer = { ok(RankingsTestData.yours()) } }
        val s = store(service)
        s.load()
        val loaded = s.phase as BoardPhase.Loaded
        assertEquals(100, loaded.board.rows.size)
        assertEquals(212, loaded.board.total)
        assertEquals("onyxia", s.filters?.encounter)
        assertEquals("us/normal/lightningbolltz", service.queries.single().character)
        assertEquals("Onyxia", s.encounterNames["onyxia"])
    }

    @Test fun `paging adds the next page once, de-duplicated (fixture mode answers the same page)`() = runTest {
        val service = FakeRankingsService()
        val s = store(service)
        s.load()
        s.loadMore()
        val loaded = s.phase as BoardPhase.Loaded
        assertEquals(100, loaded.board.rows.size)
        assertFalse(loaded.board.hasMore)
        assertEquals(listOf(1, 2), service.queries.map { it.page })
    }

    @Test fun `a failed first load is the error, a failed refresh keeps the board`() = runTest {
        val service = FakeRankingsService().apply { rankingsAnswer = { fail(ApiError.Status(500)) } }
        val s = store(service)
        s.load()
        assertEquals(BoardPhase.Failed, s.phase)
        service.rankingsAnswer = { ok(RankingsTestData.page()) }
        s.load()
        service.rankingsAnswer = { fail(ApiError.Status(500)) }
        s.load()
        assertTrue(s.phase is BoardPhase.Loaded)
    }

    @Test fun `offline shows the cached board under the strip`() = runTest {
        val cache = MemoryRaidCache()
        store(FakeRankingsService(), cache = cache).load()
        val down = FakeRankingsService().apply {
            rankingsAnswer = { fail(offline) }
            encountersAnswer = { fail(offline) }
            phasesAnswer = { fail(offline) }
            characterAnswer = { fail(offline) }
        }
        val s = store(down, cache = cache)
        s.load()
        val loaded = s.phase as BoardPhase.Loaded
        assertEquals(now, loaded.board.offlineSince)
        assertEquals(100, loaded.board.rows.size)
    }

    @Test fun `a filter change reads page one again`() = runTest {
        val service = FakeRankingsService()
        val s = store(service)
        s.load()
        s.apply(s.filters!!.copy(measure = Measure.Healing))
        assertEquals("hps", service.queries.last().metric)
        assertEquals(1, service.queries.last().page)
        s.reset()
        assertEquals("dps", service.queries.last().metric)
    }

    @Test fun `an empty board is loaded with no rows`() = runTest {
        val service = FakeRankingsService().apply { rankingsAnswer = { ok(RankingsTestData.empty()) } }
        val s = store(service)
        s.load()
        val board = (s.phase as BoardPhase.Loaded).board
        assertTrue(board.rows.isEmpty())
        assertNull(board.characterRank)
    }

    @Test fun `a row's Watch changes the list only when the write lands`() = runTest {
        val service = FakeRankingsService()
        val s = store(service)
        s.load()
        val stormwrath = RankingsTestData.page().rows[1]
        assertTrue(stormwrath.player.key in s.watched)
        assertTrue(s.toggleWatch(stormwrath))
        assertFalse(stormwrath.player.key in s.watched)
        service.watchAnswer = { fail(ApiError.Status(500)) }
        assertFalse(s.toggleWatch(stormwrath))
        assertFalse(stormwrath.player.key in s.watched)
    }

    @Test fun `links open Rankings on an encounter, or a character, on Home`() {
        val nav = RankingsNavigator()
        assertTrue(nav.open(requireNotNull(DeepLink.parse(DeepLink.Rankings("onyxia").url()))))
        assertTrue(nav.showsBoard && nav.wantsHome)
        assertEquals("onyxia", nav.encounter)
        assertTrue(nav.open(requireNotNull(DeepLink.parse(DeepLink.Player("us", "normal", "Voltaris").url()))))
        assertEquals(CharacterRoute("us", "normal", "Voltaris"), nav.character)
        assertFalse(nav.open(DeepLink.Report("x")))
    }

    @Test fun `a link's encounter aims the board`() = runTest {
        val service = FakeRankingsService()
        val s = store(service)
        s.load("227")
        assertEquals("emperor-dagran-thaurissan", service.queries.single().encounter)
    }

    @Test fun `Home's line comes from one board read, in fixture mode through the generated endpoints`() = runTest {
        val source = ClasspathFixtureSource(requireNotNull(javaClass.classLoader))
        val config = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Fixture)
        val client = ApiClient(ForeverHttpClient.create(config, NoToken, FixtureEngine.create(source)))
        val catalog = ItemCatalog(StaticDataClient(FixtureEngine.create(source)))
        val feature = RankingsFeature(LiveRankingsService(client), MemoryRaidCache(), specs, catalog, { null }, { true }, { now })
        assertEquals("14th of 212 Elemental Shaman characters on Onyxia", feature.homeLine(RankingsTestData.viewer(), resources()))
        assertTrue(catalog.items(RankingsTestData.BUILD, "shaman").orEmpty().containsKey(RankingsTestData.voltaris().gear.first().worn?.id))
    }

    private object NoToken : TokenProvider {
        override suspend fun bearerToken(): String? = null
        override suspend fun discardBearerToken() = Unit
    }
}
