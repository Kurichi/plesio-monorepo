import 'package:kiikuten/data/repository/tree_repository_impl.dart';
import 'package:kiikuten/domain/repository/tree_repository.dart';
import 'package:kiikuten/provider/datasource/tree_datasource_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'tree_repository_provider.g.dart';

@riverpod
TreeRepository treeRepository(TreeRepositoryRef ref) {
  final treeDataSource = ref.watch(treeDataSourceProvider);
  return TreeRepositoryImpl(treeDataSource);
}
