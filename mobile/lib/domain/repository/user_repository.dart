import 'package:kiikuten/domain/entity/tree.dart';
import 'package:kiikuten/domain/entity/user.dart';

abstract class UserRepository {
  Future<User> getUser(String userId);
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
}
