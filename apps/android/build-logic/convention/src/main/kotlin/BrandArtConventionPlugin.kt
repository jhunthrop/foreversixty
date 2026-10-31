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
 * The HD class crests and the LX seal from the site's own files (web/public), copied into a
 * generated res directory, so nothing is checked in twice (spec §4.3: crests are bundled):
 * `@drawable/crest_<class>` and `@drawable/brand_seal`.
 */
class BrandArtConventionPlugin : Plugin<Project> {
    override fun apply(target: Project) = with(target) {
        val web = rootProject.layout.projectDirectory.dir("../../web/public")
        val copy = tasks.register("copyBrandArt", CopyBrandArt::class.java) {
            crestDir.set(web.dir("icons/hd/crests"))
            seal.set(web.file("seal-400.png"))
            outputDir.set(layout.buildDirectory.dir("generated/brandArt"))
        }
        extensions.getByType(AndroidComponentsExtension::class.java).onVariants { variant ->
            variant.sources.res?.addGeneratedSourceDirectory(copy, CopyBrandArt::outputDir)
        }
    }
}

abstract class CopyBrandArt @Inject constructor(
    private val fs: FileSystemOperations,
) : DefaultTask() {
    @get:InputDirectory
    @get:PathSensitive(PathSensitivity.RELATIVE)
    abstract val crestDir: DirectoryProperty

    @get:InputFile
    @get:PathSensitive(PathSensitivity.RELATIVE)
    abstract val seal: RegularFileProperty

    @get:OutputDirectory
    abstract val outputDir: DirectoryProperty

    @TaskAction
    fun copy() {
        fs.delete { delete(outputDir) }
        fs.copy {
            from(crestDir) { include("*.webp") }
            into(outputDir.dir("drawable-nodpi"))
            rename { "crest_$it" }
        }
        fs.copy {
            from(seal)
            into(outputDir.dir("drawable-nodpi"))
            rename { "brand_seal.png" }
        }
    }
}
