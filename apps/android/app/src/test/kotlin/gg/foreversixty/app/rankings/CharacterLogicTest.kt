package gg.foreversixty.app.rankings

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.home.GearModel
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.home.HomePresenter
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.model.CharacterGearSlotWorn
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.paperdoll.Paperdoll
import gg.foreversixty.app.raids.MemoryRaidCache
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RecordingHaptics
import gg.foreversixty.app.raids.fail
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.rankings.logic.CharacterCard
import gg.foreversixty.app.rankings.logic.CharacterCopy
import gg.foreversixty.app.rankings.logic.CharacterPhase
import gg.foreversixty.app.rankings.logic.CharacterRoute
import gg.foreversixty.app.rankings.logic.CharacterStore
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.rankings.logic.ScoreComponent
import java.io.IOException
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** The F11.4 to F11.7 test list and the Paperdoll's slot model, mirrored in iOS's CharacterLogicTests. */
@RunWith(AndroidJUnit4::class)
class CharacterLogicTest {
    private val now = RankingsTestData.now
    private val res get() = resources()
    private fun copy() = CharacterCopy(res, RankingsTestData.zone)
    private val voltarisRoute = CharacterRoute("us", "normal", "Voltaris")
    private fun store(service: FakeRankingsService, viewer: RankingsViewer = RankingsTestData.viewer(), cache: MemoryRaidCache = MemoryRaidCache()) =
        CharacterStore(voltarisRoute, service, cache, viewer) { now }

    // Score card

    @Test fun `each bar is its value over the API's cap`() {
        val p = RankingsTestData.voltaris()
        val bars = CharacterCard.bars(p.components, p.componentCaps)
        assertEquals(ScoreComponent.entries, bars.map { it.component })
        assertEquals(listOf(412f / 500, 620f / 700, 540f / 600, 180f / 200, 78f / 100), bars.map { it.fraction })
    }

    @Test fun `no score has no bars, never a zero`() {
        val p = RankingsTestData.noLogs()
        assertFalse(CharacterCard.hasScore(p))
        assertTrue(CharacterCard.bars(p.components, p.componentCaps).isEmpty())
        assertNull(copy().scoreLabel(p))
    }

    @Test fun `the sparkline runs oldest first from lowest to highest`() {
        val points = CharacterCard.sparkline(RankingsTestData.voltaris().scoreHistory.reversed())
        assertEquals(5, points.size)
        assertEquals(0f, points.first())
        assertEquals(1f, points.last())
        assertTrue(CharacterCard.sparkline(emptyList()).isEmpty())
    }

    @Test fun `the head reads race, spec, class, level, guild, region and ruleset`() {
        val p = RankingsTestData.voltaris()
        val spec = CharacterCard.specName(p, RankingsTestData.specs())
        assertEquals("Elemental", spec)
        assertEquals("Orc Elemental Shaman · Level 60", copy().descriptor(p, spec))
        assertEquals("Hordebreakers · US Normal", copy().where(p))
        assertEquals("Mage", copy().descriptor(RankingsTestData.noLogs(), null))
        assertEquals("US PvP", copy().where(RankingsTestData.noLogs()))
    }

    @Test fun `the score, its tier and its sentence`() {
        val p = RankingsTestData.voltaris()
        assertEquals("1,830", copy().score(1830))
        assertEquals("Epic", copy().tier(p.tier))
        assertEquals(4, CharacterCard.tierQuality(p.tier))
        assertEquals("As of 17 Dec", copy().asOf(p.asOf))
        assertEquals(
            "Forever Score 1,830, Epic. Gear 412 of 500. Experience 620 of 700. Parses 540 of 600. Attendance 180 of 200. Discipline 78 of 100.",
            copy().scoreLabel(p),
        )
    }

    @Test fun `bests and recent fights`() {
        val p = RankingsTestData.voltaris()
        val best = p.best.first()
        assertEquals("462", copy().value(best))
        assertEquals("DPS", copy().unit(best.metric))
        assertEquals("Normal · 2 d ago", copy().bestMeta(p, best, now))
        assertEquals("Kill · 2 d ago", copy().fightMeta(p.history[0], now))
        assertEquals("Wipe · 2 d ago", copy().fightMeta(p.history[1], now))
        assertEquals("Elemental · first seen 10 Dec", copy().buildMeta("Elemental", p.buildsSeen.first().firstSeen))
    }

    @Test fun `parse history only on your own character with history`() {
        val viewer = RankingsTestData.viewer()
        assertTrue(CharacterCard.showsParseHistory(RankingsTestData.yours(), viewer))
        assertFalse(CharacterCard.showsParseHistory(RankingsTestData.voltaris(), viewer))
        assertFalse(CharacterCard.showsParseHistory(RankingsTestData.yours(), viewer.copy(history = false)))
    }

    @Test fun `the share link is the web character page`() {
        assertEquals("https://foreversixty.gg/character/us/normal/voltaris", CharacterCard.shareUrl(voltarisRoute))
        assertEquals("https://foreversixty.gg/character/us/normal/lightningbolltz", CharacterCard.shareUrl("us/normal/lightningbolltz"))
        assertNull(CharacterCard.shareUrl("nonsense"))
    }

    // Store: not found, offline, watch

    @Test fun `a 404 is the not-found state`() = runTest {
        val service = FakeRankingsService().apply { characterAnswer = { fail(RaidFixture.error("character.card.error-not-found", 404)) } }
        val s = store(service)
        s.load()
        assertEquals(CharacterPhase.NotFound, s.phase)
    }

    @Test fun `offline shows the cached card under the strip`() = runTest {
        val cache = MemoryRaidCache()
        store(FakeRankingsService(), cache = cache).load()
        val s = store(FakeRankingsService().apply { characterAnswer = { fail(ApiError.Network(IOException("offline"))) } }, cache = cache)
        s.load()
        assertEquals(now, (s.phase as CharacterPhase.Loaded).offlineSince)
    }

    @Test fun `the star reads the watch list, fills at once and says success`() = runTest {
        val service = FakeRankingsService().apply { watchingAnswer = { gg.foreversixty.app.raids.ok(RankingsTestData.watching().copy(rows = emptyList())) } }
        val s = store(service)
        s.load()
        assertEquals(false, s.watch?.watching)
        val haptics = RecordingHaptics()
        s.toggleWatch(haptics)
        assertEquals(true, s.watch?.watching)
        assertEquals(listOf("success"), haptics.calls)
        assertEquals("watch Voltaris", service.calls.last())
    }

    @Test fun `a failed watch reverts and says so`() = runTest {
        val service = FakeRankingsService().apply {
            watchingAnswer = { gg.foreversixty.app.raids.ok(RankingsTestData.watching().copy(rows = emptyList())) }
            watchAnswer = { fail(RaidFixture.error("watching.put.error-limit", 409)) }
        }
        val s = store(service)
        s.load()
        s.toggleWatch(RecordingHaptics())
        assertEquals(false, s.watch?.watching)
        assertEquals(true, s.watch?.failed)
    }

    @Test fun `a watched character's star is filled and stops watching`() = runTest {
        val service = FakeRankingsService()
        val s = store(service)
        s.load()
        assertEquals(true, s.watch?.watching)
        s.toggleWatch(RecordingHaptics())
        assertEquals(false, s.watch?.watching)
        assertEquals("unwatch Voltaris", service.calls.last())
    }

    @Test fun `signed out there is no star`() = runTest {
        val s = store(FakeRankingsService(), RankingsViewer.SignedOut)
        s.load()
        assertNull(s.watch)
    }

    // Paperdoll

    @Test fun `the gold dot marks only the slots whose pick gains`() {
        val items = RankingsTestData.items().items.associateBy { it.id }
        val voltaris = Paperdoll.fromProfile(RankingsTestData.voltaris().gear, items, res)
        assertEquals(17, voltaris.size)
        assertEquals(listOf("trinket2"), voltaris.filter { it.better }.map { it.slot })
        val yours = Paperdoll.fromProfile(RankingsTestData.yours().gear, items, res)
        assertEquals(listOf("head", "shoulder", "chest", "waist", "legs", "trinket2"), yours.filter { it.better }.map { it.slot })
        val head = voltaris.first()
        assertEquals("Living Crown", head.name)
        assertEquals("Head · ilvl 61", head.detail)
        assertEquals("inv_helmet_32", head.item?.icon)
        assertEquals("Head, Living Crown, Epic, item level 61.", Paperdoll.label(head, res))
        assertEquals("Head, Coif of Elements, Rare, item level 62. A better item is on your best in slot list.", Paperdoll.label(yours.first(), res))
    }

    @Test fun `the frame lays sockets out as the character frame`() {
        val slots = Paperdoll.fromProfile(RankingsTestData.voltaris().gear, emptyMap(), res)
        val columns = Paperdoll.columns(slots)
        assertEquals(Paperdoll.LEFT, columns.left.map { it.slot })
        assertEquals(Paperdoll.RIGHT, columns.right.map { it.slot })
        assertEquals(Paperdoll.WEAPONS, columns.weapons.map { it.slot })
    }

    @Test fun `an off hand under a two-hander is empty with the words`() {
        val p = RankingsTestData.voltaris()
        val items = RankingsTestData.items().items.associateBy { it.id }
        val mace = items.getValue(2243).copy(twoHand = true)
        val gear = p.gear.map { if (it.slot == "off_hand") it.copy(worn = null, pick = null, gap = null) else it }
        val off = Paperdoll.fromProfile(gear, items + (mace.id to mace), res).first { it.slot == "off_hand" }
        assertTrue(off.isEmpty)
        assertEquals("Off hand", off.name)
        assertEquals("Two-hand in main hand", off.detail)
        assertFalse(off.better)
        assertEquals("Off hand, empty.", Paperdoll.label(off, res))
    }

    @Test fun `an item the data does not know keeps its name and level`() {
        val gear = listOf(RankingsTestData.voltaris().gear.first().copy(worn = CharacterGearSlotWorn(999, "Mystery Hat", 50), pick = null))
        val head = Paperdoll.fromProfile(gear, emptyMap(), res).first()
        assertEquals("Mystery Hat", head.name)
        assertNull(head.item?.icon)
        assertFalse(head.better)
        assertEquals("Neck", Paperdoll.fromProfile(gear, emptyMap(), res)[1].detail)
    }

    @Test fun `a nameless item our data does not know reads its slot and id, printed once`() {
        val gear = listOf(RankingsTestData.voltaris().gear.first { it.slot == "trinket2" }.copy(worn = CharacterGearSlotWorn(288887, null, null)))
        val trinket = Paperdoll.fromProfile(gear, emptyMap(), res).first { it.slot == "trinket2" }
        assertTrue(trinket.isUnknown)
        assertEquals("Trinket 2 · Item 288887", trinket.name)
        assertNull(trinket.detail)
        assertEquals("Trinket 2 · Item 288887. A better item is on your best in slot list.", Paperdoll.label(trinket, res))
    }

    @Test fun `Home's unknown wrist is a nameless socket with its slot and id`() {
        val home = HomeFixtures.home()
        val rows = GearModel.rows(home, emptyMap(), HomeFixtures.tooltipSources().items, res)
        val wrist = Paperdoll.fromGear(rows, res).first { it.slot == "wrist" }
        assertEquals("Wrist · Item 999999", wrist.name)
        assertTrue(wrist.isUnknown)
        assertNull(wrist.detail)
    }

    @Test fun `Home's Your gear as sockets keeps GearModel's rows, the gold dot where Home has an upgrade`() {
        val home = HomeFixtures.home()
        val sources = HomeFixtures.tooltipSources()
        val probe = HomePresenter.tooltipRequest(ItemLook(0), home, HomeFixtures.velaren, res)
        val blocker = GearModel.offHandBlocker(home, emptyMap(), sources.items, ItemTooltipModel.band(probe, sources.bis))
        val rows = GearModel.rows(home, emptyMap(), sources.items, res, blocker)
        val slots = Paperdoll.fromGear(rows, res)
        assertEquals(rows.map { it.slot }, slots.map { it.slot })
        assertEquals(rows.filter { it.pick != null }.map { it.slot }, slots.filter { it.better }.map { it.slot })
        val off = slots.first { it.slot == "off_hand" }
        assertEquals("Off hand", off.name)
        assertEquals("Two-hand in main hand", off.detail)
        assertFalse(off.better)
        assertFalse(slots.first { it.slot == "shoulder" }.better)
    }
}
