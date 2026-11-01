package gg.foreversixty.app.system

import gg.foreversixty.app.model.Widget
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.WatchingList
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.raids.FakeRaidService
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RaidTime
import gg.foreversixty.app.raids.fail
import gg.foreversixty.app.raids.ok
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.system.logic.AppShortcut
import gg.foreversixty.app.system.logic.DynamicShortcuts
import gg.foreversixty.app.system.logic.IntentRaid
import gg.foreversixty.app.system.logic.IntentTime
import gg.foreversixty.app.system.logic.RaidPick
import gg.foreversixty.app.system.logic.ShortcutSignup
import gg.foreversixty.app.system.logic.SignupShortcutAnswer
import gg.foreversixty.app.system.logic.SignupShortcutCopy
import gg.foreversixty.app.system.logic.SignupShortcutResult
import gg.foreversixty.app.network.ApiError
import java.io.IOException
import java.time.Instant
import java.time.ZoneId
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/**
 * Lane F15's shortcut rules (spec F15.1 to F15.3): which raid, `Sign up In` end to end over a
 * scripted raid service, its toast words, the dynamic shortcuts and the Assistant's features.
 */
@RunWith(AndroidJUnit4::class)
class ShortcutLogicTest {
    private val planned = RaidFixture.load(GuildRaids.serializer(), "guild.raids").planned.orEmpty()
    private val onyxia = planned.first { it.id == ONYXIA }
    private val chicago: ZoneId = RaidTime.chicago

    private fun at(raid: RaidView, hoursBefore: Long): Instant = Instant.parse(raid.startsAt).minusSeconds(hoursBefore * HOUR)

    // Which raid

    @Test fun `the raid starting within 24 hours is the one`() {
        assertEquals(RaidPick.One(onyxia), IntentRaid.pick(planned, at(onyxia, 3)))
    }

    @Test fun `with none within 24 hours, the next planned one`() {
        assertEquals(RaidPick.One(onyxia), IntentRaid.pick(planned, at(onyxia, 50)))
    }

    @Test fun `two within 24 hours ask which`() {
        val soonAfter = onyxia.copy(id = 99, startsAt = Instant.parse(onyxia.startsAt).plusSeconds(2 * HOUR).toString())
        val pick = IntentRaid.pick(planned + soonAfter, at(onyxia, 1))
        assertEquals(RaidPick.Which(listOf(onyxia, soonAfter)), pick)
    }

    @Test fun `cancelled, done and started raids are never picked`() {
        val cancelled = onyxia.copy(status = "cancelled")
        assertEquals(RaidPick.None, IntentRaid.pick(listOf(cancelled), at(onyxia, 3)))
        assertEquals(RaidPick.None, IntentRaid.pick(listOf(onyxia), Instant.parse(onyxia.startsAt).plusSeconds(60)))
    }

    @Test fun `times read in the phone's zone`() {
        val starts = Instant.parse("2026-11-06T02:00:00Z")
        assertEquals("Thu", IntentTime.weekdayShort(starts, chicago))
        assertEquals("Thursday", IntentTime.weekday(starts, chicago))
        assertEquals("8 pm", IntentTime.time(starts, chicago))
        assertEquals("8:30 pm", IntentTime.time(starts.plusSeconds(HALF_HOUR), chicago))
    }

    // Sign up In

    private fun runner(service: FakeRaidService, signedIn: Boolean = true, on: Boolean = true, now: Instant = at(onyxia, 3)) =
        ShortcutSignup(service, { signedIn }, { null }, { on }, { now })

    @Test fun `sign up In puts In for the next raid, keeping the note`() = runTest {
        val service = FakeRaidService()
        assertEquals(SignupShortcutResult.Signed("Onyxia's Lair"), runner(service).run())
        val (key, request) = service.signups.single()
        assertEquals("us/pvp/obnoxious-yell", key)
        assertEquals("in", request.state)
        assertEquals(onyxia.signups.firstOrNull { it.characterKey == key }?.note.orEmpty(), request.note)
    }

    @Test fun `a locked raid, a failure and no connection say so`() = runTest {
        val locked = FakeRaidService().apply { signupAnswer = { fail(RaidFixture.error("guild.signup.error-raid-locked", 409)) } }
        assertEquals(SignupShortcutResult.Locked, runner(locked).run())
        val broken = FakeRaidService().apply { signupAnswer = { fail(ApiError.Status(500)) } }
        assertEquals(SignupShortcutResult.Failed, runner(broken).run())
        val offline = FakeRaidService().apply { raidsAnswer = { fail(ApiError.Network(IOException("offline"))) } }
        assertEquals(SignupShortcutResult.Offline, runner(offline).run())
    }

    @Test fun `no raid, signed out and the switch off never sign up`() = runTest {
        val empty = FakeRaidService().apply { raidsAnswer = { ok(GuildRaids(planned = emptyList())) } }
        assertEquals(SignupShortcutResult.NoRaid, runner(empty).run())
        assertEquals(SignupShortcutResult.SignedOut, runner(FakeRaidService(), signedIn = false).run())
        val off = FakeRaidService()
        assertEquals(SignupShortcutResult.Off, runner(off, on = false).run())
        assertTrue(empty.signups.isEmpty() && off.signups.isEmpty())
    }

    @Test fun `no spec saved opens the spec ask instead`() = runTest {
        val unsure = onyxia.copy(viewer = onyxia.viewer.copy(characters = onyxia.viewer.characters.map { it.copy(raidSpec = null) }))
        val service = FakeRaidService().apply { raidsAnswer = { ok(GuildRaids(planned = listOf(unsure))) } }
        assertEquals(SignupShortcutResult.AskSpec(GUILD, ONYXIA), runner(service).run())
        assertTrue(service.signups.isEmpty())
    }

    @Test fun `a character on no guild cannot sign up`() = runTest {
        val me = RaidFixture.load(Me.serializer(), "me")
        val guildless = me.copy(characters = me.characters.map { it.copy(guild = null) })
        val service = FakeRaidService().apply { meAnswer = { ok(guildless) } }
        assertTrue(runner(service).run() is SignupShortcutResult.NoGuild)
    }

    @Test fun `the toasts use the app's sentences`() {
        val res = resources()
        assertEquals("You are In for Onyxia's Lair.", SignupShortcutCopy.answer(SignupShortcutResult.Signed("Onyxia's Lair"), res).toast)
        assertEquals("Sign-ups are locked. An officer can still change yours.", SignupShortcutCopy.answer(SignupShortcutResult.Locked, res).toast)
        assertEquals("That did not save. Try again.", SignupShortcutCopy.answer(SignupShortcutResult.Failed, res).toast)
        assertEquals("No raid planned.", SignupShortcutCopy.answer(SignupShortcutResult.NoRaid, res).toast)
        assertEquals("Sign in to Forever Sixty first.", SignupShortcutCopy.answer(SignupShortcutResult.SignedOut, res).toast)
        assertEquals(
            SignupShortcutAnswer(null, DeepLink.Raid(GUILD, ONYXIA, spec = "ask")),
            SignupShortcutCopy.answer(SignupShortcutResult.AskSpec(GUILD, ONYXIA), res),
        )
        assertTrue(SignupShortcutCopy.answer(SignupShortcutResult.Off, res).openApp)
        assertNull(SignupShortcutCopy.answer(SignupShortcutResult.Off, res).toast)
    }

    // Dynamic shortcuts and the Assistant

    @Test fun `dynamic shortcuts are the character, two watched and the next raid`() {
        val me = RaidFixture.load(Me.serializer(), "me")
        val current = me.characters.first()
        val watched = RaidFixture.load(WatchingList.serializer(), "watching.list").rows
        val next = RaidFixture.load(Widget.serializer(), "widget.get").nextRaid
        val list = DynamicShortcuts.list(current, watched, next, resources(), chicago)
        assertTrue(list.size <= DynamicShortcuts.MAX)
        assertEquals("character:${current.key}", list.first().id)
        assertEquals("${current.name} · Forever Sixty · character", list.first().longLabel)
        assertEquals(current.className, list.first().crestClass)
        // Rank 0 is the current character, rank 1 the next raid, then the watched characters.
        val raid = list[1]
        assertEquals("Next raid", raid.shortLabel)
        // The widget's next raid: the story's Thursday 8 pm in the phone's zone.
        assertEquals("Onyxia's Lair · Thu 8 pm", raid.longLabel)
        assertEquals(DeepLink.Raid(GUILD, WIDGET_RAID), raid.link)
        assertEquals(minOf(watched.count { it.key != current.key }, 2), list.count { it.id.startsWith("character:") } - 1)
    }

    @Test fun `signed out there are no character shortcuts`() {
        assertTrue(DynamicShortcuts.list(null, emptyList(), null, resources(), chicago).isEmpty())
    }

    @Test fun `the Assistant's features map to the app's screens`() {
        assertEquals(AppShortcut.NextRaid, AppShortcut.ofFeature("next raid"))
        assertEquals(AppShortcut.Upgrades, AppShortcut.ofFeature("My upgrades"))
        assertEquals(AppShortcut.Bis, AppShortcut.ofFeature("best in slot"))
        assertEquals(AppShortcut.Guides, AppShortcut.ofFeature("guides"))
        assertEquals(AppShortcut.Search, AppShortcut.ofFeature("search"))
        assertNull(AppShortcut.ofFeature("sign me up"))
        assertEquals(ForeverTab.Raids, SystemNavigator.tabOf(AppShortcut.NextRaid))
        assertEquals(ForeverTab.Home, SystemNavigator.tabOf(AppShortcut.Upgrades))
        assertNull(SystemNavigator.tabOf(AppShortcut.Search))
    }

    private companion object {
        const val GUILD = 2L
        const val ONYXIA = 31L
        const val WIDGET_RAID = 7L
        const val HOUR = 3_600L
        const val HALF_HOUR = 1_800L
    }
}
