package gg.foreversixty.app.gear.flows

import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.bis.BisPhase
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** Gear with no list for the spec and band (round 2): the list line with Change and one row, never a blank page. */
@RunWith(AndroidJUnit4::class)
class GearNoticeTest {
    private val res = ApplicationProvider.getApplicationContext<android.content.Context>().resources
    private val composer = GearComposer(res)

    private fun notice(phase: BisPhase) = composer.notice(phase, "Beast Mastery Hunter", "30 to 39", "horde")

    @Test fun `no published list says so under the list line, with no retry`() {
        assertEquals(
            GearNotice("Beast Mastery Hunter", "30 to 39 · Horde", "No best in slot list published yet for this spec and level range."),
            notice(BisPhase.NoList(offlineSinceMillis = null)),
        )
    }

    @Test fun `a list that did not load says so with Try again`() {
        val failed = requireNotNull(notice(BisPhase.Failed))
        assertEquals("We could not load this list", failed.row)
        assertEquals("Check your connection", failed.sub)
        assertTrue(failed.retry)
        val notCached = requireNotNull(notice(BisPhase.NotCached))
        assertEquals("We could not load this list", notCached.row)
        assertTrue(notCached.retry)
    }

    @Test fun `loading has no notice`() {
        assertNull(notice(BisPhase.Loading))
    }
}
