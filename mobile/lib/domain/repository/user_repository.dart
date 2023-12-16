import 'package:kiikuten/domain/entity/item.dart';
import 'package:kiikuten/domain/entity/tree.dart';
import 'package:kiikuten/domain/entity/user.dart';

abstract class UserRepository {
  Future<User> getUser(String userId);
  Future<void> useItem(String userId, Item item);
}

class UserRepositoryImpl implements UserRepository {
  @override
  Future<User> getUser(String userId) async {
    return User(
      id: '1',
      username: 'chnotchy',
      avatarUrl: 'https://github.com/chnotchy.png',
      tree: Tree(
        id: '1',
        growthLevel: 1,
      ),
    );
  }

  @override
  Future<void> useItem(String userId, Item item) async {
    final user = await getUser(userId);
    user.tree.grow(item.growthEffect);
  }
}
