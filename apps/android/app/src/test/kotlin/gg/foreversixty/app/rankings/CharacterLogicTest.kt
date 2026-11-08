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
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.rankings.logic.BuildRow
import gg.foreversixty.app.rankings.logic.CellGain
import gg.foreversixty.app.rankings.logic.CharacterExtras
import gg.foreversixty.app.rankings.logic.CharacterLink
import gg.foreversixty.app.rankings.logic.CharacterPageComposer
import gg.foreversixty.app.rankings.logic.HeadAnswer
import gg.foreversixty.app.rankings.logic.HeadFact
import gg.foreversixty.app.rankings.logic.LinkedLine
import gg.foreversixty.app.rankings.logic.OwnerSections
import gg.foreversixty.app.rankings.logic.CharacterPhase
import gg.foreversixty.app.rankings.logic.CharacterRoute
import gg.foreversixty.app.rankings.logic.CharacterStore
import gg.foreversixty.app.rankings.logic.RankingsViewer
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
    private val voltarisRoute = CharacterRoute("us", "normal", "Voltaris")
    private fun store(service: FakeRankingsService, viewer: RankingsViewer = RankingsTestData.viewer(), cache: MemoryRaidCache = MemoryRaidCache()) =
        CharacterStore(voltarisRoute, service, cache, viewer) { now }

    // The page (CharacterPageComposer)

    private fun page(profile: CharacterProfile, extras: CharacterExtras, viewer: RankingsViewer) =
        CharacterPageComposer(res, RankingsTestData.specs(), RankingsTestData.zone).compose(
            profile, Paperdoll.fromProfile(profile.gear, CharacterFixtures.items, res, CharacterCard.classSlug(profile)), extras, viewer, CharacterFixtures.now,
        )

    private val elyraOwner get() = page(CharacterFixtures.elyra, CharacterFixtures.elyraOwnerExtras, CharacterFixtures.owner(CharacterFixtures.ELYRA_KEY))

    @Test fun `the sparkline runs oldest first from lowest to highest`() {
        val points = CharacterCard.sparkline(RankingsTestData.voltaris().scoreHistory.reversed())
        assertEquals(5, points.size)
        assertEquals(0f, points.first())
        assertEquals(1f, points.last())
        assertTrue(CharacterCard.sparkline(emptyList()).isEmpty())
    }

    @Test fun `recent fights keep their wipes, and the gear gap reads in points`() {
        val profile = RankingsTestData.voltaris()
        val view = page(profile, CharacterExtras(), RankingsViewer.SignedOut)
        assertEquals(profile.history.take(view.recent.size).map { it.kill != false }, view.recent.map { it.kill })
        assertTrue(view.recent.all { it.sub.startsWith("Kill · ") || it.sub.startsWith("Wipe · ") })
        assertEquals(profile.gearGapTotal?.let { "%.1f points".format(it) }, view.gearGap)
        assertEquals(5, view.score?.bars?.size)
    }

    @Test fun `the head names level, race, spec, split, guild, realm, ruleset, the supporter tag and the owner's sync`() {
        assertEquals(
            listOf("Level 60 Night Elf", "Holy Priest", "15/31/5", "<Cx66 Olympus>", "Doomhowl", "Hardcore US", "Supporter", "Addon · 1 d ago"),
            elyraOwner.facts.map {
                when (it) {
                    is HeadFact.Words -> it.text
                    is HeadFact.Link -> it.text
                    is HeadFact.Tag -> it.text
                }
            },
        )
        val visitor = page(CharacterFixtures.elyra, CharacterFixtures.elyraVisitorExtras, RankingsViewer.SignedOut)
        assertFalse(visitor.facts.any { it is HeadFact.Words && it.text.startsWith("Addon") })
    }

    @Test fun `the answer is the score beside the band's upgrades`() {
        assertEquals(HeadAnswer("1,830", "+128.9", "HPS"), elyraOwner.answer)
        assertNull(page(CharacterFixtures.elyra, CharacterExtras(), RankingsViewer.SignedOut).answer)
    }

    @Test fun `the owner's next line is the largest upgrade, or unspent talent points first`() {
        assertEquals(LinkedLine("Next: get Devout Skirt from Stratholme: Baron Rivendare.", "Open best in slot ›", CharacterLink.Bis), elyraOwner.next)
        val felsnap = page(CharacterFixtures.felsnap, CharacterFixtures.felsnapOwnerExtras, CharacterFixtures.owner(CharacterFixtures.FELSNAP_KEY))
        assertEquals("Next: spend your 2 talent points, then get Felbough Arm.", felsnap.next?.text)
        assertEquals(CharacterLink.Planner, felsnap.next?.link)
        assertEquals(2, CharacterPageComposer.unspentPoints(34, listOf(23, 0, 0)))
    }

    @Test fun `the owner's standing names the guild ladder, the spec's grade and the next raid`() {
        assertEquals(
            listOf("#3 of 14 healers in Cx66 Olympus, 60 raid", "Holy is graded A this week", "Onyxia's Lair in 33 days 8 hours, 18 in"),
            elyraOwner.standing.map { it.text },
        )
        assertTrue(page(CharacterFixtures.elyra, CharacterFixtures.elyraOwnerExtras, RankingsViewer.SignedOut).standing.isEmpty())
    }

    @Test fun `the gear runs left column, weapons, right column, the owner's gain against the pick`() {
        val cells = elyraOwner.cells
        assertEquals(CharacterPageComposer.DOLL_ORDER, cells.map { it.slot.slot })
        assertEquals(CellGain("+18.4", "HPS", ItemLook(900, "Devout Crown", 3)), cells.first().gain)
        assertEquals("Best in slot", cells.first { it.slot.slot == "wrist" }.gain?.value)
        assertEquals("Item level 62", cells.first().level)
        assertNull(cells.first { it.slot.slot == "off_hand" }.level)
        val felsnap = page(CharacterFixtures.felsnap, CharacterFixtures.felsnapOwnerExtras, CharacterFixtures.owner(CharacterFixtures.FELSNAP_KEY))
        assertEquals(CellGain("Same", null, ItemLook(904, "Garb of Florid Feathers", 3), muted = true), felsnap.cells.first { it.slot.slot == "chest" }.gain)
        assertTrue(page(CharacterFixtures.elyra, CharacterFixtures.elyraVisitorExtras, RankingsViewer.SignedOut).cells.all { it.gain == null })
    }

    @Test fun `the score names the part with most to gain and every part over its cap`() {
        val score = elyraOwner.score
        assertEquals("1,830", score?.score)
        assertEquals("Epic", score?.tier)
        assertEquals(4, score?.tierQuality)
        assertEquals("Most to gain: Experience, 180 points to the top.", score?.mostToGain)
        assertEquals("Preparation" to "100 of 100", score?.parts?.last())
        val noScore = page(CharacterFixtures.felsnap, CharacterExtras(), CharacterFixtures.owner(CharacterFixtures.FELSNAP_KEY))
        assertNull(noScore.score)
        assertEquals("No logs yet. Upload a log to get a rating.", noScore.noScoreLine)
    }

    @Test fun `the fights and builds read as the web`() {
        val best = elyraOwner.best.first()
        assertEquals(listOf("Lucifron", "3 Nov · Holy", "1,012", "HPS"), listOf(best.boss, best.sub, best.value, best.unit))
        assertEquals(98, best.percentile)
        assertEquals(7, elyraOwner.allFights)
        assertEquals(BuildRow("15/31/5", "Holy · first seen 3 Nov 2026"), elyraOwner.builds.first())
    }

    @Test fun `history and coach history only on the owner's scored character`() {
        assertEquals(OwnerSections(history = true, coach = true), elyraOwner.owner)
        assertNull(page(CharacterFixtures.elyra, CharacterExtras(), RankingsViewer.SignedOut).owner)
    }

    @Test fun `a page link opens its web page`() {
        assertEquals("https://foreversixty.gg/tiers", CharacterCard.linkUrl(CharacterLink.Web("/tiers"), voltarisRoute))
        assertEquals("https://foreversixty.gg/character/us/normal/voltaris#history", CharacterCard.linkUrl(CharacterLink.History, voltarisRoute))
        assertEquals("/guild/us/hardcore/cx66-olympus", CharacterPageComposer.guildPath("us", "hardcore", "Cx66 Olympus"))
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

    @Test fun `a relic class names the ranged socket Relic`() {
        val gear = RankingsTestData.voltaris().gear
        val shaman = Paperdoll.fromProfile(gear, emptyMap(), res, "shaman").first { it.slot == "ranged" }
        assertTrue(Paperdoll.label(shaman, res).startsWith("Relic"))
        val warrior = Paperdoll.fromProfile(gear, emptyMap(), res, "warrior").first { it.slot == "ranged" }
        assertTrue(Paperdoll.label(warrior, res).startsWith("Ranged"))
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
