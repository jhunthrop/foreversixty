package gg.foreversixty.app.system

import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Test
import java.io.File

/**
 * Labels and headings are sentence case on native (spec 2026-11-07 ruling N1; owner 2026-11-10: the
 * apps set caps where the site does not). Outside core/designsystem no source may force capitals:
 * a code written in capitals (a region, a country, a metric unit) goes through the design system's
 * `Codes`, so the one place that uppercases is reviewed once. iOS: `CaseGuardTests.swift`.
 */
class CaseGuardTest {
    @Test
    fun `no source outside the design system forces capitals`() {
        val offenders = sources().toList().flatMap { file ->
            file.readLines().mapIndexedNotNull { index, line ->
                FORCED_CAPS.find(line)?.let { "${file.relativeTo(android)}:${index + 1}: ${it.value}" }
            }
        }
        assertEquals("Forced capitals (use sentence case, or Codes for a code):\n" + offenders.joinToString("\n"), ALLOWED, offenders.toList())
    }

    @Test
    fun `the guard sees the app's sources`() {
        assertTrue(sources().count() > MIN_SOURCES)
    }

    private val android: File = checkNotNull(File(System.getProperty("user.dir").orEmpty()).absoluteFile.parentFile) { "the app module has a parent" }

    private fun sources(): Sequence<File> = SCANNED.asSequence()
        .map { File(android, it) }
        .flatMap { it.walkTopDown() }
        .filter { it.isFile && it.path.contains("/src/main/") && (it.extension == "kt" || it.extension == "xml") }
        .filterNot { it.path.contains(DESIGN_SYSTEM) || it.path.contains("/build/") }

    private companion object {
        /** The app module and the core modules, main sources only; core/designsystem owns `Codes`. */
        val SCANNED = listOf("app/src/main", "core")
        const val DESIGN_SYSTEM = "core/designsystem/"

        /** Kotlin's and Java's uppercasing, the view system's caps attribute, Compose's caps transforms. */
        val FORCED_CAPS = Regex("""\.uppercase\(|toUpperCase\(|textAllCaps|AllCaps|\.uppercaseChar\(|capitalize\(""")

        /** Empty on purpose: a new entry needs the owner's reason. */
        val ALLOWED = emptyList<String>()

        const val MIN_SOURCES = 200
    }
}
