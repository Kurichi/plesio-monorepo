import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:kiikuten/firebase_options.dart';
import 'package:kiikuten/presentation/designsystem/style/theme.dart';
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
      theme: KiikutenTheme.theme(context),
      home: const KiikutenAuth(),
    );
  }
}
