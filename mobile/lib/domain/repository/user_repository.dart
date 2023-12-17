import 'package:kiikuten/domain/entity/user.dart';

abstract class UserRepository {
  Future<User> getUser(String userId);
  Future<void> signUp(/* parameters for sign up */);
  Future<void> updateUser(/* parameters for update */);
}
