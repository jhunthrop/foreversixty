package gg.foreversixty.app.raids.flows

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.RaidCompTarget
import gg.foreversixty.app.model.RaidNotSignedRaider
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.FakeRaidService
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.resources
import kotlin.time.Duration.Companion.seconds
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.test.advanceTimeBy
import kotlinx.coroutines.test.runCurrent
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Test
import org.junit.runner.RunWith

/** Not answered's words from the officer fixture: six healers of eight, Sennah (a healer) unanswered, Varkas benched. */
@OptIn(ExperimentalCoroutinesApi::class)
@RunWith(AndroidJUnit4::class)
class NotAnsweredTest {
    private fun raider(key: String, name: String, cls: String, role: String, lastWeek: String?) =
        RaidNotSignedRaider(key, name, cls, "u:$name", discordLinked = true, role = role, lastWeek = lastWeek)

    private val raid: RaidView = RaidFixture.raid("guild.raid.officer").let { base ->
        base.copy(
            compTarget = RaidCompTarget(tanks = 3, healers = 8, melee = 6, ranged = 8, source = "guild"),
            notSignedRaiders = listOf(
                raider("us/pvp/zulmara", "Zulmara", "hunter", "dps", "missed"),
                raider("us/pvp/brokk", "Brokk", "warrior", "dps", null),
                raider("us/pvp/sennah", "Sennah", "priest", "healer", "in"),
            ),
        )
    }

    private fun content(marks: Map<String, SignupState> = emptyMap(), filter: NotAnsweredFilter = NotAnsweredFilter.NotAnswered) =
        NotAnsweredPresenter(resources(), clock(), raid, "Olympus XXVII", marks).content(filter)

    @Test
    fun `the short role leads the rows, the head and the mention`() {
        val c = content()
        assertEquals(listOf("Sennah", "Zulmara", "Brokk"), c.rows.map { it.name })
        assertEquals("Holy Priest · was In last week", c.rows[0].sub)
        assertEquals("Marksmanship Hunter · missed last week", c.rows[1].sub)
        assertEquals("Arms Warrior · new this week", c.rows[2].sub)
        assertEquals("healers short by 2", c.facts.last())
        assertEquals("Mention the 1 healer in Discord", c.mention?.role)
        assertEquals("Mention all 3", c.mention?.all)
        assertEquals("3 of 31, healers first", c.listMeta)
        assertEquals("Not answered 3", c.chips.first().title)
        assertEquals("All 31", c.chips.last().title)
    }

    @Test
    fun `a mark In moves the counts and closes the role's gap`() {
        val c = content(mapOf("us/pvp/sennah" to SignupState.In))
        assertEquals(SignupState.In, c.rows.first { it.name == "Sennah" }.marked)
        assertEquals("In 24", c.chips[1].title)
        assertEquals("healers short by 1", c.facts.last())
    }

    @Test
    fun `the other filters list the sign-ups without actions`() {
        val bench = content(filter = NotAnsweredFilter.Bench)
        assertEquals(listOf("Varkas"), bench.rows.map { it.name })
        assertEquals(false, bench.actionable)
        assertNull(bench.mention)
        assertEquals("Bench", bench.listTitle)
    }

    @Test
    fun `a mark is sent once the hold runs out, and Undo keeps it from being sent`() = runTest {
        val service = FakeRaidService().apply { raidAnswer = { gg.foreversixty.app.raids.ok(raid) } }
        val model = NotAnsweredModel(2, raid.id, service, NoFlows, resources(), backgroundScope, HOLD)
        model.load()
        val row = content().rows.first()
        model.mark(row, SignupState.In)
        advanceTimeBy(HOLD / 2)
        model.undo(row)
        advanceTimeBy(HOLD)
        runCurrent()
        assertEquals(emptyList<Pair<String, Any>>(), service.signups.toList())
        model.mark(row, SignupState.Bench)
        advanceTimeBy(HOLD + 1.seconds)
        runCurrent()
        assertEquals("us/pvp/sennah" to "bench", service.signups.single().let { it.first to it.second.state })
        assertEquals("healer", service.signups.single().second.role)
    }

    private companion object {
        val HOLD = 4.seconds
    }

    /** Not answered sends no ping in these tests. */
    private object NoFlows : RaidFlowService {
        override suspend fun applications(guild: Long) = error("unused")
        override suspend fun accept(guild: Long, application: Long) = error("unused")
        override suspend fun decline(guild: Long, application: Long) = error("unused")
        override suspend fun recruiting(spec: String, faction: String, region: String?, ruleset: String?) = error("unused")
        override suspend fun ping(guild: Long, raid: Long, role: String?) = error("unused")
        override suspend fun ready(characterKey: String, raid: Long) = error("unused")
        override suspend fun setup(guild: Long) = error("unused")
        override suspend fun myApplications(guild: Long) = error("unused")
        override suspend fun lootNights(guild: Long) = error("unused")
    }
}
