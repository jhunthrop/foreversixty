package gg.foreversixty.app.items

import android.content.Context
import androidx.compose.foundation.Image
import gg.foreversixty.app.designsystem.FsType
import androidx.compose.ui.unit.sp
import androidx.compose.ui.graphics.ColorMatrix
import androidx.compose.ui.graphics.ColorFilter
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.Alignment
import androidx.compose.runtime.setValue
import androidx.compose.runtime.remember
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.getValue
import androidx.compose.material3.Text
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.runtime.Composable
import androidx.compose.runtime.staticCompositionLocalOf
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.geometry.CornerRadius
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.PathEffect
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import coil3.ImageLoader
import coil3.compose.AsyncImage
import coil3.disk.DiskCache
import coil3.disk.directory
import coil3.request.ImageRequest
import coil3.request.crossfade
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.network.StaticDataClient

/**
 * Where game icons (items and specs) come from. Live: `/data/{build}/icons/{icon}.webp` on the site through Coil
 * with a 200 MB disk cache (spec §11). Screenshot tests provide a painter from the repo's own
 * icon files instead.
 */
fun interface ItemIconPainterSource {
    @Composable
    fun painter(icon: String, build: String): Painter?
}

/** Null in the app (Coil loads the CDN file); set by screenshot tests. */
val LocalItemIconPainter = staticCompositionLocalOf<ItemIconPainterSource?> { null }

/** The one icon loader, its disk cache under the app's cache directory. */
object ItemIconLoader {
    private const val DISK_BYTES = 200L * 1024 * 1024
    private const val FADE_MS = 120

    @Volatile
    private var loader: ImageLoader? = null

    fun get(context: Context): ImageLoader = loader ?: synchronized(this) {
        loader ?: ImageLoader.Builder(context.applicationContext)
            .diskCache {
                DiskCache.Builder()
                    .directory(context.applicationContext.cacheDir.resolve("item-icons"))
                    .maxSizeBytes(DISK_BYTES)
                    .build()
            }
            .crossfade(FADE_MS)
            .build()
            .also { loader = it }
    }

    fun url(icon: String, build: String): String = "${StaticDataClient.SITE_URL}/data/$build/icons/$icon.webp"
}

/**
 * A square game icon in its quality border (spec §4.3): 1.5 dp quality border, a 1 dp inner black
 * line and the quality glow. The tile draws at once on `#131824`; the image fades in, and a
 * missing icon leaves the tile, never a broken-image glyph (§8). Hidden from TalkBack: the name
 * beside it carries the item.
 */
@Composable
fun ItemIcon(icon: String?, quality: Int?, build: String, size: Dp, modifier: Modifier = Modifier, dashed: Boolean = false) {
    val color = ItemQuality(quality).border
    val shape = RoundedCornerShape(FsRadius.pill)
    val frame = if (dashed) {
        Modifier.drawBehind {
            drawRoundRect(
                color = FsColor.line,
                cornerRadius = CornerRadius(FsRadius.pill.toPx()),
                style = Stroke(width = 1.5.dp.toPx(), pathEffect = PathEffect.dashPathEffect(floatArrayOf(8f, 6f))),
            )
        }
    } else {
        Modifier
            .shadow(6.dp, shape, ambientColor = color.copy(alpha = GLOW), spotColor = color.copy(alpha = GLOW))
            .border(1.5.dp, color, shape)
    }
    Box(modifier.size(size).then(frame).clip(shape).background(FsColor.cardTop)) {
        if (icon != null && !dashed) GameIconImage(icon, build, size)
        if (!dashed) {
            Box(Modifier.size(size).padding(1.5.dp).border(1.dp, Color.Black, shape))
        }
    }
}

/**
 * A spec's talent-tab icon (spec §4.3): square, the pill radius, a 1 dp line; from the same loader.
 * Spell and talent icons use it too; a null or missing icon leaves the bordered tile.
 */
@Composable
fun SpecIcon(icon: String?, build: String, size: Dp, modifier: Modifier = Modifier) {
    val shape = RoundedCornerShape(FsRadius.pill)
    Box(modifier.size(size).clip(shape).background(FsColor.cardTop).border(1.dp, GameColors.iconLine, shape)) {
        GameIconImage(icon, build, size)
    }
}

/**
 * The icon file itself: the test painter when one is provided, else Coil over the CDN. With a
 * [fallback], the tile shows that short label until the icon loads, or for good when it cannot;
 * [dimmed] greys and fades only the art, so a fallback label stays readable (WCAG AA).
 */
@Composable
fun GameIconImage(icon: String?, build: String?, size: Dp, fallback: String? = null, dimmed: Boolean = false) {
    var failed by remember(icon, build) { mutableStateOf(false) }
    val art = if (dimmed) Modifier.size(size).alpha(DIMMED_ALPHA) else Modifier.size(size)
    val filter = if (dimmed) ColorFilter.colorMatrix(ColorMatrix().apply { setToSaturation(0f) }) else null
    val local = LocalItemIconPainter.current
    val painter = if (local != null && icon != null && build != null) local.painter(icon, build) else null
    val url = if (icon != null && build != null) ItemIconLoader.url(icon, build) else null
    val loads = url != null && !failed && (local == null || painter != null)
    Box(Modifier.size(size), contentAlignment = Alignment.Center) {
        if (fallback != null && !loads) GameIconLabel(fallback, dimmed)
        if (loads && painter != null) {
            Image(painter, null, art, contentScale = ContentScale.Crop, colorFilter = filter)
        } else if (loads) {
            val context = LocalContext.current
            AsyncImage(
                model = ImageRequest.Builder(context).data(url).build(),
                imageLoader = ItemIconLoader.get(context),
                contentDescription = null,
                contentScale = ContentScale.Crop,
                colorFilter = filter,
                modifier = art,
                onError = { failed = true },
            )
        }
    }
}

/**
 * A game icon's short label (`MotW`, `Bless`): strong, or muted for a dimmed (missing) tile, both
 * AA on the card. Drawn on a tile whose art cannot load, and under every raid buff tile.
 */
@Composable
fun GameIconLabel(text: String, dimmed: Boolean) {
    // In dp, not sp: the tile is a fixed size, so its label does not grow with the font scale.
    val size = with(LocalDensity.current) { FALLBACK_SIZE.toSp() }
    Text(
        text,
        style = FsType.micro.copy(letterSpacing = 0.sp, fontSize = size),
        color = if (dimmed) FsColor.muted else FsColor.strong,
        maxLines = 1,
        softWrap = false,
    )
}

private const val DIMMED_ALPHA = 0.35f
private val FALLBACK_SIZE = 11.dp

private const val GLOW = 0.22f
