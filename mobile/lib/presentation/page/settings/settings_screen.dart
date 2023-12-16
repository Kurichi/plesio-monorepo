import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/page/settings/component/setting_column.dart';
import 'package:kiikuten/presentation/page/settings/component/setting_item.dart';
import 'package:kiikuten/presentation/page/settings/component/setting_title.dart';
import 'package:url_launcher/url_launcher_string.dart';

class SettingsScreen extends StatelessWidget {
  const SettingsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('設定'),
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 20),
        child: Center(
          child: SizedBox(
            width: MediaQuery.of(context).size.width,
            child: Column(
              mainAxisSize: MainAxisSize.min,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                const SettingTitle(
                  icon: Icons.account_circle,
                  title: 'アカウント',
                ),
                SettingColumn(
                  children: [
                    SettingItem(
                      icon: Icons.badge,
                      label: 'ユーザー名',
                      value: FirebaseAuth.instance.currentUser!.displayName,
                    ),
                    SettingItem(
                      icon: Icons.onetwothree,
                      label: 'UID',
                      value: FirebaseAuth.instance.currentUser!.uid,
                    ),
                    SettingItem(
                      navigate: () {
                        launchUrlString(
                          FirebaseAuth.instance.currentUser!.photoURL ?? '',
                        );
                      },
                      icon: Icons.image,
                      label: 'アイコン画像',
                      value:
                          FirebaseAuth.instance.currentUser!.photoURL ?? '未設定',
                    ),
                    SettingItem(
                      icon: Icons.email,
                      label: 'メールアドレス',
                      value: FirebaseAuth.instance.currentUser!.email,
                    ),
                    SettingItem(
                      navigate: () {
                        launchUrlString(
                          'https://${FirebaseAuth.instance.currentUser!.providerData.first.providerId}',
                        );
                      },
                      icon: Icons.verified_user,
                      label: '認証方法',
                      value: FirebaseAuth.instance.currentUser!.providerData
                          .map((e) => e.providerId)
                          .join(', '),
                    ),
                    // SettingItem(
                    //   navigate: () {},
                    //   icon: Icons.lock_reset,
                    //   label: 'パスワードの変更',
                    // ),
                    SettingItem(
                      icon: Icons.event_available,
                      label: '登録日時',
                      value: FirebaseAuth
                          .instance.currentUser!.metadata.creationTime
                          .toString(),
                    ),
                    SettingItem(
                      icon: Icons.event_repeat,
                      label: '最終ログイン日時',
                      value: FirebaseAuth
                          .instance.currentUser!.metadata.lastSignInTime
                          .toString(),
                    ),
                    SettingItem(
                      navigate: () async {
                        Navigator.of(context).pop();
                        await FirebaseAuth.instance.signOut();
                        // final mounted = context.mounted;
                        // if (!mounted) return;
                      },
                      icon: Icons.logout,
                      label: 'ログアウト',
                      isError: true,
                    ),
                  ],
                ),
                const SizedBox(height: 40),
                const SettingTitle(
                  icon: Icons.help_outline,
                  title: 'ヘルプ',
                ),
                SettingColumn(
                  children: [
                    SettingItem(
                      navigate: () => launchUrlString(
                        'https://github.com/Kurichi/plesio-monorepo',
                      ),
                      icon: Icons.menu_book,
                      label: 'このアプリについて',
                    ),
                    SettingItem(
                      navigate: () => launchUrlString(
                        'https://github.com/Kurichi/plesio-monorepo/issues',
                      ),
                      icon: Icons.chat_bubble,
                      label: 'お問い合わせ',
                    ),
                  ],
                ),
                const SizedBox(height: 40),
                const SettingTitle(
                  icon: Icons.library_books,
                  title: '規約等',
                ),
                SettingColumn(
                  children: [
                    SettingItem(
                      navigate: () {},
                      icon: Icons.gavel,
                      label: '利用規約',
                    ),
                    SettingItem(
                      navigate: () {},
                      icon: Icons.privacy_tip,
                      label: 'プライバシーポリシー',
                    ),
                    SettingItem(
                      navigate: () {
                        Navigator.of(context).push(
                          MaterialPageRoute(
                            builder: (context) => const LicensePage(),
                          ),
                        );
                      },
                      icon: Icons.balance,
                      label: 'ライセンス',
                    ),
                  ],
                ),
                const SizedBox(height: 20),
                Text(
                  '© ${DateTime.now().year} 木育展',
                  style: TextStyle(
                    fontSize: 12,
                    color: Theme.of(context).colorScheme.outline,
                  ),
                ),
                const SizedBox(height: 20),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
