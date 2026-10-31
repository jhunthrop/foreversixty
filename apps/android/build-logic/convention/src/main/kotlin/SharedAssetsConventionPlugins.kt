import com.android.build.api.variant.AndroidComponentsExtension
import org.gradle.api.DefaultTask
import org.gradle.api.Plugin
import org.gradle.api.Project
import org.gradle.api.file.DirectoryProperty
import org.gradle.api.file.FileSystemOperations
import org.gradle.api.tasks.InputDirectory
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

/** The six faces in apps/shared/fonts, as `res/font/barlow_regular.ttf` and so on. */
class SharedFontsConventionPlugin : Plugin<Project> {
    override fun apply(target: Project) = with(target) {
        val copy = tasks.register("copySharedFonts", CopySharedFonts::class.java) {
            sourceDir.set(rootProject.layout.projectDirectory.dir("$SHARED_DIR/fonts"))
            outputDir.set(layout.buildDirectory.dir("generated/sharedFonts"))
        }
        extensions.getByType(AndroidComponentsExtension::class.java).onVariants { variant ->
            variant.sources.res?.addGeneratedSourceDirectory(copy, CopySharedFonts::outputDir)
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
            from(sourceDir) { include("*.ttf") }
            into(outputDir.dir("font"))
            rename { resourceName(it) }
        }
    }

    companion object {
        /** Android resource names are lowercase with underscores: Barlow-Regular.ttf -> barlow_regular.ttf. */
        fun resourceName(file: String): String = file.lowercase().replace('-', '_')
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
        // The adaptive icon (minSdk 28 always has it) and its two 1024 px layers, which the
        // launcher scales down from the density-independent bucket.
        fs.copy {
            from(sourceDir) { include("ic_launcher.xml") }
            into(outputDir.dir("mipmap-anydpi-v26"))
        }
        fs.copy {
            from(sourceDir) { include("*.png") }
            into(outputDir.dir("mipmap-nodpi"))
        }
    }
}
