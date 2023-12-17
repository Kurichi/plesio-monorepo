import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/designsystem/component/kiikuten_avatar.dart';
import 'package:kiikuten/presentation/designsystem/component/tree/kiikuten_seed.dart';
import 'package:kiikuten/presentation/designsystem/component/tree/kiikuten_tree.dart';
import 'package:kiikuten/presentation/page/home/section/drawer.dart';
import 'package:kiikuten/presentation/page/settings/settings_screen.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            const Icon(Icons.park_outlined),
            const SizedBox(width: 8),
            Text(
              '木育展',
              style: TextStyle(color: Theme.of(context).colorScheme.primary),
            ),
          ],
        ),
        actions: [
          IconButton(
            onPressed: () {
              Navigator.of(context).push(
                MaterialPageRoute(
                  builder: (context) => const SettingsScreen(),
                ),
              );
            },
            icon: const Hero(
              tag: 'avatar',
              child: KiikutenAvatar(),
            ),
          ),
        ],
      ),
      body: SingleChildScrollView(
        child: SizedBox(
          width: MediaQuery.of(context).size.width,
          child: Column(
            mainAxisSize: MainAxisSize.min,
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Row(
                mainAxisSize: MainAxisSize.min,
                mainAxisAlignment: MainAxisAlignment.center,
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  const Text('Mail: '),
                  Text(FirebaseAuth.instance.currentUser!.email ?? 'null'),
                ],
              ),
              const KiikutenSeed(size: 80),
              const SizedBox(height: 16),
              const KiikutenTree(size: 120),
              const KiikutenTree(size: 240),
            ],
          ),
        ),
      ),
      backgroundColor: Theme.of(context).colorScheme.inversePrimary,
      drawer: const KiikutenDrawer(),
    );
  }
}
