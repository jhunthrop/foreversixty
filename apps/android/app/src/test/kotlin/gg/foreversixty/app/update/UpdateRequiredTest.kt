package gg.foreversixty.app.update

import android.content.Context
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.MaterialTheme
import androidx.compose.ui.Modifier
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.onRoot
import androidx.compose.ui.test.performClick
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.designsystem.ForeverTheme
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNotEquals
import org.junit.Assert.assertTrue
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode
import java.util.UUID

/**
 * The blocking update screen at Pixel 9 size, and the install id behind `X-App-Install`.
 * Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav")
class UpdateRequiredTest {
    @get:Rule
    val compose = createComposeRule()

    @Test
    fun `the update screen opens the store`() {
        var opened = 0
        compose.setContent {
            ForeverTheme {
                Box(Modifier.fillMaxSize().background(MaterialTheme.colorScheme.background)) {
                    UpdateRequiredScreen(onOpenStore = { opened++ })
                }
            }
        }
        compose.onRoot().captureRoboImage(
            "src/test/screenshots/update_required.png",
            roborazziOptions = RoborazziOptions(
                compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F),
            ),
        )
        compose.onNodeWithText("Open Google Play").performClick()
        assertEquals(1, opened)
    }

    @Test
    fun `the install id is a UUID made once, and a cleared store makes a new one`() {
        val prefs = ApplicationProvider.getApplicationContext<Context>()
            .getSharedPreferences(InstallId.PREFS, Context.MODE_PRIVATE)
        val first = InstallId.current(prefs)
        assertTrue(runCatching { UUID.fromString(first) }.isSuccess)
        assertEquals(first, InstallId.current(prefs))
        prefs.edit().clear().commit()
        assertNotEquals(first, InstallId.current(prefs))
    }
}
