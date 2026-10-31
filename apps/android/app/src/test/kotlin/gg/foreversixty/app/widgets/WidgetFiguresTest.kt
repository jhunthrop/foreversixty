package gg.foreversixty.app.widgets

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.widgets.logic.RaidFigures
import gg.foreversixty.app.widgets.logic.RaidSurfaceState
import gg.foreversixty.app.widgets.logic.SignupOverlay
import gg.foreversixty.app.widgets.logic.WidgetFigures
import java.time.Instant
import java.time.ZoneId
import kotlinx.serialization.json.JsonNull
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.JsonPrimitive
import kotlinx.serialization.json.booleanOrNull
import kotlinx.serialization.json.jsonArray
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** Every surface's words against the shared golden vectors, so Android and iOS agree to the letter. */
@RunWith(AndroidJUnit4::class)
class WidgetFiguresTest {
    private fun JsonObject.text(key: String): String? = (get(key) as? JsonPrimitive)?.takeUnless { it is JsonNull }?.content

    private fun expectations(f: RaidFigures): Map<String, String?> = mapOf(
        "state" to f.state.name.replaceFirstChar { it.lowercase() },
        "tag" to f.tag,
        "tagKind" to f.tagKind?.name?.replaceFirstChar { it.lowercase() },
        "countdown" to f.countdown,
        "short" to f.short,
        "time" to f.time,
        "timeShort" to f.timeShort,
        "inline" to f.inline,
        "you" to f.you,
        "signedUp" to f.signedUp,
        "stale" to f.stale,
        "liveTitle" to f.liveTitle,
        "liveUntil" to f.liveUntil,
    )

    @Test
    fun `every raid vector matches`() {
        for (case in WidgetTestData.figures.getValue("cases").jsonArray.map { it.jsonObject }) {
            val name = case.text("name")
            val zone = ZoneId.of(case.text("zone"))
            val snapshot = WidgetTestData.snapshot(
                fixture = case.text("snapshot").orEmpty(),
                fetchedAt = case.text("fetched_at").orEmpty(),
                patch = case["patch"]?.jsonObject,
            )
            val figures = WidgetFigures(resources(), zone).raid(snapshot, Instant.parse(case.text("now")))
            val expect = case.getValue("expect").jsonObject
            val got = expectations(figures)
            for ((key, value) in expect) {
                when (key) {
                    "buttons" -> assertEquals("$name: buttons", value.jsonPrimitive.booleanOrNull, figures.buttons)
                    "liveActive" -> assertEquals("$name: liveActive", value.jsonPrimitive.booleanOrNull, figures.liveActive)
                    else -> assertEquals("$name: $key", expect.text(key), got.getValue(key))
                }
            }
        }
    }

    @Test
    fun `the upgrade vectors match`() {
        val block = WidgetTestData.figures.getValue("upgrades").jsonObject
        val expect = block.getValue("expect").jsonObject
        val figures = WidgetFigures(resources()).upgrades(WidgetTestData.snapshot(block.text("snapshot").orEmpty()))
        assertEquals(expect.text("header"), figures.header)
        assertEquals(
            expect.getValue("rows").jsonArray.map { row -> row.jsonArray.map { it.jsonPrimitive.content } },
            figures.rows.map { listOf(it.name, it.gain) },
        )
        assertEquals(expect.text("footer"), figures.footer)
        assertEquals(expect.text("rectTitle"), figures.rectTitle)
        val noGear = block.getValue("noGear").jsonObject
        assertEquals(noGear.text("sentence"), WidgetFigures(resources()).upgrades(WidgetTestData.snapshot(noGear.text("snapshot").orEmpty())).sentence)
    }

    @Test
    fun `signed out and no Home answer have words, never an empty box`() {
        val figures = WidgetFigures(resources())
        val out = figures.raid(null, WidgetTestData.boardNow)
        assertEquals(RaidSurfaceState.SignedOut, out.state)
        assertEquals("Sign in to see your next raid.", out.title)
        val noHome = WidgetTestData.snapshot().let { it.copy(widget = it.widget.copy(upgrades = null)) }
        val upgrades = figures.upgrades(noHome)
        assertTrue(upgrades.opensApp)
        assertEquals("Open the app", upgrades.sentence)
    }

    @Test
    fun `no raid planned names the last raid`() {
        val base = WidgetTestData.snapshot()
        val snapshot = base.copy(widget = base.widget.copy(nextRaid = null))
        val figures = WidgetFigures(resources(), ZoneId.of("America/Chicago")).raid(snapshot, WidgetTestData.boardNow)
        assertEquals(RaidSurfaceState.NoRaid, figures.state)
        assertEquals("No raid planned", figures.title)
        assertEquals("Last: Onyxia's Lair, 3\u00A0Dec.", figures.subtitle)
    }

    @Test
    fun `a pending press is checked at once and a refused one says Not saved`() {
        val figures = WidgetFigures(resources(), ZoneId.of("America/Chicago"))
        val pending = figures.raid(WidgetTestData.snapshot(), WidgetTestData.boardNow, SignupOverlay(7, "late"))
        assertEquals(SignupState.Late, pending.checked)
        assertEquals("Late", pending.tag)
        val failed = figures.raid(
            WidgetTestData.snapshot(), WidgetTestData.boardNow,
            SignupOverlay(7, "out", failedUntil = WidgetTestData.boardNow.plusSeconds(10).toString()),
        )
        assertEquals(SignupState.In, failed.checked)
        assertEquals(SignupState.Out, failed.notSaved)
        val expired = figures.raid(
            WidgetTestData.snapshot(), WidgetTestData.boardNow.plusSeconds(11),
            SignupOverlay(7, "out", failedUntil = WidgetTestData.boardNow.plusSeconds(10).toString()),
        )
        assertNull(expired.notSaved)
    }

    @Test
    fun `the widget reads as one sentence to TalkBack`() {
        val figures = WidgetFigures(resources(), ZoneId.of("America/Chicago")).raid(WidgetTestData.snapshot(), WidgetTestData.boardNow)
        assertEquals(
            "Next raid, Onyxia's Lair, Thursday 8 pm, in 2 days 2 hours. You are In as DPS. 24 of 40 signed up.",
            figures.description,
        )
        assertFalse(figures.asksSpec)
    }
}
