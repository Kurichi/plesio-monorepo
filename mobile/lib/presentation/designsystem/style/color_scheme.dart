import 'dart:math';
import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/designsystem/style/color.dart';

class KiikutenThemeData {
  final ColorScheme colorScheme;

  late final Color _primary;
  late final Color _background;
  late final Color leaf;
  late final Color trunk;

  KiikutenThemeData({
    required this.colorScheme,
    required Brightness brightness,
  }) {
    _primary = colorScheme.primary;
    _background = colorScheme.background;
    leaf = KiikutenColor.green;
    trunk = KiikutenColor.brown;
  }

  Color surfaceColorByAlpha(double alpha, {Color? foreground}) {
    final clr = foreground ?? _primary;
    return Color.alphaBlend(clr.withOpacity(alpha), _background);
  }

  Color surfaceColorByElevation(double elevation) {
    if (elevation <= 0.0) return _background;
    final alpha = ((4.5 * log(elevation + 1)) + 2) / 100;
    return surfaceColorByAlpha(alpha);
  }
}

extension KiikutenThemeDataExtension on ThemeData {
  KiikutenThemeData get kiikutenColorScheme =>
      KiikutenThemeData(colorScheme: colorScheme, brightness: brightness);
}
