package gg.foreversixty.app.raids

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.RaidNightSummary
import gg.foreversixty.app.model.SignupAnswer
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.raids.logic.GuildHead
import gg.foreversixty.app.raids.logic.RaidNightModel
import gg.foreversixty.app.raids.logic.RaidsTabContent
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.logic.RosterGroups
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.logic.SpecAsk
import gg.foreversixty.app.raids.ui.PlaceMenu
import gg.foreversixty.app.raids.ui.RaidNightScreen
import gg.foreversixty.app.raids.ui.RaidPageContent
import gg.foreversixty.app.raids.ui.RaidsTabContentView
import gg.foreversixty.app.raids.ui.SpecSheetContent
import java.time.Instant
import kotlinx.coroutines.test.runTest
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * Every F4 screen and state on a Pixel 9 (412 x 915 dp), at the default size and at font scale 2.0
 * (spec 15). Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 * The PNGs double as the review captures in design/reviews/2026-10-31-mobile-raids.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class RaidScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private fun capture(name: String, large: Boolean = false, screen: @Composable () -> Unit) {
        if (large) RuntimeEnvironment.setFontScale(2f)
        compose.setContent {
            ForeverTheme { Box(Modifier.fillMaxSize().background(FsColor.bg)) { screen() } }
        }
        compose.onRoot().captureRoboImage("src/test/screenshots/raids/$name.png", roborazziOptions = TOLERANT)
    }

    private val now = { RaidTime.now }

    // Raids tab (F4.1, F4.13 to F4.17)

    private fun content(offline: Instant? = null, planned: List<gg.foreversixty.app.model.RaidView>? = null): RaidsTabContent {
        val me = RaidFixture.load(Me.serializer(), "me")
        val raids = RaidFixture.load(GuildRaids.serializer(), "guild.raids")
        val character = requireNotNull(me.mainCharacter)
        return RaidsTabContent(
            head = GuildHead.of(requireNotNull(character.guild), me.guilds.firstOrNull(), resources()),
            character = character,
            planned = planned ?: raids.planned.orEmpty(),
            nights = raids.rows,
            waiting = raids.waitingForApproval,
            offlineSince = offline,
        )
    }

    @Composable
    private fun Tab(state: RaidsTabState, controller: SignupController? = null) = RaidsTabContentView(
        state, clock(), now, controller, { _, _ -> }, {}, {}, {}, {},
    )

    @Test
    fun raidsTab_loaded() {
        val c = content()
        capture("raids-tab-loaded-and") { Tab(RaidsTabState.Loaded(c), makeController(c.planned.first())) }
    }

    @Test
    @Config(qualifiers = "w412dp-h2000dp-normal-long-notround-any-420dpi-keyshidden-nonav")
    fun raidsTab_loaded_xl() {
        val c = content()
        capture("raids-tab-loaded-xl-and", large = true) { Tab(RaidsTabState.Loaded(c), makeController(c.planned.first())) }
    }

    @Test
    fun raidsTab_noGuild() = capture("raids-tab-no-guild-and") { Tab(RaidsTabState.NoGuild("Yellfury")) }

    @Test
    fun raidsTab_waiting() {
        val me = RaidFixture.load(Me.serializer(), "me.waiting")
        val character = requireNotNull(me.mainCharacter)
        val head = GuildHead.of(requireNotNull(character.guild), me.guilds.firstOrNull(), resources())
        capture("raids-tab-waiting-and") { Tab(RaidsTabState.Waiting(head, character)) }
    }

    @Test
    fun raidsTab_empty() = capture("raids-tab-no-raid-and") { Tab(RaidsTabState.Loaded(content(planned = emptyList()))) }

    @Test
    fun raidsTab_error() = capture("raids-tab-error-and") { Tab(RaidsTabState.Failed) }

    @Test
    fun raidsTab_offline() {
        val c = content(offline = RaidTime.now.minusSeconds(3600))
        val controller = makeController(c.planned.first()).also { it.offline = true }
        capture("raids-tab-offline-and") { Tab(RaidsTabState.Loaded(c), controller) }
    }

    @Test
    fun raidsTab_signedOut() = capture("raids-tab-signed-out-and") { Tab(RaidsTabState.SignedOut) }

    @Test
    fun raidsTab_loading() = capture("raids-tab-loading-and") { Tab(RaidsTabState.Loading) }

    // Raid page (F4.2, F4.6 to F4.12, F4.18 to F4.21)

    @Composable
    private fun Page(controller: SignupController, offline: Instant? = null) =
        RaidPageContent(controller, clock(), now, null, offline, "OLYMPUS XXVII", null, {}, {}, {})

    @Test
    fun raidPage_member() = capture("raid-page-member-and") { Page(makeController()) }

    @Test
    @Config(qualifiers = "w412dp-h3000dp-normal-long-notround-any-420dpi-keyshidden-nonav")
    fun raidPage_full() = capture("raid-page-full-and") { Page(makeController()) }

    @Test
    @Config(qualifiers = "w412dp-h3000dp-normal-long-notround-any-420dpi-keyshidden-nonav")
    fun raidPage_xl() = capture("raid-page-xl-and", large = true) { Page(makeController()) }

    @Test
    fun raidPage_unsigned() = capture("raid-page-unsigned-and") { Page(makeController(RaidFixture.raid("guild.raid.ask-spec"))) }

    @Test
    fun raidPage_noInvite() = capture("raid-page-no-invite-and") { Page(makeController(RaidFixture.raid("guild.raid.no-invite"))) }

    @Test
    fun raidPage_locked() = capture("raid-page-locked-and") { Page(makeController(RaidFixture.raid("guild.raid.locked"))) }

    @Test
    fun raidPage_cancelled() = capture("raid-page-cancelled-and") { Page(makeController(RaidFixture.raid("guild.raid.cancelled"))) }

    @Test
    fun raidPage_offline() {
        val controller = makeController().also { it.offline = true }
        capture("raid-page-offline-and") { Page(controller, RaidTime.now.minusSeconds(600)) }
    }

    private fun pressed(state: SignupState, answer: (FakeRaidService) -> Unit, raid: String = "guild.raid", key: String = "us/pvp/obnoxious-yell", undetermined: Boolean = false): SignupController {
        val service = FakeRaidService().also(answer)
        val controller = makeController(RaidFixture.raid(raid), service, undetermined = undetermined)
        runTest { controller.press(key, state) }
        return controller
    }

    @Test
    fun raidPage_saveFailed() {
        val controller = pressed(SignupState.Late, { it.signupAnswer = { fail(ApiError.Status(500)) } })
        capture("raid-page-save-failed-and") { Page(controller) }
    }

    @Test
    fun raidPage_rateLimited() {
        val controller = pressed(SignupState.Bench, { it.signupAnswer = { fail(RaidFixture.error("guild.signup.error-rate-limited", 429)) } })
        capture("raid-page-rate-limited-and") { Page(controller) }
    }

    @Test
    fun raidPage_raceLocked() {
        val controller = pressed(SignupState.Late, { it.signupAnswer = { fail(RaidFixture.error("guild.signup.error-raid-locked", 409)) } })
        capture("raid-page-race-locked-and") { Page(controller) }
    }

    @Test
    @Config(qualifiers = "w412dp-h1200dp-normal-long-notround-any-420dpi-keyshidden-nonav")
    fun raidPage_moved() {
        val two = RaidFixture.raid("guild.raid.two-characters")
        val controller = pressed(SignupState.In, { it.signupAnswer = { ok(SignupAnswer(two, "Yellfury")) } }, raid = "guild.raid.two-characters")
        capture("raid-page-moved-and") { Page(controller) }
    }

    @Test
    @Config(qualifiers = "w412dp-h1200dp-normal-long-notround-any-420dpi-keyshidden-nonav")
    fun raidPage_reminderAsk() {
        val controller = pressed(SignupState.In, {}, undetermined = true)
        capture("raid-page-reminder-ask-and") { Page(controller) }
    }

    @Test
    fun raidPage_done() = capture("raid-page-done-and") { Page(makeController(RaidFixture.raid("guild.raid.done"))) }

    @Test
    @Config(qualifiers = "w412dp-h3200dp-normal-long-notround-any-420dpi-keyshidden-nonav")
    fun raidPage_officer() = capture("raid-page-officer-and") { Page(makeController(RaidFixture.raid("guild.raid.officer"))) }

    @Test
    fun raidPage_officerLocked() {
        val service = FakeRaidService().also { it.lockAnswer = { ok(RaidFixture.raid("guild.raid.officer-locked")) } }
        val controller = makeController(RaidFixture.raid("guild.raid.officer"), service)
        runTest { controller.toggleLock() }
        capture("raid-page-officer-locked-and") { Page(controller) }
    }

    @Test
    fun officer_moveMenu() {
        val controller = makeController(RaidFixture.raid("guild.raid.officer"))
        val kraggor = RosterGroups(controller.raid).groups[0].rows.first { it.name == "Kraggor" }
        capture("officer-move-menu-and") { Box(Modifier.background(FsColor.raised)) { PlaceMenu(kraggor, {}, initiallyOpen = true) } }
    }

    // Sheets and Your raid night (F4.3, F4.22, F4.23)

    @Composable
    private fun Sheet(raid: String, ask: SpecAsk) {
        val controller = makeController(RaidFixture.raid(raid))
        val character = requireNotNull(controller.character(ask.characterKey))
        SpecSheetContent(ask, character, "Onyxia's Lair", "Sat 19 Dec") {}
    }

    @Test
    fun specSheet_ask() = capture("spec-sheet-ask-and") { Sheet("guild.raid.ask-spec", SpecAsk.Ask("us/pvp/lightningbolltz", SignupState.In)) }

    @Test
    fun specSheet_ask_xl() = capture("spec-sheet-ask-xl-and", large = true) { Sheet("guild.raid.ask-spec", SpecAsk.Ask("us/pvp/lightningbolltz", SignupState.In)) }

    @Test
    fun specSheet_change() = capture("spec-sheet-change-and") { Sheet("guild.raid", SpecAsk.Change("us/pvp/obnoxious-yell")) }

    @Composable
    private fun Night(fixture: String) {
        val night = RaidFixture.load(RaidNightSummary.serializer(), fixture)
        val service = FakeRaidService().also { it.nightAnswer = { ok(night) } }
        val model = RaidNightModel(29, service, gg.foreversixty.app.raids.logic.NoRaidCache, now, RaidNightModel.State.Loaded(night))
        RaidNightScreen(model, clock(), now) {}
    }

    @Test
    fun raidNight_premium() = capture("raid-night-premium-and") { Night("me.raid-night") }

    @Test
    fun raidNight_free() = capture("raid-night-free-and") { Night("me.raid-night.free") }
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
