package gg.foreversixty.app.common

import android.content.Context
import android.view.HapticFeedbackConstants
import android.view.View
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.raids.RecordingHaptics
import org.junit.Assert.assertEquals
import org.junit.Assert.assertSame
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

/** The one haptics facade (W3 §5.7 Haptics): the system constants per moment and the injection seam. */
@RunWith(AndroidJUnit4::class)
class HapticsTest {
    @get:Rule
    val compose = createComposeRule()

    /** A view that records the constants it is asked to play. */
    private class RecordingView(context: Context) : View(context) {
        val played = mutableListOf<Int>()
        override fun performHapticFeedback(feedbackConstant: Int): Boolean {
            played += feedbackConstant
            return true
        }
    }

    @Test fun `each moment plays its system constant`() {
        val view = RecordingView(ApplicationProvider.getApplicationContext())
        val haptics = ViewHaptics(view)
        haptics.selection()
        haptics.impactMedium()
        haptics.success()
        haptics.error()
        assertEquals(
            listOf(
                HapticFeedbackConstants.SEGMENT_TICK,
                HapticFeedbackConstants.LONG_PRESS,
                HapticFeedbackConstants.CONFIRM,
                HapticFeedbackConstants.REJECT,
            ),
            view.played,
        )
    }

    @Test fun `a provided facade replaces the system one`() {
        val recording = RecordingHaptics()
        var used: Haptics? = null
        compose.setContent {
            CompositionLocalProvider(LocalHaptics provides recording) { used = rememberHaptics() }
        }
        compose.waitForIdle()
        assertSame(recording, used)
    }

    @Test fun `with nothing provided the system facade plays on the view`() {
        var used: Haptics? = null
        compose.setContent { used = rememberHaptics() }
        compose.waitForIdle()
        assertEquals(ViewHaptics::class, used?.let { it::class })
    }
}
