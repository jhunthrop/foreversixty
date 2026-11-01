package gg.foreversixty.app.planner

import gg.foreversixty.app.auth.session.MemoryValue
import gg.foreversixty.app.auth.session.StoredValue
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.planner.logic.EditOutcome
import gg.foreversixty.app.planner.logic.EditorPhase
import gg.foreversixty.app.planner.logic.EditorRequest
import gg.foreversixty.app.planner.logic.EditorState
import gg.foreversixty.app.planner.logic.Fs1
import gg.foreversixty.app.planner.logic.PlannerCells
import gg.foreversixty.app.planner.logic.PlannerDraft
import gg.foreversixty.app.planner.logic.PlannerDrafts
import gg.foreversixty.app.planner.logic.PlannerEditor
import gg.foreversixty.app.planner.logic.PlannerEditorModel
import gg.foreversixty.app.planner.logic.PlannerKeys
import gg.foreversixty.app.planner.logic.PlannerRules
import gg.foreversixty.app.planner.logic.PlannerShare
import gg.foreversixty.app.planner.logic.Refusal
import gg.foreversixty.app.planner.logic.ShareOutcome
import gg.foreversixty.app.planner.logic.TalentIndex
import gg.foreversixty.app.planner.logic.TreeDirection
import gg.foreversixty.app.talents.TalentCellState
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test

/** The editor's state, haptics, draft and share (spec F13.1 to F13.3), on the JVM. */
class PlannerEditorTest {
    private val index = TalentIndex(PlannerTestData.shaman())
    private val trees = index.trees
    private fun id(name: String) = trees.flatMap { it.talents }.first { it.name == name }.id
    private fun empty(level: Int = 60) = EditorState("shaman", "orc", level, emptyList())

    /** Every tier-0 point of Elemental, repeated: enough legal adds to walk the stack. */
    private fun fill(state: EditorState, points: Int): EditorState {
        var s = state
        val order = PlannerTestData.elementalOrder()
        order.take(points).forEach { s = PlannerEditor.add(index, s, it).state }
        return s
    }

    @Test fun `a tap adds one point and touches the talent`() {
        val edit = PlannerEditor.add(index, empty(), id("Convection"))
        assertEquals(EditOutcome.Added, edit.outcome)
        assertEquals(listOf(id("Convection")), edit.state.order)
        assertEquals(id("Convection"), edit.state.touched)
        assertEquals(50, edit.state.left)
    }

    @Test fun `a refusal changes nothing but the strip`() {
        val start = empty()
        val edit = PlannerEditor.add(index, start, id("Elemental Mastery"))
        assertEquals(EditOutcome.Refused, edit.outcome)
        assertEquals(start.order, edit.state.order)
        assertEquals(start.undo, edit.state.undo)
        assertTrue(edit.state.refusal is Refusal.TierLocked)
    }

    @Test fun `the level budget refuses before the rules`() {
        val start = empty(level = 10)
        val one = PlannerEditor.add(index, start, id("Convection")).state
        val edit = PlannerEditor.add(index, one, id("Convection"))
        assertEquals(Refusal.NoneLeft, edit.state.refusal)
        assertEquals(0, one.left)
    }

    @Test fun `a long press removes the last point`() {
        val two = fill(empty(), 2)
        val edit = PlannerEditor.remove(index, two, two.order.last())
        assertEquals(EditOutcome.Removed, edit.outcome)
        assertEquals(two.order.dropLast(1), edit.state.order)
    }

    @Test fun `a removal that breaks a talent below is refused`() {
        val full = fill(empty(), PlannerTestData.elementalOrder().size)
        val first = full.order.first()
        val edit = PlannerEditor.remove(index, full, first)
        assertEquals(EditOutcome.Refused, edit.outcome)
        assertEquals(full.order, edit.state.order)
    }

    @Test fun `undo steps back one point at a time`() {
        val three = fill(empty(), 3)
        val back = PlannerEditor.undo(three)
        assertEquals(EditOutcome.Undone, back.outcome)
        assertEquals(three.order.dropLast(1), back.state.order)
    }

    @Test fun `the undo stack never grows past 100`() {
        var s = empty()
        repeat(160) { i ->
            s = if (i % 2 == 0) PlannerEditor.add(index, s, id("Convection")).state else PlannerEditor.remove(index, s, id("Convection")).state
        }
        assertEquals(PlannerEditor.UNDO_LIMIT, s.undo.size)
    }

    @Test fun `reset clears every point as one undoable step`() {
        val five = fill(empty(), 5)
        val cleared = PlannerEditor.reset(five)
        assertEquals(EditOutcome.Cleared, cleared.outcome)
        assertTrue(cleared.state.order.isEmpty())
        assertEquals(five.order, PlannerEditor.undo(cleared.state).state.order)
    }

    @Test fun `the level never drops below the spent points' level`() {
        val ten = fill(empty(), 10)
        assertEquals(19, PlannerEditor.setLevel(ten, 12).level)
        assertEquals(PlannerEditor.MIN_LEVEL, PlannerEditor.setLevel(empty(), 1).level)
        assertEquals(PlannerEditor.MAX_LEVEL, PlannerEditor.setLevel(empty(), 70).level)
    }

    @Test fun `the strip shows the shown tree's first talent before any touch`() {
        val strip = PlannerCells.strip(index, empty(), 0)
        assertEquals(PlannerCells.readingOrder(trees[0]).first().id, strip?.talent?.id)
        assertEquals(TalentCellState.Available, strip?.state)
    }

    @Test fun `arrows move in reading order and between tiers`() {
        val tree = trees[0]
        val first = PlannerCells.readingOrder(tree).first()
        assertEquals(first.id, PlannerKeys.move(tree, null, TreeDirection.Right))
        val second = PlannerKeys.move(tree, first.id, TreeDirection.Right)
        assertEquals(PlannerCells.readingOrder(tree)[1].id, second)
        val down = PlannerKeys.move(tree, first.id, TreeDirection.Down)
        assertEquals(1, tree.talents.first { it.id == down }.tier)
        assertEquals(first.id, PlannerKeys.move(tree, first.id, TreeDirection.Up))
    }

    // The model: haptics, the draft and the share

    private val stores = mutableMapOf<String, StoredValue>()
    private val drafts = PlannerDrafts { slug -> stores.getOrPut(slug) { MemoryValue() } }

    private fun model(source: FakePlannerSource = FakePlannerSource(), request: EditorRequest = EditorRequest.Draft("shaman", "orc", 60), haptics: RecordingHaptics = RecordingHaptics()) =
        PlannerEditorModel(source, drafts, request, haptics)

    private val PlannerEditorModel.view get() = (phase.value as EditorPhase.Ready).view

    @Test fun `each edit plays its haptic`() = runTest {
        val haptics = RecordingHaptics()
        val m = model(haptics = haptics).apply { load() }
        m.add(id("Convection"))
        m.remove(id("Convection"))
        m.add(id("Elemental Mastery"))
        m.undo()
        assertEquals(listOf("added", "removed", "refused"), haptics.played)
    }

    @Test fun `the draft is saved as you go and reopens after the process dies`() = runTest {
        val m = model().apply { load() }
        m.add(id("Convection"))
        m.add(id("Convection"))
        m.setLevel(40)
        val reopened = model(request = EditorRequest.Draft("shaman", "tauren", 20)).apply { load() }
        assertEquals(listOf(id("Convection"), id("Convection")), reopened.view.state.order)
        assertEquals(40, reopened.view.state.level)
        assertEquals("orc", reopened.view.state.raceSlug)
    }

    @Test fun `a new draft opens at the character's race and level`() = runTest {
        val m = model(request = EditorRequest.Draft("shaman", "tauren", 32)).apply { load() }
        assertEquals("tauren", m.view.state.raceSlug)
        assertEquals(32, m.view.state.level)
        assertEquals(23, m.view.state.left)
    }

    @Test fun `only the class's legal races are offered`() = runTest {
        val m = model().apply { load() }
        val slugs = m.view.races.map { it.slug }
        assertTrue("orc" in slugs)
        assertTrue(m.view.races.none { it.placeholder })
        assertEquals(PlannerTestData.shaman().classId, PlannerTestData.classes().first { it.slug == "shaman" }.id)
    }

    @Test fun `a guide's FS1 code loads as one legal order`() = runTest {
        val code = "FS1:${gg.foreversixty.app.bis.BisTestData.BUILD}:shaman:orc:5510331300001350/03/0533520000000000:"
        val m = model(request = EditorRequest.Code("shaman", code)).apply { load() }
        val ranks = requireNotNull(Fs1.decode(code)).treeRanks
        assertEquals(PlannerRules.orderFromRanks(index, ranks).order, m.view.state.order)
        assertTrue(PlannerRules.validate(index, m.view.state.order).isEmpty())
        assertEquals(listOf(30, 3, 18), m.view.totals)
    }

    @Test fun `edit a copy loads the shared points and leaves the record alone`() = runTest {
        val shared = PlannerTestData.sharedBuild()
        val source = FakePlannerSource()
        val m = model(source, EditorRequest.Copy("shaman", shared)).apply { load() }
        assertEquals(shared.pointOrder, m.view.state.order)
        assertEquals("orc", m.view.state.raceSlug)
        m.remove(m.view.state.order.last())
        assertTrue(source.saved.isEmpty())
        assertEquals(PlannerTestData.sharedBuild().pointOrder, shared.pointOrder)
    }

    @Test fun `share posts the build with empty gear and plays confirm`() = runTest {
        val source = FakePlannerSource()
        val haptics = RecordingHaptics()
        val m = model(source, haptics = haptics).apply { load() }
        m.add(id("Convection"))
        val outcome = m.share("  Elemental raid  ")
        assertTrue(outcome is ShareOutcome.Shared)
        val body = source.saved.single()
        assertEquals(PlannerTestData.shaman().classId, body.classId)
        assertEquals(PlannerTestData.ORC, body.raceId)
        assertEquals(gg.foreversixty.app.bis.BisTestData.BUILD, body.treeVersion)
        assertEquals(listOf(id("Convection")), body.pointOrder)
        assertEquals(emptyMap<String, Int>(), body.gear)
        assertEquals("Elemental raid", body.title)
        assertEquals("shared", haptics.played.last())
    }

    @Test fun `an empty title is left out`() {
        assertNull(PlannerShare.input(1, 2, "b", emptyList(), "   ").title)
        assertEquals(PlannerShare.TITLE_MAX, PlannerShare.cleanTitle("x".repeat(80))?.length)
    }

    @Test fun `share failures say why and are never queued`() = runTest {
        val source = FakePlannerSource(saveAnswer = FakePlannerSource.offline)
        val m = model(source).apply { load() }
        assertEquals(ShareOutcome.Offline, m.share(""))
        assertEquals(ShareOutcome.Offline, m.view.shareProblem)
        source.saveAnswer = ApiResult.Failure(ApiError.Server(429, "rate_limited", "slow down", emptyMap(), "r"))
        assertEquals(ShareOutcome.RateLimited, m.share(""))
        source.saveAnswer = ApiResult.Failure(
            ApiError.Server(400, "invalid", "bad", mapOf("point_order[10]" to "later", "point_order[2]" to "Talent 999 is not a Shaman talent"), "r"),
        )
        assertEquals(ShareOutcome.Field("Talent 999 is not a Shaman talent"), m.share(""))
        m.add(id("Convection"))
        assertNull(m.view.shareProblem)
        assertEquals(3, source.saved.size)
    }

    @Test fun `a bad draft reads as none`() {
        stores["shaman"] = MemoryValue("{not json")
        assertNull(drafts.load("shaman"))
        drafts.save(PlannerDraft("shaman", "orc", 60, listOf(1)))
        assertEquals(listOf(1), drafts.load("shaman")?.order)
        assertFalse(drafts.load("warrior") != null)
    }
}
