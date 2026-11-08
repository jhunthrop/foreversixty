package gg.foreversixty.app.designsystem

import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.remember
import androidx.compose.runtime.staticCompositionLocalOf
import androidx.compose.ui.Modifier
import androidx.compose.ui.composed
import androidx.compose.ui.geometry.Rect
import androidx.compose.ui.layout.LayoutCoordinates
import androidx.compose.ui.layout.boundsInRoot
import androidx.compose.ui.layout.onGloballyPositioned
import androidx.compose.ui.platform.testTag

/**
 * The native parity regions (design/parity.md "Native"): the web's `data-region` names
 * (web/src/lib/ui/regions.ts `REGION`, held equal by web/src/lib/mobile/parity-regions.test.ts),
 * so a region of a mock pairs with the same region of the build.
 */
object ParityRegion {
    const val HEAD = "head"
    const val HEAD_TITLE = "head-title"
    const val HEAD_META = "head-meta"
    const val HEAD_ANSWER = "head-answer"
    const val CONTROLS = "controls"
    const val SUMMARY = "summary"
    const val COUNTDOWNS = "countdowns"
    const val NOTES = "notes"
    const val PLATE = "plate"
    const val SLOTS = "slots"
    const val INSPECTOR = "inspector"
    const val INSPECTOR_INLINE = "inspector-inline"
    const val FOOTER_NOTE = "footer-note"
    const val INTRO = "intro"
    const val DIALOG = "dialog"

    /** The semantics test tag prefix a region carries: `region:head`. */
    const val TAG_PREFIX = "region:"

    /** A Section's region, `section-<slug of its title>`: the web's `sectionRegion` rule. */
    fun section(title: String): String =
        "section-" + title.lowercase().split(Regex("[^a-z0-9]+")).filter { it.isNotEmpty() }.joinToString("-")
}

/** What a region's inventory counts. */
enum class ParityItemKind(val key: String) { Row("row"), Button("button"), Segment("segment"), Field("field"), Image("image") }

/**
 * One counted element and its bounds in the content root, in px. [scroll] names the scrolling region it
 * sits in (a filter row): it counts there wherever the scroll has moved it, as the web counts a
 * scroller's content, not what shows.
 */
data class ParityItem(val kind: ParityItemKind, val bounds: Rect, val scroll: String? = null)

/**
 * Collects the regions and items a screen reports while a parity test renders it; absent (null)
 * everywhere else, so the modifiers cost nothing in the app.
 */
class ParityRecorder {
    val regions: MutableMap<String, Rect> = linkedMapOf()
    private val itemsByElement: MutableMap<Any, Pair<ParityItem, LayoutCoordinates>> = linkedMapOf()
    private var root: LayoutCoordinates? = null

    /**
     * Every counted element, once each however often it was laid out, measured against the content
     * root when read: an element placed before the root (the profile button above the page) is not
     * left in window coordinates.
     */
    val items: List<ParityItem>
        get() = itemsByElement.values.map { (item, coordinates) ->
            if (coordinates.isAttached) item.copy(bounds = boundsOf(coordinates)) else item
        }

    internal fun item(element: Any, item: ParityItem, coordinates: LayoutCoordinates) {
        itemsByElement[element] = item to coordinates
    }

    internal fun root(coordinates: LayoutCoordinates) {
        root = coordinates
    }

    internal fun boundsOf(coordinates: LayoutCoordinates): Rect =
        root?.takeIf { it.isAttached }?.localBoundingBoxOf(coordinates, clipBounds = false) ?: coordinates.boundsInRoot()
}

val LocalParityRecorder = staticCompositionLocalOf<ParityRecorder?> { null }

/** The scrolling region the items below belong to ([ParityScrollRegion]); null outside one. */
val LocalParityScroll = staticCompositionLocalOf<String?> { null }

/** Marks [content] as the scrolling region [name]'s content: its items count in that region's inventory. */
@Composable
fun ParityScrollRegion(name: String, content: @Composable () -> Unit) {
    CompositionLocalProvider(LocalParityScroll provides name, content = content)
}

/** Marks the screen's content (under the top app bar) as the space region boxes are measured in. */
fun Modifier.parityContent(): Modifier = composed {
    val recorder = LocalParityRecorder.current
    if (recorder == null) Modifier else Modifier.onGloballyPositioned { recorder.root(it) }
}

/** Reports this element's bounds as the named region, and tags it `region:<name>`. */
fun Modifier.parityRegion(name: String): Modifier = testTag(ParityRegion.TAG_PREFIX + name).composed {
    val recorder = LocalParityRecorder.current
    if (recorder == null) Modifier else Modifier.onGloballyPositioned { recorder.regions[name] = recorder.boundsOf(it) }
}

/** Counts this element in the inventory of every region that contains it. */
fun Modifier.parityItem(kind: ParityItemKind): Modifier = composed {
    val recorder = LocalParityRecorder.current
    if (recorder == null) {
        Modifier
    } else {
        val element = remember { Any() }
        val scroll = LocalParityScroll.current
        Modifier.onGloballyPositioned { coordinates ->
            recorder.item(element, ParityItem(kind, recorder.boundsOf(coordinates), scroll), coordinates)
        }
    }
}
