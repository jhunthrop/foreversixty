package gg.foreversixty.app.system

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.BuildTalents
import gg.foreversixty.app.model.CharacterBuild
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.planner.PlannerTestData
import gg.foreversixty.app.planner.logic.SharedBuildView
import gg.foreversixty.app.planner.logic.SharedCompare
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.rankings.RankingsTestData
import gg.foreversixty.app.rankings.logic.CharacterCard
import gg.foreversixty.app.talents.CompareAgainst
import gg.foreversixty.app.talents.TalentBuildKind
import gg.foreversixty.app.talents.TalentTestData
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** Round 2's planner follow-ups: F11.4's `Open in planner` and F13.5's `Compare with mine`. */
@RunWith(AndroidJUnit4::class)
class PlannerFollowUpsTest {
    private val res = resources()

    // Open in planner (F11.4, Builds seen)

    @Test fun `builds seen open the planner on the class, with no code and no race`() {
        val voltaris = RankingsTestData.voltaris()
        assertTrue(voltaris.buildsSeen.isNotEmpty())
        assertEquals(DeepLink.Planner("shaman"), CharacterCard.plannerLink(voltaris))
        assertEquals("https://foreversixty.gg/app/planner/shaman", CharacterCard.plannerLink(voltaris)?.url())
    }

    @Test fun `no builds seen or no class, no planner row`() {
        val voltaris = RankingsTestData.voltaris()
        assertNull(CharacterCard.plannerLink(voltaris.copy(buildsSeen = emptyList())))
        assertNull(CharacterCard.plannerLink(voltaris.copy(character = voltaris.character.copy(className = null))))
    }

    // Compare with mine (F13.5)

    private fun character(key: String, name: String, className: String, trees: String?) = AccountCharacter(
        key = key, region = "us", ruleset = "normal", name = name, source = "export", className = className, level = 60,
        build = trees?.let {
            CharacterBuild(source = "export", capturedAt = "2026-11-04T19:00:00Z", talents = BuildTalents(it.split("-"), listOf(1, 1, 1)))
        },
    )

    private val lightningbolltz = character("us/normal/lightningbolltz", "Lightningbolltz", "shaman", TalentTestData.SHAMAN_YOURS)
    private val obnoxiousyell = character("us/normal/obnoxiousyell", "Obnoxiousyell", "warrior", TalentTestData.BAND_60)
    private val noExport = character("us/normal/quietone", "Quietone", "shaman", null)

    @Test fun `the button shows with an export of the build's class, else not`() {
        assertEquals(lightningbolltz, SharedCompare.exportFor("shaman", lightningbolltz, listOf(lightningbolltz, obnoxiousyell)))
        // The current character is a warrior: the account's shaman with an export is compared.
        assertEquals(lightningbolltz, SharedCompare.exportFor("shaman", obnoxiousyell, listOf(obnoxiousyell, lightningbolltz)))
        assertNull(SharedCompare.exportFor("shaman", noExport, listOf(noExport, obnoxiousyell)))
        assertNull(SharedCompare.exportFor("mage", lightningbolltz, listOf(lightningbolltz, obnoxiousyell)))
        assertNull(SharedCompare.exportFor("shaman", null, emptyList()))
    }

    private fun compare() = SharedCompare.presenter(
        SharedBuildView(PlannerTestData.sharedBuild(), PlannerTestData.shaman(), "shaman", "Shaman", "Orc"),
        lightningbolltz,
    )

    @Test fun `the compare speaks of the shared build, never the guide`() {
        val presenter = compare()
        assertEquals(CompareAgainst.Shared("Lightningbolltz"), presenter.against)
        val status = presenter.status(res)
        assertTrue(status.sentence, status.sentence.endsWith("differ from the shared build."))
        assertTrue(presenter.legend(TalentBuildKind.Yours, res).startsWith("Dashed: your rank differs from the shared build."))
        assertTrue(presenter.notes(res).none { "guide" in it })
        assertTrue(presenter.showsSegments)
    }

    @Test fun `the guide compare keeps its own words`() {
        assertEquals(gg.foreversixty.app.R.string.talents_guide, CompareAgainst.Guide.side)
        assertEquals(gg.foreversixty.app.R.string.plan_cmp_side, CompareAgainst.Shared("x").side)
        assertEquals("Shared build", res.getString(CompareAgainst.Shared("x").side))
        assertEquals("Guide build", res.getString(CompareAgainst.Guide.side))
    }

    @Test fun `a differing talent's sheet line and label name the shared build`() {
        val presenter = compare()
        val tab = presenter.tabs(TalentBuildKind.Yours).first { it.differing > 0 }
        val cell = presenter.cells(TalentBuildKind.Yours, tab.index).first { it.delta != null }
        assertTrue(presenter.sheet(TalentBuildKind.Yours, cell.talent.id, res)?.guideLine.orEmpty().startsWith("Shared build: rank"))
    }
}
