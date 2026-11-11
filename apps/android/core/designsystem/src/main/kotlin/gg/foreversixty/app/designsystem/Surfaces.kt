package gg.foreversixty.app.designsystem

import androidx.annotation.DrawableRes
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Check
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.composed
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.draw.drawWithContent
import androidx.compose.ui.geometry.CornerRadius
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.geometry.Rect
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.graphics.BlendMode
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.ImageBitmap
import androidx.compose.ui.graphics.Paint
import androidx.compose.ui.graphics.Shape
import androidx.compose.ui.graphics.drawscope.DrawScope
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.graphics.drawscope.clipRect
import androidx.compose.ui.graphics.drawscope.drawIntoCanvas
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.res.imageResource
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.IntOffset
import androidx.compose.ui.unit.IntSize
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.em
import java.util.Locale
import kotlin.math.max
import kotlin.math.roundToInt

/** The site's uppercase section label (`micro`, .14em), read as a heading. */
@Composable
fun MicroLabel(text: String, modifier: Modifier = Modifier, color: Color = FsColor.muted) {
    Text(text.uppercase(Locale.ROOT), modifier.semantics { heading() }, style = ScaledType.micro, color = color)
}

/** A tag (`.pill`): 22 dp, 11 sp 700 uppercase, tinted ground and border; [checked] leads with a tick. */
@Composable
fun Tag(text: String, color: Color, modifier: Modifier = Modifier, checked: Boolean = false) {
    val shape = RoundedCornerShape(FsRadius.legacyPill)
    Row(
        modifier
            .heightIn(min = 22.dp)
            .background(color.copy(alpha = 0.12f), shape)
            .border(1.dp, color.copy(alpha = 0.35f), shape)
            .padding(horizontal = FsSpace.space2, vertical = 4.dp),
        horizontalArrangement = Arrangement.spacedBy(4.dp),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        if (checked) Icon(Icons.Filled.Check, contentDescription = null, tint = color, modifier = Modifier.size(11.dp))
        Text(text.uppercase(Locale.ROOT), style = ScaledType.micro.copy(letterSpacing = 0.08.em), color = color)
    }
}

/** A chip or pill's ground: [fill] and a 1 dp [line] border in [shape]. */
fun Modifier.pillSurface(fill: Color, line: Color, shape: Shape = RoundedCornerShape(percent = 50)): Modifier =
    background(fill, shape).border(1.dp, line, shape)

/**
 * The round-5 art (apps/shared/brand/art, tools/art/build_art.py), copied to res/drawable-nodpi by
 * the forever.shared-fonts plugin: the ground tile, the band behind the bars, Home's hero and the
 * raid card's cave. Background only, never captioned (owner, 2026-11-09). iOS: `BrandArt`.
 */
enum class BrandArt(@DrawableRes val res: Int) {
    Ground(R.drawable.art_ground),
    Band(R.drawable.art_band),
    Hero(R.drawable.art_hero),
    Cave(R.drawable.art_cave),

    /** The LX seal v4 (web/public/seal-400.png): the medallion without a portrait. */
    Seal(R.drawable.art_seal),
}

/** The ground tile's side in dp: the web's `GROUND_TILE_PX`, drawn 1:1. */
private val GROUND_TILE = 1600.dp

/**
 * The screen ground (spec Mobile, round 5): the painted brown tile at 1:1 over the earth colour,
 * edge to edge. iOS: `screenGround()`.
 */
fun Modifier.screenGround(): Modifier = composed {
    val tile = ImageBitmap.imageResource(BrandArt.Ground.res)
    drawBehind {
        drawRect(FsColor.ground)
        val side = GROUND_TILE.toPx()
        val size = IntSize(side.roundToInt(), side.roundToInt())
        var y = 0f
        while (y < this.size.height) {
            var x = 0f
            while (x < this.size.width) {
                drawImage(tile, dstOffset = IntOffset(x.roundToInt(), y.roundToInt()), dstSize = size)
                x += side
            }
            y += side
        }
    }
}

/** Which painting a band shows: the blue band on every tab but Home, the camp-road hero on Home. */
enum class BandKind(val art: BrandArt, val ground: Color, val veil: Color) {
    Band(BrandArt.Band, FsColor.band, FsColor.veilBand),
    Hero(BrandArt.Hero, FsColor.night, FsColor.veilNight),
}

/**
 * The painted art behind the bars under its veil, stretched to cover and never tiled (spec round
 * 5), with the black bottom edge when [bottomLine]: the top app bar's ground, Home's hero bands and
 * the signed-out screens' full-bleed ground (without the line). iOS: `surfaceBand` / `BandArt`.
 */
fun Modifier.surfaceBand(bottomLine: Boolean = true, kind: BandKind = BandKind.Band): Modifier = composed {
    val art = ImageBitmap.imageResource(kind.art.res)
    drawBehind {
        drawBand(art, kind, size)
        if (bottomLine) drawLine(Color.Black, Offset(0f, size.height), Offset(size.width, size.height), 1.dp.toPx())
    }
}

/**
 * The head fade (web ui/band.css, the page's last band piece): the painted [kind] behind this area,
 * run on [FsSize.bandFade] past its foot and faded there into the screen ground on the token ramp
 * ([FsGradient.bandFade]). No rule at the foot; whatever is laid out below draws over the fade. The
 * tab roots' scaffold (FlowBandPage) wears it. iOS: `FadingBand`.
 */
fun Modifier.fadingBand(kind: BandKind = BandKind.Band): Modifier = composed {
    val art = ImageBitmap.imageResource(kind.art.res)
    drawBehind {
        val area = Size(size.width, size.height + FsSize.bandFade.toPx())
        drawIntoCanvas { canvas ->
            canvas.saveLayer(Rect(Offset.Zero, area), Paint())
            drawBand(art, kind, area)
            drawRect(fadeMask(foot = size.height, end = area.height), size = area, blendMode = BlendMode.DstIn)
            canvas.restore()
        }
    }
}

/** Opaque down to [foot], then the token ramp to clear at [end]: the ramp's positions are its share of the fade. */
private fun fadeMask(foot: Float, end: Float): Brush {
    val ramp = FsGradient.bandFade.map { (position, color) -> (foot + position * (end - foot)) / end to color }
    return Brush.verticalGradient(*(listOf(0f to Color.Black) + ramp).toTypedArray(), startY = 0f, endY = end)
}

/** The band's ground colour, its art scaled to cover [area] and its veil. */
private fun DrawScope.drawBand(art: ImageBitmap, kind: BandKind, area: Size) {
    drawRect(kind.ground, size = area)
    drawCover(art, area)
    drawRect(kind.veil, size = area)
}

/** [image] scaled to cover [area] from the top left, centred, cropped at its edges. */
private fun DrawScope.drawCover(image: ImageBitmap, area: Size = size) {
    val scale = max(area.width / image.width, area.height / image.height)
    val drawn = IntSize((image.width * scale).roundToInt(), (image.height * scale).roundToInt())
    val offset = IntOffset(((area.width - drawn.width) / 2).roundToInt(), ((area.height - drawn.height) / 2).roundToInt())
    clipRect(right = area.width, bottom = area.height) { drawImage(image, dstOffset = offset, dstSize = drawn) }
}

/**
 * A panel's frame (spec round 5, "The frame"): none, the plain 1 dp cream line inset 6 dp, or the
 * lead panel's ornate frame (the line plus the owner's corner filigree). Two or three framed panels
 * a screen, one of them ornate.
 */
enum class PanelFrame { None, Plain, Ornate }

/**
 * The round-5 panel (`.ui-surface`): their card brown-black, a black hairline inside the warm outer
 * ring, the 10 dp radius; optionally framed, optionally on the cave art ([cave], under its veil).
 * Draw the ornate corners with [FrameCorners] in the panel's Box. iOS: `.panel(frame:art:)`.
 */
fun Modifier.panel(frame: PanelFrame = PanelFrame.None, cave: Boolean = false): Modifier = composed {
    val art = if (cave) ImageBitmap.imageResource(BrandArt.Cave.res) else null
    val shape = RoundedCornerShape(FsRadius.plate)
    this
        .drawBehind {
            // The warm ring outside the black hairline (`box-shadow: 0 0 0 1px`).
            val ring = 1.dp.toPx()
            drawRoundRect(
                FsColor.panelEdge,
                topLeft = Offset(-ring / 2, -ring / 2),
                size = Size(size.width + ring, size.height + ring),
                cornerRadius = CornerRadius(FsRadius.plate.toPx() + ring / 2),
                style = Stroke(ring),
            )
        }
        .clip(shape)
        .background(FsColor.raised)
        .drawBehind {
            if (art != null) {
                drawCover(art)
                drawRect(FsColor.veilCave)
            }
        }
        .border(1.dp, Color.Black, shape)
        .drawWithContent {
            drawContent()
            if (frame != PanelFrame.None) {
                val inset = FRAME_INSET.toPx()
                drawRoundRect(
                    FsColor.frameLine,
                    topLeft = Offset(inset, inset),
                    size = Size(size.width - inset * 2, size.height - inset * 2),
                    cornerRadius = CornerRadius((FsRadius.plate - CORNER_INSET).toPx()),
                    style = Stroke(1.dp.toPx()),
                )
            }
        }
}

/** The lead card: the plain-framed panel. */
fun Modifier.primaryCard(): Modifier = panel(PanelFrame.Plain)

/** A list card: the plain panel. */
fun Modifier.quietCard(): Modifier = panel()

/**
 * The lead panel's four corner ornaments: the owner's filigree (design/assets/ornaments, the
 * `ic_frame_corner` vector), cream at 80%, mirrored per corner. Place it last in the panel's Box
 * with `Modifier.matchParentSize()`. Decorative only. iOS: `FrameCorners`.
 */
@Composable
fun FrameCorners(modifier: Modifier = Modifier) {
    Box(modifier.padding(CORNER_INSET)) {
        val corner = painterResource(R.drawable.ic_frame_corner)
        val tint = FsColor.cream.copy(alpha = CORNER_ALPHA)
        listOf(
            Alignment.TopStart to (1f to 1f),
            Alignment.TopEnd to (-1f to 1f),
            Alignment.BottomEnd to (-1f to -1f),
            Alignment.BottomStart to (1f to -1f),
        ).forEach { (alignment, flip) ->
            Icon(
                corner,
                contentDescription = null,
                tint = tint,
                modifier = Modifier
                    .align(alignment)
                    .size(CORNER_SIZE)
                    .graphicsLayer(scaleX = flip.first, scaleY = flip.second),
            )
        }
    }
}

private val FRAME_INSET = 6.dp
private val CORNER_INSET = 4.dp
private val CORNER_SIZE = 44.dp
private const val CORNER_ALPHA = 0.8f
