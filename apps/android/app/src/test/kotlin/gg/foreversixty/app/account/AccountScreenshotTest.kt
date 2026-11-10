package gg.foreversixty.app.account

import androidx.compose.ui.test.swipeUp
import androidx.compose.ui.test.performTouchInput
import androidx.compose.ui.test.hasScrollAction
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.compose.ui.unit.dp
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.ExperimentalRoborazziApi
import com.github.takahirom.roborazzi.captureRoboImage
import com.github.takahirom.roborazzi.captureScreenRoboImage
import gg.foreversixty.app.auth.session.AccountActionOutcome
import gg.foreversixty.app.auth.session.DeletionConsequence
import gg.foreversixty.app.auth.session.DeviceList
import gg.foreversixty.app.auth.session.DeviceRow
import gg.foreversixty.app.auth.session.SignInFailure
import gg.foreversixty.app.auth.session.SignedOutNotice
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.Device
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.premium.PlanSummary
import kotlinx.serialization.builtins.ListSerializer
import kotlinx.serialization.json.JsonArray
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.JsonPrimitive
import kotlinx.serialization.json.jsonArray
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode
import java.time.Instant
import gg.foreversixty.app.testing.STRICT_SNAPSHOT

/**
 * F1's states at Pixel 9 size (spec §15), default and font scale 2.0 for the switcher and the
 * signed-out screen. The same images are the review captures.
 * Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class AccountScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private fun fixture(name: String): String =
        requireNotNull(javaClass.classLoader?.getResource("fixtures/app/$name")) { "no fixture $name" }.readText()

    private fun data(name: String) = requireNotNull(ApiJson.parseToJsonElement(fixture(name)).jsonObject["data"]) { "no data in $name" }

    private fun me(name: String): Me = ApiJson.decodeFromJsonElement(Me.serializer(), data(name))

    private fun deviceRows(): List<DeviceRow> {
        val devices = ApiJson.decodeFromJsonElement(ListSerializer(Device.serializer()), data("devices.list.json"))
        return DeviceList.rows(devices, "dvc_8Qk2nXw4", Instant.parse("2026-10-30T21:41:00Z"))
    }

    @Test fun `signed out`() = tab("signedout-default") { SignedOut(null) }

    @Test fun `signing in`() = tab("signin-browser-open") { SignedOutScreen(null, signingIn = true, onSignIn = {}, onPickClass = {}) }

    @Test fun `importing`() = capture("import-default") { ImportScreen() }

    @Test fun `sign-in failed`() = tab("signin-failed") { SignedOut(AccountCopy.notice(SignedOutNotice.SignInFailed(SignInFailure.BattleNet))) }

    @Test fun `signed out by a 401`() = tab("signedout-401") { SignedOut(AccountCopy.notice(SignedOutNotice.SessionEnded)) }

    @Test fun `signed out after deletion`() = tab("signedout-deleted") {
        SignedOutScreen(
            AccountCopy.notice(SignedOutNotice.AccountDeleted),
            signingIn = false,
            onSignIn = {},
            onPickClass = {},
            noticeIsFailure = false,
        )
    }

    @Test fun `no characters`() = tab("home-no-characters") { NoCharactersScreen {} }

    @Test fun `signed out at font scale 2`() {
        RuntimeEnvironment.setFontScale(2f)
        tab("signedout-ax5") { SignedOut(null) }
    }

    @Test fun `you free`() = tab("you-free", ForeverTab.Home) { You(me("me.json")) }

    @Test fun `you premium`() = tab("you-premium", ForeverTab.Home) { You(me("me.premium.json")) }

    @Test fun `sign out confirmation`() = screen("you-signout-confirm") { You(me("me.json"), confirming = true) }

    @Test fun `devices`() = tab("devices-default", ForeverTab.Home) {
        DevicesScreen(DevicesState.Loaded(deviceRows()), null, {}, {}, back = {})
    }

    @Test fun `revoke confirmation`() {
        val rows = deviceRows()
        screen("devices-revoke-confirm") { DevicesScreen(DevicesState.Loaded(rows), null, {}, {}, initiallyConfirming = rows[1], back = {}) }
    }

    @Test fun `revoke failed`() {
        val rows = deviceRows()
        tab("devices-revoke-failed", ForeverTab.Home) { DevicesScreen(DevicesState.Loaded(rows), rows[1].id, {}, {}, back = {}) }
    }

    @Test fun `delete account, free`() = delete("delete-free", "me.json")

    @Test fun `delete account, Premium founder`() = delete("delete-premium", "me.premium.json")

    @Test fun `delete account, seat holder`() = delete("delete-seat", "me.seat.json")

    @Test fun `delete account, failed`() = delete("delete-failed", "me.json", outcome = AccountActionOutcome.Failed)

    @Test fun `delete account, refused for a guild obligation`() =
        delete("delete-guild-obligation", "me.guild-payer.json", outcome = AccountActionOutcome.Refused(errorMessage("me.delete.error-guild-obligation.json")))

    @Test fun `delete account, typed BattleTag mismatch`() =
        delete("delete-mismatch", "me.json", typed = "Yell#16", outcome = AccountActionOutcome.ConfirmMismatch)

    @Test fun `delete account, BattleTag typed`() = delete("delete-ready", "me.json", typed = "yell#1604")

    @Test fun `about`() = tab("about-default", ForeverTab.Home) { AboutScreen(VERSION, back = {}) }

    private fun errorMessage(name: String): String =
        requireNotNull(ApiJson.parseToJsonElement(fixture(name)).jsonObject["error"]?.jsonObject?.get("message")) { "no message in $name" }
            .jsonPrimitive.content

    private fun delete(name: String, file: String, typed: String = "", outcome: AccountActionOutcome? = null) {
        val me = me(file)
        val consequences = DeletionConsequence.of(me)
        tab(name, ForeverTab.Home) {
            DeleteAccountScreen(consequences, me.user.battletag, typed, {}, deleting = false, outcome = outcome, onDelete = {}, back = {})
        }
        // The same state scrolled to its end: the BattleTag field, the button and the outcome are the step under review.
        compose.onNode(hasScrollAction()).performTouchInput { repeat(SCROLL_SWIPES) { swipeUp() } }
        compose.waitForIdle()
        compose.onRoot().captureRoboImage("src/test/screenshots/account_$name-end.png", roborazziOptions = STRICT_SNAPSHOT)
    }

    @Composable
    private fun SignedOut(notice: String?) = SignedOutScreen(notice, signingIn = false, onSignIn = {}, onPickClass = {})

    @Composable
    private fun You(me: Me, confirming: Boolean = false) {
        val current = me.characters.first { it.key == me.mainCharacterKey }
        YouScreen(
            summary = YouSummary(
                battletag = me.user.battletag.orEmpty(),
                plan = PlanSummary.of(me, Instant.parse("2025-11-01T00:00:00Z")),
                character = CharacterLine(current.name, current.className, AccountCopy.place(current)),
                devices = "This phone and 2 more",
                version = VERSION,
            ),
            actions = YouActions({}, {}, {}, {}, {}, {}, {}, {}),
            confirmingSignOut = confirming,
            back = {},
        )
    }

    private fun tab(name: String, tab: ForeverTab = ForeverTab.Home, screen: @Composable () -> Unit) =
        capture(name) { ForeverTabScaffold(selected = tab, onSelect = {}) { screen() } }

    private fun capture(name: String, screen: @Composable () -> Unit) {
        compose.setContent { ForeverTheme { screen() } }
        compose.onRoot().captureRoboImage("src/test/screenshots/account_$name.png", roborazziOptions = STRICT_SNAPSHOT)
    }

    /** Dialogs live in their own window, so these capture the whole screen. */
    @OptIn(ExperimentalRoborazziApi::class)
    private fun screen(name: String, content: @Composable () -> Unit) {
        compose.setContent { ForeverTheme { ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}) { content() } } }
        compose.waitForIdle()
        captureScreenRoboImage("src/test/screenshots/account_$name.png", roborazziOptions = STRICT_SNAPSHOT)
    }
}

/** Enough swipes to reach the end of the longest delete state at 1x. */
private const val SCROLL_SWIPES = 3

private const val VERSION = "Version 0.1.0 (1) · data 1.60.1.70658"

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
