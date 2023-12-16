import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:kiikuten/firebase_options.dart';
import 'package:kiikuten/presentation/page/auth/auth.dart';

Future<void> main() async {
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );

  await dotenv.load(fileName: '.env');

  runApp(
    const ProviderScope(
      child: KiikutenApp(),
    ),
  );
}

class KiikutenApp extends StatelessWidget {
  const KiikutenApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: '木育展',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        useMaterial3: true,
        colorScheme: ColorScheme.fromSeed(
          seedColor: const Color(0xFF166718),
        ),
        textTheme: kIsWeb
            ? GoogleFonts.dotGothic16TextTheme(Theme.of(context).textTheme)
            : null,
        // fontFamilyFallback:
        //     kIsWeb ? const ['${FontFamily.notoSansJP} Regular'] : null,
      ),
      home: const KiikutenAuth(),
    );
  }
}
