package gg.foreversixty.app.version

import gg.foreversixty.app.BuildConfig
import org.junit.Assert.assertEquals
import org.junit.Test
import java.io.File

/** The built app carries exactly the version in apps/VERSION, the file both apps read. */
class AppVersionTest {
    /** Unit tests run in apps/android/app, so the file is two levels up. */
    private val settings: Map<String, String> = File("../../VERSION").readLines()
        .map { it.trim() }
        .filter { it.isNotEmpty() && !it.startsWith("//") }
        .associate { line -> line.substringBefore("=").trim() to line.substringAfter("=").trim() }

    @Test
    fun `versionName and versionCode come from apps VERSION`() {
        assertEquals(settings["APP_VERSION"], BuildConfig.VERSION_NAME)
        assertEquals(settings["APP_BUILD"]?.toInt(), BuildConfig.VERSION_CODE)
    }
}
