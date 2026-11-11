package gg.foreversixty.app.common

import java.math.BigDecimal
import java.math.RoundingMode
import java.text.NumberFormat
import java.util.Locale
import kotlin.math.abs
import kotlin.math.roundToLong

/**
 * The house format for a figure (apps/shared/formats.json, owner ruling 2026-11-10): grouped and whole
 * from 100 (`182`, `1,196`), one decimal below (`18.0`, `12.2`). Gear's set figure and Activity's
 * trailing figures both print through it.
 */
object HouseFigure {
    /** From this size a figure prints whole. */
    private const val WHOLE_FROM = 100.0

    fun of(value: Double): String =
        if (abs(value) >= WHOLE_FROM) {
            NumberFormat.getIntegerInstance(Locale.US).format(value.roundToLong())
        } else {
            BigDecimal.valueOf(value).setScale(1, RoundingMode.HALF_UP).toPlainString()
        }

    /** A gain or a loss: `+158`, `+2.8`, `−0.8` (a true minus, as the server writes it). */
    fun signed(value: Double): String = if (value < 0) "$MINUS${of(-value)}" else "+${of(value)}"

    private const val MINUS = "−"
}
