import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:kiikuten/data/datasource/tree_datasource.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'tree_datasource_provider.g.dart';

@riverpod
TreeDataSource treeDataSource(TreeDataSourceRef ref) {
  return TreeDataSource(dotenv.env['BASE_URL'] ?? '');
}
