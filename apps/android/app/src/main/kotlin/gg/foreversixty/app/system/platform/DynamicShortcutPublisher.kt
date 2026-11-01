package gg.foreversixty.app.system.platform

import android.content.Context
import android.content.Intent
import android.graphics.Bitmap
import android.graphics.BitmapFactory
import android.graphics.Canvas
import android.graphics.Paint
import android.graphics.Path
import android.graphics.RectF
import androidx.core.content.pm.ShortcutInfoCompat
import androidx.core.content.pm.ShortcutManagerCompat
import androidx.core.graphics.createBitmap
import androidx.core.graphics.drawable.IconCompat
import androidx.core.graphics.toColorInt
import androidx.core.net.toUri
import gg.foreversixty.app.MainActivity
import gg.foreversixty.app.R
import gg.foreversixty.app.common.crestOf
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.system.logic.DynamicShortcut
import androidx.compose.ui.graphics.toArgb

/**
 * Publishes the dynamic shortcuts (spec F15 Android): long-press and launcher search. A character
 * carries its circular ringed HD crest (one crest language) as an adaptive icon; the next raid
 * carries the calendar shortcut icon. A tap opens the same screen as in-app, through its app link.
 */
class DynamicShortcutPublisher(private val context: Context) {
    fun publish(shortcuts: List<DynamicShortcut>) {
        val infos = shortcuts.mapIndexed { rank, shortcut ->
            ShortcutInfoCompat.Builder(context, shortcut.id)
                .setShortLabel(shortcut.shortLabel)
                .setLongLabel(shortcut.longLabel)
                .setIcon(icon(shortcut))
                .setRank(rank)
                .setIntent(Intent(Intent.ACTION_VIEW, shortcut.link.url().toUri(), context, MainActivity::class.java))
                .build()
        }
        ShortcutManagerCompat.setDynamicShortcuts(context, infos)
    }

    fun clear() = ShortcutManagerCompat.removeAllDynamicShortcuts(context)

    private fun icon(shortcut: DynamicShortcut): IconCompat {
        val crest = shortcut.crestClass?.let(::crestOf) ?: return IconCompat.createWithResource(context, R.drawable.ic_shortcut_next_raid)
        return IconCompat.createWithAdaptiveBitmap(crestBitmap(crest, classColor(shortcut.crestClass).toArgb()))
    }

    /** The crest in a circle with its class-coloured ring, centred in the adaptive icon's safe zone. */
    private fun crestBitmap(crest: Int, ring: Int): Bitmap {
        val size = ADAPTIVE_PX
        val bitmap = createBitmap(size, size)
        val canvas = Canvas(bitmap)
        canvas.drawColor(GROUND.toColorInt())
        val radius = size * SAFE_ZONE / 2f
        val centre = size / 2f
        val circle = RectF(centre - radius, centre - radius, centre + radius, centre + radius)
        BitmapFactory.decodeResource(context.resources, crest)?.let { source ->
            canvas.save()
            canvas.clipPath(Path().apply { addOval(circle, Path.Direction.CW) })
            canvas.drawBitmap(source, null, circle, Paint(Paint.FILTER_BITMAP_FLAG))
            canvas.restore()
        }
        val stroke = Paint(Paint.ANTI_ALIAS_FLAG).apply {
            style = Paint.Style.STROKE
            strokeWidth = size * RING_WIDTH
            color = ring
        }
        canvas.drawOval(circle, stroke)
        return bitmap
    }

    private companion object {
        /** 108 dp at xxxhdpi. */
        const val ADAPTIVE_PX = 432
        const val SAFE_ZONE = 0.61f
        const val RING_WIDTH = 0.02f
        const val GROUND = "#FF1A1E28"
    }
}
