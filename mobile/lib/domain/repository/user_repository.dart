import 'package:kiikuten/domain/entity/user.dart';

abstract class UserRepository {
  Future<User> getUser(String userId);
}
