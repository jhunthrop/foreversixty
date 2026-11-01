package gg.foreversixty.app.common

import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertTrue
import org.junit.Test

/** The new-data rule (W3 §5.7): the first build is shown; a newer one waits for Update. */
class BuildWatchTest {
    private var latest: String? = "1.60.1.70532"
    private val watch = BuildWatch { latest }

    @Test fun `the build on open is what is shown, with no news`() = runTest {
        watch.check()
        assertEquals("1.60.1.70532", watch.shown)
        assertFalse(watch.hasNewData)
    }

    @Test fun `a newer build waits until Update`() = runTest {
        watch.check()
        latest = "1.60.1.70658"
        watch.check()
        assertTrue(watch.hasNewData)
        assertEquals("1.60.1.70532", watch.shown)
        watch.accept()
        assertEquals("1.60.1.70658", watch.shown)
        assertFalse(watch.hasNewData)
    }

    @Test fun `offline (no fresh answer) changes nothing`() = runTest {
        watch.check()
        latest = null
        watch.check()
        assertFalse(watch.hasNewData)
        assertEquals("1.60.1.70532", watch.shown)
    }
}
