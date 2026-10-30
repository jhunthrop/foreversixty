package gg.foreversixty.app.home

import androidx.compose.runtime.Composable
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import kotlinx.serialization.json.jsonObject
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * The placeholder home inside the tab shell, at Pixel 9 size, from the contract's me fixture.
 * Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class HomeScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private val me: Me = run {
        val raw = requireNotNull(javaClass.classLoader?.getResource("fixtures/app/me.json")).readText()
        val root = ApiJson.parseToJsonElement(raw).jsonObject
        val data = root["data"] ?: root
        ApiJson.decodeFromJsonElement(Me.serializer(), data)
    }

    @Test
    fun `home with the main character`() = capture("home_character") { HomeScreen(me.mainCharacter) }

    @Test
    fun `home with no character`() = capture("home_empty") { HomeScreen(null) }

    private fun capture(name: String, screen: @Composable () -> Unit) {
        compose.setContent {
            ForeverTheme {
                ForeverTabScaffold(selected = ForeverTab.Character, onSelect = {}) { screen() }
            }
        }
        compose.onRoot().captureRoboImage("src/test/screenshots/$name.png", roborazziOptions = TOLERANT)
    }
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(
    compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F),
)

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
