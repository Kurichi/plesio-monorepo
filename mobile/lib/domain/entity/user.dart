import 'package:kiikuten/domain/entity/tree.dart';

class User {
  final String id;
  final String username;
  final String avatarUrl;
  final Tree tree;

  User({
    required this.id,
    required this.username,
    required this.avatarUrl,
    required this.tree,
  });
}
