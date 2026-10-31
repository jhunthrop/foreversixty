package gg.foreversixty.app

import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import org.junit.Assert.assertEquals
import org.junit.Test
import org.junit.runner.RunWith
import android.content.Context

/** The generated strings.xml (apps/shared/copy/en.json) formats placeholders and plurals. */
@RunWith(AndroidJUnit4::class)
class CopyResourcesTest {
    private val context: Context = ApplicationProvider.getApplicationContext()

    @Test
    fun `a plain key reads its text and the platform word`() {
        assertEquals("Home", context.getString(R.string.tab_home))
        assertEquals("This phone", context.getString(R.string.devices_this))
    }

    @Test
    fun `placeholders fill in order`() {
        assertEquals("24 of 40 signed up", context.getString(R.string.raid_count, 24, 40))
        assertEquals("You wear this · ", context.getString(R.string.tooltip_youWear))
        assertEquals("Back to Thrall's list", context.getString(R.string.bis_picker_back, "Thrall"))
    }

    @Test
    fun `plurals pick their form from the count`() {
        val res = context.resources
        assertEquals("1 point unspent.", res.getQuantityString(R.plurals.talents_unspent, 1, 1))
        assertEquals("3 points unspent.", res.getQuantityString(R.plurals.talents_unspent, 3, 3))
    }
}
