package gg.foreversixty.app.testing

import androidx.compose.ui.geometry.Rect
import gg.foreversixty.app.designsystem.ParityItem
import gg.foreversixty.app.designsystem.ParityRecorder
import gg.foreversixty.app.model.ApiJson
import kotlinx.serialization.json.float
import kotlinx.serialization.json.int
import kotlinx.serialization.json.jsonArray
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive
import java.io.File
import kotlin.math.abs
import kotlin.math.max
import kotlin.math.roundToInt

/**
 * The native parity gate's region side (design/parity.md "Native"): the regions and inventory a
 * screen's primitives report (`Modifier.parityRegion`, `parityItem`, through [ParityRecorder]),
 * compared with the mock's region file (design/mocks/app/regions/android-<screen>.json) on the
 * web gate's rules.
 */
object ParityRegions {
    /** x and y within 8 dp; width and height within 8 dp or 5%, whichever is larger. */
    const val POSITION_TOLERANCE = 8f
    const val SIZE_TOLERANCE = 8f
    const val SIZE_RATIO_TOLERANCE = 0.05f

    /** The test task runs in apps/android/app. */
    private val regionsDirectory = File("../../../design/mocks/app/regions")

    data class Region(val box: List<Float>, val inventory: Map<String, Int>)

    data class Reference(val width: Float, val height: Float, val regions: Map<String, Region>)

    /** What a render reported, in dp. */
    data class Capture(val regions: Map<String, Rect>, val items: List<ParityItem>)

    fun reference(screen: String, platform: String = "android"): Reference =
        parse(File(regionsDirectory, "$platform-$screen.json").readText())

    fun parse(json: String): Reference {
        val root = ApiJson.parseToJsonElement(json).jsonObject
        val device = root.getValue("device").jsonObject
        val regions = root.getValue("regions").jsonObject.mapValues { (_, value) ->
            val region = value.jsonObject
            Region(
                box = region.getValue("box").jsonArray.map { it.jsonPrimitive.float },
                inventory = region.getValue("inventory").jsonObject.mapValues { it.value.jsonPrimitive.int },
            )
        }
        return Reference(device.getValue("width").jsonPrimitive.float, device.getValue("height").jsonPrimitive.float, regions)
    }

    /** The recorder's px boxes in dp. */
    fun capture(recorder: ParityRecorder, density: Float): Capture {
        fun dp(rect: Rect) = Rect(rect.left / density, rect.top / density, rect.right / density, rect.bottom / density)
        return Capture(
            regions = recorder.regions.mapValues { dp(it.value) },
            items = recorder.items.map { it.copy(bounds = dp(it.bounds)) },
        )
    }

    /** Every way the capture breaks the reference; empty means it holds. */
    fun compare(capture: Capture, reference: Reference): List<String> {
        val diffs = mutableListOf<String>()
        (reference.regions.keys - capture.regions.keys).sorted().forEach { diffs += "$it: in the mock, missing from the build" }
        (capture.regions.keys - reference.regions.keys).sorted().forEach { diffs += "$it: in the build, not in the mock" }
        for ((name, want) in reference.regions.toSortedMap()) {
            val got = capture.regions[name] ?: continue
            if (want.box.size != BOX_FIELDS) continue
            diffs += boxDiffs(name, got, want.box)
            val counted = inventory(name, got, capture.items)
            if (counted != want.inventory) diffs += "$name: inventory ${describe(counted)}, the mock has ${describe(want.inventory)}"
        }
        return diffs
    }

    /**
     * The items whose centre lies inside [region], plus every item of the scrolling region [name]
     * wherever its scroll puts it, counted by kind (kinds with none left out).
     */
    fun inventory(name: String, region: Rect, items: List<ParityItem>): Map<String, Int> =
        items.filter { it.scroll == name || (it.scroll == null && region.contains(it.bounds.center)) }
            .groupingBy { it.kind.key }.eachCount()

    private fun boxDiffs(name: String, got: Rect, want: List<Float>): List<String> {
        val checks = listOf(
            Check("x", got.left, want[0], POSITION_TOLERANCE),
            Check("y", got.top, want[1], POSITION_TOLERANCE),
            Check("width", got.width, want[2], max(SIZE_TOLERANCE, want[2] * SIZE_RATIO_TOLERANCE)),
            Check("height", got.height, want[3], max(SIZE_TOLERANCE, want[3] * SIZE_RATIO_TOLERANCE)),
        )
        return checks.filter { abs(it.value - it.target) > it.tolerance }.map {
            "$name: ${it.field} ${it.value.roundToInt()}, the mock has ${it.target.roundToInt()} (within ${it.tolerance.roundToInt()})"
        }
    }

    private fun describe(counts: Map<String, Int>): String =
        if (counts.isEmpty()) "{}" else counts.toSortedMap().entries.joinToString(", ") { "${it.key} ${it.value}" }

    private data class Check(val field: String, val value: Float, val target: Float, val tolerance: Float)

    private const val BOX_FIELDS = 4
}
