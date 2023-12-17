import 'package:kiikuten/data/datasource/user_datasource.dart';
import 'package:kiikuten/domain/entity/user.dart';
import 'package:kiikuten/domain/repository/user_repository.dart';

class UserRepositoryImpl implements UserRepository {
  final UserDataSource _userDataSource;

  UserRepositoryImpl(this._userDataSource);

  @override
  Future<User> getUser(String userId) async {
    final userModel = await _userDataSource.getUser(userId);
    return userModel.toEntity();
  }

  @override
  Future<void> signUp(/* parameters for sign up */) async {
    // Delegate to UserDataSource
  }

  @override
  Future<void> updateUser(/* parameters for update */) async {
    // Delegate to UserDataSource
  }
}
