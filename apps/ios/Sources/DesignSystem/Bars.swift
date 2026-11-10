import SwiftUI
import UIKit

/// The bars in round 5 (design/specs/2026-11-09-directions.md, Mobile and Tab-bar icons): the
/// painted blue band under its veil behind every navigation bar, the inline title in Open Sans 600
/// white and the large title in Marcellus cream; the tab bar on the nav-bar teal-black with the
/// owner's icons, gold `#ffd36b` when selected and muted white otherwise. One call at launch
/// (`ForeverSixtyApp.init`) styles every bar the app draws, so screens never style their own.
@MainActor
public enum BarAppearance {
    public static func install() {
        let navigation = navigationBar()
        UINavigationBar.appearance().standardAppearance = navigation
        UINavigationBar.appearance().scrollEdgeAppearance = navigation
        UINavigationBar.appearance().compactAppearance = navigation
        let tabs = tabBar()
        UITabBar.appearance().standardAppearance = tabs
        UITabBar.appearance().scrollEdgeAppearance = tabs
    }

    static func navigationBar() -> UINavigationBarAppearance {
        let appearance = UINavigationBarAppearance()
        appearance.configureWithOpaqueBackground()
        appearance.backgroundColor = UIColor(Theme.Palette.band)
        appearance.backgroundImage = veiledBand()
        appearance.backgroundImageContentMode = .scaleAspectFill
        appearance.shadowColor = .black
        return titled(appearance)
    }

    /// A round-5 bar at rest (`BandBar`): clear, so the page's band art shows behind the back or profile
    /// button and the large title, with the same title faces. `hidesTitle`: a pushed screen's bar shows
    /// no title at rest, its head names the screen (N3).
    public static func clearEdge(hidesTitle: Bool = false) -> UINavigationBarAppearance {
        let appearance = UINavigationBarAppearance()
        appearance.configureWithTransparentBackground()
        let titled = titled(appearance)
        if hidesTitle { titled.titleTextAttributes[.foregroundColor] = UIColor.clear }
        return titled
    }

    private static func titled(_ appearance: UINavigationBarAppearance) -> UINavigationBarAppearance {
        var title: [NSAttributedString.Key: Any] = [.foregroundColor: UIColor(Theme.Palette.strong)]
        if let font = Theme.Typography.barTitle { title[.font] = font }
        appearance.titleTextAttributes = title
        var large: [NSAttributedString.Key: Any] = [.foregroundColor: UIColor(Theme.Palette.title)]
        if let font = Theme.Typography.largeTitle { large[.font] = font }
        appearance.largeTitleTextAttributes = large
        return appearance
    }

    static func tabBar() -> UITabBarAppearance {
        let appearance = UITabBarAppearance()
        appearance.configureWithOpaqueBackground()
        appearance.backgroundColor = UIColor(Theme.Palette.bar)
        appearance.shadowColor = UIColor.white.withAlphaComponent(tabBarEdgeAlpha)
        for layout in [appearance.stackedLayoutAppearance, appearance.inlineLayoutAppearance, appearance.compactInlineLayoutAppearance] {
            style(layout.normal, color: UIColor(Theme.Palette.tabUnselected))
            style(layout.selected, color: UIColor(Theme.Palette.tabSelected))
        }
        return appearance
    }

    /// The tab bar's top edge: white at 8% (the mock's `.tabbar`).
    private static let tabBarEdgeAlpha: CGFloat = 0.08

    private static func style(_ item: UITabBarItemStateAppearance, color: UIColor) {
        item.iconColor = color
        var title: [NSAttributedString.Key: Any] = [.foregroundColor: color]
        if let font = Theme.Typography.tab { title[.font] = font }
        item.titleTextAttributes = title
    }

    /// The band with its veil drawn in, since a bar's background takes one image.
    private static func veiledBand() -> UIImage? {
        guard let band = BrandArt.band else { return nil }
        let veil = UIColor(FSColor.veilBand)
        return UIGraphicsImageRenderer(size: band.size).image { context in
            band.draw(at: .zero)
            veil.setFill()
            context.fill(CGRect(origin: .zero, size: band.size), blendMode: .normal)
        }
    }
}
