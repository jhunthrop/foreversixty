package gg.foreversixty.app.premium

import androidx.compose.runtime.Composable
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.account.CharacterLine
import gg.foreversixty.app.account.YouActions
import gg.foreversixty.app.account.YouScreen
import gg.foreversixty.app.account.YouSummary
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.network.ClasspathFixtureSource
import kotlinx.serialization.json.jsonObject
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode
import java.time.Instant

/**
 * F8's states at Pixel 9 size (spec 5.8.1), the same list as iOS PremiumSnapshotTests: You free
 * in a link region and elsewhere, a seat, the Plan screen in both regions and for a seat, and
 * You and Plan at font scale 2.0. The same images are the review captures.
 * Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class PremiumScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private val source = ClasspathFixtureSource(requireNotNull(javaClass.classLoader))

    private fun me(name: String): Me {
        val text = requireNotNull(source.read(name)) { "no fixture $name" }
        val data = requireNotNull(ApiJson.parseToJsonElement(text).jsonObject["data"]) { "no data in $name" }
        return ApiJson.decodeFromJsonElement(Me.serializer(), data)
    }

    @Test fun `you free in a link region`() = capture("you-free-us") { You("me.json", mayLink = true) }

    @Test fun `you free elsewhere`() = capture("you-free-gb") { You("me.json", mayLink = false) }

    @Test fun `you with a seat`() = capture("you-seat") { You("me.seat.json", mayLink = false) }

    @Test fun `you free in a link region at font scale 2`() {
        RuntimeEnvironment.setFontScale(2f)
        capture("you-free-us-ax5") { You("me.json", mayLink = true) }
    }

    @Test fun `you premium at font scale 2`() {
        RuntimeEnvironment.setFontScale(2f)
        capture("you-premium-ax5") { You("me.premium.json", mayLink = true) }
    }

    @Test fun `plan in a link region`() = capture("plan-us") { Plan("me.premium.json", mayLink = true) }

    @Test fun `plan elsewhere`() = capture("plan-gb") { Plan("me.premium.json", mayLink = false) }

    @Test fun `plan for a seat`() = capture("plan-seat") { Plan("me.seat.json", mayLink = true) }

    @Test fun `plan at font scale 2`() {
        RuntimeEnvironment.setFontScale(2f)
        capture("plan-us-ax5") { Plan("me.premium.json", mayLink = true) }
    }

    @Composable
    private fun You(file: String, mayLink: Boolean) {
        val me = me(file)
        val current = me.characters.first { it.key == me.mainCharacterKey }
        YouScreen(
            summary = YouSummary(
                battletag = me.user.battletag.orEmpty(),
                plan = PlanSummary.of(me, NOW),
                character = CharacterLine(current.name, current.className, "US · PvP"),
                devices = "This phone and 2 more",
                version = VERSION,
            ),
            actions = YouActions({}, {}, {}, {}, {}, {}, {}, {}),
            mayLink = mayLink,
        )
    }

    @Composable
    private fun Plan(file: String, mayLink: Boolean) = PlanScreen(PlanSummary.of(me(file), NOW), mayLink) {}

    private fun capture(name: String, screen: @Composable () -> Unit) {
        compose.setContent { ForeverTheme { ForeverTabScaffold(selected = ForeverTab.You, onSelect = {}) { screen() } } }
        compose.onRoot().captureRoboImage("src/test/screenshots/premium_$name.png", roborazziOptions = TOLERANT)
    }
}

private val NOW: Instant = Instant.parse("2025-11-01T00:00:00Z")

private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))

private const val VERSION = "Version 0.1.0 (1) · data 1.60.1.70658"

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
