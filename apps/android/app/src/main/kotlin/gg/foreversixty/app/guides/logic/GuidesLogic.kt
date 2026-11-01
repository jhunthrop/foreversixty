package gg.foreversixty.app.guides.logic

import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.GuideBuild
import gg.foreversixty.app.model.GuideFile
import gg.foreversixty.app.model.GuideIndex
import gg.foreversixty.app.model.GuideIndexClass
import gg.foreversixty.app.model.GuideIndexSpec
import gg.foreversixty.app.model.GuideRotation
import gg.foreversixty.app.model.GuideSpell
import gg.foreversixty.app.model.GuideTalents
import gg.foreversixty.app.network.StaticEndpoints
import java.time.LocalDate
import java.time.format.DateTimeFormatter
import java.time.format.DateTimeParseException
import java.util.Locale
import kotlin.math.roundToLong

/** A cell of `All classes` (spec F9.1). */
data class ClassCell(val classSlug: String, val className: String)

/** `Saved on this phone: {classes} · {size}` (spec F9.1 footer). */
data class SavedSummary(val classNames: List<String>, val bytes: Long)

/** A section of the guide reader, keyed by the file's `sections` names, in the spec's order. */
enum class GuideSection(val key: String) {
    Overview("overview"),
    Talents("talents"),
    Rotation("rotation"),
    Stats("stats"),
    Gear("gear"),
    Races("races"),
    ;

    companion object {
        fun of(key: String): GuideSection? = entries.firstOrNull { it.key == key }
    }
}

enum class RotationKind { Opener, Priority, Cooldowns }

data class RotationRow(val number: Int, val spell: GuideSpell)

data class RotationGroup(val kind: RotationKind, val rows: List<RotationRow>)

enum class BuildKind { Raid, Leveling }

/** A build row: `31 / 10 / 10`, summed per tree by the file. */
data class BuildRow(val kind: BuildKind, val totals: String, val build: GuideBuild)

/** F9's rules for the Guides tab and the reader. Nothing here sorts or rewrites the guide's own data. */
object Guides {
    private const val TOTALS_SEPARATOR = " / "
    private const val KB = 1024L
    private const val MB = KB * KB

    /** Nine cells, always: the index's classes in its (classes.json) order. */
    fun classCells(index: GuideIndex): List<ClassCell> = index.classes.map { ClassCell(it.classSlug, it.className) }

    /** The current character's guide: its class, and the spec its primary tree names. */
    fun yourGuide(index: GuideIndex, character: AccountCharacter?): GuideIndexSpec? {
        val classSlug = character?.className ?: return null
        val tree = character.spec ?: return null
        return specsOf(index, classSlug).firstOrNull { it.specName.equals(tree, ignoreCase = true) }
    }

    fun specsOf(index: GuideIndex, classSlug: String): List<GuideIndexSpec> =
        index.classes.firstOrNull { it.classSlug == classSlug }?.specs.orEmpty()

    /** The class a spec id belongs to in the index. */
    fun classOf(index: GuideIndex, spec: String): GuideIndexClass? = index.classes.firstOrNull { cls -> cls.specs.any { it.spec == spec } }

    fun spec(index: GuideIndex, spec: String): GuideIndexSpec? =
        index.classes.asSequence().flatMap { it.specs }.firstOrNull { it.spec == spec }

    /** Every guide of [build] that is on the phone, with the bytes it takes. */
    fun saved(index: GuideIndex, build: String, files: GuideFiles): Map<String, Long> =
        index.classes.flatMap { it.specs }
            .mapNotNull { spec -> files.savedBytes(StaticEndpoints.getGuide(build, spec.spec))?.let { spec.spec to it } }
            .toMap()

    /** The saved classes in grid order and their total size; null with nothing saved. */
    fun footer(index: GuideIndex, saved: Map<String, Long>): SavedSummary? {
        if (saved.isEmpty()) return null
        val classes = index.classes.filter { cls -> cls.specs.any { it.spec in saved } }.map { it.className }
        return SavedSummary(classes, saved.values.sum())
    }

    /** `47 KB` under a megabyte, else `2.4 MB`; never `0 KB`. */
    fun formatSize(bytes: Long): String = if (bytes < MB) {
        "${(bytes.toDouble() / KB).roundToLong().coerceAtLeast(1)} KB"
    } else {
        String.format(Locale.US, "%.1f MB", bytes.toDouble() / MB)
    }

    /** Exactly the sections the file names, in its order; an unknown name is skipped. */
    fun sections(guide: GuideFile): List<GuideSection> = guide.sections.mapNotNull(GuideSection::of)

    /** Opener, Priority, Cooldowns, each numbered from 1 in the guide's own order; empty lists left out. */
    fun rotation(rotation: GuideRotation?): List<RotationGroup> {
        if (rotation == null) return emptyList()
        return listOf(
            RotationKind.Opener to rotation.opener,
            RotationKind.Priority to rotation.priority,
            RotationKind.Cooldowns to rotation.cooldowns,
        ).filter { it.second.isNotEmpty() }
            .map { (kind, spells) -> RotationGroup(kind, spells.mapIndexed { i, spell -> RotationRow(i + 1, spell) }) }
    }

    /** The raid build, then the leveling build when it differs; a null raid build is the leveling build. */
    fun builds(talents: GuideTalents): List<BuildRow> {
        val raid = talents.raidBuild ?: talents.levelingBuild ?: return emptyList()
        val leveling = talents.levelingBuild?.takeIf { talents.raidBuild != null }
        return listOfNotNull(
            BuildRow(BuildKind.Raid, totals(raid), raid),
            leveling?.let { BuildRow(BuildKind.Leveling, totals(it), it) },
        )
    }

    private fun totals(build: GuideBuild) = build.totals.joinToString(TOTALS_SEPARATOR)

    /** A rotation line draws a tile when it names a spell (its icon, or the neutral tile); a general line has none. */
    fun hasTile(spell: GuideSpell): Boolean = spell.spellId != null || spell.icon != null

    /** The new-data strip shows when a different build is found while a build's data is on screen. */
    fun isNewData(shown: String?, latest: String?): Boolean = shown != null && latest != null && shown != latest

    /** The site page a guide shares (spec F9.2): `https://foreversixty.gg/guides/{class}/{spec}`. */
    fun shareLink(guide: GuideFile): String = "$SITE/guides/${guide.classSlug}/${guide.specSlug}"

    /** The spec id of a `/guides/{class}/{spec}` link: `shaman-elemental`. */
    fun specId(classSlug: String, specSlug: String): String = "$classSlug-$specSlug"

    private const val SITE = "https://foreversixty.gg"
}

/** The data's ISO days in the app's words, read as calendar days (no zone shift). */
object GuideDates {
    private val LONG: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM yyyy", Locale.US)
    private val SHORT: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.US)
    private val MONTH_DAY: DateTimeFormatter = DateTimeFormatter.ofPattern("MM/dd", Locale.US)

    /** `24 Sep 2026`. */
    fun long(iso: String): String = format(iso, LONG)

    /** `24 Sep`. */
    fun short(iso: String): String = format(iso, SHORT)

    /** `10/28`, the changelog row's date. */
    fun monthDay(iso: String): String = format(iso, MONTH_DAY)

    fun parse(iso: String): LocalDate? = try {
        LocalDate.parse(iso.take(ISO_DAY_LENGTH))
    } catch (e: DateTimeParseException) {
        null
    }

    private fun format(iso: String, pattern: DateTimeFormatter): String = parse(iso)?.format(pattern) ?: iso

    private const val ISO_DAY_LENGTH = 10
}
