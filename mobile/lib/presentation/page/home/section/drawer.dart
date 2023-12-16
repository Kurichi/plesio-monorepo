import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/designsystem/component/kiikuten_avatar.dart';
import 'package:kiikuten/presentation/page/settings/settings_screen.dart';

class KiikutenDrawer extends StatelessWidget {
  const KiikutenDrawer({super.key});

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        children: [
          ListTile(
            title: Row(
              mainAxisSize: MainAxisSize.min,
              children: [
                IconButton(
                  onPressed: () {
                    Navigator.of(context).pop();
                  },
                  icon: const Icon(Icons.menu_open),
                  padding: const EdgeInsets.all(16),
                ),
                const Expanded(
                  child: Center(
                    child: Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Icon(Icons.park_outlined),
                        SizedBox(width: 8),
                        Text('木育展'),
                      ],
                    ),
                  ),
                ),
                const SizedBox(width: 64),
              ],
            ),
            contentPadding: EdgeInsets.zero,
          ),
          ListTile(
            leading: const KiikutenAvatar(size: 24),
            title: const Text('設定'),
            onTap: () {
              Navigator.pop(context);
              Navigator.of(context).push(
                MaterialPageRoute(
                  builder: (context) => const SettingsScreen(),
                ),
              );
            },
          ),
          ListTile(
            leading:
                Icon(Icons.logout, color: Theme.of(context).colorScheme.error),
            title: Text(
              'ログアウト',
              style: TextStyle(color: Theme.of(context).colorScheme.error),
            ),
            onTap: () {
              Navigator.pop(context);
              FirebaseAuth.instance.signOut();
            },
          ),
        ],
      ),
    );
  }
}
