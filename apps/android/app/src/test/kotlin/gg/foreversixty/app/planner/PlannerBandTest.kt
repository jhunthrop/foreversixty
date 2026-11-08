package gg.foreversixty.app.planner

import gg.foreversixty.app.auth.session.MemoryValue
import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.planner.logic.BandStatus
import gg.foreversixty.app.planner.logic.EditorPhase
import gg.foreversixty.app.planner.logic.EditorRequest
import gg.foreversixty.app.planner.logic.EditorNote
import gg.foreversixty.app.planner.logic.Fs1
import gg.foreversixty.app.planner.logic.ImportOutcome
import gg.foreversixty.app.planner.logic.PlannerBand
import gg.foreversixty.app.planner.logic.PlannerCodes
import gg.foreversixty.app.planner.logic.PlannerDrafts
import gg.foreversixty.app.planner.logic.PlannerEditorModel
import gg.foreversixty.app.planner.logic.PlannerRules
import gg.foreversixty.app.planner.logic.TalentIndex
import gg.foreversixty.app.planner.logic.EditorView
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test

/** The band compare, the codes a build travels as, and the editor's band load and import. */
class PlannerBandTest {
    private val hunter = PlannerPageFixtures.hunter
    private val index = TalentIndex(hunter)

    private fun ready(state: gg.foreversixty.app.planner.ui.EditorPageState): EditorView = (state.phase as EditorPhase.Ready).view

    @Test fun `a band build short of the character's points says what it covers, never Matches`() {
        val band = requireNotNull(ready(PlannerPageFixtures.you).band)
        assertEquals(30, band.build.band)
        assertEquals(BandStatus.Matches(0), band.status)
        assertEquals(21, band.bandPoints)
        assertEquals(25, band.coveredOf)
    }

    @Test fun `a build two points off the band names both sides`() {
        val band = requireNotNull(ready(PlannerPageFixtures.guide).band)
        assertEquals(60, band.build.band)
        assertEquals(BandStatus.Differs(2, 0), band.status)
        assertEquals(1, band.mine.size)
        assertEquals(1, band.theirs.size)
        assertNull(band.coveredOf)
    }

    @Test fun `the band shown is the character's while nothing is spent, else the one the build reaches`() {
        assertEquals(60, PlannerBand.bandFor(0, null))
        assertEquals(30, PlannerBand.bandFor(0, 34))
        assertEquals(30, PlannerBand.bandFor(21, 34))
        assertEquals(60, PlannerBand.bandFor(51, 60))
    }

    @Test fun `a band build loads cut at the level`() {
        val load = PlannerBand.loadAtLevel(index, "5420001505001251-0053542001000000-000000000000000000", 34)
        assertEquals(25, load.order.size)
        assertEquals(51, load.fullPoints)
    }

    @Test fun `the FS1 head reads back as the same ranks`() {
        val order = PlannerPageFixtures.guideOrder
        val code = PlannerCodes.fs1(index, order, BisTestData.BUILD, "night-elf")
        val decoded = requireNotNull(Fs1.decode(code))
        assertEquals("hunter", decoded.classSlug)
        assertEquals("night-elf", decoded.raceSlug)
        assertEquals(PlannerRules.ranks(order), PlannerRules.ranks(PlannerRules.orderFromRanks(index, decoded.treeRanks).order))
    }

    @Test fun `Copy for the game is one tab tier column triple per point`() {
        val first = index.trees[0].talents.first { it.tier == 0 }
        val code = PlannerCodes.fsb1(index, listOf(first.id, first.id), BisTestData.BUILD)
        val triple = "1" + "1" + (first.column + 1).toString()
        assertEquals("FSB1:${BisTestData.BUILD}:hunter:$triple$triple:", code)
    }

    private fun model(request: EditorRequest) = PlannerEditorModel(
        FakePlannerSource(files = mapOf("hunter" to hunter), bisFiles = bisFiles()),
        PlannerDrafts { MemoryValue() },
        request,
        RecordingHaptics(),
        BisTestData.catalog(),
    )

    private fun bisFiles() = listOf("hunter-beast-mastery", "hunter-marksmanship", "hunter-survival").associateWith {
        gg.foreversixty.app.model.ApiJson.decodeFromString(
            gg.foreversixty.app.model.BisFile.serializer(),
            java.io.File(BisTestData.repo, "data/builds/${BisTestData.BUILD}/bis/$it.json").readText(),
        )
    }

    @Test fun `a character's planner with nothing saved opens on its band build`() = runTest {
        val m = model(EditorRequest.Draft("hunter", "dwarf", 34, "Beast Mastery")).apply { load() }
        val view = (m.phase.value as EditorPhase.Ready).view
        assertEquals(21, view.state.spent)
        assertTrue(view.shownNote is EditorNote.Loaded)
    }

    @Test fun `an export of another class reopens the editor on it`() = runTest {
        val m = model(EditorRequest.Draft("hunter", "dwarf")).apply { load() }
        val outcome = m.import("FS1:${BisTestData.BUILD}:warrior:human:000/000/000:")
        assertEquals(ImportOutcome.OtherClass("warrior", "FS1:${BisTestData.BUILD}:warrior:human:000/000/000:"), outcome)
        assertEquals(ImportOutcome.Unreadable, m.import("not a code"))
    }

    @Test fun `an export of this class loads its ranks`() = runTest {
        val m = model(EditorRequest.Draft("hunter", "dwarf")).apply { load() }
        val code = PlannerCodes.fs1(index, PlannerPageFixtures.guideOrder, BisTestData.BUILD, "night-elf")
        assertEquals(ImportOutcome.Loaded, m.import(code))
        val view = (m.phase.value as EditorPhase.Ready).view
        assertEquals(51, view.state.spent)
        assertEquals("night-elf", view.state.raceSlug)
    }
}
