package gg.foreversixty.app.testing

import androidx.compose.ui.semantics.SemanticsNode
import androidx.compose.ui.semantics.SemanticsProperties
import androidx.compose.ui.semantics.getOrNull
import androidx.compose.ui.test.junit4.ComposeContentTestRule
import androidx.compose.ui.test.onRoot
import org.junit.Assert.assertEquals
import org.junit.Assert.fail
import java.io.File

/**
 * The native parity gate's text side (design/parity.md "Native"): a state's words, strictly. The
 * unmerged semantics tree's texts and content descriptions in reading order are a text golden
 * beside the image golden, so a copy change fails until it is re-recorded on purpose
 * (`recordRoborazziDebug` records both).
 */
object ParityText {
    /** Roborazzi's record switch, which `recordRoborazziDebug` sets for the test JVM. */
    private const val RECORD_PROPERTY = "roborazzi.test.record"

    /** The texts and content descriptions under the root, depth first, one line each. */
    fun lines(rule: ComposeContentTestRule): List<String> =
        lines(rule.onRoot(useUnmergedTree = true).fetchSemanticsNode())

    fun lines(node: SemanticsNode): List<String> {
        val own = buildList {
            node.config.getOrNull(SemanticsProperties.ContentDescription)?.let { addAll(it) }
            node.config.getOrNull(SemanticsProperties.Text)?.let { texts -> addAll(texts.map { it.text }) }
        }.map { it.trim() }.filter { it.isNotEmpty() }
        return own + node.children.flatMap { lines(it) }
    }

    /** Asserts the state's words against `src/test/screenshots/<path>.txt`, or records them. */
    fun assertText(rule: ComposeContentTestRule, path: String) {
        val text = lines(rule).joinToString("\n", postfix = "\n")
        val golden = File("src/test/screenshots/$path.txt")
        if (System.getProperty(RECORD_PROPERTY) == "true") {
            golden.parentFile?.mkdirs()
            golden.writeText(text)
            return
        }
        if (!golden.exists()) fail("no text golden at $golden; record it with ./gradlew :app:recordRoborazziDebug")
        assertEquals("the words of $path changed; re-record on purpose if they should", golden.readText(), text)
    }
}
