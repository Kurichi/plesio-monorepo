import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:kiikuten/presentation/designsystem/style/color.dart';

class KiikutenTheme {
  static ThemeData theme(BuildContext context) => ThemeData(
        useMaterial3: true,
        colorScheme: ColorScheme.fromSeed(
          seedColor: KiikutenColor.key,
        ),
        textTheme: kIsWeb
            ? GoogleFonts.dotGothic16TextTheme(Theme.of(context).textTheme)
            : null,
        // fontFamilyFallback:
        //     kIsWeb ? const ['${FontFamily.notoSansJP} Regular'] : null,
      );
}
