package gg.foreversixty.app.activity

import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.menu.LocalOpenMenu
import gg.foreversixty.app.menu.CharacterBar
import gg.foreversixty.app.menu.LocalPortraitPainter
import gg.foreversixty.app.menu.MenuFixtures
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.model.ActivityRow
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.push.PushGroup
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Assert.assertEquals
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode
import java.time.Instant
import java.time.LocalDate
import java.time.ZoneId

/** Activity's mocked states (design/mocks/ws_app_flows.py `activity`) inside the tab shell. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class ActivityParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun activity() = ParityCheck.state(compose, "activity") { Tab(3) { Inbox(ActivityFixtures.rows) } }

    @Test fun activityAt1_5() = ParityCheck.state(compose, "activity", fontScale = 1.5f) { Tab(3) { Inbox(ActivityFixtures.rows) } }

    @Test fun activityAt2() = ParityCheck.state(compose, "activity", fontScale = 2f) { Tab(3) { Inbox(ActivityFixtures.rows) } }

    @Test fun activityEmpty() = ParityCheck.state(compose, "activity-empty") { Tab(0) { Inbox(emptyList(), ActivityFixtures.noPortrait) } }

    @Test fun activityOut() = ParityCheck.state(compose, "activity-out") { Tab(0) { ActivitySignedOut({ CharacterBar(null, TITLE) }) {} } }

    @Test fun activityOutAt2() = ParityCheck.state(compose, "activity-out", fontScale = 2f) { Tab(0) { ActivitySignedOut({ CharacterBar(null, TITLE) }) {} } }

    @Test fun rowsFallInTheirDays() {
        val sections = Activity.sections(Activity.items(ActivityFixtures.rows), ActivityFixtures.NOW, ActivityFixtures.ZONE)
        assertEquals(listOf(ActivityDay.Today, ActivityDay.Yesterday, ActivityDay.ThisWeek), sections.map { it.day })
        assertEquals(listOf(3, 3, 2), sections.map { it.items.size })
        assertEquals(ActivityDay.Earlier, Activity.dayOf(LocalDate.of(2026, 10, 30), LocalDate.of(2026, 11, 6)))
    }

    @Test fun rowsAboutAClassWearItsCrest() {
        val specs = gg.foreversixty.app.bis.bundledSpecCatalog(androidx.test.core.app.ApplicationProvider.getApplicationContext())
        val crests = ActivityFixtures.rows.associate { it.kind to Activity.crestClass(it, specs) }
        assertEquals("warrior", crests["guildsim.passed"])
        assertEquals("warrior", crests["tier.moved"])
        assertEquals(null, crests["bis.changed"])
        assertEquals(null, crests["loot.awarded"])
        assertEquals(null, crests["report.ready"])
        val hunter = ActivityFixtures.rows.first { it.kind == "tier.moved" }.copy(title = "Beast Mastery Hunter moved to A")
        assertEquals("hunter", Activity.crestClass(hunter, specs))
    }

    @Test fun rowsLandOnTheirAnswer() {
        val (ladder, loot) = ActivityFixtures.rows
        assertEquals(ActivityLanding.Tab(ForeverTab.Guild), Activity.landing(ladder))
        assertEquals(ActivityLanding.Tab(ForeverTab.Gear), Activity.landing(loot))
        val tiers = ActivityFixtures.rows.first { it.kind == "tier.moved" }
        assertEquals(ActivityLanding.Menu(gg.foreversixty.app.menu.MenuDestination.Tiers), Activity.landing(tiers))
        assertEquals(PushGroup.Account, Activity.items(listOf(loot.copy(group = "nonsense", kind = "x.y"))).single().group)
    }

    @Composable
    private fun Tab(unread: Int, content: @Composable () -> Unit) {
        ForeverTabScaffold(selected = ForeverTab.Activity, onSelect = {}, badges = TabBadges(activity = unread)) {
            // The game icons from the repo's build data, as Home's parity test paints them.
            CompositionLocalProvider(LocalOpenMenu provides {}, LocalItemIconPainter provides HomeFixtures.repoIcons, LocalPortraitPainter provides MenuFixtures.repoPortraits) { content() }
        }
    }

    @Composable
    private fun Inbox(rows: List<ActivityRow>, me: gg.foreversixty.app.model.AccountCharacter = MenuFixtures.thandrel) {
        ActivityScreen(
            ActivityState.Loaded(rows, OLDER), ActivityFilter.All, ActivityFixtures.BUILD, ActivityFixtures.NOW, ActivityFixtures.ZONE, ActivityActions(),
        ) { CharacterBar(me, TITLE) }
    }

    private companion object {
        const val TITLE = "Activity"

        /** A page further back, so the list ends in Show older as the mock does. */
        const val OLDER = 1L
    }
}

/** The mock's rows (ws_app_flows.py ACTIVITY) at the parity clock, in Chicago. */
object ActivityFixtures {
    val NOW: Instant = Instant.parse("2026-11-06T18:00:00Z")
    val ZONE: ZoneId = ZoneId.of("America/Chicago")
    const val BUILD = "1.60.1.70790"

    /** Thandrel before his Battle.net portrait arrives: the bar draws his crest. */
    val noPortrait = MenuFixtures.thandrel.copy(avatarUrl = null)

    @Suppress("LongParameterList")
    private fun row(id: Long, kind: String, group: String, title: String, body: String, at: String, read: Boolean, icon: String, figure: Double? = null) =
        ActivityRow(id = id, kind = kind, group = group, title = title, body = body, link = "", occurredAt = at, read = read, icon = icon, figure = figure)

    val rows = listOf(
        row(8, "guildsim.passed", "ladder", "Kaelra passed you in the guild sims", "You are #7 of 14", "2026-11-06T10:00:00Z", false, "ability_warrior_innerrage", 1196.2),
        row(7, "loot.awarded", "gear", "You got Deathbringer", "Onyxia · main hand, +12.6 DPS", "2026-11-06T09:00:00Z", false, "inv_axe_09"),
        row(6, "report.ready", "logs", "Onyxia's Lair report ready", "87th percentile, your best on Onyxia so far · the coach has 2 notes", "2026-11-06T09:00:00Z", false, "inv_misc_head_dragon_black"),
        row(5, "raid.signup_nudge", "raids", "You're In for Onyxia", "As Thandrel, Damage · 8:00 pm CST", "2026-11-06T00:02:00Z", true, "inv_misc_head_dragon_black"),
        row(4, "sim.finished", "gear", "Sim done: 949 DPS ±6", "+21 vs your last", "2026-11-05T20:40:00Z", true, "inv_misc_pocketwatch_01"),
        row(3, "raid.short", "raids", "Healers short for Onyxia tonight", "3 of 8 missing", "2026-11-05T19:15:00Z", true, "inv_misc_head_dragon_black"),
        row(2, "tier.moved", "tiers", "Fury Warrior moved to S", "Raid tiers, 60", "2026-11-03T18:00:00Z", true, "inv_misc_book_09"),
        row(1, "bis.changed", "gear", "Fury Warrior BiS changed", "A new pick for Hands: Theramore Gauntlets", "2026-11-03T17:00:00Z", true, "inv_misc_book_09"),
    )
}
