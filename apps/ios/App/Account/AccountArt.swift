import DesignSystem
import SwiftUI
import UIKit

/// The circular ringed HD crest (spec §4.3, apps/shared/README.md "one crest language"): the
/// bundled crest clipped to a circle over the raised backing, a 2 pt ring in the class colour
/// at 60%, and its glow. Never a square class icon.
struct ClassCrest: View {
    let className: String?
    let size: CGFloat

    var body: some View {
        let color = ClassColor.of(className)
        ZStack {
            Circle().fill(FSColor.raised)
            if let image = Self.image(className) {
                Image(uiImage: image).resizable().scaledToFill()
            }
        }
        .frame(width: size, height: size)
        .clipShape(.circle)
        .overlay(Circle().stroke(color.opacity(0.6), lineWidth: 2))
        .shadow(color: color.opacity(0.35), radius: 7)
        .accessibilityHidden(true)
    }

    /// `web/public/icons/hd/crests/<class>.webp`, copied into the bundle as `crests` (project.yml).
    static func image(_ className: String?) -> UIImage? {
        guard let slug = className?.lowercased(),
              let url = Bundle.main.url(forResource: slug, withExtension: "webp", subdirectory: "crests")
        else { return nil }
        return UIImage(contentsOfFile: url.path)
    }
}

/// The LX seal (`web/public/seal-400.png`, copied to the bundle root by project.yml).
struct Seal: View {
    let size: CGFloat

    var body: some View {
        Group {
            if let url = Bundle.main.url(forResource: "seal-400", withExtension: "png"),
               let image = UIImage(contentsOfFile: url.path) {
                Image(uiImage: image).resizable().scaledToFit()
            } else {
                Circle().stroke(FSColor.gold, lineWidth: 2)
            }
        }
        .frame(width: size, height: size)
        .accessibilityHidden(true)
    }
}

/// The hero band ground (`--surface-band`): the ember radial over the warm dark.
struct BandBackground: View {
    var body: some View {
        ZStack {
            FSColor.bg
            RadialGradient(
                colors: [FSColor.ember.opacity(0.28), .clear],
                center: .top,
                startRadius: 0,
                endRadius: 520
            )
        }
        .ignoresSafeArea()
    }
}

/// The toast at the top of F1.1: F1.4 and the 401 line in wipe red, a finished deletion neutral.
struct NoticeToast: View {
    let text: String
    /// A failure is wipe-red with the alert glyph; anything else is neutral with the check.
    var isFailure = true

    private var tint: Color { isFailure ? FSColor.wipe : FSColor.line }

    var body: some View {
        HStack(alignment: .top, spacing: FSSpacing.space3) {
            Image(systemName: isFailure ? "exclamationmark.circle" : "checkmark.circle")
                .foregroundStyle(isFailure ? FSColor.wipe : FSColor.gold)
            Text(text)
                .font(ScaledType.body)
                .foregroundStyle(FSColor.strong)
                .frame(maxWidth: .infinity, alignment: .leading)
                .fixedSize(horizontal: false, vertical: true)
        }
        .padding(FSSpacing.space4)
        .background(FSColor.raised, in: .rect(cornerRadius: FSRadius.panel))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.panel).stroke(tint.opacity(0.45)))
        .accessibilityElement(children: .combine)
        .accessibilityAddTraits(.isStaticText)
    }
}
