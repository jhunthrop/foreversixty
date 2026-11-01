package gg.foreversixty.app.planner

import gg.foreversixty.app.model.BuildList
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.planner.logic.BuildsProblem
import gg.foreversixty.app.planner.logic.EditorRequest
import gg.foreversixty.app.planner.logic.PlannerNavigator
import gg.foreversixty.app.planner.logic.PlannerScreen
import gg.foreversixty.app.planner.logic.SharedBuildModel
import gg.foreversixty.app.planner.logic.SharedBuildPhase
import gg.foreversixty.app.planner.logic.YourBuildsModel
import gg.foreversixty.app.planner.logic.YourBuildsPhase
import gg.foreversixty.app.network.CachedFile
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test

/** Your builds' remove and Undo, the viewer's states and the planner links (spec F13.5, F13.6, §3.4). */
class PlannerBuildsTest {
    private val YourBuildsModel.loaded get() = phase.value as YourBuildsPhase.Loaded

    @Test fun `rows carry names and their own trees' totals`() = runTest {
        val m = YourBuildsModel(FakePlannerSource()).apply { load() }
        val row = m.loaded.rows.single()
        assertEquals("Shaman", row.className)
        assertEquals("Orc", row.raceName)
        assertEquals(listOf(30, 3, 18), row.totals)
    }

    @Test fun `a swipe removes the row, calls DELETE and holds it for Undo`() = runTest {
        val source = FakePlannerSource()
        val m = YourBuildsModel(source).apply { load() }
        m.remove("k7q2m9xa")
        assertTrue(m.loaded.rows.isEmpty())
        assertEquals(listOf("k7q2m9xa"), source.removed)
        assertEquals("k7q2m9xa", m.loaded.pending?.row?.build?.id)
    }

    @Test fun `Undo saves the same build again and puts the row back`() = runTest {
        val source = FakePlannerSource()
        val m = YourBuildsModel(source).apply { load() }
        val build = m.loaded.rows.single().build
        m.remove(build.id)
        m.undo()
        val again = source.saved.single()
        assertEquals(build.classId, again.classId)
        assertEquals(build.raceId, again.raceId)
        assertEquals(build.treeVersion, again.treeVersion)
        assertEquals(build.pointOrder, again.pointOrder)
        assertEquals(build.gear, again.gear)
        assertEquals(build.title, again.title)
        assertEquals(listOf(build.id), m.loaded.rows.map { it.build.id })
        assertNull(m.loaded.pending)
    }

    @Test fun `a failed DELETE puts the row back and says so`() = runTest {
        val source = FakePlannerSource(removeAnswer = FakePlannerSource.offline)
        val m = YourBuildsModel(source).apply { load() }
        m.remove("k7q2m9xa")
        assertEquals(1, m.loaded.rows.size)
        assertEquals(BuildsProblem.RemoveFailed, m.loaded.problem)
    }

    @Test fun `a 404 on DELETE means it is already off the list`() = runTest {
        val source = FakePlannerSource(removeAnswer = ApiResult.Failure(ApiError.Server(404, "not_found", "x", emptyMap(), "r")))
        val m = YourBuildsModel(source).apply { load() }
        m.remove("k7q2m9xa")
        assertTrue(m.loaded.rows.isEmpty())
        m.expire()
        assertNull(m.loaded.pending)
    }

    @Test fun `empty and failed lists`() = runTest {
        val empty = YourBuildsModel(FakePlannerSource(listAnswer = ApiResult.Success(BuildList(emptyList(), 0, 1, 100), "r"))).apply { load() }
        assertTrue((empty.phase.value as YourBuildsPhase.Loaded).rows.isEmpty())
        val failed = YourBuildsModel(FakePlannerSource(listAnswer = FakePlannerSource.offline)).apply { load() }
        assertEquals(YourBuildsPhase.Failed, failed.phase.value)
    }

    @Test fun `the viewer draws the build's own tree version`() = runTest {
        val m = SharedBuildModel(FakePlannerSource(), "k7q2m9xa").apply { load() }
        val view = (m.phase.value as SharedBuildPhase.Loaded).view
        assertEquals("shaman", view.classSlug)
        assertEquals("Orc", view.raceName)
        assertEquals(listOf(30, 3, 18), view.totals)
        assertNull(view.offlineSinceMillis)
    }

    @Test fun `the viewer's unknown id, offline and saved copy`() = runTest {
        val notFound = FakePlannerSource(buildAnswer = ApiResult.Failure(ApiError.Server(404, "not_found", "x", emptyMap(), "r")))
        assertEquals(SharedBuildPhase.NotFound, SharedBuildModel(notFound, "x").apply { load() }.phase.value)
        val offline = FakePlannerSource(buildAnswer = FakePlannerSource.offline)
        assertEquals(SharedBuildPhase.Offline, SharedBuildModel(offline, "x").apply { load() }.phase.value)
        val cached = FakePlannerSource(buildAnswer = ApiResult.Success(CachedFile(PlannerTestData.sharedBuild(), 42, true), "r"))
        assertEquals(42L, (SharedBuildModel(cached, "x").apply { load() }.phase.value as SharedBuildPhase.Loaded).view.offlineSinceMillis)
    }

    @Test fun `planner links open the right screen`() {
        val nav = PlannerNavigator()
        assertTrue(nav.open(DeepLink.Planner("Shaman", code = "FS1:b:shaman:orc:5/0/0:", race = "orc"), signedIn = false))
        assertEquals(PlannerScreen.Editor(EditorRequest.Code("shaman", "FS1:b:shaman:orc:5/0/0:", "orc")), nav.top)
        assertTrue(nav.open(DeepLink.Planner("warrior"), signedIn = false))
        assertEquals(PlannerScreen.Editor(EditorRequest.Draft("warrior")), nav.top)
        assertTrue(nav.open(DeepLink.SharedBuild("k7q2m9xa"), signedIn = false))
        assertEquals(PlannerScreen.Viewer("k7q2m9xa"), nav.top)
        assertFalse(nav.open(DeepLink.Builds, signedIn = false))
        assertTrue(nav.open(DeepLink.Builds, signedIn = true))
        assertEquals(listOf(PlannerScreen.Builds), nav.stack)
        assertFalse(nav.open(DeepLink.Guides, signedIn = true))
        assertEquals(DeepLink.SharedBuild("k7q2m9xa"), DeepLink.parse("https://foreversixty.gg/b/k7q2m9xa"))
    }
}
