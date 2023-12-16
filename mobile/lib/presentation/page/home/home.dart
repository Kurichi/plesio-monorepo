import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/designsystem/component/kiikuten_avatar.dart';
import 'package:kiikuten/presentation/page/home/section/drawer.dart';

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
            onPressed: () {},
            icon: const KiikutenAvatar(),
          ),
        ],
      ),
      body: Center(
        child: Row(
          children: [
            const Text('GitHub ID: '),
            Text(FirebaseAuth.instance.currentUser!.displayName ?? 'null'),
            Text(FirebaseAuth.instance.currentUser!.email ?? 'null'),
          ],
        ),
      ),
      backgroundColor: Theme.of(context).colorScheme.inversePrimary,
      drawer: const KiikutenDrawer(),
    );
  }
}
