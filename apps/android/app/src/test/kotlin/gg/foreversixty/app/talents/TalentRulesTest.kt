package gg.foreversixty.app.talents

import android.content.Context
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/**
 * The planner's rules on the app (spec F6 §4; web/src/lib/planner band-compare, connectors,
 * rules, styles), over the real talent files. The same list runs on iOS (TalentRulesTests.swift).
 */
@RunWith(AndroidJUnit4::class)
class TalentRulesTest {
    private val warrior = TalentTestData.warrior()
    private val shaman = TalentTestData.shaman()
    private val arms = TalentRules.trees(warrior).first()

    @Test
    fun `a tree string reads one base-36 digit per talent, a bad digit as 0`() {
        assertEquals(listOf(5, 0, 3, 10), TalentRules.ranksOfTree("503a"))
        assertEquals(listOf(5, 0, 1), TalentRules.ranksOfTree("5?1"))
    }

    @Test
    fun `a band string splits into three trees, empty trees kept`() {
        assertEquals(listOf(listOf(3, 2), listOf(0, 5), emptyList()), TalentRules.ranksOfBand("32-05-"))
    }

    @Test
    fun `ranks clamp to each talent's maximum, a short string and a missing tree read as 0`() {
        val ranks = TalentRules.rankById(warrior, listOf(listOf(9)))
        val heroicStrike = arms.talents.first()
        assertEquals(heroicStrike.maxRank, ranks[heroicStrike.id])
        assertEquals(heroicStrike.maxRank, ranks.values.sum())
        assertEquals(TalentRules.trees(warrior).sumOf { it.talents.size }, ranks.size)
    }

    @Test
    fun `the spec's vector - diffCount 2, missingCount 5, four dashed cells, totals 28 3 17`() {
        val guide = TalentRules.rankById(shaman, TalentRules.ranksOfBand(TalentTestData.SHAMAN_GUIDE))
        val yours = TalentRules.rankById(shaman, TalentRules.ranksOfBand(TalentTestData.SHAMAN_YOURS))
        val diff = TalentRules.compare(shaman, yours, guide)
        assertEquals(2, diff.diffCount)
        assertEquals(5, diff.missingCount)
        val names = TalentRules.trees(shaman).flatMap { it.talents }.filter { it.id in diff.differing }.map { it.name }.toSet()
        assertEquals(setOf("Reverberation", "Call of Flame", "Elemental Warding", "Tidal Focus"), names)
        assertEquals(TalentStatus.Differs(2), diff.status)
        assertEquals(listOf(28, 3, 17), TalentRules.pointsPerTree(shaman, yours))
        assertEquals(3, TalentRules.pointsAvailable(60) - yours.values.sum())
    }

    @Test
    fun `the same build matches, with nothing marked, totals 30 3 18`() {
        val guide = TalentRules.rankById(shaman, TalentRules.ranksOfBand(TalentTestData.SHAMAN_GUIDE))
        val diff = TalentRules.compare(shaman, guide, guide)
        assertEquals(TalentDiff(0, 0, 0, emptySet()), diff)
        assertEquals(TalentStatus.Matches, diff.status)
        assertEquals(listOf(30, 3, 18), TalentRules.pointsPerTree(shaman, guide))
    }

    @Test
    fun `a build inside the guide's and short of it reads as points still to take`() {
        val guide = TalentRules.rankById(warrior, TalentRules.ranksOfBand(TalentTestData.BAND_60))
        val yours = TalentRules.rankById(warrior, TalentRules.ranksOfBand(TalentTestData.BAND_50))
        val diff = TalentRules.compare(warrior, yours, guide)
        assertEquals(0, diff.diffCount)
        assertEquals(0, diff.extraCount)
        assertEquals(TalentStatus.ToTake(10), diff.status)
    }

    @Test
    fun `level 58 - 49 points holding the whole 41-point guide match, 8 beyond it`() {
        val guide = TalentRules.rankById(warrior, TalentRules.ranksOfBand(TalentTestData.BAND_50))
        val yours = TalentRules.rankById(warrior, TalentTestData.yoursTrees)
        val diff = TalentRules.compare(warrior, yours, guide)
        assertEquals(0, diff.diffCount)
        assertEquals(0, diff.missingCount)
        assertEquals(8, diff.extraCount)
        assertEquals(TalentStatus.Matches, diff.status)
    }

    @Test
    fun `stray points differ only up to the guide's total - 10 stray minus (46 - 41) = 5`() {
        val guide = TalentRules.rankById(warrior, TalentRules.ranksOfBand(TalentTestData.BAND_50))
        val yours = TalentRules.rankById(warrior, TalentRules.ranksOfBand(TalentTestData.STRAY_AND_SHORT))
        val diff = TalentRules.compare(warrior, yours, guide)
        assertEquals(5, diff.diffCount)
        assertEquals(5, diff.missingCount)
        assertEquals(5, diff.extraCount)
        assertEquals(TalentStatus.Differs(5), diff.status)
    }

    @Test
    fun `cells read maxed, filled, open and locked`() {
        val ranks = TalentRules.rankById(warrior, TalentTestData.yoursTrees)
        val byName = arms.talents.associateBy { it.name }
        fun state(name: String) = TalentRules.cellState(requireNotNull(byName[name]), arms, ranks)
        assertEquals(TalentCellState.Maxed, state("Improved Heroic Strike"))
        assertEquals(TalentCellState.Filled, state("Deflection"))
        assertEquals(TalentCellState.Open, state("Improved Tactical Mastery"))
        assertEquals(TalentCellState.Locked, state("Anger Management")) // prerequisite unmet
        assertEquals(TalentCellState.Locked, state("Bloodthrill")) // tier 4 needs 20 points
    }

    @Test
    fun `prerequisite lines follow connectorsFor, gold once met`() {
        val ranks = TalentRules.rankById(warrior, TalentTestData.yoursTrees)
        val lines = TalentRules.connectors(arms, ranks)
        assertEquals(4, lines.size)
        assertTrue(lines.all { it.shape == TalentConnectorShape.Down })
        assertEquals(listOf(105_947, 105_950), lines.filter { it.met }.map { it.toId }.sorted())
    }

    @Test
    fun `a line is down a column, across a row, or down then across`() {
        fun line(from: TalentCell, to: TalentCell) = TalentConnector(from, to, 1, 2, false).shape
        assertEquals(TalentConnectorShape.Down, line(TalentCell(0, 1), TalentCell(2, 1)))
        assertEquals(TalentConnectorShape.Across, line(TalentCell(1, 0), TalentCell(1, 1)))
        assertEquals(TalentConnectorShape.Elbow, line(TalentCell(1, 0), TalentCell(2, 2)))
    }

    @Test
    fun `points available follow the level, capped at 51`() {
        assertEquals(0, TalentRules.pointsAvailable(9))
        assertEquals(49, TalentRules.pointsAvailable(58))
        assertEquals(51, TalentRules.pointsAvailable(70))
    }

    @Test
    fun `the grid is as large as the tree's last tier and column`() {
        assertEquals(TalentCell(7, 4), TalentRules.gridSize(arms))
    }

    @Test
    fun `the tree with the most points opens, ties to the first`() {
        assertEquals(1, TalentRules.openingTree(listOf(17, 32, 0)))
        assertEquals(0, TalentRules.openingTree(listOf(10, 10, 0)))
        assertEquals(0, TalentRules.openingTree(listOf(0, 0, 0)))
    }

    @Test
    fun `the band follows the level - 60 from 60, the decade from 20, none under 20`() {
        assertNull(TalentGuide.bandForLevel(19))
        assertEquals(20, TalentGuide.bandForLevel(20))
        assertEquals(50, TalentGuide.bandForLevel(59))
        assertEquals(60, TalentGuide.bandForLevel(60))
        assertNull(TalentGuide.bandForLevel(null))
    }

    @Test
    fun `the entry is the raid preset when the band has one, else bare, for the faction`() {
        val file = TalentTestData.fury()
        assertEquals("raid", TalentGuide.entry(file, 60, "horde")?.preset)
        assertEquals("bare", TalentGuide.entry(file, 50, "horde")?.preset)
        assertNull(TalentGuide.entry(file, 40, "horde"))
        assertNull(TalentGuide.entry(file, 60, "alliance"))
    }

    @Test
    fun `the band label is 60, or the span below it`() {
        val res = ApplicationProvider.getApplicationContext<Context>().resources
        assertEquals("60", TalentGuide.label(60, res))
        assertEquals("50 to 59", TalentGuide.label(50, res))
    }

    @Test
    fun `an export reads on another build only when the layouts match`() {
        assertTrue(TalentGuide.sameLayout(warrior, warrior))
        assertFalse(TalentGuide.sameLayout(warrior, shaman))
    }
}
