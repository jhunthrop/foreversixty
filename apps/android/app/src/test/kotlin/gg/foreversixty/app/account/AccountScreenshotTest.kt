package gg.foreversixty.app.account

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
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import com.github.takahirom.roborazzi.captureScreenRoboImage
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
import kotlinx.serialization.builtins.ListSerializer
import kotlinx.serialization.json.JsonArray
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.JsonPrimitive
import kotlinx.serialization.json.jsonArray
import kotlinx.serialization.json.jsonObject
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode
import java.time.Instant

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

    /** me.json plus a third character with no guild, as the spec's F1.5 capture asks. */
    private fun threeCharacters(): Me {
        val root = data("me.json").jsonObject
        val extra = JsonObject(
            mapOf(
                "key" to JsonPrimitive("us/pvp/sablewhisper"), "region" to JsonPrimitive("us"),
                "ruleset" to JsonPrimitive("pvp"), "name" to JsonPrimitive("Sablewhisper"),
                "class" to JsonPrimitive("priest"), "spec" to JsonPrimitive("Shadow"),
                "race" to JsonPrimitive("undead"), "level" to JsonPrimitive(23), "source" to JsonPrimitive("bnet"),
            ),
        )
        val characters = JsonArray(requireNotNull(root["characters"]).jsonArray + extra)
        return ApiJson.decodeFromJsonElement(Me.serializer(), JsonObject(root + ("characters" to characters)))
    }

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
        SignedOut(AccountCopy.notice(SignedOutNotice.AccountDeleted(Instant.parse("2026-11-29T21:41:00Z"))))
    }

    @Test fun `no characters`() = tab("home-no-characters") { NoCharactersScreen {} }

    @Test fun `signed out at font scale 2`() {
        RuntimeEnvironment.setFontScale(2f)
        tab("signedout-xl") { SignedOut(null) }
    }

    @Test fun `switcher`() = capture("switcher-default") { Switcher(threeCharacters()) }

    @Test fun `switcher at font scale 2`() {
        RuntimeEnvironment.setFontScale(2f)
        capture("switcher-xl") { Switcher(threeCharacters()) }
    }

    @Test fun `you free`() = tab("you-free", ForeverTab.Account) { You(me("me.json")) }

    @Test fun `you premium`() = tab("you-premium", ForeverTab.Account) { You(me("me.premium.json")) }

    @Test fun `sign out confirmation`() = screen("you-signout-confirm") { You(me("me.json"), confirming = true) }

    @Test fun `devices`() = tab("devices-default", ForeverTab.Account) {
        DevicesScreen(DevicesState.Loaded(deviceRows()), null, {}, {})
    }

    @Test fun `revoke confirmation`() {
        val rows = deviceRows()
        screen("devices-revoke-confirm") { DevicesScreen(DevicesState.Loaded(rows), null, {}, {}, initiallyConfirming = rows[1]) }
    }

    @Test fun `revoke failed`() {
        val rows = deviceRows()
        tab("devices-revoke-failed", ForeverTab.Account) { DevicesScreen(DevicesState.Loaded(rows), rows[1].id, {}, {}) }
    }

    @Test fun `delete account`() = tab("delete-default", ForeverTab.Account) { DeleteAccountScreen(false, false, false, {}) }

    @Test fun `delete account with a plan, failed`() = tab("delete-plan-failed", ForeverTab.Account) { DeleteAccountScreen(true, false, true, {}) }

    @Test fun `delete confirmation`() = screen("delete-confirm") { DeleteAccountScreen(false, false, false, {}, initiallyConfirming = true) }

    @Composable
    private fun SignedOut(notice: String?) = SignedOutScreen(notice, signingIn = false, onSignIn = {}, onPickClass = {})

    @Composable
    private fun You(me: Me, confirming: Boolean = false) = YouScreen(
        summary = YouSummary(
            battletag = me.user.battletag.orEmpty(),
            plan = AccountCopy.planLine(me.entitlements.billing),
            founder = me.entitlements.billing?.founder == true,
            devices = "This phone and 2 more",
        ),
        onDevices = {},
        onDeleteAccount = {},
        onSignOut = {},
        confirmingSignOut = confirming,
    )

    /** The selector open, and the sheet drawn as it presents. */
    @Composable
    private fun Switcher(me: Me) {
        val current = me.characters.first()
        Column(Modifier.fillMaxSize().background(FsColor.bg)) {
            Column(Modifier.padding(16.dp)) { CharacterSelector(current, open = true) {} }
            Column(Modifier.clip(RoundedCornerShape(topStart = 28.dp, topEnd = 28.dp)).background(FsColor.raised)) {
                SwitcherSheetContent(me, current.key) {}
            }
        }
    }

    private fun tab(name: String, tab: ForeverTab = ForeverTab.Character, screen: @Composable () -> Unit) =
        capture(name) { ForeverTabScaffold(selected = tab, onSelect = {}) { screen() } }

    private fun capture(name: String, screen: @Composable () -> Unit) {
        compose.setContent { ForeverTheme { screen() } }
        compose.onRoot().captureRoboImage("src/test/screenshots/account_$name.png", roborazziOptions = TOLERANT)
    }

    /** Dialogs live in their own window, so these capture the whole screen. */
    @OptIn(ExperimentalRoborazziApi::class)
    private fun screen(name: String, content: @Composable () -> Unit) {
        compose.setContent { ForeverTheme { ForeverTabScaffold(selected = ForeverTab.Account, onSelect = {}) { content() } } }
        compose.waitForIdle()
        captureScreenRoboImage("src/test/screenshots/account_$name.png", roborazziOptions = TOLERANT)
    }
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(
    compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F),
)

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
