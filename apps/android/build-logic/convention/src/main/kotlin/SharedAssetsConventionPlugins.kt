import com.android.build.api.variant.AndroidComponentsExtension
import org.gradle.api.DefaultTask
import org.gradle.api.Plugin
import org.gradle.api.Project
import org.gradle.api.file.DirectoryProperty
import org.gradle.api.file.FileSystemOperations
import org.gradle.api.file.RegularFileProperty
import org.gradle.api.tasks.InputDirectory
import org.gradle.api.tasks.InputFile
import org.gradle.api.tasks.OutputDirectory
import org.gradle.api.tasks.PathSensitive
import org.gradle.api.tasks.PathSensitivity
import org.gradle.api.tasks.TaskAction
import javax.inject.Inject

/**
 * apps/shared is the one source of the brand fonts and the launcher icon (lane T,
 * web/scripts). Nothing is checked in twice: these plugins copy the files into a generated
 * res directory of every variant, so a regenerated font or icon reaches the app on the next build.
 */
private const val SHARED_DIR = "../shared"

/** From apps/android, the Gradle root, to the repository root. */
private const val REPO_ROOT = "../.."

/**
 * The apps' faces from apps/shared/fonts, as `res/font/opensans_regular.ttf` and so on: Open Sans
 * 400, 600 and 700, Archivo 600 and Marcellus 400 (design/tokens/tokens.json `font`, Tokens.kt
 * `FsFont`; `make brand-fonts` writes the Open Sans and Archivo files from the site's subsets).
 * Barlow, Barlow Condensed and Cinzel stay in apps/shared/fonts for the older design mocks and the
 * Discord assets only.
 */
internal val APP_FONTS = listOf(
    "OpenSans-Regular.ttf",
    "OpenSans-SemiBold.ttf",
    "OpenSans-Bold.ttf",
    "ArchivoSemiBold-SemiBold.ttf",
    "Marcellus-Regular.ttf",
)

/**
 * The design system's shared files: the faces above as `res/font`, and the round-5 art from
 * apps/shared/brand/art (tools/art/build_art.py: the ground tile, the band behind the bars, Home's
 * hero and the raid card's cave) as `res/drawable-nodpi/art_<name>.jpg`, plus the LX seal
 * (web/public/seal-400.png, seal v4) as `art_seal.png` for the paperdoll medallion, so the design
 * system's surfaces (Surfaces.kt `screenGround`, `surfaceBand`, `panel`) draw the same files the
 * iOS app bundles.
 */
class SharedFontsConventionPlugin : Plugin<Project> {
    override fun apply(target: Project) = with(target) {
        val shared = rootProject.layout.projectDirectory.dir(SHARED_DIR)
        val fonts = tasks.register("copySharedFonts", CopySharedFonts::class.java) {
            sourceDir.set(shared.dir("fonts"))
            outputDir.set(layout.buildDirectory.dir("generated/sharedFonts"))
        }
        val art = tasks.register("copySharedArt", CopySharedArt::class.java) {
            sourceDir.set(shared.dir("brand/art"))
            seal.set(rootProject.layout.projectDirectory.file("$REPO_ROOT/web/public/seal-400.png"))
            outputDir.set(layout.buildDirectory.dir("generated/sharedArt"))
        }
        extensions.getByType(AndroidComponentsExtension::class.java).onVariants { variant ->
            variant.sources.res?.addGeneratedSourceDirectory(fonts, CopySharedFonts::outputDir)
            variant.sources.res?.addGeneratedSourceDirectory(art, CopySharedArt::outputDir)
        }
    }
}

/** The adaptive launcher icon in apps/shared/brand/android, as `@mipmap/ic_launcher`. */
class AppIconConventionPlugin : Plugin<Project> {
    override fun apply(target: Project) = with(target) {
        val copy = tasks.register("copyAppIcon", CopyAppIcon::class.java) {
            sourceDir.set(rootProject.layout.projectDirectory.dir("$SHARED_DIR/brand/android"))
            outputDir.set(layout.buildDirectory.dir("generated/appIcon"))
        }
        extensions.getByType(AndroidComponentsExtension::class.java).onVariants { variant ->
            variant.sources.res?.addGeneratedSourceDirectory(copy, CopyAppIcon::outputDir)
        }
    }
}

abstract class CopySharedFonts @Inject constructor(
    private val fs: FileSystemOperations,
) : DefaultTask() {
    @get:InputDirectory
    @get:PathSensitive(PathSensitivity.RELATIVE)
    abstract val sourceDir: DirectoryProperty

    @get:OutputDirectory
    abstract val outputDir: DirectoryProperty

    @TaskAction
    fun copy() {
        fs.sync {
            from(sourceDir) { include(APP_FONTS) }
            into(outputDir.dir("font"))
            rename { resourceName(it) }
        }
    }

    companion object {
        /** Android resource names are lowercase with underscores: OpenSans-Regular.ttf -> opensans_regular.ttf. */
        fun resourceName(file: String): String = file.lowercase().replace('-', '_')
    }
}

abstract class CopySharedArt @Inject constructor(
    private val fs: FileSystemOperations,
) : DefaultTask() {
    @get:InputDirectory
    @get:PathSensitive(PathSensitivity.RELATIVE)
    abstract val sourceDir: DirectoryProperty

    @get:InputFile
    @get:PathSensitive(PathSensitivity.RELATIVE)
    abstract val seal: RegularFileProperty

    @get:OutputDirectory
    abstract val outputDir: DirectoryProperty

    @TaskAction
    fun copy() {
        fs.sync {
            from(sourceDir) {
                include("*.jpg")
                rename { "art_$it" }
            }
            from(seal) { rename { "art_seal.png" } }
            into(outputDir.dir("drawable-nodpi"))
        }
    }
}

abstract class CopyAppIcon @Inject constructor(
    private val fs: FileSystemOperations,
) : DefaultTask() {
    @get:InputDirectory
    @get:PathSensitive(PathSensitivity.RELATIVE)
    abstract val sourceDir: DirectoryProperty

    @get:OutputDirectory
    abstract val outputDir: DirectoryProperty

    @TaskAction
    fun copy() {
        fs.delete { delete(outputDir) }
        // The adaptive icon and its round alias (minSdk 28 always has them) and their two 432 px
        // layers (108 dp at xxxhdpi), which the launcher scales down from the density-independent
        // bucket; then the legacy 48 dp PNGs per density, for any launcher that asks for a bitmap.
        fs.copy {
            from(sourceDir) { include("ic_launcher*.xml") }
            into(outputDir.dir("mipmap-anydpi-v26"))
        }
        fs.copy {
            from(sourceDir) { include("*.png") }
            into(outputDir.dir("mipmap-nodpi"))
        }
        fs.copy {
            from(sourceDir.dir("legacy"))
            into(outputDir)
        }
    }
}
