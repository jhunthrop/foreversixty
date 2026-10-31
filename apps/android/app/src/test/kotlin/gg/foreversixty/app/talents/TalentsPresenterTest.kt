package gg.foreversixty.app.talents

import android.content.Context
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.R
import gg.foreversixty.app.model.TalentsFile
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** What the Talents screen says and draws (spec F6 §4 to §7). iOS: TalentsPresenterTests.swift. */
@RunWith(AndroidJUnit4::class)
class TalentsPresenterTest {
    private val res = ApplicationProvider.getApplicationContext<Context>().resources
    private val shaman = TalentTestData.shaman()
    private val warrior = TalentTestData.warrior()

    private fun presenter(
        yours: String? = TalentTestData.SHAMAN_YOURS,
        guide: TalentGuideBuild = TalentGuideBuild.Build(TalentTestData.SHAMAN_GUIDE, 60),
        export: TalentExport? = null,
        level: Int? = 60,
    ): TalentsPresenter {
        val mine = export ?: yours?.let { TalentExport.Ranks(TalentRules.ranksOfBand(it)) } ?: TalentExport.None
        return TalentsPresenter(shaman, guide, mine, level, "Elemental")
    }

    private fun id(name: String, file: TalentsFile = shaman) = TalentRules.trees(file).flatMap { it.talents }.first { it.name == name }.id

    @Test
    fun `differs - the count sentence, unspent and missing lines, tab totals and dots`() {
        val model = presenter()
        assertEquals(
            TalentStatusLine("Differs", TalentTagKind.Differs, "2 of your 48 points differ from the 60 guide build."),
            model.status(res),
        )
        assertEquals(listOf("3 points unspent.", "The guide build takes 5 points you have not."), model.notes(res))
        assertEquals(listOf(28, 3, 17), model.tabs(TalentBuildKind.Yours).map { it.points })
        assertEquals(listOf(3, 0, 1), model.tabs(TalentBuildKind.Yours).map { it.differing })
        assertEquals(listOf(30, 3, 18), model.tabs(TalentBuildKind.Guide).map { it.points })
        assertTrue(model.tabs(TalentBuildKind.Guide).all { it.differing == 0 })
        assertTrue(model.showsSegments)
        assertEquals(TalentBuildKind.Yours, model.initialBuild)
        assertEquals(0, model.openingTree)
        assertEquals(res.getString(R.string.talents_legend), model.legend(TalentBuildKind.Yours, res))
        assertEquals(res.getString(R.string.talents_tapHint), model.legend(TalentBuildKind.Guide, res))
    }

    @Test
    fun `matches - the guide sentence, no notes, the tap hint`() {
        val model = presenter(yours = TalentTestData.SHAMAN_GUIDE)
        assertEquals(TalentStatusLine("Matches", TalentTagKind.Matches, "Your build is the Elemental 60 guide build."), model.status(res))
        assertTrue(model.notes(res).isEmpty())
        assertEquals(res.getString(R.string.talents_tapHint), model.legend(TalentBuildKind.Yours, res))
    }

    @Test
    fun `a build inside the guide's reads as points still to take`() {
        val model = TalentsPresenter(
            warrior, TalentGuideBuild.Build(TalentTestData.BAND_60, 60),
            TalentExport.Ranks(TalentRules.ranksOfBand(TalentTestData.BAND_50)), 60, "Fury",
        )
        assertEquals("10 of the 60 guide build's points still to take.", model.status(res).sentence)
        assertEquals(listOf("10 points unspent."), model.notes(res))
    }

    @Test
    fun `a level past the band's build says why it differs`() {
        val model = TalentsPresenter(
            warrior, TalentGuideBuild.Build(TalentTestData.BAND_50, 50), TalentExport.Ranks(TalentTestData.yoursTrees), 58, "Fury",
        )
        assertEquals("8 of your 49 points differ from the 50 to 59 guide build.", model.status(res).sentence)
        assertEquals(
            listOf("Level 58 is past the 50 to 59 guide build's 41 points, so points past it show as differing."),
            model.notes(res),
        )
    }

    @Test
    fun `no export - the guide alone, no segments, the addon line`() {
        val model = presenter(yours = null)
        assertEquals(
            TalentStatusLine("Guide build", TalentTagKind.Muted, res.getString(R.string.home_next_talentsNone)),
            model.status(res),
        )
        assertEquals(listOf(TalentBuildKind.Guide), model.builds)
        assertFalse(model.showsSegments)
        assertEquals(TalentBuildKind.Guide, model.initialBuild)
        assertTrue(model.notes(res).isEmpty())
    }

    @Test
    fun `no export and no guide - both lines and no tree`() {
        val model = presenter(yours = null, guide = TalentGuideBuild.None)
        assertNull(model.status(res).tag)
        assertEquals(listOf(res.getString(R.string.talents_noList)), model.notes(res))
        assertFalse(model.hasTree)
    }

    @Test
    fun `a stale export draws the guide with the older-build line`() {
        val model = presenter(export = TalentExport.Stale)
        assertEquals(res.getString(R.string.talents_staleExport), model.status(res).sentence)
        assertEquals(listOf(TalentBuildKind.Guide), model.builds)
    }

    @Test
    fun `no build - your build alone with the no-list line and the unspent line`() {
        val model = presenter(guide = TalentGuideBuild.None)
        assertEquals(TalentStatusLine("Your build", TalentTagKind.Muted, res.getString(R.string.talents_noList)), model.status(res))
        assertEquals(listOf(TalentBuildKind.Yours), model.builds)
        assertEquals(listOf("3 points unspent."), model.notes(res))
        assertTrue(model.cells(TalentBuildKind.Yours, 0).all { it.delta == null })
    }

    @Test
    fun `no spec - your build alone with the spec line`() {
        val model = presenter(guide = TalentGuideBuild.NoSpec)
        assertEquals(res.getString(R.string.talents_pickSpec), model.status(res).sentence)
        assertFalse(model.showsSegments)
    }

    @Test
    fun `the guide failed - your build with the failed line and no compare`() {
        val model = presenter(guide = TalentGuideBuild.Failed)
        assertTrue(model.status(res).isError)
        assertEquals(res.getString(R.string.talents_guideFailed), model.status(res).sentence)
        assertFalse(model.showsSegments)
    }

    @Test
    fun `cells on Your build carry the signed chip and the guide rank, the guide's carry none`() {
        val model = presenter()
        val cells = model.cells(TalentBuildKind.Yours, 0).associateBy { it.talent.name }
        assertEquals("+2", cells.getValue("Reverberation").chip)
        assertEquals("−2", cells.getValue("Call of Flame").chip)
        assertEquals("Call of Flame, rank 1 of 3, guide rank 3", cells.getValue("Call of Flame").accessibilityLabel(res))
        assertEquals("Convection, rank 5 of 5", cells.getValue("Convection").accessibilityLabel(res))
        assertTrue(model.cells(TalentBuildKind.Guide, 0).all { it.chip == null })
    }

    @Test
    fun `the sheet - rank, current and next text, and the guide line only where ranks differ`() {
        val model = presenter()
        val sheet = requireNotNull(model.sheet(TalentBuildKind.Yours, id("Call of Flame"), res))
        assertEquals(1, sheet.rank)
        assertEquals(sheet.talent.ranks[0].description, sheet.current)
        assertEquals(sheet.talent.ranks[1].description, sheet.next)
        assertEquals("Guide build: rank 3/3", sheet.guideLine)
        assertFalse(sheet.isOver)
        assertTrue(requireNotNull(model.sheet(TalentBuildKind.Yours, id("Reverberation"), res)).isOver)
        assertNull(model.sheet(TalentBuildKind.Guide, id("Call of Flame"), res)?.guideLine)
        val convection = requireNotNull(model.sheet(TalentBuildKind.Yours, id("Convection"), res))
        assertNull(convection.next)
        assertNull(convection.guideLine)
    }

    @Test
    fun `the sheet names unmet requirements in the game's words`() {
        val model = presenter(guide = TalentGuideBuild.None, export = TalentExport.Ranks(listOf(emptyList(), emptyList(), emptyList())))
        val tree = TalentRules.trees(shaman).first()
        val gated = tree.talents.first { it.tier >= 2 && it.prereqTalentId != null }
        val sheet = requireNotNull(model.sheet(TalentBuildKind.Yours, gated.id, res))
        assertEquals("Requires ${gated.tier * 5} points in ${tree.name} Talents", sheet.requirements.first())
        assertEquals(2, sheet.requirements.size)
        val open = tree.talents.first { it.tier == 0 }
        assertTrue(requireNotNull(model.sheet(TalentBuildKind.Yours, open.id, res)).requirements.isEmpty())
    }
}
