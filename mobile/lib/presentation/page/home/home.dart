import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/page/home/section/drawer.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: const Text('木育展'),
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
      drawer: const KiikutenDrawer(),
    );
  }
}
