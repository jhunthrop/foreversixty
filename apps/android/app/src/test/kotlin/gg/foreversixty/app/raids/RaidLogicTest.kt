package gg.foreversixty.app.raids

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.SignupAnswer
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.raids.logic.CompositionRead
import gg.foreversixty.app.raids.logic.NotesText
import gg.foreversixty.app.raids.logic.PlannedRaids
import gg.foreversixty.app.raids.logic.RaidChip
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidFacts
import gg.foreversixty.app.raids.logic.RaidOpenAction
import gg.foreversixty.app.raids.logic.RaidsTabModel
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.logic.ReminderAskPolicy
import gg.foreversixty.app.raids.logic.RosterCopy
import gg.foreversixty.app.raids.logic.RosterGroups
import gg.foreversixty.app.raids.logic.RosterKind
import gg.foreversixty.app.raids.logic.RowChip
import gg.foreversixty.app.raids.logic.SignupRefusal
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.logic.SpecAsk
import gg.foreversixty.app.raids.logic.StatusLine
import gg.foreversixty.app.raids.push.RaidPushAnswerer
import gg.foreversixty.app.raids.push.RaidPushCategory
import gg.foreversixty.app.raids.push.RaidPushOutcome
import gg.foreversixty.app.raids.push.RaidPushPayload
import gg.foreversixty.app.raids.ui.placeMenuItems
import java.time.Duration
import java.time.Instant
import java.time.ZoneId
import kotlinx.coroutines.test.runTest
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** The F4 test list, one per behaviour; iOS's RaidLogicTests carries the same names. */
@RunWith(AndroidJUnit4::class)
class RaidLogicTest {
    private val key = "us/pvp/obnoxious-yell"

    // Raid times

    @Test
    fun `clocks read in the viewer's zone, and the date block`() {
        val invite = Instant.parse("2026-12-20T00:45:00Z")
        assertEquals("6:45 pm", clock().clock(invite))
        assertEquals(Triple("SAT", "19", "DEC"), clock().dateBlock(invite))
        assertEquals("Saturday 19 December", clock().longDate(invite))
    }

    @Test
    fun `the countdown reads like the site's`() {
        val start = Instant.parse("2026-12-20T00:45:00Z")
        val end = Instant.parse("2026-12-20T04:00:00Z")
        val c = clock()
        assertEquals("in 2 days 2 h", c.countdown(start, end, RaidTime.now))
        assertEquals("starting now", c.countdown(start, end, start.minusSeconds(30)))
        assertEquals("started 20 min ago", c.countdown(start, end, start.plusSeconds(1200)))
        assertEquals("ended 1 h ago", c.countdown(start, end, end.plusSeconds(3600)))
        assertEquals("5 h 20 min", c.span(Duration.ofMinutes(320)))
    }

    @Test
    fun `TalkBack hears whole words, never h`() {
        assertEquals("2 days 3 hours", clock().spanWords(Duration.ofHours(51)))
        assertEquals("1 hour 1 minute", clock().spanWords(Duration.ofMinutes(61)))
    }

    @Test
    fun `a planned row's when - tonight, tomorrow, weekday, then the date`() {
        val now = Instant.parse("2026-12-17T18:00:00Z")
        val c = clock()
        assertEquals("Tonight 8:00 pm", c.whenPhrase(Instant.parse("2026-12-18T02:00:00Z"), now))
        assertEquals("Tomorrow 8:00 pm", c.whenPhrase(Instant.parse("2026-12-19T02:00:00Z"), now))
        assertEquals("Saturday 7:00 pm", c.whenPhrase(Instant.parse("2026-12-20T01:00:00Z"), now))
        assertEquals("Thu 7 Jan, 8:00 pm", c.whenPhrase(Instant.parse("2027-01-08T02:00:00Z"), now))
    }

    @Test
    fun `the guild-time line shows only when the offsets differ`() {
        val start = Instant.parse("2026-12-20T01:00:00Z")
        assertFalse(clock().offsetDiffers(start, RaidTime.chicago))
        val berlin = RaidClock(resources(), ZoneId.of("Europe/Berlin"))
        assertTrue(berlin.offsetDiffers(start, RaidTime.chicago))
        assertTrue(berlin.guildClock(start, RaidTime.chicago).startsWith("7:00 pm"))
    }

    // Raid facts

    @Test
    fun `the count is In plus Late, with the bar only when the size is known`() {
        val facts = RaidFacts(RaidFixture.raid(), RaidTime.now)
        assertEquals("24 of 40 signed up", facts.countLine(resources()))
        assertEquals("24/40", facts.ratio)
        assertEquals("3 tanks · 6 healers · 14 DPS · 1 late", facts.breakdown(resources()))
        val barrow = RaidFixture.load(GuildRaids.serializer(), "guild.raids").planned.orEmpty()[1]
        assertEquals("11 signed up", RaidFacts(barrow, RaidTime.now).countLine(resources()))
        assertNull(RaidFacts(barrow, RaidTime.now).fill)
    }

    @Test
    fun `one tank reads 1 tank, no late clause when nobody is late`() {
        val ony = RaidFixture.load(GuildRaids.serializer(), "guild.raids").planned.orEmpty()[2]
        assertEquals("1 tank · 2 healers · 5 DPS", RaidFacts(ony, RaidTime.now).breakdown(resources()))
    }

    @Test
    fun `a raid past its end reads Ended and is no longer signable`() {
        val facts = RaidFacts(RaidFixture.raid(), Instant.parse("2026-12-21T00:00:00Z"))
        assertEquals(RaidChip.Ended, facts.chip)
        assertFalse(facts.isSignable)
    }

    @Test
    fun `the countdown runs to invites, else to the pull`() {
        assertEquals(Instant.parse("2026-12-20T00:45:00Z"), RaidFacts(RaidFixture.raid(), RaidTime.now).countdownTarget)
        assertEquals(Instant.parse("2026-12-20T01:00:00Z"), RaidFacts(RaidFixture.raid("guild.raid.no-invite"), RaidTime.now).countdownTarget)
    }

    @Test
    fun `officer tools come from viewer can only`() {
        assertFalse(RaidFacts(RaidFixture.raid(), RaidTime.now).showsOfficerTools)
        val officer = RaidFacts(RaidFixture.raid("guild.raid.officer"), RaidTime.now)
        assertTrue(officer.canLock && officer.canMove && officer.showsOfficerTools)
    }

    @Test
    fun `a planned row's chip - Cancelled, Locked, Ended, else the viewer's answer, else Not signed up`() {
        val planned = RaidFixture.load(GuildRaids.serializer(), "guild.raids").planned.orEmpty()
        assertEquals(RowChip.Signed(SignupState.In, "dps"), RaidFacts(planned[0], RaidTime.now).rowChip(null))
        assertEquals(RowChip.NotSigned, RaidFacts(planned[1], RaidTime.now).rowChip(null))
        assertEquals(RowChip.Status(RaidChip.Locked), RaidFacts(planned[2], RaidTime.now).rowChip(null))
        assertEquals(RowChip.Status(RaidChip.Cancelled), RaidFacts(RaidFixture.raid("guild.raid.cancelled"), RaidTime.now).rowChip(null))
    }

    @Test
    fun `the next raid is the earliest signable one, Also planned lists the rest`() {
        val planned = RaidFixture.load(GuildRaids.serializer(), "guild.raids").planned.orEmpty()
        val split = PlannedRaids(planned.reversed(), RaidTime.now)
        assertEquals(31L, split.next?.id)
        assertEquals(listOf(32L, 33L), split.also.map { it.id })
    }

    // Roster

    @Test
    fun `groups in order, class order, then name`() {
        val roster = RosterGroups(RaidFixture.raid())
        assertEquals(RosterKind.entries.toList(), roster.groups.map { it.kind })
        assertEquals(listOf("Ironhide", "Kraggor", "Bramblehoof"), roster.groups[0].rows.map { it.name })
        assertEquals("Draknar", roster.groups[2].rows.first().name)
        assertEquals(listOf("Brokk", "Zulmara", "Sennah"), roster.notSigned.map { it.name })
    }

    @Test
    fun `an empty role says so, an empty Late, Bench or Out is left out`() {
        val barrow = RaidFixture.load(GuildRaids.serializer(), "guild.raids").planned.orEmpty()[1]
        assertEquals(listOf(RosterKind.Tanks, RosterKind.Healers, RosterKind.Dps), RosterGroups(barrow).groups.map { it.kind })
    }

    @Test
    fun `an officer's move shows at once in the new group`() {
        val roster = RosterGroups(RaidFixture.raid(), mapOf("us/pvp/kraggor" to SignupState.Bench))
        assertTrue(roster.groups.first { it.kind == RosterKind.Bench }.rows.any { it.name == "Kraggor" })
    }

    @Test
    fun `the composition block words the API's counts, targets and buff coverage`() {
        val read = requireNotNull(CompositionRead.of(RaidFixture.raid().composition, resources()))
        assertEquals(listOf(5, 0, 2, 2, 4, 4, 2, 2, 3), read.classes.map { it.count })
        assertEquals("Paladin", read.classes[1].name)
        assertEquals(listOf("Tanks 3", "Healers 6", "DPS 14"), read.roles)
        assertEquals("Missing: Blessings", read.buffSummary)
        assertEquals("Blessings, missing", read.buffs.first { it.key == "blessings" }.accessibilityLabel)
        assertEquals("MotW", read.buffs.first { it.key == "mark_of_the_wild" }.label)
        assertTrue(read.buffs.any { it.key == "windfury_totem" && it.present })
        assertTrue(read.classesAccessibilityLabel.startsWith("Classes coming: Warrior 5, Paladin 0, Hunter 2"))
        val officer = requireNotNull(CompositionRead.of(RaidFixture.raid("guild.raid.officer").composition, resources()))
        assertEquals(listOf("Tanks 3/4", "Healers 6/8", "DPS 14/28"), officer.roles)
        assertNull(CompositionRead.of(null, resources()))
    }

    @Test
    fun `Copy the list - header, count, then each group's names`() {
        val lines = RosterCopy.text(RaidFixture.raid(), clock(), resources(), RaidTime.now).lines()
        assertTrue(lines[0], lines[0].startsWith("Onyxia's Lair · Sat 19 Dec, 7:00 pm CST · <t:"))
        assertEquals("24 of 40: 3 tanks, 6 healers, 14 DPS, 1 late", lines[1])
        assertEquals("Tanks: Ironhide, Kraggor, Bramblehoof", lines[2])
        assertEquals("Out: Wyrmtalon, Snaggle, Felgrin", lines.last())
    }

    @Test
    fun `fixture mode reads the active build from the bundled fixtures, so icon URLs resolve`() = runTest {
        val engine = gg.foreversixty.app.network.FixtureEngine.create(
            gg.foreversixty.app.network.ClasspathFixtureSource(requireNotNull(javaClass.classLoader)),
        )
        val build = gg.foreversixty.app.network.StaticDataClient(engine).activeBuild().let { (it as gg.foreversixty.app.network.ApiResult.Success).data.build }
        assertEquals("1.60.1.70658", build)
        assertEquals(
            "https://foreversixty.gg/data/1.60.1.70658/icons/spell_nature_windfury.webp",
            gg.foreversixty.app.raids.ui.RaidIcons(build).url("spell_nature_windfury"),
        )
    }

    @Test
    fun `the move menu offers the spec's items per place`() {
        assertEquals(listOf(SignupState.Bench, SignupState.Out), placeMenuItems(SignupState.In).map { it.second })
        assertEquals(listOf(SignupState.In, SignupState.Bench), placeMenuItems(null).map { it.second })
        assertTrue(placeMenuItems(SignupState.Bench)[1].third)
    }

    // Sign-up

    @Test
    fun `a press draws at once, saves, says Saved and plays tick then confirm`() = runTest {
        val service = FakeRaidService()
        val haptics = RecordingHaptics()
        val controller = makeController(service = service, haptics = haptics)
        controller.press(key, SignupState.Late)
        assertEquals("late", service.signups.first().second.state)
        assertEquals(StatusLine.Saved("Saved."), controller.block(key).status)
        assertEquals(listOf("selection", "success"), haptics.calls)
    }

    @Test
    fun `the PUT carries the stored note, since the PUT writes it`() = runTest {
        val service = FakeRaidService()
        val controller = makeController(service = service)
        controller.editNote(key, "Late from work")
        controller.press(key, SignupState.Bench)
        assertEquals("Late from work", service.signups.first().second.note)
    }

    @Test
    fun `a refused press restores the previous answer with the API's sentence (F4_11)`() = runTest {
        val service = FakeRaidService()
        service.signupAnswer = { fail(RaidFixture.error("guild.signup.error-raid-locked", 409)) }
        val haptics = RecordingHaptics()
        val controller = makeController(service = service, haptics = haptics)
        controller.press(key, SignupState.Late)
        assertEquals(SignupState.In, controller.block(key).shown)
        assertEquals(StatusLine.Error("This raid was locked while you were signing up."), controller.block(key).status)
        assertEquals(listOf("selection", "error"), haptics.calls)
    }

    @Test
    fun `every refusal maps to its sentence`() {
        val res = resources()
        assertEquals("This raid was cancelled while you were signing up.", SignupRefusal.sentence(RaidFixture.error("guild.signup.error-raid-cancelled", 409), res))
        assertEquals("Too many sign-up changes. Wait a minute, then try again.", SignupRefusal.sentence(RaidFixture.error("guild.signup.error-rate-limited", 429), res))
        assertEquals("Not saved. This guild's officer tools have ended.", SignupRefusal.sentence(RaidFixture.error("guild.signup.error-plan-required", 403), res))
        assertEquals("You can only sign up your own characters.", SignupRefusal.sentence(RaidFixture.error("guild.signup.error-forbidden", 403), res))
        assertEquals("You are offline. Sign-ups need a connection.", SignupRefusal.sentence(ApiError.Network(java.io.IOException()), res))
        assertEquals("That did not save. Try again.", SignupRefusal.sentence(ApiError.Status(500), res))
    }

    @Test
    fun `a lost connection disables the segments with the offline line`() = runTest {
        val service = FakeRaidService()
        service.signupAnswer = { fail(ApiError.Network(java.io.IOException())) }
        val controller = makeController(service = service)
        controller.press(key, SignupState.Out)
        assertTrue(controller.offline)
        assertFalse(controller.segmentsEnabled(key))
    }

    @Test
    fun `In with no saved spec asks which spec, the pick sends it and says it is kept`() = runTest {
        val service = FakeRaidService()
        val ask = "us/pvp/lightningbolltz"
        val controller = makeController(RaidFixture.raid("guild.raid.ask-spec"), service)
        controller.press(ask, SignupState.In)
        assertEquals(SpecAsk.Ask(ask, SignupState.In), controller.specAsk)
        assertEquals(SignupState.In, controller.block(ask).pending)
        assertTrue(service.signups.isEmpty())
        val elemental = controller.character(ask)?.specOptions.orEmpty().first { it.spec == "shaman-elemental" }
        controller.pickSpec(elemental)
        assertEquals("shaman-elemental", service.signups.first().second.spec)
        assertEquals(StatusLine.Saved("Saved. Elemental is kept for Lightningbolltz."), controller.block(ask).status)
    }

    @Test
    fun `closing the spec sheet drops the press, Out never asks`() = runTest {
        val service = FakeRaidService()
        val ask = "us/pvp/lightningbolltz"
        val controller = makeController(RaidFixture.raid("guild.raid.ask-spec"), service)
        controller.press(ask, SignupState.Bench)
        controller.closeSpecSheet()
        assertNull(controller.block(ask).pending)
        assertNull(controller.specAsk)
        assertTrue(service.signups.isEmpty())
        controller.press(ask, SignupState.Out)
        assertNull(controller.specAsk)
        assertEquals("out", service.signups.first().second.state)
    }

    @Test
    fun `Change spec resends the current answer with the new spec`() = runTest {
        val service = FakeRaidService()
        val controller = makeController(service = service)
        controller.changeSpec(key)
        assertEquals(SpecAsk.Change(key), controller.specAsk)
        controller.pickSpec(controller.character(key)?.specOptions.orEmpty().first { it.spec == "warrior-arms" })
        assertEquals("in", service.signups.first().second.state)
        assertEquals("warrior-arms", service.signups.first().second.spec)
    }

    @Test
    fun `moved draws Bench on the other character's block, with its line (R3, F4_12)`() = runTest {
        val service = FakeRaidService()
        val answer = RaidFixture.load(SignupAnswer.serializer(), "guild.signup.moved")
        service.signupAnswer = { ok(answer) }
        val controller = makeController(RaidFixture.raid("guild.raid.two-characters"), service)
        controller.press("us/pvp/yellfury", SignupState.In)
        assertEquals(SignupState.In, controller.block("us/pvp/yellfury").shown)
        assertEquals(SignupState.Bench, controller.block(key).shown)
        assertEquals(StatusLine.Info("Obnoxious Yell moved to Bench."), controller.block(key).status)
    }

    @Test
    fun `a locked raid disables a member's segments and keeps an officer's`() {
        val member = makeController(RaidFixture.raid("guild.raid.locked"))
        assertFalse(member.segmentsEnabled(key))
        assertEquals(StatusLine.Info("Sign-ups are locked. An officer can still change yours."), member.restingStatus(key))
        assertFalse(makeController(RaidFixture.raid("guild.raid.cancelled")).segmentsEnabled(key))
        assertTrue(makeController(RaidFixture.raid("guild.raid.officer-locked")).segmentsEnabled(key))
    }

    @Test
    fun `an unverified character has no controls`() {
        assertFalse(makeController(RaidFixture.raid("guild.raid.two-characters")).segmentsEnabled("us/pvp/vorgash"))
    }

    @Test
    fun `the reminder ask shows after a first In, at most twice, 30 days apart`() = runTest {
        val store = MemoryAskStore()
        val controller = makeController(undetermined = true, askStore = store)
        controller.press(key, SignupState.In)
        assertTrue(controller.showsReminderAsk)
        assertEquals(1, store.shown)
        assertFalse(ReminderAskPolicy(1, RaidTime.now).mayAsk(RaidTime.now.plus(Duration.ofDays(1))))
        assertTrue(ReminderAskPolicy(1, RaidTime.now).mayAsk(RaidTime.now.plus(Duration.ofDays(31))))
        assertFalse(ReminderAskPolicy(2, null).mayAsk(RaidTime.now))
        val denied = makeController(undetermined = false)
        denied.press(key, SignupState.In)
        assertFalse(denied.showsReminderAsk)
    }

    @Test
    fun `an officer's move sends the row's state and role, a failure restores the row`() = runTest {
        val service = FakeRaidService()
        val officer = RaidFixture.raid("guild.raid.officer")
        service.signupAnswer = { ok(SignupAnswer(officer, null)) }
        val controller = makeController(officer, service)
        val kraggor = RosterGroups(controller.raid).groups[0].rows.first { it.name == "Kraggor" }
        controller.move(kraggor, SignupState.Bench)
        assertEquals("us/pvp/kraggor", service.signups.first().first)
        assertEquals("bench", service.signups.first().second.state)
        assertEquals("tank", service.signups.first().second.role)
        service.signupAnswer = { fail(ApiError.Status(500)) }
        controller.move(kraggor, SignupState.Out)
        assertTrue(controller.rosterOverrides.isEmpty())
        assertEquals("That did not save. Try again.", controller.rowErrors["us/pvp/kraggor"])
    }

    @Test
    fun `a member cannot move anyone`() = runTest {
        val service = FakeRaidService()
        val controller = makeController(service = service)
        controller.move(RosterGroups(controller.raid).groups[0].rows[0], SignupState.Bench)
        assertTrue(service.signups.isEmpty())
    }

    @Test
    fun `Lock sign-ups locks and toasts who was told`() = runTest {
        val service = FakeRaidService()
        val controller = makeController(RaidFixture.raid("guild.raid.officer"), service)
        controller.toggleLock()
        assertEquals(listOf("lock"), service.locks)
        assertEquals("Sign-ups locked. 23 In and 1 on the Bench were told.", controller.toast)
    }

    // Raids tab

    private fun tab(service: FakeRaidService, cache: gg.foreversixty.app.raids.logic.RaidCache = gg.foreversixty.app.raids.logic.NoRaidCache) =
        RaidsTabModel(service, cache, resources(), { RaidTime.now })

    @Test
    fun `signed out on a 401`() = runTest {
        val service = FakeRaidService()
        service.meAnswer = { fail(RaidFixture.error("error-unauthorized", 401)) }
        val model = tab(service)
        model.load()
        assertEquals(RaidsTabState.SignedOut, model.state)
    }

    @Test
    fun `a character in no guild shows No guild yet with its name`() = runTest {
        val model = tab(FakeRaidService())
        model.currentKey = "us/pvp/yellfury"
        model.load()
        assertEquals(RaidsTabState.NoGuild("Yellfury"), model.state)
    }

    @Test
    fun `an unverified member waits for approval`() = runTest {
        val service = FakeRaidService()
        service.meAnswer = { ok(RaidFixture.load(Me.serializer(), "me.waiting")) }
        val model = tab(service)
        model.load()
        assertTrue(model.state is RaidsTabState.Waiting)
    }

    @Test
    fun `loaded - the planned raids, three raid nights at most, the officer's waiting count`() = runTest {
        val model = tab(FakeRaidService())
        model.load()
        val content = (model.state as RaidsTabState.Loaded).content
        assertEquals(3, content.planned.size)
        assertEquals(2, content.waiting)
        assertEquals("Horde · US PvP · 24 on the roster", content.head.line)
    }

    @Test
    fun `offline with a cached answer shows it under the strip, with none, the error`() = runTest {
        val cache = MemoryRaidCache()
        val service = FakeRaidService()
        val warm = tab(service, cache)
        warm.load()
        service.raidsAnswer = { fail(ApiError.Network(java.io.IOException())) }
        warm.load()
        assertEquals(RaidTime.now, (warm.state as RaidsTabState.Loaded).content.offlineSince)
        val cold = tab(service)
        cold.load()
        assertEquals(RaidsTabState.Failed, cold.state)
    }

    // Raid notifications

    private fun payload(extra: Map<String, String> = emptyMap()): RaidPushPayload {
        val data = RaidFixture.push("push.raid-posted.fcm").getValue("message").jsonObject.getValue("data").jsonObject
            .mapValues { it.value.jsonPrimitive.content } + mapOf("raid_id" to "31") + extra
        return requireNotNull(RaidPushPayload.parse(data))
    }

    @Test
    fun `the payload names the raid, the character and the answers`() {
        val push = payload()
        assertEquals(2L, push.guild)
        assertEquals(31L, push.raid)
        assertEquals("us/pvp/obnoxious-yell", push.characterKey)
        assertFalse(push.askSpec)
        assertEquals("You are In for Onyxia's Lair", push.answers["in"]?.title)
        assertEquals(SignupState.Late, RaidPushCategory.stateFor("signup_late"))
    }

    @Test
    fun `In from the shade sends the PUT for the chosen character, keeping its note`() = runTest {
        val service = FakeRaidService()
        val outcome = RaidPushAnswerer(service, resources()).answer(payload(), SignupState.In)
        assertEquals("us/pvp/obnoxious-yell", service.signups.first().first)
        assertEquals("in", service.signups.first().second.state)
        assertEquals("You are In for Onyxia's Lair", (outcome as RaidPushOutcome.Answered).answer.title)
    }

    @Test
    fun `a moved character is named after the answer (R3)`() = runTest {
        val service = FakeRaidService()
        service.signupAnswer = { ok(RaidFixture.load(SignupAnswer.serializer(), "guild.signup.moved")) }
        val outcome = RaidPushAnswerer(service, resources()).answer(payload(), SignupState.In) as RaidPushOutcome.Answered
        assertTrue(outcome.answer.body.endsWith(" Obnoxious Yell moved to Bench."))
    }

    @Test
    fun `a locked raid posts the locked answer, a failure posts the failed one`() = runTest {
        val service = FakeRaidService()
        service.signupAnswer = { fail(RaidFixture.error("guild.signup.error-raid-locked", 409)) }
        val locked = RaidPushAnswerer(service, resources()).answer(payload(), SignupState.Late) as RaidPushOutcome.Answered
        assertEquals(RaidPushPayload.Answer("Sign-ups for Onyxia's Lair are locked", "An officer can still change yours."), locked.answer)
        service.raidAnswer = { fail(ApiError.Network(java.io.IOException())) }
        val failed = RaidPushAnswerer(service, resources()).answer(payload(), SignupState.Late) as RaidPushOutcome.Answered
        assertEquals("That did not save", failed.answer.title)
    }

    @Test
    fun `a character with no spec opens the raid page with the spec sheet`() = runTest {
        val service = FakeRaidService()
        service.raidAnswer = { ok(RaidFixture.raid("guild.raid.ask-spec")) }
        val outcome = RaidPushAnswerer(service, resources()).answer(payload(mapOf("character_key" to "us/pvp/lightningbolltz")), SignupState.In)
        assertEquals(RaidPushOutcome.Open(DeepLink.Raid(2, 31, spec = "ask")), outcome)
        assertTrue(service.signups.isEmpty())
    }

    @Test
    fun `with no pick and two unsigned characters, the tap opens the page`() {
        val push = payload(mapOf("character_key" to ""))
        assertEquals("us/pvp/obnoxious-yell", RaidPushAnswerer.character(push, RaidFixture.raid("guild.raid.two-characters"))?.characterKey)
        assertEquals("us/pvp/lightningbolltz", RaidPushAnswerer.character(push, RaidFixture.raid("guild.raid.ask-spec"))?.characterKey)
    }

    // Raid links

    @Test
    fun `a raid link pushes the raid page, night pushes the page then the night, spec-ask presses In`() {
        val navigator = RaidsNavigator()
        assertTrue(navigator.open(DeepLink.Raid(2, 41, spec = "ask")))
        assertEquals(listOf(RaidsDestination.Raid(2, 41, open = RaidOpenAction(SignupState.In))), navigator.stack.toList())
        navigator.open(DeepLink.RaidNight(2, 41))
        assertEquals(listOf(RaidsDestination.Raid(2, 41), RaidsDestination.Night(41)), navigator.stack.toList())
        navigator.open(DeepLink.Guild(2))
        assertTrue(navigator.stack.isEmpty())
        assertFalse(navigator.open(DeepLink.Plan))
    }

    // Notes

    @Test
    fun `notes_html keeps bold and https links, drops other tags`() {
        val html = "<p>Bring <strong>fire resistance</strong>. See <a href=\"https://foreversixty.gg/guides/onyxia\">the guide</a>.</p>\n<p>Second &amp; last</p>"
        val text = NotesText.annotated(html, androidx.compose.ui.graphics.Color.Yellow)
        assertEquals("Bring fire resistance. See the guide.\nSecond & last", text.text)
        assertTrue(text.getLinkAnnotations(0, text.length).isNotEmpty())
    }
}
