package gg.foreversixty.app.guides

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.guides.logic.BuildKind
import gg.foreversixty.app.guides.logic.ChangeTag
import gg.foreversixty.app.guides.logic.Changes
import gg.foreversixty.app.guides.logic.GuideDates
import gg.foreversixty.app.guides.logic.GuideSection
import gg.foreversixty.app.guides.logic.Guides
import gg.foreversixty.app.guides.logic.GuidesCopy
import gg.foreversixty.app.guides.logic.KeyDates
import gg.foreversixty.app.guides.logic.RotationKind
import gg.foreversixty.app.guides.logic.TierChoice
import gg.foreversixty.app.guides.logic.TierTone
import gg.foreversixty.app.guides.logic.Tiers
import gg.foreversixty.app.network.StaticEndpoints
import gg.foreversixty.app.raids.resources
import java.time.LocalDate
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** F9's logic (spec §5.1), read from the W3-C1 fixtures; iOS's GuidesLogicTests keep the same list. */
@RunWith(AndroidJUnit4::class)
class GuidesLogicTest {
    private val index = GuidesTestData.index()
    private fun copy() = GuidesCopy(resources())

    // F9.1 the Guides tab

    @Test
    fun `the class grid is the index's nine classes in its order`() {
        val cells = Guides.classCells(index)
        assertEquals(9, cells.size)
        assertEquals("warrior", cells.first().classSlug)
        assertEquals(listOf("Warrior", "Paladin", "Hunter", "Rogue", "Priest", "Shaman", "Mage", "Warlock", "Druid"), cells.map { it.className })
    }

    @Test
    fun `your guide is the current character's class and spec, and absent without one`() {
        assertEquals(GuidesTestData.ELEMENTAL, Guides.yourGuide(index, GuidesTestData.lightningbolltz)?.spec)
        assertEquals(GuidesTestData.FIRE, Guides.yourGuide(index, GuidesTestData.emberquill)?.spec)
        assertNull(Guides.yourGuide(index, null))
        assertNull(Guides.yourGuide(index, GuidesTestData.lightningbolltz.copy(spec = null)))
    }

    @Test
    fun `saved guides are the ones in the cache, sized by their saved bytes`() = runTest {
        val files = FakeGuideFiles()
        files.load(StaticEndpoints.getGuide(index.build, GuidesTestData.ELEMENTAL))
        files.load(StaticEndpoints.getGuide(index.build, GuidesTestData.FIRE))
        val saved = Guides.saved(index, index.build, files)
        assertEquals(setOf(GuidesTestData.ELEMENTAL, GuidesTestData.FIRE), saved.keys)
        val elementalBytes = files.cache.read(StaticEndpoints.getGuide(index.build, GuidesTestData.ELEMENTAL).path)?.body?.toByteArray()?.size?.toLong()
        assertEquals(elementalBytes, saved[GuidesTestData.ELEMENTAL])
    }

    @Test
    fun `the footer names the saved classes in grid order with their size, or says nothing is saved`() {
        val saved = mapOf(GuidesTestData.FIRE to 30_000L, GuidesTestData.ELEMENTAL to 22_000L)
        val footer = Guides.footer(index, saved)
        assertEquals(listOf("Shaman", "Mage"), footer?.classNames)
        assertEquals(52_000L, footer?.bytes)
        assertEquals(
            "Saved on this phone: Shaman, Mage · 51 KB. Guides update by themselves when the game data does.",
            copy().footer(footer),
        )
        assertNull(Guides.footer(index, emptyMap()))
        assertEquals("Nothing is saved on this phone yet. Open a guide once and it stays.", copy().footer(null))
    }

    @Test
    fun `sizes read in KB under a megabyte and in MB with one decimal over`() {
        assertEquals("47 KB", Guides.formatSize(48_000))
        assertEquals("2.4 MB", Guides.formatSize(2_516_582))
        assertEquals("1 KB", Guides.formatSize(10))
    }

    @Test
    fun `the card meta is the role and the guide's updated day`() {
        val spec = requireNotNull(Guides.yourGuide(index, GuidesTestData.lightningbolltz))
        assertEquals("DPS · Updated 1 Nov 2026", copy().cardMeta(spec.role, GuideDates.long(spec.updated)))
        assertEquals("1 Nov", GuideDates.short("2026-11-01"))
        assertEquals("Healer", copy().role("healer"))
        assertEquals("Tank", copy().role("tank"))
    }

    // F9.2 the guide

    @Test
    fun `the strip draws exactly the guide's sections, in its order, with their TalkBack labels`() {
        val sections = Guides.sections(GuidesTestData.elemental())
        assertEquals(GuideSection.entries.toList(), sections)
        assertEquals("Talents, section 2 of 6", copy().sectionLabel(sections, GuideSection.Talents))
        val noRotation = GuidesTestData.elemental().copy(sections = listOf("overview", "talents", "unknown", "races"))
        assertEquals(listOf(GuideSection.Overview, GuideSection.Talents, GuideSection.Races), Guides.sections(noRotation))
    }

    @Test
    fun `the head names the recommended races, and a stat chip reads its place`() {
        val guide = GuidesTestData.elemental()
        assertEquals("Updated 1 Nov 2026 · Recommended races: Dwarf, Orc", copy().headMeta(guide))
        assertEquals("Updated 1 Nov 2026", copy().headMeta(guide.copy(recommendedRaces = emptyList())))
        assertEquals("Priority 1, Spell power", copy().statLabel(0, guide.statPriority.first()))
    }

    @Test
    fun `rotation rows are numbered per list in the guide's own order, each with its rank`() {
        val groups = Guides.rotation(GuidesTestData.elemental().rotation)
        assertEquals(listOf(RotationKind.Opener, RotationKind.Priority, RotationKind.Cooldowns), groups.map { it.kind })
        val priority = groups.single { it.kind == RotationKind.Priority }.rows
        assertEquals((1..priority.size).toList(), priority.map { it.number })
        assertEquals("Lava Burst", priority.first().spell.name)
        assertEquals("Rank 4", copy().rank(priority.first().spell.rank))
        assertNull(copy().rank(null))
        assertTrue(Guides.rotation(null).isEmpty())
    }

    @Test
    fun `a spell with no icon keeps the neutral tile and a general line has none`() {
        val rows = Guides.rotation(GuidesTestData.elemental().rotation).flatMap { it.rows }.map { it.spell }
        val fireNova = rows.first { it.name == "Fire Nova" }
        assertNull(fireNova.icon)
        assertTrue(Guides.hasTile(fireNova))
        val general = rows.first { it.spellId == null && it.icon == null }
        assertFalse(Guides.hasTile(general))
        assertTrue(Guides.hasTile(rows.first { it.name == "Lava Burst" }))
    }

    @Test
    fun `talents read their ranks with the plural rule`() {
        val key = GuidesTestData.elemental().talents.key
        assertEquals("5 ranks", copy().ranks(key.first { it.name == "Elemental Fury" }.maxRank))
        assertEquals("1 rank", copy().ranks(key.first { it.name == "Lava Burst" }.maxRank))
    }

    @Test
    fun `build rows show the raid build then the leveling build, totals summed per tree`() {
        val rows = Guides.builds(GuidesTestData.elemental().talents)
        assertEquals(listOf(BuildKind.Raid, BuildKind.Leveling), rows.map { it.kind })
        assertEquals(listOf("31 / 10 / 10", "30 / 3 / 18"), rows.map { it.totals })
    }

    @Test
    fun `a guide whose raid build is its leveling build shows one raid row`() {
        val rows = Guides.builds(GuidesTestData.protection().talents)
        assertEquals(listOf(BuildKind.Raid), rows.map { it.kind })
        assertEquals(GuidesTestData.protection().talents.levelingBuild?.totals?.joinToString(" / "), rows.single().totals)
    }

    // F9.4 tier lists

    @Test
    fun `the tier list opens on the character's role and faction, else DPS and Horde`() {
        assertEquals(TierChoice("dps", "horde"), Tiers.defaultChoice(index, GuidesTestData.emberquill))
        assertEquals(TierChoice("tank", "horde"), Tiers.defaultChoice(index, GuidesTestData.kraggor))
        assertEquals(TierChoice("tank", "alliance"), Tiers.defaultChoice(index, GuidesTestData.kraggor.copy(faction = "alliance")))
        assertEquals(TierChoice("dps", "horde"), Tiers.defaultChoice(index, null))
    }

    @Test
    fun `rows keep the file's order and letters, mark ties and your row`() {
        val list = Tiers.list(GuidesTestData.tiers(), TierChoice("dps", "horde"))
        val rows = Tiers.rows(list, GuidesTestData.FIRE)
        assertEquals(list.rows.map { it.rank }, rows.map { it.row.rank })
        assertEquals("Top", rows.first().gap)
        assertEquals(TierTone.Legendary, rows.first().tone)
        val tied = rows.first { it.row.tie != null }
        assertEquals("≈ ${tied.row.gapText}", tied.gap)
        assertEquals(listOf(GuidesTestData.FIRE), rows.filter { it.isYours }.map { it.row.spec })
        assertEquals(TierTone.Muted, Tiers.tone("D"))
        assertEquals(TierTone.Epic, Tiers.tone("A"))
        assertEquals(TierTone.Rare, Tiers.tone("B"))
        assertEquals(TierTone.Uncommon, Tiers.tone("C"))
        assertNull(Tiers.tone(null))
    }

    @Test
    fun `the callout is your spec's row with the list's last note, and absent without one`() {
        val list = Tiers.list(GuidesTestData.tiers(), TierChoice("dps", "horde"))
        val callout = requireNotNull(Tiers.callout(list, GuidesTestData.FIRE))
        assertTrue(callout.row.title.startsWith("Fire Mage is "))
        assertEquals(list.notes.last(), callout.note)
        assertNull(Tiers.callout(list, null))
        assertNull(Tiers.callout(list, GuidesTestData.PROTECTION))
    }

    @Test
    fun `tier rows read rank, letter, gap and race to TalkBack`() {
        val dps = Tiers.list(GuidesTestData.tiers(), TierChoice("dps", "horde"))
        val rows = Tiers.rows(dps, null)
        assertEquals("Fury Warrior, rank 1 of 20, tier S, top spec, Troll", copy().tierLabel(rows[0], rows.size))
        val second = rows[1].row
        val gap = String.format(java.util.Locale.US, "%.1f", second.gapPercent)
        assertEquals("${second.name}, rank 2 of 20, tier ${second.tier}, $gap percent behind the top, ${second.race}", copy().tierLabel(rows[1], rows.size))
        val tied = rows.first { it.row.tie != null }
        assertTrue(copy().tierLabel(tied, rows.size).endsWith(", Within 1 percent of the next spec"))
        val tank = Tiers.rows(Tiers.list(GuidesTestData.tiers(), TierChoice("tank", "horde")), null)
        assertEquals("Protection Warrior, rank 1 of 3, top spec, Tauren", copy().tierLabel(tank[0], tank.size))
    }

    // F9.5 what changed

    @Test
    fun `key dates drop the past, keep the order of iso, and run while in progress`() {
        val rows = KeyDates.upcoming(GuidesTestData.dates(), LocalDate.of(2026, 11, 1))
        assertEquals(listOf("Oct 27", "Dec 9"), rows.map { it.date.key })
        assertEquals("Name reservation open through Nov 3", rows[0].running)
        assertNull(rows[0].note)
        assertNull(rows[1].running)
        assertEquals(rows[1].date.note, rows[1].note)
        val shuffled = KeyDates.upcoming(GuidesTestData.dates().reversed(), LocalDate.of(2026, 9, 1))
        assertEquals(shuffled.map { it.date.iso }, shuffled.map { it.date.iso }.sorted())
    }

    @Test
    fun `changelog rows carry their month and day and a Site or Game tag`() {
        val entry = GuidesTestData.changelog().entries.first()
        assertEquals("10/28", GuideDates.monthDay(entry.date))
        assertEquals(ChangeTag.Site, Changes.tag(entry))
        assertEquals(ChangeTag.Game, Changes.tag(entry.copy(kind = "patch")))
        assertEquals("Site", copy().tag(ChangeTag.Site))
        assertEquals("Game", copy().tag(ChangeTag.Game))
    }

    // Sync with the nightly data

    @Test
    fun `new data shows only when a different build is found while one is shown`() {
        assertTrue(Guides.isNewData(shown = "1.60.1.70658", latest = "1.60.1.70707"))
        assertFalse(Guides.isNewData(shown = "1.60.1.70707", latest = "1.60.1.70707"))
        assertFalse(Guides.isNewData(shown = "1.60.1.70707", latest = null))
        assertFalse(Guides.isNewData(shown = null, latest = "1.60.1.70707"))
    }
}
