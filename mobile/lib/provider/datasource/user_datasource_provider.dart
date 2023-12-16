import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:kiikuten/data/datasource/user_datasource.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'user_datasource_provider.g.dart';

@riverpod
UserDataSource userDataSource(UserDataSourceRef ref) {
  return UserDataSource(dotenv.env['BASE_URL'] ?? '');
}
