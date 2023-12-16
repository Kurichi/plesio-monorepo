import 'package:kiikuten/data/repository/user_repository_impl.dart';
import 'package:kiikuten/domain/repository/user_repository.dart';
import 'package:kiikuten/provider/datasource/user_datasource_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'user_repository_provider.g.dart';

@riverpod
UserRepository userRepository(UserRepositoryRef ref) {
  final userDataSource = ref.watch(userDataSourceProvider);
  return UserRepositoryImpl(userDataSource);
}
